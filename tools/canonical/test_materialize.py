#!/usr/bin/env python3
from __future__ import annotations
import importlib.util, json, tempfile, unittest
from pathlib import Path
MODULE_PATH=Path(__file__).with_name("materialize.py")
spec=importlib.util.spec_from_file_location("canonical_materializer",MODULE_PATH); assert spec and spec.loader
module=importlib.util.module_from_spec(spec); spec.loader.exec_module(module)
class MaterializerTests(unittest.TestCase):
    def test_forbidden_sources_cannot_be_authority(self):
        config=json.loads(module.SOURCE_MAP_PATH.read_text()); config["registries"].append({"type":"fake","owner":"research/fake.md","section":"# Fake","kind":"markdown_table","ids":["R-F04"]})
        with self.assertRaisesRegex(ValueError,"forbidden authority source"): module.validate_source_map(config)
    def test_id_looking_research_text_is_not_discovered(self):
        with tempfile.TemporaryDirectory() as tmp:
            fake=Path(tmp)/"research.md"; fake.write_text("| ID | Thing |\n|---|---|\n| `R-F04` | fake |\n")
            owners={entry["owner"] for entry in json.loads(module.SOURCE_MAP_PATH.read_text())["registries"]}
            self.assertNotIn(str(fake),owners); self.assertTrue(all(not owner.startswith("research/") for owner in owners))
    def test_duplicate_configured_id_fails(self):
        config=json.loads(module.SOURCE_MAP_PATH.read_text()); config["registries"].append(dict(config["registries"][0]))
        with self.assertRaisesRegex(ValueError,"duplicate configured canonical id"): module.validate_source_map(config)
if __name__=="__main__": unittest.main()
