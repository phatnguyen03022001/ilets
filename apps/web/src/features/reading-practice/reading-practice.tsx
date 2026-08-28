"use client";

import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { useTranslations } from "next-intl";
import { useEffect } from "react";
import { Controller, useForm } from "react-hook-form";
import type { components } from "@/generated/public-v1";
import { Alert } from "@/components/ui/alert";
import { Button } from "@/components/ui/button";
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { RadioGroup, RadioGroupItem } from "@/components/ui/radio-group";
import { api } from "@/lib/api";
import { newIdempotencyKey } from "@/lib/idempotency";

type TargetProfile = components["schemas"]["TargetProfile"];
type TestVariant = components["schemas"]["TestVariant"];
type PracticeActivity = components["schemas"]["PracticeActivity"];
type Attempt = components["schemas"]["Attempt"];
type Choice = components["schemas"]["Choice"];

type TargetForm = {
  testVariant: "" | TestVariant;
  targetOverallBand: string;
  minimumListeningBand: string;
  minimumReadingBand: string;
  minimumWritingBand: string;
  minimumSpeakingBand: string;
};

type AnswerForm = {
  answers: Record<string, Choice>;
};

const targetQueryKey = ["target-profile"] as const;

export default function ReadingPractice() {
  const t = useTranslations("Reading");
  const queryClient = useQueryClient();
  const targetForm = useForm<TargetForm>({
    defaultValues: {
      testVariant: "",
      targetOverallBand: "",
      minimumListeningBand: "",
      minimumReadingBand: "",
      minimumWritingBand: "",
      minimumSpeakingBand: "",
    },
  });
  const answerForm = useForm<AnswerForm>({
    defaultValues: { answers: {} },
  });

  const sessionQuery = useQuery({
    queryKey: ["learner-session"],
    staleTime: Infinity,
    queryFn: async () => {
      const response = await api.POST("/v1/session");
      if (response.error)
        throw new Error(response.error.error.message || t("errors.session"));
      return response.data;
    },
  });

  const targetQuery = useQuery<TargetProfile | null>({
    queryKey: targetQueryKey,
    enabled: sessionQuery.isSuccess,
    queryFn: async () => {
      const response = await api.GET("/v1/target-profile");
      if (response.response.status === 404) return null;
      if (response.error)
        throw new Error(response.error.error.message || t("errors.targetRead"));
      return response.data ?? null;
    },
  });

  const targetMutation = useMutation({
    mutationFn: async ({
      testVariant,
      targetOverallBand,
      minimumListeningBand,
      minimumReadingBand,
      minimumWritingBand,
      minimumSpeakingBand,
    }: TargetForm) => {
      if (!testVariant) throw new Error(t("errors.targetVariant"));

      const overallBand = targetOverallBand
        ? Number(targetOverallBand)
        : undefined;
      const listeningBand = minimumListeningBand
        ? Number(minimumListeningBand)
        : undefined;
      const readingBand = minimumReadingBand
        ? Number(minimumReadingBand)
        : undefined;
      const writingBand = minimumWritingBand
        ? Number(minimumWritingBand)
        : undefined;
      const speakingBand = minimumSpeakingBand
        ? Number(minimumSpeakingBand)
        : undefined;
      const hasAnyBand = Boolean(
        overallBand ??
        listeningBand ??
        readingBand ??
        writingBand ??
        speakingBand,
      );
      if (!hasAnyBand) throw new Error(t("errors.targetBand"));

      const response = await api.PUT("/v1/target-profile", {
        body: {
          test_variant: testVariant,
          target_overall_band: overallBand,
          minimum_listening_band: listeningBand,
          minimum_reading_band: readingBand,
          minimum_writing_band: writingBand,
          minimum_speaking_band: speakingBand,
          expected_resource_revision: targetQuery.data?.resource_revision ?? 0,
        },
      });
      if (response.error || !response.data) {
        throw new Error(
          response.error?.error.message ?? t("errors.targetSave"),
        );
      }
      return response.data;
    },
    onSuccess: (target) => {
      queryClient.setQueryData(targetQueryKey, target);
      targetForm.reset({
        testVariant: target.test_variant,
        targetOverallBand:
          target.target_overall_band === undefined
            ? ""
            : String(target.target_overall_band),
        minimumListeningBand:
          target.minimum_listening_band === undefined
            ? ""
            : String(target.minimum_listening_band),
        minimumReadingBand:
          target.minimum_reading_band === undefined
            ? ""
            : String(target.minimum_reading_band),
        minimumWritingBand:
          target.minimum_writing_band === undefined
            ? ""
            : String(target.minimum_writing_band),
        minimumSpeakingBand:
          target.minimum_speaking_band === undefined
            ? ""
            : String(target.minimum_speaking_band),
      });
    },
  });

  useEffect(() => {
    if (!targetQuery.isSuccess || targetForm.formState.isDirty) return;

    targetForm.reset({
      testVariant: targetQuery.data?.test_variant ?? "",
      targetOverallBand:
        targetQuery.data?.target_overall_band === undefined
          ? ""
          : String(targetQuery.data.target_overall_band),
      minimumListeningBand:
        targetQuery.data?.minimum_listening_band === undefined
          ? ""
          : String(targetQuery.data.minimum_listening_band),
      minimumReadingBand:
        targetQuery.data?.minimum_reading_band === undefined
          ? ""
          : String(targetQuery.data.minimum_reading_band),
      minimumWritingBand:
        targetQuery.data?.minimum_writing_band === undefined
          ? ""
          : String(targetQuery.data.minimum_writing_band),
      minimumSpeakingBand:
        targetQuery.data?.minimum_speaking_band === undefined
          ? ""
          : String(targetQuery.data.minimum_speaking_band),
    });
  }, [targetForm, targetQuery.data, targetQuery.isSuccess]);

  const activityMutation = useMutation({
    mutationFn: async () => {
      const response = await api.POST("/v1/practice-activities", {
        params: {
          header: { "Idempotency-Key": newIdempotencyKey("activity") },
        },
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
        params: {
          header: { "Idempotency-Key": newIdempotencyKey("attempt") },
        },
        body: { practice_activity_id: activity.practice_activity_id },
      });
      if (response.error || !response.data) {
        throw new Error(response.error?.error.message ?? t("errors.attempt"));
      }
      return response.data;
    },
  });

  const submissionMutation = useMutation({
    mutationFn: async ({
      attempt,
      activity,
      answers,
    }: {
      attempt: Attempt;
      activity: PracticeActivity;
      answers: Record<string, Choice>;
    }) => {
      const response = await api.POST("/v1/attempts/{attempt_id}/submissions", {
        params: {
          path: { attempt_id: attempt.attempt_id },
          header: { "Idempotency-Key": newIdempotencyKey("submit") },
        },
        body: {
          expected_resource_revision: attempt.resource_revision,
          answers: activity.items.map((item) => ({
            item_id: item.item_id,
            choice: answers[item.item_id],
          })),
        },
      });
      if (response.error || !response.data) {
        throw new Error(
          response.error?.error.message ?? t("errors.submission"),
        );
      }
      return response.data;
    },
  });

  const activity = activityMutation.data;
  const attempt = submissionMutation.data ?? attemptMutation.data;
  const answers = answerForm.watch("answers");
  const allAnswered =
    Boolean(activity) &&
    activity!.items.every((item) => Boolean(answers[item.item_id]));
  const ready = sessionQuery.isSuccess && !targetQuery.isPending;
  const academicPracticeAvailable =
    targetQuery.data?.test_variant === "ACADEMIC";

  const errors = [
    sessionQuery.error,
    targetQuery.error,
    targetMutation.error,
    activityMutation.error,
    attemptMutation.error,
    submissionMutation.error,
  ];
  const currentError = errors.find(
    (error): error is Error => error instanceof Error,
  );

  function startActivity() {
    answerForm.reset({ answers: {} });
    attemptMutation.reset();
    submissionMutation.reset();
    activityMutation.mutate();
  }

  async function submitAnswers(values: AnswerForm) {
    if (!activity || !allAnswered || submissionMutation.isPending) return;

    try {
      const currentAttempt =
        attemptMutation.data ?? (await attemptMutation.mutateAsync(activity));
      await submissionMutation.mutateAsync({
        attempt: currentAttempt,
        activity,
        answers: values.answers,
      });
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

      {currentError && (
        <Alert variant="destructive">{currentError.message}</Alert>
      )}

      <Card aria-labelledby="target-heading">
        <CardHeader>
          <CardTitle id="target-heading">{t("targetProfile")}</CardTitle>
        </CardHeader>
        <CardContent>
          <form
            className="grid gap-5"
            onSubmit={targetForm.handleSubmit((values) =>
              targetMutation.mutate(values),
            )}
          >
            <div className="grid gap-2">
              <Label htmlFor="variant">{t("variant")}</Label>
              <select
                id="variant"
                aria-label={t("variant")}
                className="flex h-10 w-full rounded-md border border-input bg-background px-3 py-2 text-sm outline-none transition-shadow focus-visible:ring-2 focus-visible:ring-ring"
                required
                {...targetForm.register("testVariant", { required: true })}
              >
                <option value="">{t("selectVariant")}</option>
                <option value="ACADEMIC">{t("academic")}</option>
                <option value="GENERAL_TRAINING">{t("generalTraining")}</option>
              </select>
            </div>
            <div className="grid gap-2">
              <Label htmlFor="target-overall-band">
                {t("targetOverallBand")}
              </Label>
              <Input
                id="target-overall-band"
                aria-label={t("targetOverallBand")}
                type="number"
                min="3"
                max="9"
                step="0.5"
                {...targetForm.register("targetOverallBand")}
              />
            </div>
            <div className="grid gap-2">
              <Label htmlFor="minimum-listening-band">
                {t("minimumListeningBand")}
              </Label>
              <Input
                id="minimum-listening-band"
                aria-label={t("minimumListeningBand")}
                type="number"
                min="3"
                max="9"
                step="0.5"
                {...targetForm.register("minimumListeningBand")}
              />
            </div>
            <div className="grid gap-2">
              <Label htmlFor="minimum-reading-band">
                {t("minimumReadingBand")}
              </Label>
              <Input
                id="minimum-reading-band"
                aria-label={t("minimumReadingBand")}
                type="number"
                min="3"
                max="9"
                step="0.5"
                {...targetForm.register("minimumReadingBand")}
              />
            </div>
            <div className="grid gap-2">
              <Label htmlFor="minimum-writing-band">
                {t("minimumWritingBand")}
              </Label>
              <Input
                id="minimum-writing-band"
                aria-label={t("minimumWritingBand")}
                type="number"
                min="3"
                max="9"
                step="0.5"
                {...targetForm.register("minimumWritingBand")}
              />
            </div>
            <div className="grid gap-2">
              <Label htmlFor="minimum-speaking-band">
                {t("minimumSpeakingBand")}
              </Label>
              <Input
                id="minimum-speaking-band"
                aria-label={t("minimumSpeakingBand")}
                type="number"
                min="3"
                max="9"
                step="0.5"
                {...targetForm.register("minimumSpeakingBand")}
              />
            </div>
            <Button type="submit" disabled={!ready || targetMutation.isPending}>
              {t("saveTarget")}
            </Button>
            {targetQuery.data && (
              <p
                className="text-sm text-muted-foreground"
                data-testid="target-saved"
              >
                {t("targetSaved", {
                  revision: targetQuery.data.resource_revision,
                })}
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
          <Button
            onClick={startActivity}
            disabled={!academicPracticeAvailable || activityMutation.isPending}
          >
            {t("startActivity")}
          </Button>
          {targetQuery.data?.test_variant === "GENERAL_TRAINING" && (
            <p className="mt-3 text-sm text-muted-foreground">
              {t("academicPracticeOnly")}
            </p>
          )}
        </CardContent>
      </Card>

      {activity && (
        <Card aria-labelledby="activity-heading">
          <CardHeader>
            <CardTitle id="activity-heading">
              {activity.stimulus.title}
            </CardTitle>
            <CardDescription className="text-base leading-7 text-foreground">
              {activity.stimulus.text}
            </CardDescription>
          </CardHeader>
          <CardContent>
            <form
              className="grid gap-6"
              onSubmit={answerForm.handleSubmit(submitAnswers)}
            >
              {activity.items.map((item) => (
                <fieldset
                  className="grid gap-3 rounded-lg border p-4"
                  key={item.item_id}
                  data-testid={`item-${item.item_id}`}
                >
                  <legend className="px-1 font-medium">{item.statement}</legend>
                  <Controller
                    control={answerForm.control}
                    name={`answers.${item.item_id}`}
                    render={({ field }) => (
                      <RadioGroup
                        value={field.value}
                        onValueChange={(value) =>
                          field.onChange(value as Choice)
                        }
                      >
                        {item.choices.map((choice) => {
                          const id = `${item.item_id}-${choice}`;
                          return (
                            <div
                              className="flex items-center gap-2"
                              key={choice}
                            >
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
              <Button
                type="submit"
                disabled={
                  !allAnswered ||
                  submissionMutation.isPending ||
                  attempt?.status === "EVALUATED"
                }
              >
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
              {t("score", {
                raw: attempt.observation.raw_score,
                max: attempt.observation.max_score,
              })}
            </p>
            <p className="text-sm text-muted-foreground">{t("trainingOnly")}</p>
            {attempt.feedback?.map((feedback) => (
              <div className="rounded-lg border p-4" key={feedback.item_id}>
                <p>
                  <strong>
                    {feedback.correct ? t("correct") : t("review")}
                  </strong>
                  : {feedback.explanation}
                </p>
                <p className="mt-1 text-sm text-muted-foreground">
                  {t("correctAnswer", {
                    answer: choiceLabel(feedback.correct_choice),
                  })}
                </p>
              </div>
            ))}
          </CardContent>
        </Card>
      )}
    </main>
  );
}
