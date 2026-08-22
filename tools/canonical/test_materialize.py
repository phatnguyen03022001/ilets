#!/usr/bin/env python3
from __future__ import annotations

import importlib.util
import json
import tempfile
import unittest
from pathlib import Path

MODULE_PATH = Path(__file__).with_name("materialize.py")
spec = importlib.util.spec_from_file_location("canonical_materializer", MODULE_PATH)
assert spec and spec.loader
module = importlib.util.module_from_spec(spec)
spec.loader.exec_module(module)


class MaterializerTests(unittest.TestCase):
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
                "owner": "research/fake.md",
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
            ("spec/a.md", "aaa"),
            ("design/b.md", "bbb"),
        ]
        first = module.canonical_source_revision_from_inputs(source_map, owners)
        second = module.canonical_source_revision_from_inputs(source_map, owners)
        self.assertEqual(first, second)

    def test_owner_order_does_not_change_revision(self):
        source_map = b'{"version":1}'
        owners = [
            ("spec/a.md", "aaa"),
            ("design/b.md", "bbb"),
        ]
        forward = module.canonical_source_revision_from_inputs(source_map, owners)
        reverse = module.canonical_source_revision_from_inputs(
            source_map, list(reversed(owners))
        )
        self.assertEqual(forward, reverse)

    def test_owner_blob_change_changes_revision(self):
        source_map = b'{"version":1}'
        before = module.canonical_source_revision_from_inputs(
            source_map, [("spec/a.md", "aaa")]
        )
        after = module.canonical_source_revision_from_inputs(
            source_map, [("spec/a.md", "bbb")]
        )
        self.assertNotEqual(before, after)

    def test_source_map_change_changes_revision(self):
        owners = [("spec/a.md", "aaa")]
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
