# LEARNER-EXPERIENCE-CONTRACT Learner-facing product-quality contract

> **CANONICAL PRODUCT AUTHORITY**
>
> Canonical within the PRODUCT authority domain under `docs/catalog/project.json`. This shard owns durable learner-facing experience semantics only; it does not freeze pixel tokens, a final brand system, a component library, or implementation-specific layout.

## Mental model and progressive disclosure

The default learner experience is a simple surface over a deep system. It exposes the learner's declared target, current evidence-backed state, a small number of material blockers or unknowns, and the next eligible action without requiring the learner to understand internal ontology, provider, schema, or policy machinery.

Primary navigation centers on **Today, Skills, Practice, Review, Media, Progress, and Mock**, with Target/Diagnostic, Vocabulary, and Profile available where context requires them. The default view should answer what matters now. Progressive disclosure may reveal skill/gap maps, Knowledge Objects, attempt/evidence history, policy, provenance, or deeper explanation on demand.

Internal IDs, graph mechanics, evaluator/provider plumbing, implementation state, and migration/catalog details are not normal learner chrome. A surface must not invent a global completeness percentage when material evidence, product coverage, or applicability remains unresolved.

## Visual character and learner language

The durable character is typography-led, content-first, calm, and legible. Hierarchy, spacing, readable text, state clarity, and learner action outrank decorative density.

Avoid product-wide patterns that obscure the learning task or make the system feel like spectacle: childish edtech treatment, decorative gamification, excessive gradients/glass, futuristic AI glow, dashboard mosaics, badge/icon clutter, ontology-shaped navigation, and generic AI-branded chrome.

Learner-facing copy names the **capability, action, result, state, blocker, uncertainty, or recovery path first**. AI or provider branding does not replace a meaningful product label. Privacy, provenance, source, evaluator/AI use, or other disclosures still appear when legally or materially required; capability-first language is not a reason to hide consequential provenance.

## Iconography, imagery, and depth

Icons are functional, restrained, coherent, and paired with text or another accessible fallback where meaning would otherwise be ambiguous. Do not use emoji, cartoon iconography, mixed visual packs, or icons on every surface merely as decoration.

Imagery is sparse and high quality. Background depth may support hierarchy, but content remains foreground and no required meaning is carried only by background treatment. Stock-tech spectacle, futuristic robots, and hero imagery that competes with the learner's task are excluded from the durable product language.

## Motion

Motion is used only when it clarifies state change, hierarchy, feedback, or navigation. It is not a spectacle layer. Avoid parallax, ambient decorative video, long blocking animation, and movement that destabilizes reading or controls.

Reduced-motion preference, accessibility, responsive behavior, stable layout, and task continuity outrank animation. A no-motion or reduced-motion path must preserve the same product meaning and actionability.

## Interaction micro-quality

Learner-facing state distinctions remain explicit and usable:

- loading and skeleton states preserve layout and do not pretend unknown content is already known;
- an empty state states what is absent and the next valid action when one exists;
- error, degraded, disabled, pending, and unavailable are distinct states rather than one generic failure;
- hover, focus, pressed, selected, expanded, and disabled controls are perceivable and consistent where the input modality supports them;
- keyboard use, focus order/visibility, and screen-reader semantics remain valid; color or an icon alone never carries a required distinction;
- concise copy preserves uncertainty rather than manufacturing certainty for visual cleanliness;
- primary, secondary, and recovery actions remain distinguishable; and
- responsive/mobile layouts preserve hierarchy, readability, stable controls, and safe capture/input behavior instead of merely shrinking a desktop dashboard.

Micro-quality is part of product truth when a missing or ambiguous interaction state could cause a learner to misunderstand evidence, product availability, eligibility, capture success, or the next valid action.

## Simple-surface / deep-system invariant

The system may maintain deep evidence, prerequisite, coverage, content, provider, lifecycle, and provenance machinery internally, but the learner-facing default remains small and consequential: target, current state, material blocker/unknown, and next eligible action.

Deeper explanation is available progressively when it helps the learner understand why a recommendation exists, what evidence supports a state, what remains unresolved, or how to recover. Progressive disclosure never hides a blocker that changes what the learner may validly do, and simplicity never collapses materially different states into one friendly-looking score.

This contract governs durable product-quality semantics. Exact font sizes, colors, radii, spacing scales, animation timings, breakpoints, component implementation, final artwork, and brand tokens remain implementation/design-system choices so long as they preserve these semantics.
