STATUS: CANONICAL
OWNS: media-learning product semantics, YouTube eligibility, authorized transcript paths, media lesson creation, player-compliance rules, and media-to-practice mapping
DEPENDS_ON: ../spec/07-PRACTICE.md, ../spec/10-CONTENT-MODEL.md, 01-skill-features.md, 02-practice-catalog.md
DOES_NOT_OWN: YouTube platform policy itself, Skill/Band truth, Assessment sufficiency, arbitrary media-downloading rights, API wire shape, provider legal terms, or framework selection

# Media and YouTube Learning

## Purpose

Use authentic media as a learning source without turning the product into a media clone, bypassing external platform restrictions, or creating a second learning/practice taxonomy.

# Core rule

```text
Media Source
  ↓
eligibility + rights/provenance
  ↓
authorized transcript/segment when available
  ↓
canonical target mapping
  ↓
existing Practice Mode
  ↓
Attempt
```

A media item is a **source**, not a Skill, Knowledge Object, Learning Mechanism, or Practice Type.

# Supported learning uses

Eligible media may instantiate existing practice semantics for:

1. **Dictation** — trusted transcript/answer reference required for automated correctness scoring;
2. **Shadowing** — pronunciation/prosody/fluency imitation where relevant;
3. **Retell/comprehension** — listening followed by meaning reconstruction/discussion;
4. **Vocabulary/collocation mining** — only from text the product is permitted to process or learner-entered phrases.

Do not create a parallel “YouTube practice” taxonomy. Source type and learning action remain separate dimensions.

# YouTube source flow

```text
learner supplies URL
  ↓
resolve video identity + permitted metadata
  ↓
check availability / embeddability / language / suitability
  ↓
resolve transcript/rights state
  ↓
map to canonical target + existing practice mode
  ↓
optional bounded analysis of permitted material
  ↓
validate generated lesson candidate
  ↓
learner preview/save
```

# Transcript states

## `AUTHORIZED_CREATOR_CAPTION`

Caption access is obtained through an authorized account/path with the permissions required by the external platform.

May support segmentation, trusted dictation keys, vocabulary extraction, and lesson generation within the permitted use.

## `LICENSED_TRANSCRIPT`

Transcript/timed text is supplied independently through an owned or licensed content source. Provenance/rights metadata identifies that source.

## `USER_PROVIDED_TRANSCRIPT`

The learner/content author supplies text under the product's applicable rights/usage terms. It must not be represented as official platform caption data unless that provenance is independently established.

## `NO_AUTHORIZED_TRANSCRIPT`

Playback may still be used where embedding is permitted, but the product cannot fabricate a trustworthy caption-derived answer key.

Possible uses include:

- shadowing from playback;
- learner notes/bookmarks;
- retell/discussion prompts that do not require copied transcript text;
- learner-entered vocabulary.

Automated transcript-dependent scoring is disabled.

# Explicit non-goal: arbitrary extraction

The initial architecture does not authorize:

```text
arbitrary public URL
→ unofficial downloader
→ server copies audio/video
→ unrestricted transcription/storage
```

A future broader extraction path requires a new current decision after legal/platform, rights, storage, quota, deletion, and privacy review.

# Embedded-player compliance boundary

When YouTube playback is used, implementation follows the live supported IFrame/API requirements rather than a copied local interpretation.

Product-level invariants include:

- use a supported embed/player integration path;
- preserve required controls/branding/visibility behavior;
- do not overlay/obscure controls in a prohibited way;
- do not suppress required origin/referrer behavior;
- do not remove advertising or standard player behavior;
- identify the external content/source truthfully;
- provide independent learning value rather than mirror the media platform.

External policy pages remain external authority and must be rechecked before implementation/release.

# Caption API constraint

Current YouTube Data API caption-download behavior requires appropriate authorization/permission for the video. Therefore arbitrary public-video caption download is not a product assumption.

This is why transcript eligibility is a first-class state rather than a hidden implementation detail.

# `MediaSource` product contract

Conceptual fields:

```text
id
provider
provider_source_id
canonical_url
source_title
source_author_or_channel
language
duration
playability_or_embeddability
rights_state
transcript_state
transcript_ref optional
metadata_snapshot
```

Exact wire/storage shape belongs to machine/persistence implementation.

# `MediaLesson` product contract

Conceptual fields:

```text
media_source_id
segment_start / segment_end
practice_mode_id
canonical_target_ids
transcript_ref when permitted/required
prompt_set
scaffold_policy
feedback_policy
difficulty_metadata
provenance
```

Generated proposals do not become valid lessons until canonical targets, practice mode, source rights, and transcript eligibility validate.

# Segmentation defaults

For authorized transcript-backed media, bounded analysis may propose chunk boundaries.

Useful initial UX ranges include:

- short focused chunks for dictation/shadowing;
- longer segments for gist/detail/retell;
- longer continuous material only when the selected Practice Mode requires integrated performance.

Exact seconds are mutable content/product calibration, not learning truth.

# Difficulty

Media difficulty may consider:

- speech rate;
- intelligibility/accent demand;
- lexical rarity;
- syntactic complexity;
- discourse density;
- background noise;
- number of speakers;
- abstraction/topic familiarity;
- transcript reliability.

A media difficulty label is not an IELTS Band score.

# AI boundary

AI may:

- segment permitted transcript text;
- propose vocabulary/collocations;
- generate comprehension/retell/discussion prompts;
- propose canonical target candidates;
- generate bounded feedback under normal Practice/Assessment rules.

AI may not:

- claim inaccessible captions are official;
- silently download/copy restricted media;
- upgrade ordinary media practice into certification evidence;
- rewrite canonical target semantics;
- override rights/eligibility validation.

# Evidence boundary

Media-based practice is **non-certifying by default**.

A media-derived Observation may become formal evidence only when it independently passes normal Assessment validity, target fit, independence, provenance, scoring-quality, and inference-scope rules.

# Storage rule

For an embedded YouTube source, preferred persisted product state is reference-based:

```text
provider video identity / URL
allowed metadata
segment timestamps
lesson configuration
authorized transcript reference when permitted
learner attempts / derived feedback/evidence
```

Do not mirror audiovisual content merely for implementation convenience.

# Source removal/failure

A source may disappear, become private, lose embedding permission, or otherwise become unavailable.

The product must:

- preserve historical learner attempts/evidence at their valid historical scope;
- mark the source unavailable/ineligible as appropriate;
- stop scheduling inaccessible lesson instances;
- substitute an equivalent eligible source when possible;
- never reinterpret source loss as learner regression.