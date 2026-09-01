import AxeBuilder from "@axe-core/playwright";
import { expect, test } from "@playwright/test";

test("Today consumes the real Core plan through sampled AT-02 then shows the authoritative blocker", async ({
  page,
}) => {
  const activityRequests: Array<Record<string, unknown>> = [];
  const activityResponses: Array<Record<string, any>> = [];
  const submissionRequests: Array<Record<string, any>> = [];
  const planResponses: Array<Record<string, any>> = [];

  page.on("request", (request) => {
    const url = request.url();
    if (
      url.includes("/v1/practice-activities") &&
      request.method() === "POST"
    ) {
      activityRequests.push(request.postDataJSON());
    }
    if (url.includes("/submissions") && request.method() === "POST") {
      submissionRequests.push(request.postDataJSON());
    }
  });
  page.on("response", async (response) => {
    if (
      response.url().includes("/v1/practice-activities") &&
      response.request().method() === "POST" &&
      response.ok()
    ) {
      activityResponses.push(await response.json());
    }
    if (
      response.url().includes("/v1/daily-plan") &&
      response.request().method() === "GET" &&
      response.ok()
    ) {
      planResponses.push(await response.json());
    }
  });

  await page.goto("/");
  await expect(
    page.getByRole("heading", { name: "Set your target" }),
  ).toBeVisible();
  await expect(page.getByLabel("Variant")).toHaveValue("");

  const saveTarget = page.getByRole("button", { name: "Save target" });
  await expect(saveTarget).toBeEnabled();
  await page.getByLabel("Variant").selectOption("Academic");
  await page.getByLabel("Minimum Reading Band").fill("6.5");
  await saveTarget.click();

  await expect(
    page.getByRole("heading", { name: "Recommended next" }),
  ).toBeVisible();
  await expect(page.getByText("T/F/NG + Y/N/NG")).toBeVisible();
  await expect(
    page.getByText(
      "We need more independent evidence for this part of Reading.",
    ),
  ).toBeVisible();
  await expect(page.getByText("INSUFFICIENT_EVIDENCE")).toHaveCount(0);
  await expect(page.getByText("PM-R03")).toHaveCount(0);

  await page
    .getByRole("button", { name: "Start recommended activity" })
    .click();
  await expect(
    page.getByRole("heading", { name: "Local history archive" }),
  ).toBeVisible();

  expect(activityRequests[0]).toEqual(
    expect.objectContaining({ daily_plan_item_id: expect.any(String) }),
  );
  expect(activityRequests[0]).not.toHaveProperty("practice_mode_id");
  const assessmentAssignment = activityResponses[0];
  expect(assessmentAssignment.outcome).toBe("ASSIGNED");
  expect(assessmentAssignment.activity.primary_activity_purpose).toBe(
    "ASSESSMENT",
  );
  expect(assessmentAssignment.activity.evidence_candidacy).toBe(
    "ASSESSMENT_MAY_ADMIT",
  );
  expect(JSON.stringify(assessmentAssignment)).not.toContain("correct_choice");
  expect(JSON.stringify(assessmentAssignment)).not.toContain("explanation");
  expect(JSON.stringify(assessmentAssignment)).not.toContain(
    "validation_policy",
  );

  const choices = ["TRUE", "FALSE", "NOT_GIVEN", "YES", "NO", "NOT_GIVEN"];
  const items = page.locator("[data-testid^='item-']");
  await expect(items).toHaveCount(6);
  for (let index = 0; index < choices.length; index += 1) {
    await items
      .nth(index)
      .getByRole("radio", { name: choices[index], exact: true })
      .click();
  }

  await page.getByRole("button", { name: "Submit answers" }).click();
  await expect(page.getByTestId("result")).toBeVisible();
  await expect(
    page.getByText(
      "Assessment submitted. This sampled activity does not by itself establish a Reading Band or readiness.",
    ),
  ).toBeVisible();
  await expect(page.getByText(/evidence admitted/i)).toHaveCount(0);
  await expect(page.getByText(/Band improved/i)).toHaveCount(0);
  await expect(page.getByText(/^ready$/i)).toHaveCount(0);
  await expect(page.getByText(/^passed$/i)).toHaveCount(0);

  expect(submissionRequests).toHaveLength(1);
  expect(submissionRequests[0].actual_conditions).toEqual({
    delivery: assessmentAssignment.activity.delivery_mode,
    assistance: assessmentAssignment.activity.assistance_conditions,
    exposure: assessmentAssignment.activity.exposure_conditions,
    input: [],
    timing: [],
  });

  await expect
    .poll(() =>
      [...planResponses]
        .reverse()
        .find((plan) =>
          plan.coverage_gaps?.some(
            (gap: Record<string, any>) =>
              gap.gap_class === "TRANSITION" &&
              gap.condition_status === "BLOCKED",
          ),
        ),
    )
    .toBeTruthy();
  const postSubmissionPlan = [...planResponses]
    .reverse()
    .find((plan) =>
      plan.coverage_gaps?.some(
        (gap: Record<string, any>) =>
          gap.gap_class === "TRANSITION" && gap.condition_status === "BLOCKED",
      ),
    );
  expect(postSubmissionPlan?.items).toHaveLength(0);
  const transitionBlocker = postSubmissionPlan?.coverage_gaps.find(
    (gap: Record<string, any>) =>
      gap.gap_class === "TRANSITION" && gap.condition_status === "BLOCKED",
  );
  expect(transitionBlocker).toEqual(
    expect.objectContaining({ blocking_consequence: expect.any(String) }),
  );
  await expect(
    page.getByText(transitionBlocker.blocking_consequence),
  ).toBeVisible();
  await expect(
    page.getByRole("heading", { name: "Recommended next" }),
  ).toHaveCount(0);
  await expect(page.getByText("progression_transition")).toHaveCount(0);
  await expect(page.getByText("TRANSITION", { exact: true })).toHaveCount(0);

  await expect(
    page.getByRole("heading", { name: "Extra practice" }),
  ).toBeVisible();
  await page.getByRole("button", { name: "Practice directly" }).click();
  await expect(
    page.getByText(
      "This is training. It does not count as Reading Band evidence.",
    ),
  ).toBeVisible();
  expect(activityRequests[1]).toEqual({ practice_mode_id: "PM-R03" });
  expect(activityResponses[1].activity.primary_activity_purpose).toBe(
    "TRAINING",
  );
  expect(activityResponses[1].activity.evidence_candidacy).toBe(
    "NOT_EVIDENCE_CANDIDATE",
  );

  const accessibility = await new AxeBuilder({ page }).analyze();
  const severe = accessibility.violations.filter(
    (violation) =>
      violation.impact === "critical" || violation.impact === "serious",
  );
  expect(severe).toEqual([]);
});
