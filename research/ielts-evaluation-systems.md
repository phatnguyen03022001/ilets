# IELTS Evaluation Systems Research

## Overview
Research on best practices for automated IELTS evaluation systems across all four skills.

## Research Findings

### AI Evaluation Standards
**Source: Official IELTS Band Descriptors (2026)**
- Band scores 0-9 for each skill
- Four main criteria per skill (e.g., TA/CC/LR/GRA for writing)
- Per-criterion grading required for detailed feedback
- Rubric alignment is critical for credibility

**Best Practices Identified:**
1. Use official band descriptors as ground truth
2. Implement per-criterion scoring (not just overall scores)
3. Provide detailed feedback with examples
4. Track progression over time (redraft analysis)
5. Handle edge cases and uncertainty

### AI Writing Evaluation
**Source: IELTS Writing Task 1 & 2 Band Descriptors**
- Task 1: Graph/diagram description, TA focus
- Task 2: Essay writing, TA/CC/LR/GRA all weighted
- Maximum 40 minutes per task
- Word count: 150-250 (Task 1), 250-300 (Task 2)

**Recommended Tech Stack:**
- **LLM**: GPT-4o or Claude 3.5 Sonnet (2026 models)
  - Reasoning capabilities for rubric alignment
  - Strong grammar and coherence evaluation
- **Tools**:
  - LangChain for prompt orchestration
  - Promptfoo for evaluation harness
  - Rag-stack for context management

**Common Pitfalls:**
- Over-consistency (all essays get 6.0)
- Over-critical feedback
- Band score inflation
- Context drift in long essays

### AI Speaking Evaluation
**Source: IELTS Speaking Band Descriptors**
- 4 parts: Introduction, Cue Card, Part 3 discussion
- Fluency & coherence (25%), Lexical resource (25%)
- Grammatical range (25%), Pronunciation (25%)
- 11-14 minutes total

**Recommended Tech Stack:**
- **Audio Processing**:
  - OpenAI Whisper v3-large (latest 2026 version)
  - Faster Whisper for batch processing
- **Evaluation**:
  - GPT-4o audio understanding
  - Claude 3.5 Sonnet for speech analysis
- **Metrics**:
  - Fluency rate (words per minute)
  - Hesitation detection
  - Pronunciation scoring (phoneme accuracy)

**Common Pitfalls:**
- Microphone quality issues
- Background noise interference
- Time limit enforcement
- Cue card topic coverage

### AI Listening Evaluation
**Source: IELTS Listening Band Descriptors**
- 4 sections, 40 questions
- Multiple choice, completion, matching
- 30 minutes total + 10 transfer time
- Various accents (British, Australian, etc.)

**Recommended Tech Stack:**
- **Audio Processing**:
  - FFmpeg for format conversion
  - Waveform analysis for timing
- **Evaluation**:
  - GPT-4o multimodal for transcription
  - Answer matching algorithms
- **Features**:
  - Accent recognition
  - Speed adjustment
  - Replay functionality

**Common Pitfalls:**
- Transcript availability (need for full transcripts)
- Audio quality degradation
- Time limit enforcement
- Section transitions

### AI Reading Evaluation
**Source: IELTS Reading Band Descriptors**
- 3 passages, 40 questions
- True/False/Not Given, Matching, etc.
- 60 minutes total

**Recommended Tech Stack:**
- **Content Management**:
  - Vector database for passages (pgvector)
  - RAG for answer validation
- **Evaluation**:
  - GPT-4o for question analysis
  - Rule-based validation for specific question types
- **Metrics**:
  - Reading speed (words per minute)
  - Answer accuracy by question type
  - Time management analysis

**Common Pitfalls:**
- Passage authenticity
- Question type variation
- Answer key management
- Time limit enforcement

## Implementation Recommendations

### Architecture
1. **Separate evaluation pipelines per skill**
2. **Shared AI infrastructure** (same models)
3. **Standardized scoring protocol**
4. **Custom scoring utilities** (band score calculation)

### Data Requirements
- Full transcripts for listening/reading
- Audio recordings for speaking
- Writing samples with band reference
- Per-criterion rubrics

### Quality Assurance
- Human-in-the-loop for initial evaluation
- Regular calibration with official samples
- Feedback from certified IELTS trainers
- Progression tracking for validation

## References
- [IELTS Official Band Descriptors](https://ielts.org/for-organisations/ielts-scoring-in-detail)
- [Cambridge IELTS Past Papers](https://www.cambridge.org/elt/english-for-educators/ielts/)
- [British Council IELTS Resources](https://www.britishcouncil.org/take-ielts/prepare)
