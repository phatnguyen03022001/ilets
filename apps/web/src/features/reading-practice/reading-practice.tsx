"use client";

import { SignInButton, useAuth } from "@clerk/nextjs";
import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { useTranslations } from "next-intl";
import { useEffect, useMemo } from "react";
import { Controller, useForm } from "react-hook-form";
import {
  createAttempt,
  createPracticeActivity,
  getTargetProfile,
  putTargetProfile,
  submitAttempt,
  type Attempt,
  type AttemptSubmissionResult,
  type PracticeActivity,
  type TargetProfile,
  type TargetProfileReadResult,
  type TestVariant,
} from "@/generated/public";
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
import { createPublicApi } from "@/lib/api";
import { newIdempotencyKey } from "@/lib/idempotency";

type TargetForm = {
  testVariant: "" | TestVariant;
  targetOverallBand: string;
  minimumListeningBand: string;
  minimumReadingBand: string;
  minimumWritingBand: string;
  minimumSpeakingBand: string;
};

type AnswerForm = { answers: Record<string, string> };

const targetQueryKey = ["target-profile"] as const;

function apiError(error: unknown, fallback: string) {
  if (
    error &&
    typeof error === "object" &&
    "error" in error &&
    error.error &&
    typeof error.error === "object" &&
    "message" in error.error &&
    typeof error.error.message === "string"
  ) {
    return error.error.message;
  }
  return fallback;
}

function configuredProfile(result: TargetProfileReadResult | undefined) {
  return result?.state === "CONFIGURED" ? result.profile : undefined;
}

