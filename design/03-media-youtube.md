STATUS: CANONICAL
OWNS: media-learning product semantics, YouTube eligibility, authorized transcript paths, media lesson creation, player-compliance rules, and media-to-practice mapping
DEPENDS_ON: ../spec/07-PRACTICE.md, ../spec/10-CONTENT-MODEL.md, 01-skill-features.md, 02-practice-catalog.md
DOES_NOT_OWN: YouTube platform policy itself, Skill/Band truth, assessment sufficiency, arbitrary media-downloading rights, API wire shape, or framework selection

# Media and YouTube Learning

## Purpose

Use authentic media as a learning substrate without turning the product into a YouTube clone, bypassing platform restrictions, or creating a second practice taxonomy.

The product inspiration is the useful interaction pattern seen in products such as Parroto: short video-driven dictation, shadowing, and vocabulary review. The IELTS product extends that pattern with explicit Skill targets, gap/action semantics, transfer, and evidence boundaries.

## Core rule

```text
Media Source
  ↓
eligibility + rights/provenance
  ↓
segment / transcript when permitted
  ↓
canonical target mapping
  ↓
existing Practice Mode
  ↓
Attempt
```

A YouTube video is a **source**, not a learning objective and not a Practice Type.

## Initial media modes

Eligible media supports four primary learning uses:

1. **Dictation** — `PM-L01`; transcript/answer key required;
2. **Shadowing** — `PM-S02`; transcript useful but not always required for basic imitation;
3. **Retell / comprehension** — listening first, then spoken/written summary or discussion using existing Speaking/Listening practice semantics;
4. **Vocabulary mining** — extract candidate vocabulary/collocation items from an authorized transcript or learner-selected phrase, then move suitable items to review.

A fifth user-facing "YouTube practice" taxonomy is forbidden. Media composes existing modes.

## YouTube source flow

```text
paste YouTube URL
  ↓
parse video ID
  ↓
fetch permitted metadata
  ↓
check availability / embeddability / language / duration
  ↓
resolve transcript path
  ↓
map to target skill + practice mode
  ↓
AI may segment / level / annotate permitted text
  ↓
learner previews generated lesson
  ↓
publish to personal library
```

## Transcript eligibility

The system recognizes four transcript states.

### `AUTHORIZED_CREATOR_CAPTION`

The authenticated user/channel has the rights required by YouTube to retrieve/edit the caption track through an authorized API flow.

This may support automated segmentation, dictation answer keys, vocabulary extraction, and lesson generation.

### `LICENSED_TRANSCRIPT`

A transcript or timed text is provided by a licensed/owned content source independently of YouTube playback.

The provenance/rights metadata must identify the source.

### `USER_PROVIDED_TRANSCRIPT`

The learner or content author supplies transcript text and accepts the product's rights/usage terms.

The system may use the text for lesson generation but should not represent it as official YouTube caption data.

### `NO_AUTHORIZED_TRANSCRIPT`

The app may still embed and play the video where permitted, but it must not fabricate an authoritative answer key from inaccessible captions.

Allowed uses can include:

- shadowing from playback;
- learner-created notes;
- manual segment bookmarks;
- retell/discussion prompts that do not require copied transcript text;
- learner-selected vocabulary typed by the learner.

Automated Dictation scoring is disabled unless a trustworthy answer transcript exists.

## Explicit non-goal: arbitrary media extraction

The initial architecture does **not** authorize this flow:

```text
arbitrary YouTube URL
→ unofficial downloader
→ server copies audio/video
→ unrestricted transcription
```

If a future legal/platform review establishes a compliant way to support a broader flow, it requires a new decision with rights, storage, policy, quota, and deletion consequences.

## Embedded-player rules

When YouTube playback is used, the implementation must preserve the standard YouTube experience required by YouTube's current API policies.

Product requirements include:

- use the supported YouTube embedded/IFrame player path;
- preserve required player controls and branding;
- do not place overlays over the player or obscure controls;
- do not suppress required playback-origin/referrer signals;
- respect minimum player-size and visibility/autoplay requirements;
- do not remove ads or standard player behavior;
- clearly identify the content as YouTube content;
- add independent learning value rather than presenting a YouTube clone.

Policy references to re-check before implementation/release:

- `https://developers.google.com/youtube/terms/developer-policies`
- `https://developers.google.com/youtube/terms/developer-policies-guide`
- `https://developers.google.com/youtube/terms/required-minimum-functionality`
- `https://developers.google.com/youtube/terms/api-services-terms-of-service`

The policy pages themselves remain external authority and may change.

## Caption API constraint

The YouTube Data API caption-download operation currently requires authorization associated with permission to edit the video. Therefore, public-video caption download must not be assumed to work for arbitrary learner-pasted links.

Current reference:

- `https://developers.google.com/youtube/v3/docs/captions/download`

This constraint is why transcript eligibility is explicit rather than hidden behind a "generate lesson" button.

## `MediaSource` product contract

A media source conceptually contains:

```text
id
provider                  # youtube | owned_audio | owned_video | other_licensed
provider_source_id
canonical_url
source_title
source_author_or_channel
language
duration
embeddable_or_playable
rights_state
transcript_state
transcript_ref optional
metadata_snapshot
```

This is product-level representation. Exact storage/API schema belongs to implementation contracts.

## `MediaLesson` contract

A generated lesson contains:

```text
media_source_id
segment_start / segment_end
practice_mode_id
canonical_target_ids
transcript_ref when authorized/required
prompt_set
scaffold_policy
feedback_policy
difficulty_label
provenance
```

AI may propose these fields, but canonical target IDs and rights/eligibility must validate before publication.

## Segmentation

For authorized transcript-backed media, AI may propose sentence/chunk boundaries using timing and linguistic structure.

Default preferred practice segment duration:

- 3–12 seconds for focused dictation/shadowing chunks;
- 20–90 seconds for gist/detail or retell segments;
- longer material only when the selected Practice Mode requires integrated performance.

These are product defaults, not learning laws.

## Media difficulty

Difficulty may consider:

- speech rate;
- accent/intelligibility;
- lexical rarity;
- syntactic complexity;
- discourse density;
- background noise;
- number of speakers;
- abstraction/topic familiarity;
- transcript reliability.

A media difficulty label is not an IELTS Band score.

## AI use

AI may:

- segment authorized transcript text;
- propose vocabulary/collocations;
- generate comprehension prompts;
- classify candidate practice targets;
- generate retell/discussion prompts;
- produce feedback from learner attempts when the owning Assessment/Practice rules permit it.

AI must not:

- claim inaccessible captions are official;
- silently download/copy restricted media;
- convert a media lesson into IELTS certification evidence by itself;
- rewrite canonical target semantics.

## Certification boundary

Media-based practice is **non-certifying by default**.

A media-derived activity contributes formal evidence only if it independently satisfies the normal Assessment validity contract, including target fit, conditions, provenance, independence, scoring validity, and scope. Ordinary YouTube shadowing/dictation does not certify IELTS Speaking or Listening Band.

## Storage rule

For YouTube sources, the preferred persisted representation is:

```text
video ID / canonical URL
metadata required for the product
segment timestamps
lesson configuration
authorized transcript reference when permitted
learner attempts / feedback
```

Do not mirror the audiovisual content merely for convenience.

## Removal / source failure

Media can disappear, become private, lose embedding permission, or change availability. A MediaLesson must therefore fail gracefully:

- preserve learner attempt/progress history;
- mark source unavailable;
- stop scheduling inaccessible lesson instances;
- replace with an equivalent eligible source when possible;
- never reinterpret source loss as learner regression.