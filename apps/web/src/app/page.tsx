"use client";

import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { useTranslations } from "next-intl";
import { Controller, useForm } from "react-hook-form";
import type { components } from "@/generated/public-v1";
import { Alert } from "@/components/ui/alert";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { RadioGroup, RadioGroupItem } from "@/components/ui/radio-group";
import { api } from "@/lib/api";
import { newIdempotencyKey } from "@/lib/idempotency";

type TargetProfile = components["schemas"]["TargetProfile"];
type PracticeActivity = components["schemas"]["PracticeActivity"];
type Attempt = components["schemas"]["Attempt"];
type Choice = components["schemas"]["Choice"];

type TargetForm = {
  minimumReadingBand: string;
};

type AnswerForm = {
  answers: Record<string, Choice>;
};

const targetQueryKey = ["target-profile"] as const;

export default function Page() {
  const t = useTranslations("Reading");
  const queryClient = useQueryClient();
  const targetForm = useForm<TargetForm>({ defaultValues: { minimumReadingBand: "6.5" } });
  const answerForm = useForm<AnswerForm>({ defaultValues: { answers: {} } });

  const sessionQuery = useQuery({
    queryKey: ["learner-session"],
    staleTime: Infinity,
    queryFn: async () => {
      const response = await api.POST("/v1/session");
      if (response.error) throw new Error(response.error.error.message || t("errors.session"));
      return response.data;
    },
  });

  const targetQuery = useQuery<TargetProfile | null>({
    queryKey: targetQueryKey,
    enabled: sessionQuery.isSuccess,
    queryFn: async () => {
      const response = await api.GET("/v1/target-profile");
      if (response.response.status === 404) return null;
      if (response.error) throw new Error(response.error.error.message || t("errors.targetRead"));
      return response.data ?? null;
    },
  });

  const targetMutation = useMutation({
    mutationFn: async ({ minimumReadingBand }: TargetForm) => {
      const response = await api.PUT("/v1/target-profile", {
        body: {
          test_variant: "ACADEMIC",
          minimum_reading_band: Number(minimumReadingBand),
          expected_resource_revision: targetQuery.data?.resource_revision ?? 0,
        },
      });
      if (response.error || !response.data) {
        throw new Error(response.error?.error.message ?? t("errors.targetSave"));
      }
      return response.data;
    },
    onSuccess: (target) => queryClient.setQueryData(targetQueryKey, target),
  });

  const activityMutation = useMutation({
    mutationFn: async () => {
      const response = await api.POST("/v1/practice-activities", {
        params: { header: { "Idempotency-Key": newIdempotencyKey("activity") } },
        body: { practice_mode_id: "PM-R03" },
      });
      if (response.error || !response.data) {
        throw new Error(response.error?.error.message ?? t("errors.activity"));
      }
      return response.data;
    },
  });

  const attemptMutation = useMutation({
    mutationFn: async (activity: PracticeActivity) => {
      const response = await api.POST("/v1/attempts", {
        params: { header: { "Idempotency-Key": newIdempotencyKey("attempt") } },
        body: { practice_activity_id: activity.practice_activity_id },
      });
      if (response.error || !response.data) {
        throw new Error(response.error?.error.message ?? t("errors.attempt"));
      }
      return response.data;
    },
  });

  const submissionMutation = useMutation({
    mutationFn: async ({ attempt, activity, answers }: { attempt: Attempt; activity: PracticeActivity; answers: Record<string, Choice> }) => {
      const response = await api.POST("/v1/attempts/{attempt_id}/submissions", {
        params: {
          path: { attempt_id: attempt.attempt_id },
          header: { "Idempotency-Key": newIdempotencyKey("submit") },
        },
        body: {
          expected_resource_revision: attempt.resource_revision,
          answers: activity.items.map((item) => ({ item_id: item.item_id, choice: answers[item.item_id] })),
        },
      });
      if (response.error || !response.data) {
        throw new Error(response.error?.error.message ?? t("errors.submission"));
      }
      return response.data;
    },
  });

  const activity = activityMutation.data;
  const attempt = submissionMutation.data ?? attemptMutation.data;
  const answers = answerForm.watch("answers");
  const allAnswered = Boolean(activity) && activity!.items.every((item) => Boolean(answers[item.item_id]));
  const ready = sessionQuery.isSuccess && !targetQuery.isPending;

  const errors = [
    sessionQuery.error,
    targetQuery.error,
    targetMutation.error,
    activityMutation.error,
    attemptMutation.error,
    submissionMutation.error,
  ];
  const currentError = errors.find((error): error is Error => error instanceof Error);

  function startActivity() {
    answerForm.reset({ answers: {} });
    attemptMutation.reset();
    submissionMutation.reset();
    activityMutation.mutate();
  }

  async function submitAnswers(values: AnswerForm) {
    if (!activity || !allAnswered || submissionMutation.isPending) return;

    try {
      const currentAttempt = attemptMutation.data ?? (await attemptMutation.mutateAsync(activity));
      await submissionMutation.mutateAsync({ attempt: currentAttempt, activity, answers: values.answers });
    } catch {
      // Mutation state owns the rendered error.
    }
  }

  function choiceLabel(choice: Choice) {
    return t(`choices.${choice}`);
  }

  return (
    <main className="mx-auto grid min-h-screen max-w-3xl gap-6 px-4 py-10 sm:px-6">
      <header className="space-y-2">
        <h1 className="text-3xl font-semibold tracking-tight">{t("title")}</h1>
        <p className="text-sm text-muted-foreground">{t("syntheticNotice")}</p>
      </header>

      {currentError && <Alert variant="destructive">{currentError.message}</Alert>}

      <Card aria-labelledby="target-heading">
        <CardHeader>
          <CardTitle id="target-heading">{t("targetProfile")}</CardTitle>
        </CardHeader>
        <CardContent>
          <form className="grid gap-5" onSubmit={targetForm.handleSubmit((values) => targetMutation.mutate(values))}>
            <div className="grid gap-2">
              <Label htmlFor="variant">{t("variant")}</Label>
              <Input id="variant" value={t("academic")} disabled readOnly />
            </div>
            <div className="grid gap-2">
              <Label htmlFor="minimum-reading-band">{t("minimumReadingBand")}</Label>
              <Input
                id="minimum-reading-band"
                aria-label={t("minimumReadingBand")}
                type="number"
                min="3"
                max="9"
                step="0.5"
                {...targetForm.register("minimumReadingBand", { required: true })}
              />
            </div>
            <Button type="submit" disabled={!ready || targetMutation.isPending}>
              {t("saveTarget")}
            </Button>
            {targetQuery.data && (
              <p className="text-sm text-muted-foreground" data-testid="target-saved">
                {t("targetSaved", { revision: targetQuery.data.resource_revision })}
              </p>
            )}
          </form>
        </CardContent>
      </Card>

      <Card aria-labelledby="practice-heading">
        <CardHeader>
          <CardTitle id="practice-heading">{t("practiceHeading")}</CardTitle>
          <CardDescription>{t("practiceDescription")}</CardDescription>
        </CardHeader>
        <CardContent>
          <Button onClick={startActivity} disabled={!targetQuery.data || activityMutation.isPending}>
            {t("startActivity")}
          </Button>
        </CardContent>
      </Card>

      {activity && (
        <Card aria-labelledby="activity-heading">
          <CardHeader>
            <CardTitle id="activity-heading">{activity.stimulus.title}</CardTitle>
            <CardDescription className="text-base leading-7 text-foreground">{activity.stimulus.text}</CardDescription>
          </CardHeader>
          <CardContent>
            <form className="grid gap-6" onSubmit={answerForm.handleSubmit(submitAnswers)}>
              {activity.items.map((item) => (
                <fieldset className="grid gap-3 rounded-lg border p-4" key={item.item_id} data-testid={`item-${item.item_id}`}>
                  <legend className="px-1 font-medium">{item.statement}</legend>
                  <Controller
                    control={answerForm.control}
                    name={`answers.${item.item_id}`}
                    render={({ field }) => (
                      <RadioGroup value={field.value} onValueChange={(value) => field.onChange(value as Choice)}>
                        {item.choices.map((choice) => {
                          const id = `${item.item_id}-${choice}`;
                          return (
                            <div className="flex items-center gap-2" key={choice}>
                              <RadioGroupItem id={id} value={choice} />
                              <Label htmlFor={id}>{choiceLabel(choice)}</Label>
                            </div>
                          );
                        })}
                      </RadioGroup>
                    )}
                  />
                </fieldset>
              ))}
              <Button type="submit" disabled={!allAnswered || submissionMutation.isPending || attempt?.status === "EVALUATED"}>
                {t("submitAnswers")}
              </Button>
            </form>
          </CardContent>
        </Card>
      )}

      {attempt?.status === "EVALUATED" && attempt.observation && (
        <Card aria-labelledby="result-heading" data-testid="result">
          <CardHeader>
            <CardTitle id="result-heading">{t("result")}</CardTitle>
          </CardHeader>
          <CardContent className="grid gap-4">
            <p className="text-lg font-semibold" data-testid="score">
              {t("score", { raw: attempt.observation.raw_score, max: attempt.observation.max_score })}
            </p>
            <p className="text-sm text-muted-foreground">{t("trainingOnly")}</p>
            {attempt.feedback?.map((feedback) => (
              <div className="rounded-lg border p-4" key={feedback.item_id}>
                <p>
                  <strong>{feedback.correct ? t("correct") : t("review")}</strong>: {feedback.explanation}
                </p>
                <p className="mt-1 text-sm text-muted-foreground">
                  {t("correctAnswer", { answer: choiceLabel(feedback.correct_choice) })}
                </p>
              </div>
            ))}
          </CardContent>
        </Card>
      )}
    </main>
  );
}
