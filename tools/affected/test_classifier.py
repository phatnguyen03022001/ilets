#!/usr/bin/env python3
from __future__ import annotations

import os
import shutil
import stat
import subprocess
import sys
import tempfile
import unittest
from pathlib import Path

HERE = Path(__file__).resolve().parent
TOOLS = HERE.parent
ROOT = TOOLS.parent
sys.path.insert(0, str(TOOLS))


def load_classifier():
    try:
        from affected import classifier
    except (ImportError, ModuleNotFoundError) as exc:
        raise AssertionError("affected classifier implementation is missing") from exc
    return classifier


class ParserTests(unittest.TestCase):
    def test_parses_nul_delimited_records(self) -> None:
        classifier = load_classifier()
        changes = classifier.parse_name_status_z(b"M\0README.md\0A\0apps/web/src/app/page.tsx\0")
        self.assertEqual(
            changes,
            [
                classifier.Change("M", ("README.md",)),
                classifier.Change("A", ("apps/web/src/app/page.tsx",)),
            ],
        )

    def test_rename_classifies_old_and_new_paths(self) -> None:
        classifier = load_classifier()
        changes = classifier.parse_name_status_z(b"R100\0README.md\0contracts/http/public-v2.json\0")
        result = classifier.classify_changes(changes)
        self.assertEqual(result.mode, "full")

    def test_copy_classifies_old_and_new_paths(self) -> None:
        classifier = load_classifier()
        changes = classifier.parse_name_status_z(
            b"C87\0apps/web/src/lib/a.ts\0services/core-api/internal/httpapi/a.go\0"
        )
        result = classifier.classify_changes(changes)
        self.assertEqual(result.mode, "web+go")

    def test_malformed_stream_is_rejected(self) -> None:
        classifier = load_classifier()
        with self.assertRaises(classifier.ParseError):
            classifier.parse_name_status_z(b"M\0README.md")

    def test_unsupported_status_is_rejected(self) -> None:
        classifier = load_classifier()
        with self.assertRaises(classifier.ParseError):
            classifier.parse_name_status_z(b"U\0README.md\0")


class ClassificationTests(unittest.TestCase):
    def classify(self, *paths: str):
        classifier = load_classifier()
        return classifier.classify_changes([classifier.Change("M", (path,)) for path in paths])

    def test_docs_only(self) -> None:
        self.assertEqual(self.classify("README.md", "docs/notes/fast-path.md").mode, "docs")

    def test_web_only_runs_web_lane(self) -> None:
        self.assertEqual(self.classify("apps/web/src/features/reading/a.tsx").mode, "web")

    def test_go_only_runs_go_lane(self) -> None:
        self.assertEqual(self.classify("services/core-api/internal/httpapi/attempt.go").mode, "go")

    def test_mixed_web_and_go_runs_both_lanes(self) -> None:
        self.assertEqual(
            self.classify(
                "apps/web/src/features/reading/a.tsx",
                "services/core-api/internal/httpapi/attempt.go",
            ).mode,
            "web+go",
        )

    def test_go_integration_test_change_forces_full(self) -> None:
        self.assertEqual(
            self.classify("services/core-api/internal/httpapi/idempotency_integration_test.go").mode,
            "full",
        )

    def test_db_change_forces_full(self) -> None:
        self.assertEqual(self.classify("services/core-api/migrations/0002_more.sql").mode, "full")

    def test_contract_change_forces_full(self) -> None:
        self.assertEqual(self.classify("contracts/http/public.openapi.yaml").mode, "full")

    def test_unknown_path_forces_full(self) -> None:
        self.assertEqual(self.classify("evidence/new-note.md").mode, "full")

    def test_classifier_self_change_forces_full(self) -> None:
        self.assertEqual(self.classify("tools/affected/classifier.py").mode, "full")

    def test_type_change_forces_full(self) -> None:
        classifier = load_classifier()
        result = classifier.classify_changes([classifier.Change("T", ("README.md",))])
        self.assertEqual(result.mode, "full")

    def test_empty_explicit_diff_is_no_changes(self) -> None:
        classifier = load_classifier()
        self.assertEqual(classifier.classify_changes([]).mode, "no-changes")


