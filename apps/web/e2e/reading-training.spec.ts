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
  await page.getByLabel("Minimum Reading Band").fill("6.5");
  await page.getByRole("button", { name: "Save target" }).click();
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
