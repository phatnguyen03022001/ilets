"use client";

import { Controller, type UseFormReturn } from "react-hook-form";
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
  onSubmit: (values: AnswerForm) => void | Promise<void>;
};

export default function ActivityWorkspace({
  activity,
  attempt,
  form,
  isPending,
  submitLabel,
  fallbackTitle,
  onSubmit,
}: Props) {
  const answers = form.watch("answers");
  const allAnswered = activity.material.tasks.every((task) =>
    Boolean(answers[task.task_id]),
  );
  const locked = isPending || attempt?.status === "evaluated";

  return (
    <Card aria-labelledby="activity-heading">
      <CardHeader>
        <CardTitle id="activity-heading">
          {activity.material.stimuli[0]?.title ?? fallbackTitle}
        </CardTitle>
        {activity.material.stimuli[0]?.text && (
          <CardDescription className="whitespace-pre-wrap text-base leading-7 text-foreground">
            {activity.material.stimuli[0].text}
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
          <Button type="submit" disabled={!allAnswered || locked}>
            {submitLabel}
          </Button>
        </form>
      </CardContent>
    </Card>
  );
}