class EntrypointIntegrationTests(unittest.TestCase):
    def setUp(self) -> None:
        self.temp = tempfile.TemporaryDirectory()
        self.repo = Path(self.temp.name)
        subprocess.run(["git", "init", "-q"], cwd=self.repo, check=True)
        subprocess.run(["git", "config", "user.email", "test@example.com"], cwd=self.repo, check=True)
        subprocess.run(["git", "config", "user.name", "Test"], cwd=self.repo, check=True)
        (self.repo / "tools" / "affected").mkdir(parents=True)
        script_source = TOOLS / "check-affected"
        classifier_source = TOOLS / "affected" / "classifier.py"
        if not script_source.is_file() or not classifier_source.is_file():
            self.fail("check-affected implementation is missing")
        shutil.copy2(script_source, self.repo / "tools" / "check-affected")
        shutil.copy2(classifier_source, self.repo / "tools" / "affected" / "classifier.py")
        verify = self.repo / "verify"
        verify.write_text("#!/usr/bin/env bash\necho full >> .full-called\nexit 0\n", encoding="utf-8")
        verify.chmod(verify.stat().st_mode | stat.S_IXUSR)
        (self.repo / "README.md").write_text("base\n", encoding="utf-8")
        subprocess.run(["git", "add", "."], cwd=self.repo, check=True)
        subprocess.run(["git", "commit", "-qm", "base"], cwd=self.repo, check=True)
        self.base = subprocess.check_output(["git", "rev-parse", "HEAD"], cwd=self.repo, text=True).strip()

    def tearDown(self) -> None:
        self.temp.cleanup()

    def run_check(self, *args: str, env: dict[str, str] | None = None) -> subprocess.CompletedProcess[str]:
        merged_env = os.environ.copy()
        merged_env.pop("AFFECTED_BASE", None)
        merged_env.pop("PYTHONDONTWRITEBYTECODE", None)
        if env:
            merged_env.update(env)
        return subprocess.run(
            [str(self.repo / "tools" / "check-affected"), *args],
            cwd=self.repo,
            env=merged_env,
            text=True,
            stdout=subprocess.PIPE,
            stderr=subprocess.STDOUT,
            check=False,
        )

    def commit(self, path: str, content: str) -> None:
        target = self.repo / path
        target.parent.mkdir(parents=True, exist_ok=True)
        target.write_text(content, encoding="utf-8")
        subprocess.run(["git", "add", path], cwd=self.repo, check=True)
        subprocess.run(["git", "commit", "-qm", f"change {path}"], cwd=self.repo, check=True)

    def test_missing_baseline_falls_back_full(self) -> None:
        result = self.run_check()
        self.assertEqual(result.returncode, 0, result.stdout)
        self.assertIn("CHECK_PASS mode=full-fallback", result.stdout)
        self.assertTrue((self.repo / ".full-called").exists())

    def test_explicit_same_head_can_report_no_changes(self) -> None:
        head = subprocess.check_output(["git", "rev-parse", "HEAD"], cwd=self.repo, text=True).strip()
        result = self.run_check("--base", head)
        self.assertEqual(result.returncode, 0, result.stdout)
        self.assertIn("CHECK_PASS mode=no-changes", result.stdout)
        self.assertFalse((self.repo / ".full-called").exists())

    def test_direct_invocation_does_not_create_bytecode(self) -> None:
        head = subprocess.check_output(["git", "rev-parse", "HEAD"], cwd=self.repo, text=True).strip()
        result = self.run_check("--base", head)
        self.assertEqual(result.returncode, 0, result.stdout)
        self.assertFalse((self.repo / "tools" / "affected" / "__pycache__").exists())

    def test_environment_baseline_supports_docs_only(self) -> None:
        self.commit("README.md", "changed\n")
        result = self.run_check(env={"AFFECTED_BASE": self.base})
        self.assertEqual(result.returncode, 0, result.stdout)
        self.assertIn("CHECK_PASS mode=docs", result.stdout)
        self.assertFalse((self.repo / ".full-called").exists())

    def test_dirty_worktree_falls_back_full(self) -> None:
        head = subprocess.check_output(["git", "rev-parse", "HEAD"], cwd=self.repo, text=True).strip()
        (self.repo / "README.md").write_text("dirty\n", encoding="utf-8")
        result = self.run_check("--base", head)
        self.assertEqual(result.returncode, 0, result.stdout)
        self.assertIn("CHECK_PASS mode=full-fallback", result.stdout)
        self.assertTrue((self.repo / ".full-called").exists())

    def test_submodule_change_falls_back_full(self) -> None:
        source = self.repo.parent / f"{self.repo.name}-submodule-source"
        source.mkdir()
        subprocess.run(["git", "init", "-q"], cwd=source, check=True)
        subprocess.run(["git", "config", "user.email", "test@example.com"], cwd=source, check=True)
        subprocess.run(["git", "config", "user.name", "Test"], cwd=source, check=True)
        (source / "data.txt").write_text("submodule\n", encoding="utf-8")
        subprocess.run(["git", "add", "."], cwd=source, check=True)
        subprocess.run(["git", "commit", "-qm", "submodule base"], cwd=source, check=True)
        subprocess.run(
            ["git", "-c", "protocol.file.allow=always", "submodule", "add", "-q", str(source), "vendor/module"],
            cwd=self.repo,
            check=True,
        )
        subprocess.run(["git", "commit", "-am", "add submodule", "-q"], cwd=self.repo, check=True)
        result = self.run_check("--base", self.base)
        self.assertEqual(result.returncode, 0, result.stdout)
        self.assertIn("CHECK_PASS mode=full-fallback", result.stdout)
        self.assertTrue((self.repo / ".full-called").exists())

    def test_unknown_path_falls_back_full(self) -> None:
        self.commit("evidence/new-note.md", "unknown\n")
        result = self.run_check("--base", self.base)
        self.assertEqual(result.returncode, 0, result.stdout)
        self.assertIn("CHECK_PASS mode=full-fallback", result.stdout)
        self.assertTrue((self.repo / ".full-called").exists())


