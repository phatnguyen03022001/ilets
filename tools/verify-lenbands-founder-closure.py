#!/usr/bin/env python3
from __future__ import annotations

import re
import sys
from collections import Counter
from pathlib import Path

ROOT = Path(__file__).resolve().parents[1]
RESEARCH = ROOT / "research" / "lenbands-founder-decisions"
CLOSURE = RESEARCH / "CLOSURE.md"

SOURCE_FILES = [
    "platform-and-reliability.md",
    "identity-privacy-and-access.md",
    "ai-evaluation-and-cost.md",
    "product-and-requirements.md",
    "learning-interventions.md",
    "evidence-and-readiness.md",
    "learner-experience.md",
    "economics-and-entitlements.md",
    "coverage-and-support.md",
]

ALLOWED = {
    "ADOPTED",
    "SUPERSEDED",
    "REJECTED",
    "DEFERRED",
    "NOT_APPLICABLE",
    "UNRESOLVED",
}
BASELINE_COUNT = 325
NUMBERED_ID = re.compile(r"^(?:V[1-9]\.\d+|10[A-F]\.\d+)$")
V7 = {
    "Identity": "V7/Identity",
    "API / evaluation workers": "V7/API-evaluation-workers",
    "Frontend": "V7/Frontend",
    "Region": "V7/Region",
    "Postgres": "V7/Postgres",
    "Independent backup": "V7/Independent-backup",
    "Object storage": "V7/Object-storage",
    "Redis": "V7/Redis",
    "Admin auth": "V7/Admin-auth",
    "Staging": "V7/Staging",
    "IaC": "V7/IaC",
    "Provider rule": "V7/Provider-rule",
}


def table_cells(line: str) -> list[str]:
    if not line.startswith("|"):
        return []
    return [cell.strip() for cell in line.strip().strip("|").split("|")]


def source_ids() -> list[str]:
    ids: list[str] = []
    for name in SOURCE_FILES:
        in_v7 = False
        for line in (RESEARCH / name).read_text(encoding="utf-8").splitlines():
            if line.startswith("## V7 "):
                in_v7 = True
                continue
            if in_v7 and line.startswith("## ") and not line.startswith("## V7 "):
                in_v7 = False
            cells = table_cells(line)
            if not cells:
                continue
            first = cells[0]
            if NUMBERED_ID.fullmatch(first):
                ids.append(first)
            elif in_v7 and first in V7:
                ids.append(V7[first])
    return ids


def closure_rows(expected: set[str]) -> list[tuple[str, str, str, str]]:
    rows: list[tuple[str, str, str, str]] = []
    for line in CLOSURE.read_text(encoding="utf-8").splitlines():
        cells = table_cells(line)
        if len(cells) != 4 or cells[0] not in expected:
            continue
        rows.append((cells[0], cells[1], cells[2], cells[3]))
    return rows


def fail(message: str) -> None:
    print(f"LenBands founder closure check failed: {message}", file=sys.stderr)
    raise SystemExit(1)


def main() -> None:
    if not CLOSURE.exists():
        fail("CLOSURE.md is missing")

    closure_text = CLOSURE.read_text(encoding="utf-8")
    if "AUTHORITY: NONE" not in closure_text.splitlines()[:8]:
        fail("CLOSURE.md must declare AUTHORITY: NONE")

    imported = source_ids()
    source_counts = Counter(imported)
    source_duplicates = sorted(key for key, count in source_counts.items() if count != 1)
    if source_duplicates:
        fail(f"source IDs are not unique: {source_duplicates}")
    if len(imported) != BASELINE_COUNT:
        fail(
            f"authoritative imported source count changed: expected closure baseline "
            f"{BASELINE_COUNT}, found {len(imported)}; re-evaluate and close the current set"
        )

    expected = set(imported)
    closed = closure_rows(expected)
    closed_counts = Counter(row[0] for row in closed)
    duplicates = sorted(key for key, count in closed_counts.items() if count != 1)
    missing = sorted(expected - set(closed_counts))
    extras = sorted(set(closed_counts) - expected)

    if duplicates:
        fail(f"duplicate closure IDs: {duplicates}")
    if missing:
        fail(f"missing closure IDs: {missing}")
    if extras:
        fail(f"unexpected closure IDs: {extras}")

    for row_id, disposition, target, rationale in closed:
        if disposition not in ALLOWED:
            fail(f"{row_id} has invalid disposition {disposition!r}")
        if disposition == "UNRESOLVED":
            fail(f"{row_id} remains UNRESOLVED")
        if not target or target == "-":
            fail(f"{row_id} has no current owner/trigger/superseding rule")
        if not rationale or rationale == "-":
            fail(f"{row_id} has no rationale")

    print(
        "LenBands founder closure: "
        f"expected={len(imported)} dispositioned={len(closed)} "
        "duplicates=0 missing=0 unresolved=0"
    )


if __name__ == "__main__":
    main()
