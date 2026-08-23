#!/usr/bin/env python3
from __future__ import annotations

import re
from dataclasses import dataclass
from pathlib import PurePosixPath
from typing import Sequence


class ParseError(ValueError):
    pass


@dataclass(frozen=True)
class Change:
    status: str
    paths: tuple[str, ...]


@dataclass(frozen=True)
class Classification:
    mode: str
    reasons: tuple[str, ...] = ()


_RENAME_COPY = re.compile(r"^[RC](\d{1,3})$")
_SINGLE_PATH_STATUSES = {"A", "D", "M", "T"}

_WEB_CRITICAL_EXACT = {
    "apps/web/.prettierignore",
    "apps/web/.prettierrc.json",
    "apps/web/components.json",
    "apps/web/eslint.config.mjs",
    "apps/web/next-env.d.ts",
    "apps/web/next.config.ts",
    "apps/web/package.json",
    "apps/web/playwright.config.ts",
    "apps/web/pnpm-lock.yaml",
    "apps/web/pnpm-workspace.yaml",
    "apps/web/postcss.config.mjs",
    "apps/web/tsconfig.json",
    "apps/web/vitest.config.ts",
}

_GO_CRITICAL_EXACT = {
    "services/core-api/go.mod",
    "services/core-api/go.sum",
    "services/core-api/sqlc.yaml",
}

_ROOT_CRITICAL_EXACT = {
    ".go-version",
    ".node-version",
    "CONSTITUTION.md",
    "OBJECTIVE.md",
    "tools/check-affected",
    "tools/test_verify_lenbands_founder_closure.py",
    "tools/verify-lenbands-founder-closure.py",
    "tools/verify-local",
    "verify",
}

_CRITICAL_PREFIXES = (
    ".github/workflows/",
    "contracts/",
    "design/",
    "research/lenbands-founder-decisions/",
    "services/core-api/cmd/seed/",
    "services/core-api/internal/bootstrap/",
    "services/core-api/internal/db/",
    "services/core-api/internal/httpapi/generated/",
    "services/core-api/migrations/",
    "spec/",
    "tools/affected/",
    "tools/canonical/",
    "tools/contracts/",
    "tools/db/",
    "tools/local-verify/",
    "tools/slice/",
)


def _decode_path(raw: bytes) -> str:
    try:
        path = raw.decode("utf-8", errors="strict")
    except UnicodeDecodeError as exc:
        raise ParseError("git path is not valid UTF-8") from exc
    if not path:
        raise ParseError("git path is empty")
    pure = PurePosixPath(path)
    if pure.is_absolute() or any(part in {".", ".."} for part in pure.parts):
        raise ParseError(f"unsafe git path: {path!r}")
    return path


def parse_name_status_z(data: bytes) -> list[Change]:
    if not data:
        return []
    if not data.endswith(b"\0"):
        raise ParseError("git name-status stream is not NUL terminated")

    fields = data[:-1].split(b"\0")
    changes: list[Change] = []
    index = 0
    while index < len(fields):
        raw_status = fields[index]
        index += 1
        try:
            status = raw_status.decode("ascii", errors="strict")
        except UnicodeDecodeError as exc:
            raise ParseError("git status token is not ASCII") from exc

        if status in _SINGLE_PATH_STATUSES:
            if index >= len(fields):
                raise ParseError(f"missing path for git status {status}")
            path = _decode_path(fields[index])
            index += 1
            changes.append(Change(status, (path,)))
            continue

        match = _RENAME_COPY.fullmatch(status)
        if match:
            score = int(match.group(1))
            if score > 100:
                raise ParseError(f"invalid rename/copy score: {status}")
            if index + 1 >= len(fields):
                raise ParseError(f"missing old/new path for git status {status}")
            old_path = _decode_path(fields[index])
            new_path = _decode_path(fields[index + 1])
            index += 2
            changes.append(Change(status, (old_path, new_path)))
            continue

        raise ParseError(f"unsupported git status: {status!r}")

    return changes


def _critical_reason(path: str) -> str | None:
    if path in _ROOT_CRITICAL_EXACT:
        return f"critical:{path}"
    if path.startswith("compose") and path.endswith(".yml") and "/" not in path:
        return f"critical:{path}"
    if path in _WEB_CRITICAL_EXACT or path in _GO_CRITICAL_EXACT:
        return f"critical:{path}"
    if path.startswith("apps/web/e2e/"):
        return f"critical:{path}"
    if path.startswith("apps/web/src/generated/"):
        return f"critical:{path}"
    for prefix in _CRITICAL_PREFIXES:
        if path.startswith(prefix):
            return f"critical:{path}"
    return None


def _is_docs(path: str) -> bool:
    return path == "README.md" or (path.startswith("docs/") and path.endswith(".md"))


def classify_changes(changes: Sequence[Change]) -> Classification:
    if not changes:
        return Classification("no-changes")

    lanes: set[str] = set()
    full_reasons: list[str] = []

    for change in changes:
        if change.status == "T":
            full_reasons.append("git-status:type-change")
        for path in change.paths:
            critical = _critical_reason(path)
            if critical:
                full_reasons.append(critical)
                continue
            if _is_docs(path):
                continue
            if path.startswith("apps/web/"):
                lanes.add("web")
                continue
            if path.startswith("services/core-api/"):
                lanes.add("go")
                continue
            full_reasons.append(f"unknown:{path}")

    if full_reasons:
        return Classification("full", tuple(dict.fromkeys(full_reasons)))
    if lanes == {"web", "go"}:
        return Classification("web+go")
    if lanes == {"web"}:
        return Classification("web")
    if lanes == {"go"}:
        return Classification("go")
    return Classification("docs")
