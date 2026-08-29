#!/usr/bin/env python3
from __future__ import annotations

import subprocess
import tempfile
import textwrap
import unittest
from pathlib import Path

ROOT = Path(__file__).resolve().parents[2]
CORE = ROOT / "services" / "core-api"
VALIDATOR = ROOT / "tools" / "contracts" / "validate.go"
PUBLIC = ROOT / "contracts" / "http" / "public.openapi.yaml"
EVALUATOR = ROOT / "contracts" / "http" / "evaluator.openapi.yaml"


def minimal_spec(path: str, operation_id: str, security_scheme: str) -> str:
    return f"""openapi: 3.0.3
info:
  title: test
  version: 1.0.0
security:
  - {security_scheme}: []
paths:
  {path}:
    get:
      operationId: {operation_id}
      responses:
        '200':
          description: ok
components:
  securitySchemes:
    {security_scheme}:
      type: http
      scheme: bearer
      bearerFormat: JWT
"""


class ContractValidatorTests(unittest.TestCase):
    def run_validator(
        self, public_text: str, evaluator_text: str, *, repository: bool = False
    ) -> subprocess.CompletedProcess[str]:
        with tempfile.TemporaryDirectory() as temp_dir:
            temp = Path(temp_dir)
            public = temp / "public.openapi.yaml"
            evaluator = temp / "evaluator.openapi.yaml"
            public.write_text(public_text, encoding="utf-8")
            evaluator.write_text(evaluator_text, encoding="utf-8")
            argv = ["go", "run", str(VALIDATOR)]
            if repository:
                argv.append("--repository")
            argv.extend([str(public), str(evaluator)])
            return subprocess.run(
                argv,
                cwd=CORE,
                text=True,
                stdout=subprocess.PIPE,
                stderr=subprocess.PIPE,
                check=False,
            )

    def test_repository_contracts_validate_and_match_authority_shape(self) -> None:
        result = subprocess.run(
            ["go", "run", str(VALIDATOR), "--repository", str(PUBLIC), str(EVALUATOR)],
            cwd=CORE,
            text=True,
            stdout=subprocess.PIPE,
            stderr=subprocess.PIPE,
            check=False,
        )
        self.assertEqual(result.returncode, 0, result.stdout + result.stderr)
        self.assertIn("HTTP_CONTRACTS_VALID", result.stdout)
        self.assertIn("HTTP_CONTRACT_AUTHORITIES_AUDITED", result.stdout)

    def test_rejects_public_internal_route(self) -> None:
        public = minimal_spec("/internal/v1/oops", "oops", "ClerkBearer")
        evaluator = minimal_spec("/internal/v1/health", "getEvaluatorHealth", "GoogleOidcBearer")
        result = self.run_validator(public, evaluator)
        self.assertNotEqual(result.returncode, 0)
        self.assertIn("public contract contains internal route", result.stderr)

    def test_rejects_stale_public_session_authority(self) -> None:
        public = minimal_spec("/v1/me", "getMe", "ClerkBearer") + "\n# ilets_session\n"
        evaluator = minimal_spec("/internal/v1/health", "getEvaluatorHealth", "GoogleOidcBearer")
        result = self.run_validator(public, evaluator)
        self.assertNotEqual(result.returncode, 0)
        self.assertIn("stale public session authority", result.stderr)

    def test_rejects_duplicate_operation_ids(self) -> None:
        public = textwrap.dedent(
            """\
            openapi: 3.0.3
            info: {title: test, version: 1.0.0}
            security:
              - ClerkBearer: []
            paths:
              /v1/me:
                get:
                  operationId: duplicate
                  responses: {'200': {description: ok}}
              /v1/progress:
                get:
                  operationId: duplicate
                  responses: {'200': {description: ok}}
            components:
              securitySchemes:
                ClerkBearer: {type: http, scheme: bearer, bearerFormat: JWT}
            """
        )
        evaluator = minimal_spec("/internal/v1/health", "getEvaluatorHealth", "GoogleOidcBearer")
        result = self.run_validator(public, evaluator)
        self.assertNotEqual(result.returncode, 0)
        self.assertIn("duplicate operationId", result.stderr)

    def test_bootstrap_authority_is_retired_from_generation(self) -> None:
        self.assertFalse((ROOT / "contracts" / "http" / "public-v1.json").exists())
        generator = (ROOT / "tools" / "contracts" / "generate.sh").read_text(encoding="utf-8")
        self.assertIn("contracts/http/public.openapi.yaml", generator)
        self.assertIn("contracts/http/evaluator.openapi.yaml", generator)
        self.assertNotIn("public-v1", generator)
        package = (ROOT / "apps" / "web" / "package.json").read_text(encoding="utf-8")
        self.assertNotIn("contract:generate", package)
        self.assertNotIn('"openapi-typescript"', package)
        self.assertIn('"@hey-api/openapi-ts": "0.99.0"', package)


if __name__ == "__main__":
    unittest.main()
