"use client";

import { useEffect, useMemo, useState } from "react";
import { api } from "@/lib/api";
import { newIdempotencyKey } from "@/lib/idempotency";
import type { components } from "@/generated/public-v1";

type TargetProfile = components["schemas"]["TargetProfile"];
type PracticeActivity = components["schemas"]["PracticeActivity"];
type Attempt = components["schemas"]["Attempt"];
type Choice = components["schemas"]["Choice"];

export default function Page() {
  const [ready, setReady] = useState(false);
  const [error, setError] = useState("");
  const [target, setTarget] = useState<TargetProfile | null>(null);
  const [band, setBand] = useState("6.5");
  const [activity, setActivity] = useState<PracticeActivity | null>(null);
  const [attempt, setAttempt] = useState<Attempt | null>(null);
  const [answers, setAnswers] = useState<Record<string, Choice>>({});

  useEffect(() => {
    void (async () => {
      const session = await api.POST("/v1/session");
      if (session.error) return setError("Could not establish learner session.");
      const current = await api.GET("/v1/target-profile");
      if (current.data) setTarget(current.data);
      setReady(true);
    })();
  }, []);

  const allAnswered = useMemo(
    () => Boolean(activity) && activity!.items.every((item) => Boolean(answers[item.item_id])),
    [activity, answers],
  );

  async function saveTarget() {
    setError("");
    const response = await api.PUT("/v1/target-profile", {
      body: {
        test_variant: "ACADEMIC",
        minimum_reading_band: Number(band),
        expected_resource_revision: target?.resource_revision ?? 0,
      },
    });
    if (response.error) return setError(response.error.error.message);
    setTarget(response.data ?? null);
  }

  async function startActivity() {
    setError("");
    const response = await api.POST("/v1/practice-activities", {
      params: { header: { "Idempotency-Key": newIdempotencyKey("activity") } },
      body: { practice_mode_id: "PM-R03" },
    });
    if (response.error) return setError(response.error.error.message);
    setActivity(response.data ?? null);
    setAttempt(null);
    setAnswers({});
  }

  async function submit() {
    if (!activity || !allAnswered) return;
    setError("");
    let currentAttempt = attempt;
    if (!currentAttempt) {
      const created = await api.POST("/v1/attempts", {
        params: { header: { "Idempotency-Key": newIdempotencyKey("attempt") } },
        body: { practice_activity_id: activity.practice_activity_id },
      });
      if (created.error || !created.data) return setError(created.error?.error.message ?? "Attempt creation failed.");
      currentAttempt = created.data;
      setAttempt(created.data);
    }
    const submitted = await api.POST("/v1/attempts/{attempt_id}/submissions", {
      params: {
        path: { attempt_id: currentAttempt.attempt_id },
        header: { "Idempotency-Key": newIdempotencyKey("submit") },
      },
      body: {
        expected_resource_revision: currentAttempt.resource_revision,
        answers: activity.items.map((item) => ({ item_id: item.item_id, choice: answers[item.item_id] })),
      },
    });
    if (submitted.error) return setError(submitted.error.error.message);
    setAttempt(submitted.data ?? null);
  }

  return (
    <main>
      <h1>Reading classification training</h1>
      <p className="muted">Synthetic training only. This activity does not create Band readiness or certification evidence.</p>
      {error && <p className="error" role="alert">{error}</p>}

      <section aria-labelledby="target-heading">
        <h2 id="target-heading">Target profile</h2>
        <label>
          Variant
          <select value="ACADEMIC" disabled><option>Academic</option></select>
        </label>
        <label>
          Minimum Reading Band
          <input aria-label="Minimum Reading Band" type="number" min="3" max="9" step="0.5" value={band} onChange={(event) => setBand(event.target.value)} />
        </label>
        <button onClick={saveTarget} disabled={!ready}>Save target</button>
        {target && <p data-testid="target-saved">Academic target saved at revision {target.resource_revision}.</p>}
      </section>

      <section aria-labelledby="practice-heading">
        <h2 id="practice-heading">T/F/NG + Y/N/NG</h2>
        <p>Deterministic Reading classification practice backed by PM-R03.</p>
        <button onClick={startActivity} disabled={!target}>Start activity</button>
      </section>

      {activity && (
        <section aria-labelledby="activity-heading">
          <h2 id="activity-heading">{activity.stimulus.title}</h2>
          <p>{activity.stimulus.text}</p>
          {activity.items.map((item) => (
            <div className="item" key={item.item_id} data-testid={`item-${item.item_id}`}>
              <p>{item.statement}</p>
              {item.choices.map((choice) => (
                <label className="choice" key={choice}>
                  <input type="radio" name={item.item_id} value={choice} checked={answers[item.item_id] === choice} onChange={() => setAnswers((old) => ({ ...old, [item.item_id]: choice }))} />
                  {choice.replace("_", " ")}
                </label>
              ))}
            </div>
          ))}
          <button onClick={submit} disabled={!allAnswered || attempt?.status === "EVALUATED"}>Submit answers</button>
        </section>
      )}

      {attempt?.status === "EVALUATED" && attempt.observation && (
        <section aria-labelledby="result-heading" data-testid="result">
          <h2 id="result-heading">Result</h2>
          <p data-testid="score">{attempt.observation.raw_score} / {attempt.observation.max_score} correct</p>
          <p className="muted">Training observation only — NOT_EVIDENCE_CANDIDATE.</p>
          {attempt.feedback?.map((feedback) => (
            <div className="feedback" key={feedback.item_id}>
              <strong>{feedback.correct ? "Correct" : "Review"}</strong>: {feedback.explanation}
              <div>Correct answer: {feedback.correct_choice.replace("_", " ")}</div>
            </div>
          ))}
        </section>
      )}
    </main>
  );
}
