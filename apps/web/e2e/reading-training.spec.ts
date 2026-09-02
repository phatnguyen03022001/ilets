import AxeBuilder from "@axe-core/playwright";
import { expect, test } from "@playwright/test";

test("Today consumes fresh sampled AT-02 supply before showing the authoritative blocker", async ({
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

  await expect(
    page.getByRole("heading", { name: "Recommended next" }),
  ).toBeVisible();
  await expect(page.getByText("Headings & Structure")).toBeVisible();
  await expect(page.getByText("content_assets")).toHaveCount(0);
  await expect(page.getByText("CONTENT_OR_ASSET", { exact: true })).toHaveCount(
    0,
  );

  const postFirstSubmissionPlan = [...planResponses]
    .reverse()
    .find(
      (plan) =>
        plan.items?.length === 1 &&
        plan.items[0]?.practice_mode_id === "PM-R04",
    );
  expect(postFirstSubmissionPlan).toBeTruthy();
  expect(postFirstSubmissionPlan?.items[0]).toEqual(
    expect.objectContaining({
      practice_mode_id: "PM-R04",
      canonical_target_ids: ["R-QT-01"],
    }),
  );
  expect(postFirstSubmissionPlan?.coverage_gaps).toEqual([]);

  await page
    .getByRole("button", { name: "Start recommended activity" })
    .click();
  await expect(
    page.getByRole("heading", { name: "Cooling the block" }),
  ).toBeVisible();

  const headingsAssignment = activityResponses[1];
  expect(headingsAssignment.outcome).toBe("ASSIGNED");
  expect(headingsAssignment.activity.content_revision_id).toBe(
    "reading-bootstrap-assessment-002-r1",
  );
  expect(headingsAssignment.activity.practice_mode_id).toBe("PM-R04");
  expect(headingsAssignment.activity.canonical_target_ids).toEqual(["R-QT-01"]);
  expect(headingsAssignment.activity.official_family_ids).toEqual({
    state: "PRESENT",
    values: ["IELTS-R-QF-05"],
  });
  expect(JSON.stringify(headingsAssignment)).not.toContain("correct_choice");
  expect(JSON.stringify(headingsAssignment)).not.toContain("explanation");

  const headingItems = page.locator("[data-testid^='item-']");
  await expect(headingItems).toHaveCount(2);
  await headingItems
    .nth(0)
    .getByRole("radio", {
      name: "Cooling roofs without rebuilding",
      exact: true,
    })
    .click();
  await headingItems
    .nth(1)
    .getByRole("radio", { name: "Shade where passengers wait", exact: true })
    .click();
  await page.getByRole("button", { name: "Submit answers" }).click();
  await expect(page.getByTestId("result")).toBeVisible();

  await expect
    .poll(() =>
      [...planResponses]
        .reverse()
        .find((plan) =>
          plan.coverage_gaps?.some(
            (gap: Record<string, any>) =>
              gap.gap_class === "CONTENT_OR_ASSET" &&
              gap.condition_id === "content_assets" &&
              gap.condition_status === "BLOCKED",
          ),
        ),
    )
    .toBeTruthy();
  const exhaustedPlan = [...planResponses]
    .reverse()
    .find((plan) =>
      plan.coverage_gaps?.some(
        (gap: Record<string, any>) =>
          gap.gap_class === "CONTENT_OR_ASSET" &&
          gap.condition_id === "content_assets" &&
          gap.condition_status === "BLOCKED",
      ),
    );
  expect(exhaustedPlan?.items).toHaveLength(0);
  const supplyBlocker = exhaustedPlan?.coverage_gaps.find(
    (gap: Record<string, any>) =>
      gap.gap_class === "CONTENT_OR_ASSET" &&
      gap.condition_id === "content_assets" &&
      gap.condition_status === "BLOCKED",
  );
  expect(supplyBlocker).toEqual(
    expect.objectContaining({
      gap_class: "CONTENT_OR_ASSET",
      condition_id: "content_assets",
      condition_status: "BLOCKED",
      blocking_consequence: expect.any(String),
      dependencies: ["fresh eligible sampled Reading assessment content"],
      demand_class: "content/assets/supply route",
    }),
  );
  expect(supplyBlocker.blocking_consequence).toMatch(/\S/);
  await expect(
    page.getByText(supplyBlocker.blocking_consequence),
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
  expect(activityRequests[2]).toEqual({ practice_mode_id: "PM-R03" });
  expect(activityResponses[2].activity.primary_activity_purpose).toBe(
    "TRAINING",
  );
  expect(activityResponses[2].activity.evidence_candidacy).toBe(
    "NOT_EVIDENCE_CANDIDATE",
  );

  const accessibility = await new AxeBuilder({ page }).analyze();
  const severe = accessibility.violations.filter(
    (violation) =>
      violation.impact === "critical" || violation.impact === "serious",
  );
  expect(severe).toEqual([]);
});
