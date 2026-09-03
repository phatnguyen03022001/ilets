#!/usr/bin/env python3
from __future__ import annotations

import importlib.util
import json
import shutil
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
        source_slice = module.SLICE_INPUT_DIR / "reading-training.json"
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


if __name__ == "__main__":
    unittest.main()
