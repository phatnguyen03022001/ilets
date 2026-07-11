# IELTS

A serious IELTS learning platform. Solo-founder · free-tier · AI-Runtime-friendly.

**Read first:**

- `.runtime/` — project context + laws + decision format (`CLAUDE.md` is generated from it by `script/bootstrap`)
- `project/founder-intent.md` — direction (Founder-owned)
- `project/engineering-strategy.md` — 10-year strategy
- `project/decisions.md` — ADR log

## Quickstart

```bash
cp .env.example .env.local      # fill DATABASE_URL (Neon) + ZAI_API_KEY
npm install
npm run db:generate && npm run db:migrate
npm run dev
```

## Common commands

```bash
npm run dev        # local dev server
npm run typecheck   # tsc --noEmit
npm run lint        # biome check (lint + format)
npm run build       # next build
```

## Status

The wedge-invariant engineering shell is **built and validated** (web · db · llm · env-validation · lint · ci + structural seams).
The wedge-dependent domain awaits the Founder's wedge decision — see `.runtime/`.