class ShallowHistoryIntegrationTests(unittest.TestCase):
    def test_unresolvable_shallow_baseline_falls_back_full(self) -> None:
        with tempfile.TemporaryDirectory() as temp_name:
            temp = Path(temp_name)
            source = temp / "source"
            source.mkdir()
            subprocess.run(["git", "init", "-q"], cwd=source, check=True)
            subprocess.run(["git", "config", "user.email", "test@example.com"], cwd=source, check=True)
            subprocess.run(["git", "config", "user.name", "Test"], cwd=source, check=True)
            (source / "tools" / "affected").mkdir(parents=True)
            shutil.copy2(TOOLS / "check-affected", source / "tools" / "check-affected")
            shutil.copy2(TOOLS / "affected" / "classifier.py", source / "tools" / "affected" / "classifier.py")
            verify = source / "verify"
            verify.write_text("#!/usr/bin/env bash\necho full >> .full-called\nexit 0\n", encoding="utf-8")
            verify.chmod(verify.stat().st_mode | stat.S_IXUSR)
            (source / "README.md").write_text("base\n", encoding="utf-8")
            subprocess.run(["git", "add", "."], cwd=source, check=True)
            subprocess.run(["git", "commit", "-qm", "base"], cwd=source, check=True)
            baseline = subprocess.check_output(["git", "rev-parse", "HEAD"], cwd=source, text=True).strip()
            (source / "README.md").write_text("head\n", encoding="utf-8")
            subprocess.run(["git", "add", "README.md"], cwd=source, check=True)
            subprocess.run(["git", "commit", "-qm", "head"], cwd=source, check=True)

            clone = temp / "clone"
            subprocess.run(["git", "clone", "-q", "--depth", "1", f"file://{source}", str(clone)], check=True)
            env = os.environ.copy()
            env.pop("PYTHONDONTWRITEBYTECODE", None)
            result = subprocess.run(
                [str(clone / "tools" / "check-affected"), "--base", baseline],
                cwd=clone,
                env=env,
                text=True,
                stdout=subprocess.PIPE,
                stderr=subprocess.STDOUT,
                check=False,
            )
            self.assertEqual(result.returncode, 0, result.stdout)
            self.assertIn("CHECK_PASS mode=full-fallback", result.stdout)
            self.assertTrue((clone / ".full-called").exists())


if __name__ == "__main__":
    unittest.main(verbosity=2)
