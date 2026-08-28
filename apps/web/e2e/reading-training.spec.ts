import AxeBuilder from "@axe-core/playwright";
import { expect, test } from "@playwright/test";

test("Reading training remains non-evidence and does not leak answer keys", async ({
  page,
}) => {
  let activityPayload = "";
  page.on("response", async (response) => {
    if (
      response.url().includes("/v1/practice-activities") &&
      response.request().method() === "POST" &&
      response.ok()
    ) {
      activityPayload = await response.text();
    }
  });

  await page.goto("/");
  const saveTarget = page.getByRole("button", { name: "Save target" });
  await expect(saveTarget).toBeEnabled();
  await page.getByLabel("Variant").selectOption("ACADEMIC");
  await page.getByLabel("Minimum Reading Band").fill("6.5");
  await saveTarget.click();
  await expect(page.getByTestId("target-saved")).toBeVisible();

  await page.getByRole("button", { name: "Start activity" }).click();
  await expect(
    page.getByRole("heading", { name: "Rooftop garden pilot" }),
  ).toBeVisible();
  expect(activityPayload).not.toContain("correct_choice");
  expect(activityPayload).not.toContain("explanation");

  const items = page.locator("[data-testid^='item-']");
  const count = await items.count();
  expect(count).toBe(6);
  for (let i = 0; i < count; i++) {
    await items.nth(i).getByRole("radio").first().click();
  }
  await page.getByRole("button", { name: "Submit answers" }).click();
  await expect(page.getByTestId("result")).toBeVisible();
  await expect(
    page.getByText("Training observation only — NOT_EVIDENCE_CANDIDATE."),
  ).toBeVisible();
  await expect(page.getByText(/Correct answer:/).first()).toBeVisible();

  const accessibility = await new AxeBuilder({ page }).analyze();
  const severe = accessibility.violations.filter(
    (violation) =>
      violation.impact === "critical" || violation.impact === "serious",
  );
  expect(severe).toEqual([]);
});

test("Academic Reading assessment admits sampled evidence without Band claims", async ({
  page,
}) => {
  let assessmentPayload = "";
  page.on("response", async (response) => {
    if (
      response.url().includes("/v1/assessment-activities") &&
      response.request().method() === "POST" &&
      response.ok()
    ) {
      assessmentPayload = await response.text();
    }
  });

  await page.goto("/");
  const saveTarget = page.getByRole("button", { name: "Save target" });
  await expect(saveTarget).toBeEnabled();
  await page.getByLabel("Variant").selectOption("ACADEMIC");
  await page.getByLabel("Minimum Reading Band").fill("6.5");
  await saveTarget.click();
  await expect(page.getByTestId("target-saved")).toBeVisible();

  await page.getByRole("button", { name: "Start assessment" }).click();
  await expect(
    page.getByRole("heading", { name: "Local history archive" }),
  ).toBeVisible();
  expect(assessmentPayload).not.toContain("correct_choice");
  expect(assessmentPayload).not.toContain("explanation");

  const items = page.locator("[data-testid^='item-']");
  await expect(items).toHaveCount(6);
  for (let i = 0; i < 6; i++) {
    await items.nth(i).getByRole("radio").first().click();
  }
  await page.getByRole("button", { name: "Submit answers" }).click();

  await expect(page.getByTestId("result")).toBeVisible();
  await expect(
    page.getByText(
      "Evidence admitted only for this sampled Reading classification performance. This does not establish a Reading Band, readiness, mastery, or certification.",
    ),
  ).toBeVisible();
  await expect(
    page.getByRole("button", { name: "Practice again" }),
  ).toHaveCount(0);
});