export default function ReadingPractice() {
  const t = useTranslations("Reading");
  const queryClient = useQueryClient();
  const { getToken, isLoaded, isSignedIn } = useAuth();
  const api = useMemo(() => createPublicApi(getToken), [getToken]);
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
  const answerForm = useForm<AnswerForm>({ defaultValues: { answers: {} } });

  const targetQuery = useQuery({
    queryKey: targetQueryKey,
    enabled: isLoaded && isSignedIn,
    queryFn: async () => {
      const response = await getTargetProfile({ client: api });
      if (response.error)
        throw new Error(apiError(response.error, t("errors.targetRead")));
      if (!response.data) throw new Error(t("errors.targetRead"));
      return response.data;
    },
  });

  const targetMutation = useMutation({
    mutationFn: async (values: TargetForm) => {
      if (!values.testVariant) throw new Error(t("errors.targetVariant"));
      const request = {
        test_variant: values.testVariant,
        target_overall_band: optionalBand(values.targetOverallBand),
        minimum_listening_band: optionalBand(values.minimumListeningBand),
        minimum_reading_band: optionalBand(values.minimumReadingBand),
        minimum_writing_band: optionalBand(values.minimumWritingBand),
        minimum_speaking_band: optionalBand(values.minimumSpeakingBand),
      };
      if (
        !Object.values(request).some(
          (value, index) => index > 0 && value !== undefined,
        )
      ) {
        throw new Error(t("errors.targetBand"));
      }
      const current = configuredProfile(targetQuery.data);
      const response = await putTargetProfile({
        client: api,
        body: request,
        headers: {
          "Expected-Resource-Revision": current?.resource_revision ?? 0,
        },
      });
      if (response.error)
        throw new Error(apiError(response.error, t("errors.targetSave")));
      if (!response.data) throw new Error(t("errors.targetSave"));
      return response.data;
    },
    onSuccess: (profile) => {
      const result: TargetProfileReadResult = { state: "CONFIGURED", profile };
      queryClient.setQueryData(targetQueryKey, result);
      resetTargetForm(targetForm, profile);
    },
  });

  useEffect(() => {
    const profile = configuredProfile(targetQuery.data);
    if (!targetQuery.isSuccess || targetForm.formState.isDirty) return;
    resetTargetForm(targetForm, profile);
  }, [targetForm, targetQuery.data, targetQuery.isSuccess]);

  const activityMutation = useMutation({
    mutationFn: async () => {
      const response = await createPracticeActivity({
        client: api,
        headers: { "Idempotency-Key": newIdempotencyKey("activity") },
        body: { practice_mode_id: "PM-R03" },
      });
      if (response.error)
        throw new Error(apiError(response.error, t("errors.activity")));
      if (!response.data?.activity || response.data.outcome !== "ASSIGNED") {
        throw new Error(
          response.data?.unavailability?.explanation ?? t("errors.activity"),
        );
      }
      return response.data.activity;
    },
  });

  const attemptMutation = useMutation({
    mutationFn: async (activity: PracticeActivity) => {
      const response = await createAttempt({
        client: api,
        headers: { "Idempotency-Key": newIdempotencyKey("attempt") },
        body: { practice_activity_id: activity.practice_activity_id },
      });
      if (response.error)
        throw new Error(apiError(response.error, t("errors.attempt")));
      if (!response.data) throw new Error(t("errors.attempt"));
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
      answers: Record<string, string>;
    }) => {
      const response = await submitAttempt({
        client: api,
        path: { attempt_id: attempt.attempt_id },
        headers: { "Idempotency-Key": newIdempotencyKey("submit") },
        body: {
          response: {
            parts: activity.material.tasks.map((task) => ({
              task_id: task.task_id,
              selected_values: [answers[task.task_id]],
            })),
          },
          actual_conditions: {
            delivery: { state: "UNKNOWN" },
            assistance: [],
            exposure: [],
            input: [],
            timing: [],
          },
        },
      });
      if (response.error)
        throw new Error(apiError(response.error, t("errors.submission")));
      if (!response.data) throw new Error(t("errors.submission"));
      return response.data;
    },
  });

  const activity = activityMutation.data;
  const attempt: Attempt | undefined =
    submissionMutation.data?.attempt ?? attemptMutation.data;
  const submission: AttemptSubmissionResult | undefined =
    submissionMutation.data;
  const answers = answerForm.watch("answers");
  const allAnswered = activity
    ? activity.material.tasks.every((task) => Boolean(answers[task.task_id]))
    : false;
  const profile = configuredProfile(targetQuery.data);
  const academicPracticeAvailable =
    profile?.test_variant.state === "PRESENT" &&
    profile.test_variant.value === "Academic";

  const currentError = [
    targetQuery.error,
    targetMutation.error,
    activityMutation.error,
    attemptMutation.error,
    submissionMutation.error,
  ].find((error): error is Error => error instanceof Error);

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

  if (!isLoaded) {
    return (
      <main className="mx-auto max-w-3xl px-4 py-10">{t("authLoading")}</main>
    );
  }
  if (!isSignedIn) {
    return (
      <main className="mx-auto grid max-w-3xl gap-4 px-4 py-10">
        <Alert>{t("signInRequired")}</Alert>
        <SignInButton mode="modal">
          <Button>{t("signIn")}</Button>
        </SignInButton>
      </main>
    );
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
                <option value="Academic">{t("academic")}</option>
                <option value="General Training">{t("generalTraining")}</option>
              </select>
            </div>
            <BandInput
              id="target-overall-band"
              label={t("targetOverallBand")}
              registration={targetForm.register("targetOverallBand")}
            />
            <BandInput
              id="minimum-listening-band"
              label={t("minimumListeningBand")}
              registration={targetForm.register("minimumListeningBand")}
            />
            <BandInput
              id="minimum-reading-band"
              label={t("minimumReadingBand")}
              registration={targetForm.register("minimumReadingBand")}
            />
            <BandInput
              id="minimum-writing-band"
              label={t("minimumWritingBand")}
              registration={targetForm.register("minimumWritingBand")}
            />
            <BandInput
              id="minimum-speaking-band"
              label={t("minimumSpeakingBand")}
              registration={targetForm.register("minimumSpeakingBand")}
            />
            <Button
              type="submit"
              disabled={targetQuery.isPending || targetMutation.isPending}
            >
              {t("saveTarget")}
            </Button>
            {profile && (
              <p
                className="text-sm text-muted-foreground"
                data-testid="target-saved"
              >
                {t("targetSaved", { revision: profile.resource_revision })}
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
          {profile?.test_variant.value === "General Training" && (
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
              {activity.material.stimuli[0]?.title ?? t("practiceHeading")}
            </CardTitle>
            <CardDescription className="text-base leading-7 text-foreground">
              {activity.material.stimuli[0]?.text}
            </CardDescription>
          </CardHeader>
          <CardContent>
            <form
              className="grid gap-6"
              onSubmit={answerForm.handleSubmit(submitAnswers)}
            >
              {activity.material.tasks.map((task) => (
                <fieldset
                  className="grid gap-3 rounded-lg border p-4"
                  key={task.task_id}
                  data-testid={`item-${task.task_id}`}
                >
                  <legend className="px-1 font-medium">{task.prompt}</legend>
                  <Controller
                    control={answerForm.control}
                    name={`answers.${task.task_id}`}
                    render={({ field }) => (
                      <RadioGroup
                        value={field.value}
                        disabled={
                          attemptMutation.isPending ||
                          submissionMutation.isPending ||
                          attempt?.status === "evaluated"
                        }
                        onValueChange={field.onChange}
                      >
                        {(task.response_contract.options ?? []).map(
                          (option) => {
                            const id = `${task.task_id}-${option.value}`;
                            return (
                              <div
                                className="flex items-center gap-2"
                                key={option.value}
                              >
                                <RadioGroupItem id={id} value={option.value} />
                                <Label htmlFor={id}>{option.label}</Label>
                              </div>
                            );
                          },
                        )}
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
                  attempt?.status === "evaluated"
                }
              >
                {t("submitAnswers")}
              </Button>
            </form>
          </CardContent>
        </Card>
      )}

      {attempt?.status === "evaluated" && submission && (
        <Card aria-labelledby="result-heading" data-testid="result">
          <CardHeader>
            <CardTitle id="result-heading">{t("result")}</CardTitle>
          </CardHeader>
          <CardContent className="grid gap-4">
            <p className="text-sm text-muted-foreground">
              {t("trainingCompleted")}
            </p>
            <Button
              type="button"
              onClick={startActivity}
              disabled={activityMutation.isPending}
            >
              {t("practiceAgain")}
            </Button>
          </CardContent>
        </Card>
      )}
    </main>
  );
}

function optionalBand(value: string) {
  return value === "" ? undefined : Number(value);
}

function resetTargetForm(
  form: ReturnType<typeof useForm<TargetForm>>,
  profile: TargetProfile | undefined,
) {
  form.reset({
    testVariant:
      profile?.test_variant.state === "PRESENT"
        ? (profile.test_variant.value ?? "")
        : "",
    targetOverallBand:
      profile?.target_overall_band === undefined
        ? ""
        : String(profile.target_overall_band),
    minimumListeningBand:
      profile?.minimum_listening_band === undefined
        ? ""
        : String(profile.minimum_listening_band),
    minimumReadingBand:
      profile?.minimum_reading_band === undefined
        ? ""
        : String(profile.minimum_reading_band),
    minimumWritingBand:
      profile?.minimum_writing_band === undefined
        ? ""
        : String(profile.minimum_writing_band),
    minimumSpeakingBand:
      profile?.minimum_speaking_band === undefined
        ? ""
        : String(profile.minimum_speaking_band),
  });
}

type BandInputProps = {
  id: string;
  label: string;
  registration: ReturnType<ReturnType<typeof useForm<TargetForm>>["register"]>;
};

function BandInput({ id, label, registration }: BandInputProps) {
  return (
    <div className="grid gap-2">
      <Label htmlFor={id}>{label}</Label>
      <Input
        id={id}
        aria-label={label}
        type="number"
        min="3"
        max="9"
        step="0.5"
        {...registration}
      />
    </div>
  );
}
