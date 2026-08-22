#!/usr/bin/env python3
from __future__ import annotations

import re
import sys
from collections import Counter
from pathlib import Path
from typing import NamedTuple

ROOT = Path(__file__).resolve().parents[1]
RESEARCH_DIR = Path("research/lenbands-founder-decisions")
CLOSURE_NAME = "CLOSURE.md"

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
RIGHTS_EXPECTED = {f"RIGHTS-{index}" for index in range(1, 12)}
SOURCE_NUMBERED_ID = re.compile(r"^(?:V[1-9]\.\d+|10[A-F]\.\d+)$")
FOUNDER_CLOSURE_ID = re.compile(r"^(?:V\d+\.\d+|10[A-Z]\.\d+|V7/[A-Za-z0-9-]+)$")
RIGHTS_CLOSURE_ID = re.compile(r"^RIGHTS-[A-Za-z0-9._-]+$")
CANONICAL_PATH = re.compile(
    r"(?<![A-Za-z0-9_./-])"
    r"(CONSTITUTION\.md|OBJECTIVE\.md|(?:spec|design)/[A-Za-z0-9._-]+\.md)"
    r"(?![A-Za-z0-9_./-])"
)
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


class ClosureRow(NamedTuple):
    row_id: str
    disposition: str
    target: str
    rationale: str


class VerificationError(ValueError):
    pass


def table_cells(line: str) -> list[str]:
    if not line.startswith("|"):
        return []
    return [cell.strip() for cell in line.strip().strip("|").split("|")]


def source_ids(research: Path) -> list[str]:
    ids: list[str] = []
    for name in SOURCE_FILES:
        path = research / name
        if not path.is_file():
            raise VerificationError(f"imported source file is missing: {path}")
        in_v7 = False
        for line in path.read_text(encoding="utf-8").splitlines():
            if line.startswith("## V7 "):
                in_v7 = True
                continue
            if in_v7 and line.startswith("## ") and not line.startswith("## V7 "):
                in_v7 = False
            cells = table_cells(line)
            if not cells:
                continue
            first = cells[0]
            if SOURCE_NUMBERED_ID.fullmatch(first):
                ids.append(first)
            elif in_v7 and first in V7:
                ids.append(V7[first])
    return ids


def closure_rows(closure: Path) -> tuple[list[ClosureRow], list[ClosureRow]]:
    numbered: list[ClosureRow] = []
    rights: list[ClosureRow] = []
    section = "none"

    for line_number, line in enumerate(closure.read_text(encoding="utf-8").splitlines(), start=1):
        if line.startswith("## Rights/provenance closure"):
            section = "rights"
            continue
        if line.startswith("## Mechanical closure result"):
            section = "none"
            continue
        if line.startswith("## "):
            section = "numbered"
            continue

        cells = table_cells(line)
        if not cells or section == "none":
            continue
        first = cells[0]
        is_founder = FOUNDER_CLOSURE_ID.fullmatch(first) is not None
        is_rights = RIGHTS_CLOSURE_ID.fullmatch(first) is not None
        if not is_founder and not is_rights:
            continue
        if len(cells) != 4:
            raise VerificationError(f"{first} has malformed closure row at line {line_number}")

        row = ClosureRow(cells[0], cells[1], cells[2], cells[3])
        if is_founder:
            if section != "numbered":
                raise VerificationError(f"{first} founder closure row is outside numbered closure sections")
            numbered.append(row)
        else:
            if section != "rights":
                raise VerificationError(f"{first} rights closure row is outside rights closure section")
            rights.append(row)

    return numbered, rights


def canonical_paths(target: str) -> list[str]:
    return [match.group(1) for match in CANONICAL_PATH.finditer(target)]


def validate_target_paths(row: ClosureRow, repo_root: Path) -> None:
    root = repo_root.resolve()
    for relative in canonical_paths(row.target):
        resolved = (root / relative).resolve()
        if not resolved.is_relative_to(root):
            raise VerificationError(f"{row.row_id} references path outside repository root: {relative}")
        if not resolved.is_file():
            raise VerificationError(f"{row.row_id} references missing canonical file: {relative}")


