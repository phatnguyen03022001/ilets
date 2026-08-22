#!/usr/bin/env python3
from __future__ import annotations

import importlib.util
import tempfile
import unittest
from pathlib import Path

MODULE_PATH = Path(__file__).with_name("verify-lenbands-founder-closure.py")
SPEC = importlib.util.spec_from_file_location("verify_lenbands_founder_closure", MODULE_PATH)
assert SPEC and SPEC.loader
verifier = importlib.util.module_from_spec(SPEC)
SPEC.loader.exec_module(verifier)


class VerifierMutationTests(unittest.TestCase):
    def setUp(self) -> None:
        self.temp = tempfile.TemporaryDirectory()
        self.root = Path(self.temp.name)
        (self.root / "design").mkdir()
        (self.root / "spec").mkdir()
        (self.root / "design" / "01-test.md").write_text("test\n", encoding="utf-8")
        (self.root / "spec" / "01-test.md").write_text("test\n", encoding="utf-8")
        (self.root / "CONSTITUTION.md").write_text("test\n", encoding="utf-8")
        (self.root / "OBJECTIVE.md").write_text("test\n", encoding="utf-8")
        self.expected = {"V1.1", "V1.2", "V7/Identity"}

    def tearDown(self) -> None:
        self.temp.cleanup()

    def row(
        self,
        row_id: str,
        disposition: str = "ADOPTED",
        target: str = "design/01-test.md owner",
        rationale: str = "bounded rationale",
    ):
        return verifier.ClosureRow(row_id, disposition, target, rationale)

    def numbered(self):
        return [self.row("V1.1"), self.row("V1.2"), self.row("V7/Identity")]

    def rights(self):
        return [self.row(f"RIGHTS-{index}", target="spec/01-test.md rights") for index in range(1, 12)]

    def validate(self, numbered=None, rights=None):
        return verifier.validate_rows(
            self.expected,
            self.numbered() if numbered is None else numbered,
            self.rights() if rights is None else rights,
            self.root,
        )

    def assert_rejected(self, pattern: str, numbered=None, rights=None) -> None:
        with self.assertRaisesRegex(verifier.VerificationError, pattern):
            self.validate(numbered=numbered, rights=rights)

    def test_missing_numbered_id_rejected(self) -> None:
        self.assert_rejected("missing closure IDs", numbered=self.numbered()[:-1])

    def test_duplicate_numbered_id_rejected(self) -> None:
        numbered = self.numbered() + [self.row("V1.1")]
        self.assert_rejected("duplicate closure IDs", numbered=numbered)

    def test_unexpected_numbered_id_rejected(self) -> None:
        numbered = self.numbered() + [self.row("V99.1")]
        self.assert_rejected("unexpected closure IDs", numbered=numbered)

    def test_unresolved_numbered_row_rejected(self) -> None:
        numbered = self.numbered()
        numbered[0] = self.row("V1.1", disposition="UNRESOLVED")
        self.assert_rejected("V1.1 remains UNRESOLVED", numbered=numbered)

    def test_invalid_disposition_rejected(self) -> None:
        numbered = self.numbered()
        numbered[0] = self.row("V1.1", disposition="DONE")
        self.assert_rejected("invalid disposition", numbered=numbered)

    def test_missing_numbered_target_rejected(self) -> None:
        numbered = self.numbered()
        numbered[0] = self.row("V1.1", target="-")
        self.assert_rejected("has no current owner", numbered=numbered)

    def test_missing_numbered_rationale_rejected(self) -> None:
        numbered = self.numbered()
        numbered[0] = self.row("V1.1", rationale="-")
        self.assert_rejected("has no rationale", numbered=numbered)

    def test_missing_rights_row_rejected(self) -> None:
        self.assert_rejected("missing rights closure IDs", rights=self.rights()[:-1])

    def test_duplicate_rights_row_rejected(self) -> None:
        rights = self.rights() + [self.row("RIGHTS-1")]
        self.assert_rejected("duplicate rights closure IDs", rights=rights)

    def test_unexpected_rights_row_rejected(self) -> None:
        rights = self.rights() + [self.row("RIGHTS-12")]
        self.assert_rejected("unexpected rights closure IDs", rights=rights)

    def test_unresolved_rights_row_rejected(self) -> None:
        rights = self.rights()
        rights[0] = self.row("RIGHTS-1", disposition="UNRESOLVED")
        self.assert_rejected("RIGHTS-1 remains UNRESOLVED", rights=rights)

    def test_missing_rights_target_rejected(self) -> None:
        rights = self.rights()
        rights[0] = self.row("RIGHTS-1", target="-")
        self.assert_rejected("has no current owner", rights=rights)

    def test_missing_rights_rationale_rejected(self) -> None:
        rights = self.rights()
        rights[0] = self.row("RIGHTS-1", rationale="-")
        self.assert_rejected("has no rationale", rights=rights)

    def test_missing_canonical_reference_rejected(self) -> None:
        numbered = self.numbered()
        numbered[0] = self.row("V1.1", target="design/missing.md production gate")
        self.assert_rejected("references missing canonical file", numbered=numbered)

    def test_canonical_path_extraction_is_conservative(self) -> None:
        self.assertEqual(
            verifier.canonical_paths(
                "design/01-test.md + spec/01-test.md validation + CONSTITUTION.md authority + OBJECTIVE.md"
            ),
            ["design/01-test.md", "spec/01-test.md", "CONSTITUTION.md", "OBJECTIVE.md"],
        )
        self.assertEqual(
            verifier.canonical_paths(
                "Auth0 selection order + https://example.com/spec/not-a-repo-file.md + archive/design/history.md"
            ),
            [],
        )

    def test_v7_source_identity_normalization_preserved(self) -> None:
        research = self.root / "research"
        research.mkdir()
        for name in verifier.SOURCE_FILES:
            (research / name).write_text("", encoding="utf-8")
        (research / "platform-and-reliability.md").write_text(
            "## V7 — Provider architecture selection\n"
            "| Boundary | Founder decision/direction | Lifecycle | Rationale |\n"
            "|---|---|---|---|\n"
            "| API / evaluation workers | Cloud Run | historical | rationale |\n",
            encoding="utf-8",
        )
        self.assertEqual(verifier.source_ids(research), ["V7/API-evaluation-workers"])

    def test_parser_discovers_unexpected_founder_id_without_expected_filter(self) -> None:
        closure = self.root / "CLOSURE-unexpected.md"
        closure.write_text(
            "## Platform and reliability\n"
            "| ID | Disposition | Target | Rationale |\n"
            "|---|---|---|---|\n"
            "| V99.1 | ADOPTED | design/01-test.md | rationale |\n",
            encoding="utf-8",
        )
        numbered, rights = verifier.closure_rows(closure)
        self.assertEqual([row.row_id for row in numbered], ["V99.1"])
        self.assertEqual(rights, [])

    def test_section_aware_closure_parser_keeps_rights_distinct(self) -> None:
        closure = self.root / "CLOSURE.md"
        closure.write_text(
            "## Platform and reliability\n"
            "| ID | Disposition | Target | Rationale |\n"
            "|---|---|---|---|\n"
            "| V1.1 | ADOPTED | design/01-test.md | rationale |\n"
            "## Rights/provenance closure — outside the 325 rows\n"
            "| ID | Disposition | Target | Rationale |\n"
            "|---|---|---|---|\n"
            "| RIGHTS-1 | ADOPTED | spec/01-test.md | rationale |\n"
            "## Mechanical closure result\n"
            "| ID | Disposition | Target | Rationale |\n"
            "| V99.1 | ADOPTED | design/01-test.md | summary example |\n",
            encoding="utf-8",
        )
        numbered, rights = verifier.closure_rows(closure)
        self.assertEqual([row.row_id for row in numbered], ["V1.1"])
        self.assertEqual([row.row_id for row in rights], ["RIGHTS-1"])


if __name__ == "__main__":
    unittest.main()
