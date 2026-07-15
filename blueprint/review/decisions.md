# Review Decisions (Governance)

Founder decisions defining the review/governance section. Indexed from [product/foundational-decisions.md](../product/foundational-decisions.md).

---

## GV-001 — review/ is the Blueprint Governance layer
**Date:** 2026-07-16 · **Status:** Decided (Founder)

**Decision.** `review/` is the **Blueprint Governance layer** — it defines and executes repository-wide quality assurance. It is **not** per-document review.
- **Verifies (at minimum):** structural consistency; terminology consistency; traceability completeness; dependency correctness; prerequisite consistency; canonical object uniqueness; duplicate detection; orphan detection; Academic vs General Training separation; evidence labeling consistency; decision-record consistency; reference integrity across all canonical layers.
- **Output:** a formal **Blueprint Health Report** ([health-report.md](health-report.md)) describing every issue with **severity, rationale, and recommended resolution**.
- **Severity levels:** `Critical` (blocks freeze) · `High` · `Medium` · `Low`.
- The QA check catalog is [governance.md](governance.md); execution results in [health-report.md](health-report.md).

---

## GV-002 — Academic vs General Training variant-separation convention
**Date:** 2026-07-16 · **Status:** Decided (Founder) · **Resolves:** task #5

**Decision.** Per [FD-001](../product/foundational-decisions.md) (shared concepts modeled once; variant-specific content separated), the variant-separation convention is:
- **Shared content** lives at the category/skill level with **no variant tag** (default = `shared`). This includes Listening, Speaking, Writing Task 2, and all shared Knowledge / Practice / Assessment types.
- **Variant-specific content** (GT Writing Task 1 letter; GT Reading passages + conversion) lives in a clearly separated sub-location and/or carries a `variant` field: `academic` | `general-training` | `shared`.
- **GT is added later without restructuring** existing Academic artifacts ([FD-001](../product/foundational-decisions.md)).
- **Current state:** Academic-only (v1); GT deferred. The convention guarantees GT-addition is non-breaking.

**Implication.** When GT is built, only variant-specific docs/objects are added (tagged `general-training`); shared objects are reused as-is.