def validate_rows(
    expected_numbered: set[str],
    numbered: list[ClosureRow],
    rights: list[ClosureRow],
    repo_root: Path,
) -> dict[str, int]:
    numbered_counts = Counter(row.row_id for row in numbered)
    numbered_duplicates = sorted(row_id for row_id, count in numbered_counts.items() if count > 1)
    numbered_missing = sorted(expected_numbered - set(numbered_counts))
    numbered_unexpected = sorted(set(numbered_counts) - expected_numbered)

    if numbered_duplicates:
        raise VerificationError(f"duplicate closure IDs: {numbered_duplicates}")
    if numbered_missing:
        raise VerificationError(f"missing closure IDs: {numbered_missing}")
    if numbered_unexpected:
        raise VerificationError(f"unexpected closure IDs: {numbered_unexpected}")

    rights_counts = Counter(row.row_id for row in rights)
    rights_duplicates = sorted(row_id for row_id, count in rights_counts.items() if count > 1)
    rights_missing = sorted(RIGHTS_EXPECTED - set(rights_counts))
    rights_unexpected = sorted(set(rights_counts) - RIGHTS_EXPECTED)

    if rights_duplicates:
        raise VerificationError(f"duplicate rights closure IDs: {rights_duplicates}")
    if rights_missing:
        raise VerificationError(f"missing rights closure IDs: {rights_missing}")
    if rights_unexpected:
        raise VerificationError(f"unexpected rights closure IDs: {rights_unexpected}")

    numbered_unresolved = 0
    rights_unresolved = 0
    for row, category in [(row, "numbered") for row in numbered] + [(row, "rights") for row in rights]:
        if row.disposition not in ALLOWED:
            raise VerificationError(f"{row.row_id} has invalid disposition {row.disposition!r}")
        if row.disposition == "UNRESOLVED":
            if category == "numbered":
                numbered_unresolved += 1
            else:
                rights_unresolved += 1
            raise VerificationError(f"{row.row_id} remains UNRESOLVED")
        if not row.target or row.target == "-":
            raise VerificationError(f"{row.row_id} has no current owner/trigger/superseding rule")
        if not row.rationale or row.rationale == "-":
            raise VerificationError(f"{row.row_id} has no rationale")
        validate_target_paths(row, repo_root)

    return {
        "numbered_closure": len(numbered),
        "numbered_duplicate": len(numbered_duplicates),
        "numbered_missing": len(numbered_missing),
        "numbered_unexpected": len(numbered_unexpected),
        "numbered_unresolved": numbered_unresolved,
        "rights_closure": len(rights),
        "rights_duplicate": len(rights_duplicates),
        "rights_missing": len(rights_missing),
        "rights_unexpected": len(rights_unexpected),
        "rights_unresolved": rights_unresolved,
    }


def verify_repository(repo_root: Path = ROOT) -> dict[str, int]:
    research = repo_root / RESEARCH_DIR
    closure = research / CLOSURE_NAME
    if not closure.is_file():
        raise VerificationError("CLOSURE.md is missing")

    closure_text = closure.read_text(encoding="utf-8")
    if "AUTHORITY: NONE" not in closure_text.splitlines()[:8]:
        raise VerificationError("CLOSURE.md must declare AUTHORITY: NONE")

    imported = source_ids(research)
    source_counts = Counter(imported)
    source_duplicates = sorted(row_id for row_id, count in source_counts.items() if count > 1)
    if source_duplicates:
        raise VerificationError(f"source IDs are not unique: {source_duplicates}")
    if len(imported) != BASELINE_COUNT:
        raise VerificationError(
            f"authoritative imported source count changed: expected closure baseline "
            f"{BASELINE_COUNT}, found {len(imported)}; re-evaluate and close the current set"
        )

    numbered, rights = closure_rows(closure)
    result = validate_rows(set(imported), numbered, rights, repo_root)
    result.update(
        {
            "numbered_source": len(imported),
            "source_duplicate": len(source_duplicates),
            "rights_expected": len(RIGHTS_EXPECTED),
        }
    )
    return result


def fail(message: str) -> None:
    print(f"LenBands founder closure check failed: {message}", file=sys.stderr)
    raise SystemExit(1)


def main() -> None:
    try:
        result = verify_repository()
    except VerificationError as exc:
        fail(str(exc))

    print(
        "LenBands founder closure: "
        f"source={result['numbered_source']} source_duplicates={result['source_duplicate']} "
        f"closure={result['numbered_closure']} duplicates={result['numbered_duplicate']} "
        f"missing={result['numbered_missing']} "
        f"unexpected={result['numbered_unexpected']} unresolved={result['numbered_unresolved']}; "
        f"rights_expected={result['rights_expected']} rights_closure={result['rights_closure']} "
        f"rights_duplicates={result['rights_duplicate']} rights_missing={result['rights_missing']} "
        f"rights_unexpected={result['rights_unexpected']} rights_unresolved={result['rights_unresolved']}"
    )


if __name__ == "__main__":
    main()
