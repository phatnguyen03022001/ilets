#!/usr/bin/env python3
from __future__ import annotations

import hashlib
import os
from pathlib import Path
import re
import subprocess

ROOT = Path(__file__).resolve().parents[2]
GENERATOR = ROOT / "tools/contracts/generate.sh"
PUBLIC_SPEC = "contracts/http/public.openapi.yaml"
EVALUATOR_SPEC = "contracts/http/evaluator.openapi.yaml"
LEGACY_INPUT_MARKERS = ("public-v1.json", "public-v1.yaml", "public_v1")
RETIRED_CONSUMERS = (
    ROOT / "apps/web/src/generated/public-v1.ts",
    ROOT / "services/core-api/internal/httpapi/generated/public_v1.gen.go",
)
PUBLIC_GO = ROOT / "services/core-api/internal/generated/openapi/public/public.gen.go"
EVALUATOR_GO = ROOT / "services/core-api/internal/generated/openapi/evaluator/evaluator.gen.go"
PUBLIC_TS = ROOT / "apps/web/src/generated/public"

PUBLIC_OPERATIONS = (
    "getCoreHealth",
    "getMe",
    "getTargetProfile",
    "putTargetProfile",
    "getDailyPlan",
    "listPracticeModes",
    "createPracticeActivity",
    "getPracticeActivity",
    "createAttempt",
    "getAttempt",
    "patchAttempt",
    "submitAttempt",
    "getEvaluation",
    "getProgress",
    "listGaps",
    "getReviewQueue",
    "streamResourceEvents",
)
EVALUATOR_OPERATIONS = ("executeEvaluation", "getEvaluatorHealth")
STALE_AUTH_MARKERS = (
    "/v1/session",
    "ilets_session",
    "opaqueSession",
    "opaque-server-session",
    "assessment-activity",
)


def sha256(path: Path) -> str:
    return hashlib.sha256(path.read_bytes()).hexdigest()


def tree_digest(path: Path) -> str:
    digest = hashlib.sha256()
    for child in sorted(p for p in path.rglob("*") if p.is_file()):
        digest.update(child.relative_to(path).as_posix().encode())
        digest.update(b"\0")
        digest.update(child.read_bytes())
        digest.update(b"\0")
    return digest.hexdigest()


def go_name(operation_id: str) -> str:
    return operation_id[0].upper() + operation_id[1:]


def assert_contains_all(text: str, needles: tuple[str, ...], label: str) -> None:
    missing = [needle for needle in needles if needle not in text]
    if missing:
        raise AssertionError(f"{label} missing expected values: {missing}")


def main() -> None:
    if not GENERATOR.is_file():
        raise AssertionError(f"missing generation entrypoint: {GENERATOR.relative_to(ROOT)}")
    if not os.access(GENERATOR, os.X_OK):
        raise AssertionError("generation entrypoint must be executable")

    generator_text = GENERATOR.read_text()
    assert_contains_all(generator_text, (PUBLIC_SPEC, EVALUATOR_SPEC), "generator")
    for marker in LEGACY_INPUT_MARKERS:
        if marker in generator_text:
            raise AssertionError(f"legacy contract marker cannot be generator input: {marker}")

    for path in RETIRED_CONSUMERS:
        if path.exists():
            raise AssertionError(f"retired generated consumer still exists: {path.relative_to(ROOT)}")

    subprocess.run([str(GENERATOR)], cwd=ROOT, check=True)
    for path in RETIRED_CONSUMERS:
        if path.exists():
            raise AssertionError(f"generation recreated retired consumer: {path.relative_to(ROOT)}")

    if not PUBLIC_GO.is_file() or not EVALUATOR_GO.is_file() or not PUBLIC_TS.is_dir():
        raise AssertionError("generation did not materialize all expected output surfaces")

    public_go_text = PUBLIC_GO.read_text()
    evaluator_go_text = EVALUATOR_GO.read_text()
    public_ts_text = "\n".join(
        path.read_text() for path in sorted(PUBLIC_TS.rglob("*.ts")) if path.is_file()
    )
    public_sdk_text = (PUBLIC_TS / "sdk.gen.ts").read_text()
    public_types_text = (PUBLIC_TS / "types.gen.ts").read_text()
    combined_public = public_go_text + "\n" + public_ts_text

    for marker in STALE_AUTH_MARKERS:
        if marker in combined_public:
            raise AssertionError(f"new public generated artifacts contain stale marker: {marker}")

    assert_contains_all(
        public_go_text,
        tuple(go_name(operation) for operation in PUBLIC_OPERATIONS),
        "public Go bindings",
    )
    assert_contains_all(public_ts_text, PUBLIC_OPERATIONS, "public TypeScript bindings")
    ts_operations = set(re.findall(r"^export const ([A-Za-z][A-Za-z0-9_]*) =", public_sdk_text, re.MULTILINE))
    if ts_operations != set(PUBLIC_OPERATIONS):
        raise AssertionError(
            f"public TypeScript operation surface differs: got {sorted(ts_operations)}, "
            f"want {sorted(PUBLIC_OPERATIONS)}"
        )
    if public_sdk_text.count("security: [{ scheme: 'bearer', type: 'http' }]") != len(PUBLIC_OPERATIONS) - 1:
        raise AssertionError("public TypeScript bearer security metadata count is not exact")
    assert_contains_all(
        public_types_text,
        ("'Idempotency-Key'", "'Expected-Resource-Revision'", "'Last-Event-ID'", "ResourceChangedEvent"),
        "public TypeScript contract types",
    )
    if ".sse.get<StreamResourceEventsResponses" not in public_sdk_text:
        raise AssertionError("public TypeScript SSE operation lost generated SSE transport")

    assert_contains_all(
        evaluator_go_text,
        tuple(go_name(operation) for operation in EVALUATOR_OPERATIONS),
        "evaluator Go bindings",
    )

    for operation in PUBLIC_OPERATIONS:
        if go_name(operation) in evaluator_go_text:
            raise AssertionError(f"evaluator bindings leaked public operation: {operation}")

    first_digest = hashlib.sha256(
        (tree_digest(PUBLIC_GO.parent) + tree_digest(EVALUATOR_GO.parent) + tree_digest(PUBLIC_TS)).encode()
    ).hexdigest()
    subprocess.run([str(GENERATOR)], cwd=ROOT, check=True)
    second_digest = hashlib.sha256(
        (tree_digest(PUBLIC_GO.parent) + tree_digest(EVALUATOR_GO.parent) + tree_digest(PUBLIC_TS)).encode()
    ).hexdigest()
    if first_digest != second_digest:
        raise AssertionError("repeated generation is not deterministic")

    for path in RETIRED_CONSUMERS:
        if path.exists():
            raise AssertionError(f"repeated generation recreated retired consumer: {path.relative_to(ROOT)}")

    print("generated binding checks passed")


if __name__ == "__main__":
    main()
