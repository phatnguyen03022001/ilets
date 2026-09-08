#!/usr/bin/env python3
from __future__ import annotations

import contextlib
import importlib.util
import io
import json
import shutil
import subprocess
import tempfile
import unittest
from pathlib import Path

MODULE_PATH = Path(__file__).with_name("materialize.py")
spec = importlib.util.spec_from_file_location("canonical_materializer", MODULE_PATH)
assert spec and spec.loader
module = importlib.util.module_from_spec(spec)
spec.loader.exec_module(module)


class MaterializerTests(unittest.TestCase):
    def test_slice_discovery_is_sorted_and_ignores_generated_output(self):
        with tempfile.TemporaryDirectory() as tmp:
            slice_dir = Path(tmp)
            (slice_dir / "zeta.json").write_text("{}\n")
            (slice_dir / "alpha.json").write_text("{}\n")
            (slice_dir / "generated").mkdir()
            (slice_dir / "generated" / "ignored.json").write_text("{}\n")

            original_slice_dir = module.SLICE_INPUT_DIR
            module.SLICE_INPUT_DIR = slice_dir
            try:
                self.assertEqual(
                    [path.name for path in module.slice_input_paths()],
                    ["alpha.json", "zeta.json"],
                )
            finally:
                module.SLICE_INPUT_DIR = original_slice_dir

    def test_each_slice_writes_a_stem_based_trace_path(self):
        self.assertEqual(
            module.trace_output_path(Path("tools/slice/reading-training.json")),
            module.ROOT / "tools/slice/generated/reading-training-trace.json",
        )

    def test_multiple_slices_share_one_global_registry(self):
        source_slice = module.SLICE_INPUT_DIR / "reading-training.json"
        with tempfile.TemporaryDirectory() as tmp:
            slice_dir = Path(tmp)
            first = slice_dir / "zeta.json"
            second = slice_dir / "alpha.json"
            shutil.copyfile(source_slice, first)
            shutil.copyfile(source_slice, second)

            original_slice_dir = module.SLICE_INPUT_DIR
            module.SLICE_INPUT_DIR = slice_dir
            try:
                first = module.materialize()
                second = module.materialize()
            finally:
                module.SLICE_INPUT_DIR = original_slice_dir

            self.assertEqual(first, second)
            registry, traces = first
            self.assertEqual([path.name for path, _ in traces], ["alpha.json", "zeta.json"])
            self.assertEqual(
                [module.trace_output_path(path).name for path, _ in traces],
                ["alpha-trace.json", "zeta-trace.json"],
            )
            self.assertEqual(traces[0][1], traces[1][1])
            self.assertEqual(
                traces[0][1]["canonical_registry_sha256"],
                module.registry_sha256(registry),
            )
            self.assertEqual(
                {entry["id"] for entry in registry["entries"]},
                {
                    "ASSESSMENT",
                    "ASSESSMENT_MAY_ADMIT",
                    "AT-02",
                    "CTX-READING-ACADEMIC",
                    "CTX-LISTENING-SHARED",
                    "IELTS-R-QF-02",
                    "IELTS-R-QF-03",
                    "IELTS-R-QF-05",
                    "IELTS-L-QF-01",
                    "IELTS-L-QF-04",
                    "NOT_EVIDENCE_CANDIDATE",
                    "L-COMP-01",
                    "L-COMP-02",
                    "L-F03",
                    "L-F04",
                    "L-QT-01",
                    "PM-L02",
                    "PM-L03",
                    "PM-R03",
                    "PM-R04",
                    "PT-12",
                    "PT-13",
                    "PT-16",
                    "R-F04",
                    "R-F05",
                    "R-QT-01",
                    "R-QT-02",
                    "R-QT-03",
                    "TRAINING",
                },
            )

    def test_unknown_reference_in_any_slice_fails_against_shared_registry(self):
        source_slice = module.SLICE_INPUT_DIR / "listening-detail-completion.json"
        with tempfile.TemporaryDirectory() as tmp:
            slice_path = Path(tmp) / "listening-training.json"
            slice_config = json.loads(source_slice.read_text())
            slice_config["slice_id"] = "listening-training-bootstrap-v1"
            slice_config["feature_id"] = "L-F01"
            slice_path.write_text(json.dumps(slice_config) + "\n")

            original_slice_dir = module.SLICE_INPUT_DIR
            module.SLICE_INPUT_DIR = Path(tmp)
            try:
                with self.assertRaisesRegex(
                    ValueError,
                    "slice trace listening-training.json contains unknown canonical references: L-F01",
                ):
                    module.materialize()
            finally:
                module.SLICE_INPUT_DIR = original_slice_dir

    def test_strip_cell_unwraps_one_code_span(self):
        self.assertEqual(module.strip_cell("`PT-13`"), "PT-13")

    def test_strip_cell_preserves_multiple_code_spans(self):
        self.assertEqual(
            module.strip_cell("`LM-02`, `LM-04`"),
            "`LM-02`, `LM-04`",
        )

    def test_forbidden_sources_cannot_be_authority(self):
        config = json.loads(module.SOURCE_MAP_PATH.read_text())
        config["registries"].append(
            {
                "type": "fake",
                "owner": "archive/fake.md",
                "section": "# Fake",
                "kind": "markdown_table",
                "ids": ["R-F04"],
            }
        )
        with self.assertRaisesRegex(ValueError, "forbidden authority source"):
            module.validate_source_map(config)

    def test_id_looking_research_text_is_not_discovered(self):
        with tempfile.TemporaryDirectory() as tmp:
            fake = Path(tmp) / "research.md"
            fake.write_text("| ID | Thing |\n|---|---|\n| `R-F04` | fake |\n")
            owners = {
                entry["owner"]
                for entry in json.loads(module.SOURCE_MAP_PATH.read_text())[
                    "registries"
                ]
            }
            self.assertNotIn(str(fake), owners)
            self.assertTrue(
                all(not owner.startswith("research/") for owner in owners)
            )

    def test_duplicate_configured_id_fails(self):
        config = json.loads(module.SOURCE_MAP_PATH.read_text())
        config["registries"].append(dict(config["registries"][0]))
        with self.assertRaisesRegex(ValueError, "duplicate configured canonical id"):
            module.validate_source_map(config)

    def test_revision_is_deterministic(self):
        source_map = b'{"version":1}'
        owners = [
            ("docs/a.md", "aaa"),
            ("docs/b.md", "bbb"),
        ]
        first = module.canonical_source_revision_from_inputs(source_map, owners)
        second = module.canonical_source_revision_from_inputs(source_map, owners)
        self.assertEqual(first, second)

    def test_owner_order_does_not_change_revision(self):
        source_map = b'{"version":1}'
        owners = [
            ("docs/a.md", "aaa"),
            ("docs/b.md", "bbb"),
        ]
        forward = module.canonical_source_revision_from_inputs(source_map, owners)
        reverse = module.canonical_source_revision_from_inputs(
            source_map, list(reversed(owners))
        )
        self.assertEqual(forward, reverse)

    def test_owner_blob_change_changes_revision(self):
        source_map = b'{"version":1}'
        before = module.canonical_source_revision_from_inputs(
            source_map, [("docs/a.md", "aaa")]
        )
        after = module.canonical_source_revision_from_inputs(
            source_map, [("docs/a.md", "bbb")]
        )
        self.assertNotEqual(before, after)

    def test_source_map_change_changes_revision(self):
        owners = [("docs/a.md", "aaa")]
        before = module.canonical_source_revision_from_inputs(
            b'{"version":1}', owners
        )
        after = module.canonical_source_revision_from_inputs(
            b'{"version":2}', owners
        )
        self.assertNotEqual(before, after)

    def test_materialize_is_deterministic_for_current_sources(self):
        first = module.materialize()
        second = module.materialize()
        self.assertEqual(first, second)

    def test_materialize_does_not_read_git_history(self):
        original_git = module.git
        calls = []

        def recording_git(*args):
            calls.append(args)
            return original_git(*args)

        module.git = recording_git
        try:
            module.materialize()
        finally:
            module.git = original_git

        self.assertFalse(
            any(args and args[0] == "rev-list" for args in calls),
            calls,
        )

    def _commit_fixture(self, root: Path, message: str = "fixture") -> str:
        subprocess.run(["git", "add", "."], cwd=root, check=True, capture_output=True)
        subprocess.run(
            ["git", "commit", "-m", message],
            cwd=root,
            check=True,
            capture_output=True,
        )
        return subprocess.check_output(
            ["git", "rev-parse", "HEAD"], cwd=root, text=True
        ).strip()

    def _trace_fixture(self):
        tmp = tempfile.TemporaryDirectory()
        root = Path(tmp.name)
        subprocess.run(["git", "init", "-q"], cwd=root, check=True)
        subprocess.run(
            ["git", "config", "user.email", "trace@example.invalid"],
            cwd=root,
            check=True,
        )
        subprocess.run(
            ["git", "config", "user.name", "Trace Fixture"], cwd=root, check=True
        )
        for relative in [
            "docs/catalog",
            "tools/slice",
            "apps/web/src/features/today",
            "services/core-api/internal/httpapi",
            "apps/web/e2e",
        ]:
            (root / relative).mkdir(parents=True, exist_ok=True)
        (root / "docs/catalog/project.json").write_text(
            json.dumps(
                {
                    "features": [
                        {
                            "id": "FTR-018",
                            "name": "R-F04 — T/F/NG + Y/N/NG Lab",
                            "spec_ref": "docs/BEHAVIOR.md#FTR-018",
                        },
                        {
                            "id": "FTR-019",
                            "name": "R-F05 — Headings & Structure Lab",
                            "spec_ref": "docs/BEHAVIOR.md#FTR-019",
                        },
                        {
                            "id": "FTR-009",
                            "name": "L-F03 — Gist & Main-Idea Drill",
                            "spec_ref": "docs/BEHAVIOR.md#FTR-009",
                        },
                        {
                            "id": "FTR-030",
                            "name": "W-F08 — Rubric Feedback & Revision Diff",
                            "spec_ref": "docs/BEHAVIOR.md#FTR-030",
                        },
                        {
                            "id": "FTR-041",
                            "name": "X-F01 — Daily Plan",
                            "spec_ref": "docs/BEHAVIOR.md#FTR-041",
                        },
                    ]
                }
            )
            + "\n"
        )
        (root / "docs/BEHAVIOR.md").write_text(
            """### FTR-009 L-F03 — Gist & Main-Idea Drill

Bounded Listening gist behavior.

### FTR-018 R-F04 — T/F/NG + Y/N/NG Lab

| ID | Feature |
|---|---|
| `R-F04` | Classification |

### FTR-019 R-F05 — Headings & Structure Lab

| ID | Feature |
|---|---|
| `R-F05` | Headings |

### FTR-030 W-F08 — Rubric Feedback & Revision Diff

Bounded evaluator-heavy behavior without implementation proof.

### FTR-041 X-F01 — Daily Plan

Current eligible recommendation snapshot.
"""
        )
        slice_config = {
            "slice_id": "reading-training-bootstrap-v1",
            "feature_root_id": "FTR-018",
            "feature_id": "R-F04",
            "implementation_refs": [
                {
                    "path": "apps/web/src/features/today/today.tsx",
                    "required_literals": ["PM-R03"],
                },
                {
                    "path": "services/core-api/internal/httpapi/activity.go",
                    "required_literals": ["R-F04", "PM-R03"],
                },
            ],
            "verification_refs": [
                {
                    "path": "apps/web/e2e/reading-training.spec.ts",
                    "required_literals": ["T/F/NG + Y/N/NG", "PM-R03"],
                },
                {
                    "path": "services/core-api/internal/httpapi/integration_test.go",
                    "required_literals": ["PM-R03", "NOT_EVIDENCE_CANDIDATE"],
                },
            ],
        }
        (root / "tools/slice/reading-training.json").write_text(
            json.dumps(slice_config, indent=2) + "\n"
        )
        (root / "apps/web/src/features/today/today.tsx").write_text(
            'const reading = "PM-R03"; const gist = "PM-L02"; const dailyPlanQueryKey = ["daily-plan"]; const source = "TODAY";\n'
        )
        (root / "services/core-api/internal/httpapi/activity.go").write_text(
            'const readingFeature = "R-F04"\nconst readingMode = "PM-R03"\nconst listeningFeature = "L-F03"\nconst listeningMode = "PM-L02"\nconst dailyPlanItem = "DailyPlanItemId"\nconst recheck = "cannot recheck current content eligibility"\n'
        )
        (root / "apps/web/e2e/reading-training.spec.ts").write_text(
            'const label = "T/F/NG + Y/N/NG"; const readingMode = "PM-R03"; const gist = "Start Gist Sprint PM-L02"; const dailyPlan = "/v1/daily-plan Start recommended activity";\n'
        )
        (root / "services/core-api/internal/httpapi/integration_test.go").write_text(
            'const readingMode = "PM-R03"\nconst gist = "L-F03 PM-L02"\nconst candidacy = "NOT_EVIDENCE_CANDIDATE"\n'
        )
        (root / "services/core-api/internal/httpapi/planner.go").write_text(
            'const decision = "plannercore.Decide"\nconst persistence = "InsertDailyPlan"\n'
        )
        (root / "services/core-api/internal/httpapi/planner_integration_test.go").write_text(
            'const route = "/v1/daily-plan"\nconst freshness = "freshness is exact revision scoped across the bounded supply"\n'
        )
        subject = self._commit_fixture(root)
        return tmp, root, subject


    def test_feature_root_slice_materialization_rejects_conflicting_alias(self):
        config = json.loads((module.SLICE_INPUT_DIR / "reading-training.json").read_text())
        config["feature_id"] = "R-F99"
        with tempfile.TemporaryDirectory() as tmp:
            slice_dir = Path(tmp)
            (slice_dir / "conflict.json").write_text(json.dumps(config) + "\n")
            original_slice_dir = module.SLICE_INPUT_DIR
            module.SLICE_INPUT_DIR = slice_dir
            try:
                with self.assertRaisesRegex(ValueError, "conflicting feature mapping"):
                    module.materialize()
            finally:
                module.SLICE_INPUT_DIR = original_slice_dir

    def test_relation_only_slice_materializes_without_practice_semantics(self):
        relation_only = {
            "slice_id": "daily-plan-v1",
            "feature_root_id": "FTR-041",
            "feature_id": "X-F01",
            "implementation_refs": [
                {"path": "services/core-api/internal/httpapi/planner.go", "required_literals": ["plannercore.Decide"]}
            ],
            "verification_refs": [
                {"path": "services/core-api/internal/httpapi/planner_integration_test.go", "required_literals": ["/v1/daily-plan"]}
            ],
        }
        with tempfile.TemporaryDirectory() as tmp:
            slice_dir = Path(tmp)
            path = slice_dir / "daily-plan.json"
            path.write_text(json.dumps(relation_only, indent=2) + "\n")
            original_slice_dir = module.SLICE_INPUT_DIR
            module.SLICE_INPUT_DIR = slice_dir
            try:
                _, traces = module.materialize()
            finally:
                module.SLICE_INPUT_DIR = original_slice_dir
        self.assertEqual(len(traces), 1)
        self.assertEqual(traces[0][1]["slice"], relation_only)

    def test_legacy_practice_slice_without_feature_root_still_materializes(self):
        legacy = module.SLICE_INPUT_DIR / "listening-detail-completion.json"
        with tempfile.TemporaryDirectory() as tmp:
            slice_dir = Path(tmp)
            copied = slice_dir / legacy.name
            shutil.copyfile(legacy, copied)
            original_slice_dir = module.SLICE_INPUT_DIR
            module.SLICE_INPUT_DIR = slice_dir
            try:
                _, traces = module.materialize()
            finally:
                module.SLICE_INPUT_DIR = original_slice_dir
        self.assertEqual(traces[0][1]["slice"], json.loads(legacy.read_text()))

    def test_partial_practice_semantics_fail_closed(self):
        config = json.loads((module.SLICE_INPUT_DIR / "reading-training.json").read_text())
        config.pop("practice_mode_id")
        with tempfile.TemporaryDirectory() as tmp:
            slice_dir = Path(tmp)
            (slice_dir / "partial.json").write_text(json.dumps(config) + "\n")
            original_slice_dir = module.SLICE_INPUT_DIR
            module.SLICE_INPUT_DIR = slice_dir
            try:
                with self.assertRaisesRegex(ValueError, "incomplete specialized practice semantics"):
                    module.materialize()
            finally:
                module.SLICE_INPUT_DIR = original_slice_dir

    def test_representative_queries_preserve_truthful_positive_and_negative_cases(self):
        tmp, root, _ = self._trace_fixture()
        self.addCleanup(tmp.cleanup)
        listening = {
            "slice_id": "listening-gist-sprint-v1",
            "feature_root_id": "FTR-009",
            "feature_id": "L-F03",
            "implementation_refs": [
                {"path": "apps/web/src/features/today/today.tsx", "required_literals": ["PM-L02"]},
                {"path": "services/core-api/internal/httpapi/activity.go", "required_literals": ["L-F03", "PM-L02"]},
            ],
            "verification_refs": [
                {"path": "apps/web/e2e/reading-training.spec.ts", "required_literals": ["Start Gist Sprint", "PM-L02"]},
                {"path": "services/core-api/internal/httpapi/integration_test.go", "required_literals": ["L-F03", "PM-L02", "NOT_EVIDENCE_CANDIDATE"]},
            ],
        }
        daily_plan = {
            "slice_id": "daily-plan-v1",
            "feature_root_id": "FTR-041",
            "feature_id": "X-F01",
            "implementation_refs": [
                {"path": "apps/web/src/features/today/today.tsx", "required_literals": ["dailyPlanQueryKey", "TODAY"]},
                {"path": "services/core-api/internal/httpapi/activity.go", "required_literals": ["DailyPlanItemId", "cannot recheck current content eligibility"]},
                {"path": "services/core-api/internal/httpapi/planner.go", "required_literals": ["plannercore.Decide", "InsertDailyPlan"]},
            ],
            "verification_refs": [
                {"path": "apps/web/e2e/reading-training.spec.ts", "required_literals": ["/v1/daily-plan", "Start recommended activity"]},
                {"path": "services/core-api/internal/httpapi/planner_integration_test.go", "required_literals": ["/v1/daily-plan", "freshness is exact revision scoped across the bounded supply"]},
            ],
        }
        (root / "tools/slice/listening-gist-sprint.json").write_text(json.dumps(listening, indent=2) + "\n")
        (root / "tools/slice/daily-plan.json").write_text(json.dumps(daily_plan, indent=2) + "\n")
        subject = self._commit_fixture(root, "representative-traces")

        expected = {
            "FTR-009": ("L-F03", "TRACE_PRESENT"),
            "FTR-041": ("X-F01", "TRACE_PRESENT"),
            "FTR-030": ("W-F08", "IMPLEMENTATION_PROOF_NOT_ESTABLISHED"),
            "FTR-019": ("R-F05", "IMPLEMENTATION_PROOF_NOT_ESTABLISHED"),
        }
        for feature_root_id, (alias, condition) in expected.items():
            first = module.query_feature_trace(feature_root_id, subject, root=root)
            second = module.query_feature_trace(feature_root_id, subject, root=root)
            self.assertEqual(module.render(first), module.render(second))
            self.assertEqual(first["authority"], "NONE")
            self.assertEqual(first["subject_revision"], subject)
            self.assertEqual(first["aliases"], [alias])
            self.assertEqual(first["condition"], condition)
            self.assertNotIn("status", first)
            self.assertNotIn("release_ready", first)
            if condition == "IMPLEMENTATION_PROOF_NOT_ESTABLISHED":
                self.assertEqual(first["implementation_refs"], [])
                self.assertEqual(first["verification_refs"], [])

    def test_feature_trace_query_binds_ftr_root_to_alias_and_exact_subject(self):
        tmp, root, subject = self._trace_fixture()
        self.addCleanup(tmp.cleanup)
        result = module.query_feature_trace("FTR-018", subject, root=root)

        self.assertEqual(result["authority"], "NONE")
        self.assertEqual(result["condition"], "TRACE_PRESENT")
        self.assertEqual(result["subject_revision"], subject)
        self.assertEqual(result["feature_root_id"], "FTR-018")
        self.assertEqual(result["aliases"], ["R-F04"])
        self.assertEqual(
            [item["path"] for item in result["implementation_refs"]],
            [
                "apps/web/src/features/today/today.tsx",
                "services/core-api/internal/httpapi/activity.go",
            ],
        )
        self.assertEqual(
            [item["path"] for item in result["verification_refs"]],
            [
                "apps/web/e2e/reading-training.spec.ts",
                "services/core-api/internal/httpapi/integration_test.go",
            ],
        )

    def test_trace_query_rejects_stale_expected_revision(self):
        tmp, root, stale = self._trace_fixture()
        self.addCleanup(tmp.cleanup)
        (root / "marker.txt").write_text("new head\n")
        self._commit_fixture(root, "advance")
        with self.assertRaisesRegex(ValueError, "stale expected revision"):
            module.query_feature_trace("FTR-018", stale, root=root)


    def test_duplicate_canonical_feature_root_fails_closed(self):
        tmp, root, _ = self._trace_fixture()
        self.addCleanup(tmp.cleanup)
        path = root / "docs/catalog/project.json"
        catalog = json.loads(path.read_text())
        catalog["features"].append(dict(catalog["features"][0]))
        path.write_text(json.dumps(catalog) + "\n")
        subject = self._commit_fixture(root, "duplicate-canonical-root")
        with self.assertRaisesRegex(ValueError, "expected exactly one canonical feature root"):
            module.query_feature_trace("FTR-018", subject, root=root)

    def test_heading_alias_resolution_does_not_require_table_anchor(self):
        tmp, root, subject = self._trace_fixture()
        self.addCleanup(tmp.cleanup)
        result = module.query_feature_trace("FTR-041", subject, root=root)
        self.assertEqual(result["aliases"], ["X-F01"])
        self.assertEqual(result["condition"], "IMPLEMENTATION_PROOF_NOT_ESTABLISHED")

    def test_conflicting_behavior_heading_alias_fails_closed(self):
        tmp, root, _ = self._trace_fixture()
        self.addCleanup(tmp.cleanup)
        path = root / "docs/BEHAVIOR.md"
        path.write_text(path.read_text().replace("FTR-041 X-F01", "FTR-041 X-F99", 1))
        subject = self._commit_fixture(root, "conflicting-heading-alias")
        with self.assertRaisesRegex(ValueError, "conflicting behavior alias"):
            module.query_feature_trace("FTR-041", subject, root=root)

    def test_malformed_catalog_feature_name_fails_closed(self):
        tmp, root, _ = self._trace_fixture()
        self.addCleanup(tmp.cleanup)
        path = root / "docs/catalog/project.json"
        catalog = json.loads(path.read_text())
        feature = next(item for item in catalog["features"] if item["id"] == "FTR-041")
        feature["name"] = "Daily Plan"
        path.write_text(json.dumps(catalog) + "\n")
        subject = self._commit_fixture(root, "malformed-feature-name")
        with self.assertRaisesRegex(ValueError, "has no lower-level feature alias"):
            module.query_feature_trace("FTR-041", subject, root=root)

    def test_duplicate_behavior_heading_fails_closed(self):
        tmp, root, _ = self._trace_fixture()
        self.addCleanup(tmp.cleanup)
        path = root / "docs/BEHAVIOR.md"
        path.write_text(path.read_text() + "\n### FTR-041 X-F01 — Duplicate\n")
        subject = self._commit_fixture(root, "duplicate-behavior-heading")
        with self.assertRaisesRegex(ValueError, "expected exactly one behavior section"):
            module.query_feature_trace("FTR-041", subject, root=root)

    def test_untraced_feature_returns_truthful_missing_proof_condition(self):
        tmp, root, subject = self._trace_fixture()
        self.addCleanup(tmp.cleanup)
        result = module.query_feature_trace("FTR-019", subject, root=root)
        self.assertEqual(result["authority"], "NONE")
        self.assertEqual(result["condition"], "IMPLEMENTATION_PROOF_NOT_ESTABLISHED")
        self.assertEqual(result["feature_root_id"], "FTR-019")
        self.assertEqual(result["aliases"], ["R-F05"])
        self.assertEqual(result["implementation_refs"], [])
        self.assertEqual(result["verification_refs"], [])
        self.assertNotIn("status", result)
        self.assertNotIn("release_ready", result)

    def test_trace_query_reads_committed_subject_not_dirty_worktree(self):
        tmp, root, subject = self._trace_fixture()
        self.addCleanup(tmp.cleanup)
        path = root / "tools/slice/reading-training.json"
        config = json.loads(path.read_text())
        config["feature_id"] = "R-F99"
        path.write_text(json.dumps(config, indent=2) + "\n")
        result = module.query_feature_trace("FTR-018", subject, root=root)
        self.assertEqual(result["condition"], "TRACE_PRESENT")
        self.assertEqual(result["aliases"], ["R-F04"])

    def test_conflicting_alias_mapping_fails_closed(self):
        tmp, root, _ = self._trace_fixture()
        self.addCleanup(tmp.cleanup)
        path = root / "tools/slice/reading-training.json"
        config = json.loads(path.read_text())
        config["feature_id"] = "R-F99"
        path.write_text(json.dumps(config, indent=2) + "\n")
        subject = self._commit_fixture(root, "conflict")
        with self.assertRaisesRegex(ValueError, "conflicting feature mapping"):
            module.query_feature_trace("FTR-018", subject, root=root)

    def test_duplicate_feature_root_trace_fails_closed(self):
        tmp, root, _ = self._trace_fixture()
        self.addCleanup(tmp.cleanup)
        shutil.copyfile(
            root / "tools/slice/reading-training.json",
            root / "tools/slice/reading-training-copy.json",
        )
        subject = self._commit_fixture(root, "duplicate")
        with self.assertRaisesRegex(ValueError, "multiple slice traces"):
            module.query_feature_trace("FTR-018", subject, root=root)

    def test_missing_implementation_path_fails_closed(self):
        tmp, root, _ = self._trace_fixture()
        self.addCleanup(tmp.cleanup)
        path = root / "tools/slice/reading-training.json"
        config = json.loads(path.read_text())
        config["implementation_refs"][0]["path"] = "apps/web/src/missing.tsx"
        path.write_text(json.dumps(config, indent=2) + "\n")
        subject = self._commit_fixture(root, "missing-ref")
        with self.assertRaisesRegex(ValueError, "missing referenced path"):
            module.query_feature_trace("FTR-018", subject, root=root)

    def test_unattributable_verification_ref_fails_closed(self):
        tmp, root, _ = self._trace_fixture()
        self.addCleanup(tmp.cleanup)
        path = root / "tools/slice/reading-training.json"
        config = json.loads(path.read_text())
        config["verification_refs"][0]["required_literals"].append("R-F99")
        path.write_text(json.dumps(config, indent=2) + "\n")
        subject = self._commit_fixture(root, "unattributable")
        with self.assertRaisesRegex(ValueError, "unattributable verification relation"):
            module.query_feature_trace("FTR-018", subject, root=root)

    def test_malformed_trace_relation_fails_closed(self):
        tmp, root, _ = self._trace_fixture()
        self.addCleanup(tmp.cleanup)
        path = root / "tools/slice/reading-training.json"
        config = json.loads(path.read_text())
        config["implementation_refs"] = [{"path": "apps/web/src/features/today/today.tsx"}]
        path.write_text(json.dumps(config, indent=2) + "\n")
        subject = self._commit_fixture(root, "malformed")
        with self.assertRaisesRegex(ValueError, "malformed implementation relation"):
            module.query_feature_trace("FTR-018", subject, root=root)

    def test_trace_query_render_is_deterministic(self):
        tmp, root, subject = self._trace_fixture()
        self.addCleanup(tmp.cleanup)
        first = module.render(module.query_feature_trace("FTR-018", subject, root=root))
        second = module.render(module.query_feature_trace("FTR-018", subject, root=root))
        self.assertEqual(first, second)


    def test_cli_query_emits_derived_trace(self):
        tmp, root, subject = self._trace_fixture()
        self.addCleanup(tmp.cleanup)
        stdout = io.StringIO()
        with contextlib.redirect_stdout(stdout):
            code = module.main(
                [
                    "--query-feature",
                    "FTR-018",
                    "--expected-revision",
                    subject,
                ],
                root=root,
            )
        self.assertEqual(code, 0)
        result = json.loads(stdout.getvalue())
        self.assertEqual(result["condition"], "TRACE_PRESENT")
        self.assertEqual(result["subject_revision"], subject)


if __name__ == "__main__":
    unittest.main()
