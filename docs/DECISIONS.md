# Decisions

> **MIGRATION DRAFT — AUTHORITY NONE**
>
> Created only for `TASK-0001` revision 1. This file is non-authoritative until a later explicit canonical cutover. It cannot override `CONSTITUTION.md`, `OBJECTIVE.md`, `spec/**`, or `design/**`; those existing owners remain canonical.
>
> Target source snapshot: `phatnguyen03022001/ilets@23fa05c8586a9f295a3c0fe90774b78b248d61f7`.

This draft records only the bounded migration/adoption rationale authorized by `TASK-0001`. It does not instantiate `DEC-*` or `UNK-*` catalog identities because `docs/catalog/project.json` is explicitly outside this slice.

## Documentation-model adoption

**Draft outcome:** IELTS intends to adopt `phatnguyen03022001/agent-documents@e6728594b0371e2e941c1457fc8efdc14a90deee` as the exact support authority for the target repository's documentation structure and documentation-closure model.

The adopted revision defines the eight authority domains and their ownership boundaries, including `PRODUCT` for objective, actor/role semantics, scope, non-goals, and domain/external constraints, and `DECISIONS` for material decision explanation plus unknown/open-question context. It also keeps generic model meaning in `agent-documents` while the target repository owns actual project truth and instantiated documentation.

The rationale for this target migration is to make `agent-documents` materially shape IELTS documentation rather than remain disconnected from the target, while avoiding a one-step rewrite or split canonical authority. This first slice therefore creates only two non-authoritative migration drafts and leaves all current IELTS canonical owners untouched.

**Candidate status is material:** `agent-documents@e6728594b0371e2e941c1457fc8efdc14a90deee` is an **unreleased V1 candidate**, not a globally stable release. IELTS adopting this exact candidate for its migration pilot does not promote, freeze, release, or otherwise make `agent-documents` globally stable.

**Consequences for this slice:** no remaining root authority-domain files, catalog, coverage state, generated program, or closure claim are created. A later explicitly authorized task must perform any further migration and eventual cutover.

**Sources:** `.agent/tasks/TASK-0001/task.yaml@23fa05c8586a9f295a3c0fe90774b78b248d61f7`; `phatnguyen03022001/agent-documents@e6728594b0371e2e941c1457fc8efdc14a90deee` — `DOCUMENT_MODEL.md`, `templates/docs/PRODUCT.md`, `templates/docs/DECISIONS.md`.

## Engineering claim and evidence semantics adoption

**Draft outcome:** IELTS intends to adopt `phatnguyen03022001/agent-standards@3f4950f280a3a35fee81471d4b83715fa72cf9ee` as the exact frozen V1 support authority for engineering claim and evidence semantics used by later target-owned assurance declarations.

The adopted revision's `STANDARD_MODEL.md` states that V1 is frozen at that revision and defines status/evidence semantics such as `PASS`, `PARTIAL`, `GAP`, `UNKNOWN`, `N/A`, and `EXCEPTION`, together with evidence sufficiency and cumulative achievement rules. Those generic meanings remain owned by `agent-standards`; this target draft only records the exact immutable adoption identity.

This slice performs **no** IELTS standards assessment. It does not claim any requirement `PASS`, achieved maturity level, `N/A`, exception, release readiness, or other assurance result. Any later target-owned assessment must be supported by sufficient attributable evidence under the adopted frozen semantics.

**Sources:** `.agent/tasks/TASK-0001/task.yaml@23fa05c8586a9f295a3c0fe90774b78b248d61f7`; `phatnguyen03022001/agent-standards@3f4950f280a3a35fee81471d4b83715fa72cf9ee` — `STANDARD_MODEL.md`.

## Ownership and cutover boundary

For this migration pilot:

- IELTS owns target-specific product truth and target-owned future documentation/assurance declarations.
- `agent-documents@e6728594b0371e2e941c1457fc8efdc14a90deee` remains the external owner of generic documentation-model, catalog, coverage, unknown, build/buy, and closure semantics adopted by the target.
- `agent-standards@3f4950f280a3a35fee81471d4b83715fa72cf9ee` remains the external owner of generic engineering requirement, status, level, and evidence semantics adopted by the target.
- Existing `CONSTITUTION.md`, `OBJECTIVE.md`, `spec/**`, and `design/**` remain canonical for IELTS at this stage.
- `docs/PRODUCT.md` and this file remain migration drafts with authority NONE and cannot resolve conflicts by overriding current owners.

No generic taxonomy or standards requirement catalog is copied into IELTS. No canonical cutover, `DOCS_READY`, documentation lock, implementation readiness, promotion, or release is declared here.

**Source:** `.agent/tasks/TASK-0001/task.yaml@23fa05c8586a9f295a3c0fe90774b78b248d61f7`.

## Migration unknowns

No unresolved authority conflict or contradiction was identified for the adoption decisions authorized in this slice.

Catalog decision identities, reversibility classifications, formal unknown records, and closure state are intentionally not instantiated because the catalog and remaining migration structure are explicitly reserved for later bounded work. Their absence in this draft is not silently converted into a resolved decision or readiness claim.

**Sources:** `.agent/tasks/TASK-0001/task.yaml@23fa05c8586a9f295a3c0fe90774b78b248d61f7`; `phatnguyen03022001/agent-documents@e6728594b0371e2e941c1457fc8efdc14a90deee` — `DOCUMENT_MODEL.md`.