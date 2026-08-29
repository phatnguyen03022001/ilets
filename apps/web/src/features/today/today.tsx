"use client";

import { SignInButton, useAuth } from "@clerk/nextjs";
import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { useTranslations } from "next-intl";
import { useEffect, useMemo, useState } from "react";
import { useForm } from "react-hook-form";
import {
  createAttempt,
  createPracticeActivity,
  getDailyPlan,
  listPracticeModes,
  putTargetProfile,
  submitAttempt,
  type Attempt,
  type AttemptSubmissionResult,
  type DailyPlanItem,
  type PlanReasonCode,
  type PracticeActivity,
  type PracticeActivityCreationResult,
  type TargetProfile,
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
import { createPublicApi } from "@/lib/api";
import { newIdempotencyKey } from "@/lib/idempotency";
import ActivityWorkspace, {
  type AnswerForm,
} from "@/features/reading-practice/activity-workspace";

type TargetForm = {
  testVariant: "" | TestVariant;
  targetOverallBand: string;
  minimumListeningBand: string;
  minimumReadingBand: string;
  minimumWritingBand: string;
  minimumSpeakingBand: string;
};

type ActivitySource = "TODAY" | "DIRECT";
type Assignment = {
  source: ActivitySource;
  result: PracticeActivityCreationResult;
};

const dailyPlanQueryKey = ["daily-plan"] as const;
const practiceModesQueryKey = ["practice-modes"] as const;

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

function configuredProfile(
  plan: Awaited<ReturnType<typeof getDailyPlan>>["data"],
) {
  const context = plan?.target_context;
  return context?.state === "CONFIGURED" ? context.profile : undefined;
}

function reasonMessageKey(code: PlanReasonCode) {
  switch (code) {
    case "PREREQUISITE_GAP":
      return "why.prerequisiteGap";
    case "ABILITY_GAP":
      return "why.abilityGap";
    case "INSUFFICIENT_EVIDENCE":
      return "why.insufficientEvidence";
    case "CONFLICTING_EVIDENCE":
      return "why.conflictingEvidence";
    case "STALE_EVIDENCE":
      return "why.staleEvidence";
    case "SCAFFOLD_DEPENDENCE":
      return "why.scaffoldDependence";
    case "TRANSFER_GAP":
      return "why.transferGap";
    case "FLUENCY_GAP":
      return "why.fluencyGap";
    case "REVIEW_DUE":
      return "why.reviewDue";
    case "EXAM_CONDITION_GAP":
      return "why.examConditionGap";
    case "DELIVERY_MODE_PREPARATION":
      return "why.deliveryModePreparation";
    case "PRODUCT_COVERAGE_BLOCKED":
      return "why.productCoverageBlocked";
  }
}

export default function Today() {
  const t = useTranslations("Today");
  const queryClient = useQueryClient();
  const { getToken, isLoaded, isSignedIn } = useAuth();
  const api = useMemo(() => createPublicApi(getToken), [getToken]);
  const [assignment, setAssignment] = useState<Assignment>();
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

  const dailyPlanQuery = useQuery({
    queryKey: dailyPlanQueryKey,
    enabled: isLoaded && isSignedIn,
    queryFn: async () => {
      const response = await getDailyPlan({ client: api });
      if (response.error)
        throw new Error(apiError(response.error, t("errors.dailyPlan")));
      if (!response.data) throw new Error(t("errors.dailyPlan"));
      return response.data;
    },
  });

  const practiceModesQuery = useQuery({
    queryKey: practiceModesQueryKey,
    enabled: isLoaded && isSignedIn,
    queryFn: async () => {
      const response = await listPracticeModes({ client: api });
      if (response.error)
        throw new Error(apiError(response.error, t("errors.practiceModes")));
      if (!response.data) throw new Error(t("errors.practiceModes"));
      return response.data;
    },
  });

  const profile = configuredProfile(dailyPlanQuery.data);
  const academicDirectPracticeAvailable =
    profile?.test_variant.state === "PRESENT" &&
    profile.test_variant.value === "Academic";

  useEffect(() => {
    if (!dailyPlanQuery.isSuccess || targetForm.formState.isDirty) return;
    resetTargetForm(targetForm, profile);
  }, [dailyPlanQuery.isSuccess, profile, targetForm]);

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
      const response = await putTargetProfile({
        client: api,
        body: request,
        headers: {
          "Expected-Resource-Revision": profile?.resource_revision ?? 0,
        },
      });
      if (response.error)
        throw new Error(apiError(response.error, t("errors.targetSave")));
      if (!response.data) throw new Error(t("errors.targetSave"));
      return response.data;
    },
    onSuccess: async (savedProfile) => {
      resetTargetForm(targetForm, savedProfile);
      await queryClient.invalidateQueries({ queryKey: dailyPlanQueryKey });
    },
  });

  const activityMutation = useMutation({
    mutationFn: async ({
      source,
      item,
    }: {
      source: ActivitySource;
      item?: DailyPlanItem;
    }): Promise<Assignment> => {
      const body =
        source === "TODAY"
          ? { daily_plan_item_id: item?.plan_item_id }
          : { practice_mode_id: "PM-R03" };
      if (source === "TODAY" && !body.daily_plan_item_id) {
        throw new Error(t("errors.activity"));
      }
      const response = await createPracticeActivity({
        client: api,
        headers: { "Idempotency-Key": newIdempotencyKey("activity") },
        body,
      });
      if (response.error)
        throw new Error(apiError(response.error, t("errors.activity")));
      if (!response.data) throw new Error(t("errors.activity"));
      return { source, result: response.data };
    },
    onSuccess: async (nextAssignment) => {
      setAssignment(nextAssignment);
      if (nextAssignment.result.outcome === "UNAVAILABLE") {
        await queryClient.invalidateQueries({ queryKey: dailyPlanQueryKey });
      }
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
            delivery: activity.delivery_mode,
            assistance: activity.assistance_conditions,
            exposure: activity.exposure_conditions,
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

  const assignedActivity =
    assignment?.result.outcome === "ASSIGNED"
      ? assignment.result.activity
      : undefined;
  const activity = assignedActivity;
  const attempt: Attempt | undefined =
    submissionMutation.data?.attempt ?? attemptMutation.data;
  const submission: AttemptSubmissionResult | undefined =
    submissionMutation.data;
  const recommendation = dailyPlanQuery.data?.items[0];
  const modeLabels = new Map(
    practiceModesQuery.data?.modes.map((mode) => [
      mode.practice_mode_id,
      mode.label,
    ]) ?? [],
  );
  const unavailable =
    assignment?.result.outcome === "UNAVAILABLE"
      ? assignment.result.unavailability
      : undefined;

  const currentError = [
    dailyPlanQuery.error,
    practiceModesQuery.error,
    targetMutation.error,
    activityMutation.error,
    attemptMutation.error,
    submissionMutation.error,
  ].find((error): error is Error => error instanceof Error);

  function resetActivityState() {
    answerForm.reset({ answers: {} });
    attemptMutation.reset();
    submissionMutation.reset();
    setAssignment(undefined);
  }

  function startRecommendation() {
    if (!recommendation) return;
    resetActivityState();
    activityMutation.mutate({ source: "TODAY", item: recommendation });
  }

  function startDirectPractice() {
    resetActivityState();
    activityMutation.mutate({ source: "DIRECT" });
  }

  async function submitAnswers(values: AnswerForm) {
    if (!activity || submissionMutation.isPending) return;
    try {
      const currentAttempt =
        attemptMutation.data ?? (await attemptMutation.mutateAsync(activity));
      await submissionMutation.mutateAsync({
        attempt: currentAttempt,
        activity,
        answers: values.answers,
      });
      await queryClient.invalidateQueries({ queryKey: dailyPlanQueryKey });
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
        <p className="text-sm font-medium text-muted-foreground">
          {t("eyebrow")}
        </p>
        <h1 className="text-3xl font-semibold tracking-tight">{t("title")}</h1>
        <p className="text-sm text-muted-foreground">{t("description")}</p>
      </header>

      {currentError && (
        <Alert variant="destructive" role="alert">
          {currentError.message}
        </Alert>
      )}

      {unavailable && (
        <Alert role="status" data-testid="unavailable">
          {unavailable.explanation ?? t("unavailable")}
        </Alert>
      )}

      <TargetCard
        t={t}
        form={targetForm}
        profile={profile}
        isPending={dailyPlanQuery.isPending || targetMutation.isPending}
        onSave={(values) => targetMutation.mutate(values)}
      />

      {(dailyPlanQuery.data?.unresolved_target_conditions.length ?? 0) > 0 && (
        <Card aria-labelledby="unresolved-heading">
          <CardHeader>
            <CardTitle id="unresolved-heading">
              {t("unresolvedHeading")}
            </CardTitle>
            <CardDescription>{t("unresolvedDescription")}</CardDescription>
          </CardHeader>
          <CardContent>
            <ul className="grid gap-2 text-sm">
              {dailyPlanQuery.data?.unresolved_target_conditions.map(
                (condition) => (
                  <li key={condition.condition_id}>{condition.explanation}</li>
                ),
              )}
            </ul>
          </CardContent>
        </Card>
      )}

      {profile && recommendation && (
        <Card aria-labelledby="recommendation-heading">
          <CardHeader>
            <CardTitle id="recommendation-heading">
              {t("recommendationHeading")}
            </CardTitle>
            <CardDescription>
              {modeLabels.get(recommendation.practice_mode_id) ??
                t("activityFallback")}
            </CardDescription>
          </CardHeader>
          <CardContent className="grid gap-4">
            <div className="space-y-2">
              <p className="text-sm font-medium">{t("whyHeading")}</p>
              {recommendation.reason_codes.map((reason) => (
                <p className="text-sm text-muted-foreground" key={reason}>
                  {t(reasonMessageKey(reason))}
                </p>
              ))}
            </div>
            <Button
              onClick={startRecommendation}
              disabled={activityMutation.isPending}
            >
              {t("startRecommendation")}
            </Button>
          </CardContent>
        </Card>
      )}

      {profile && !dailyPlanQuery.isPending && !recommendation && (
        <p className="text-sm text-muted-foreground" role="status">
          {t("noRecommendation")}
        </p>
      )}

      {(dailyPlanQuery.data?.coverage_gaps.length ?? 0) > 0 && (
        <Card aria-labelledby="coverage-heading">
          <CardHeader>
            <CardTitle id="coverage-heading">{t("coverageHeading")}</CardTitle>
            <CardDescription>{t("coverageDescription")}</CardDescription>
          </CardHeader>
          <CardContent>
            <ul className="grid gap-2 text-sm">
              {dailyPlanQuery.data?.coverage_gaps.map((gap, index) => (
                <li key={`${index}-${gap.blocking_consequence}`}>
                  {gap.blocking_consequence}
                </li>
              ))}
            </ul>
          </CardContent>
        </Card>
      )}

      {activity && (
        <>
          <Alert role="status">
            {activity.primary_activity_purpose === "TRAINING" &&
            activity.evidence_candidacy === "NOT_EVIDENCE_CANDIDATE"
              ? t("trainingBoundary")
              : t("assessmentBoundary")}
          </Alert>
          <ActivityWorkspace
            activity={activity}
            attempt={attempt}
            form={answerForm}
            isPending={
              attemptMutation.isPending || submissionMutation.isPending
            }
            submitLabel={t("submitAnswers")}
            fallbackTitle={t("activityFallback")}
            onSubmit={submitAnswers}
          />
        </>
      )}

      {attempt?.status === "evaluated" && submission && activity && (
        <Card aria-labelledby="result-heading" data-testid="result">
          <CardHeader>
            <CardTitle id="result-heading">{t("result")}</CardTitle>
          </CardHeader>
          <CardContent>
            <p className="text-sm text-muted-foreground" role="status">
              {activity.primary_activity_purpose === "TRAINING" &&
              activity.evidence_candidacy === "NOT_EVIDENCE_CANDIDATE"
                ? t("trainingCompleted")
                : t("assessmentCompleted")}
            </p>
          </CardContent>
        </Card>
      )}

      <Card aria-labelledby="direct-heading">
        <CardHeader>
          <CardTitle id="direct-heading">{t("directHeading")}</CardTitle>
          <CardDescription>{t("directDescription")}</CardDescription>
        </CardHeader>
        <CardContent>
          <Button
            variant="outline"
            onClick={startDirectPractice}
            disabled={
              !academicDirectPracticeAvailable || activityMutation.isPending
            }
          >
            {t("startDirect")}
          </Button>
        </CardContent>
      </Card>
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

type TFunction = ReturnType<typeof useTranslations>;
type TargetCardProps = {
  t: TFunction;
  form: ReturnType<typeof useForm<TargetForm>>;
  profile?: TargetProfile;
  isPending: boolean;
  onSave: (values: TargetForm) => void;
};

function TargetCard({ t, form, profile, isPending, onSave }: TargetCardProps) {
  return (
    <Card aria-labelledby="target-heading">
      <CardHeader>
        <CardTitle id="target-heading">
          {profile ? t("targetHeading") : t("targetSetupHeading")}
        </CardTitle>
        <CardDescription>
          {profile ? t("targetDescription") : t("targetSetupDescription")}
        </CardDescription>
      </CardHeader>
      <CardContent>
        <form className="grid gap-5" onSubmit={form.handleSubmit(onSave)}>
          <div className="grid gap-2">
            <Label htmlFor="variant">{t("variant")}</Label>
            <select
              id="variant"
              aria-label={t("variant")}
              className="flex h-10 w-full rounded-md border border-input bg-background px-3 py-2 text-sm outline-none transition-shadow focus-visible:ring-2 focus-visible:ring-ring"
              required
              {...form.register("testVariant", { required: true })}
            >
              <option value="">{t("selectVariant")}</option>
              <option value="Academic">{t("academic")}</option>
              <option value="General Training">{t("generalTraining")}</option>
            </select>
          </div>
          <BandInput
            id="target-overall-band"
            label={t("targetOverallBand")}
            registration={form.register("targetOverallBand")}
          />
          <BandInput
            id="minimum-listening-band"
            label={t("minimumListeningBand")}
            registration={form.register("minimumListeningBand")}
          />
          <BandInput
            id="minimum-reading-band"
            label={t("minimumReadingBand")}
            registration={form.register("minimumReadingBand")}
          />
          <BandInput
            id="minimum-writing-band"
            label={t("minimumWritingBand")}
            registration={form.register("minimumWritingBand")}
          />
          <BandInput
            id="minimum-speaking-band"
            label={t("minimumSpeakingBand")}
            registration={form.register("minimumSpeakingBand")}
          />
          <Button type="submit" disabled={isPending}>
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
  );
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
