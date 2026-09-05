"use client";

import { Controller, type UseFormReturn } from "react-hook-form";
import { useCallback, useEffect, useRef, useState } from "react";
import type { Attempt, PracticeActivity } from "@/generated/public";
import { Button } from "@/components/ui/button";
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Label } from "@/components/ui/label";
import { RadioGroup, RadioGroupItem } from "@/components/ui/radio-group";

export type AnswerForm = { answers: Record<string, string> };

type Props = {
  activity: PracticeActivity;
  attempt?: Attempt;
  form: UseFormReturn<AnswerForm>;
  isPending: boolean;
  submitLabel: string;
  fallbackTitle: string;
  loadMedia: (activityId: string, mediaReference: string) => Promise<Blob>;
  mediaLoadingLabel: string;
  mediaErrorLabel: string;
  onSubmit: (values: AnswerForm) => void | Promise<void>;
};

type MediaState =
  | { status: "idle" }
  | { status: "loading"; reference: string }
  | { status: "ready"; reference: string; url: string }
  | { status: "error"; reference: string };

export default function ActivityWorkspace({
  activity,
  attempt,
  form,
  isPending,
  submitLabel,
  fallbackTitle,
  loadMedia,
  mediaLoadingLabel,
  mediaErrorLabel,
  onSubmit,
}: Props) {
  const answers = form.watch("answers");
  const allAnswered = activity.material.tasks.every((task) =>
    Boolean(answers[task.task_id]),
  );
  const locked = isPending || attempt?.status === "evaluated";
  const stimulus = activity.material.stimuli[0];
  const mediaReference =
    stimulus?.kind === "MEDIA_REFERENCE" ? stimulus.media_reference : undefined;
  const objectUrlRef = useRef<string | undefined>(undefined);
  const [mediaState, setMediaState] = useState<MediaState>({ status: "idle" });
  const releaseMediaUrl = useCallback(() => {
    if (objectUrlRef.current) {
      URL.revokeObjectURL(objectUrlRef.current);
      objectUrlRef.current = undefined;
    }
  }, []);

  useEffect(() => {
    let disposed = false;
    releaseMediaUrl();
    if (!mediaReference) {
      return () => {
        disposed = true;
        releaseMediaUrl();
      };
    }

    void loadMedia(activity.practice_activity_id, mediaReference)
      .then((blob) => {
        const url = URL.createObjectURL(blob);
        if (disposed) {
          URL.revokeObjectURL(url);
          return;
        }
        objectUrlRef.current = url;
        setMediaState({ status: "ready", reference: mediaReference, url });
      })
      .catch(() => {
        if (!disposed)
          setMediaState({ status: "error", reference: mediaReference });
      });

    return () => {
      disposed = true;
      releaseMediaUrl();
    };
  }, [
    activity.practice_activity_id,
    loadMedia,
    mediaReference,
    releaseMediaUrl,
  ]);

  const currentMediaState =
    mediaReference &&
    "reference" in mediaState &&
    mediaState.reference === mediaReference
      ? mediaState
      : undefined;
  const mediaReady = currentMediaState?.status === "ready";
  const mediaLoading =
    Boolean(mediaReference) &&
    (!currentMediaState || currentMediaState.status === "loading");
  const mediaError = currentMediaState?.status === "error";
  const mediaBlocked = Boolean(mediaReference) && !mediaReady;

  function handleMediaError() {
    releaseMediaUrl();
    if (mediaReference) {
      setMediaState({ status: "error", reference: mediaReference });
    }
  }

  return (
    <Card aria-labelledby="activity-heading">
      <CardHeader>
        <CardTitle id="activity-heading">
          {stimulus?.title ?? fallbackTitle}
        </CardTitle>
        {mediaLoading && (
          <p className="text-sm text-muted-foreground" role="status">
            {mediaLoadingLabel}
          </p>
        )}
        {mediaError && (
          <p className="text-sm text-destructive" role="alert">
            {mediaErrorLabel}
          </p>
        )}
        {mediaReady && currentMediaState.status === "ready" && (
          <audio
            aria-label={stimulus?.title ?? fallbackTitle}
            controls
            data-testid="activity-audio"
            onError={handleMediaError}
            preload="metadata"
            src={currentMediaState.url}
          />
        )}
        {stimulus?.text && (
          <CardDescription className="whitespace-pre-wrap text-base leading-7 text-foreground">
            {stimulus.text}
          </CardDescription>
        )}
      </CardHeader>
      <CardContent>
        <form className="grid gap-6" onSubmit={form.handleSubmit(onSubmit)}>
          {activity.material.tasks.map((task) => (
            <fieldset
              className="grid gap-3 rounded-lg border p-4"
              key={task.task_id}
              data-testid={`item-${task.task_id}`}
            >
              <legend className="px-1 font-medium">{task.prompt}</legend>
              <Controller
                control={form.control}
                name={`answers.${task.task_id}`}
                render={({ field }) => (
                  <RadioGroup
                    value={field.value ?? ""}
                    disabled={locked}
                    onValueChange={field.onChange}
                  >
                    {(task.response_contract.options ?? []).map((option) => {
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
                    })}
                  </RadioGroup>
                )}
              />
            </fieldset>
          ))}
          <Button
            type="submit"
            disabled={!allAnswered || locked || mediaBlocked}
          >
            {submitLabel}
          </Button>
        </form>
      </CardContent>
    </Card>
  );
}
