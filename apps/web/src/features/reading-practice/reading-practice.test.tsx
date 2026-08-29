import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import {
  cleanup,
  fireEvent,
  render,
  screen,
  waitFor,
} from "@testing-library/react";
import { afterEach, beforeEach, describe, expect, it, vi } from "vitest";

const sdk = vi.hoisted(() => ({
  getTargetProfile: vi.fn(),
  putTargetProfile: vi.fn(),
  createPracticeActivity: vi.fn(),
  createAttempt: vi.fn(),
  submitAttempt: vi.fn(),
}));
const getToken = vi.hoisted(() => vi.fn(async () => "clerk-token"));
const apiClient = vi.hoisted(() => ({ marker: "generated-client" }));
const authState = vi.hoisted(() => ({ isLoaded: true, isSignedIn: true }));

vi.mock("@clerk/nextjs", () => ({
  useAuth: () => ({
    getToken,
    isLoaded: authState.isLoaded,
    isSignedIn: authState.isSignedIn,
  }),
  SignInButton: ({ children }: { children: React.ReactNode }) => (
    <>{children}</>
  ),
}));
vi.mock("next-intl", () => ({
  useTranslations: () => (key: string, values?: Record<string, unknown>) =>
    values ? `${key}:${JSON.stringify(values)}` : key,
}));
vi.mock("@/lib/api", () => ({ createPublicApi: () => apiClient }));
vi.mock("@/generated/public", async (importOriginal) => ({
  ...(await importOriginal<typeof import("@/generated/public")>()),
  ...sdk,
}));

import ReadingPractice from "./reading-practice";

function renderSubject() {
  const queryClient = new QueryClient({
    defaultOptions: { queries: { retry: false }, mutations: { retry: false } },
  });
  return render(
    <QueryClientProvider client={queryClient}>
      <ReadingPractice />
    </QueryClientProvider>,
  );
}

function profile(variant: "Academic" | "General Training", revision = 1) {
  return {
    state: "CONFIGURED",
    profile: {
      test_variant: { state: "PRESENT", value: variant },
      delivery_mode: { state: "UNKNOWN" },
      purpose_or_receiving_rule: { state: "UNKNOWN" },
      selected_skill_retake: { state: "UNKNOWN" },
      target_overall_band: 7,
      minimum_listening_band: 6.5,
      minimum_reading_band: 7.5,
      minimum_writing_band: 6,
      minimum_speaking_band: 6.5,
      resolution: { state: "RESOLVED", unresolved_conditions: [] },
      resource_revision: revision,
      updated_at: "2026-08-29T00:00:00Z",
    },
  };
}

describe("ReadingPractice canonical consumer", () => {
  afterEach(cleanup);
  beforeEach(() => {
    Object.values(sdk).forEach((mock) => mock.mockReset());
    authState.isLoaded = true;
    authState.isSignedIn = true;
    sdk.getTargetProfile.mockResolvedValue({ data: profile("Academic") });
  });

  it("offers Clerk sign-in when the learner is signed out", () => {
    authState.isSignedIn = false;
    renderSubject();
    expect(screen.getByText("signInRequired")).toBeVisible();
    expect(screen.getByRole("button", { name: "signIn" })).toBeVisible();
  });

  it("hydrates the canonical TargetProfile projection", async () => {
    sdk.getTargetProfile.mockResolvedValue({
      data: profile("General Training", 3),
    });
    renderSubject();
    await waitFor(() =>
      expect(screen.getByLabelText("variant")).toHaveValue("General Training"),
    );
    expect(screen.getByLabelText("minimumReadingBand")).toHaveValue(7.5);
    expect(
      screen.getByRole("button", { name: "startActivity" }),
    ).toBeDisabled();
    expect(screen.getByText("academicPracticeOnly")).toBeVisible();
  });

  it("uses the generated PUT operation and revision header", async () => {
    sdk.getTargetProfile.mockResolvedValue({ data: profile("Academic", 4) });
    sdk.putTargetProfile.mockResolvedValue({
      data: profile("Academic", 5).profile,
    });
    renderSubject();
    const reading = await screen.findByLabelText("minimumReadingBand");
    await waitFor(() => expect(reading).toHaveValue(7.5));
    fireEvent.change(reading, { target: { value: "8" } });
    fireEvent.click(screen.getByRole("button", { name: "saveTarget" }));
    await waitFor(() =>
      expect(sdk.putTargetProfile).toHaveBeenCalledWith(
        expect.objectContaining({
          client: apiClient,
          headers: { "Expected-Resource-Revision": 4 },
          body: expect.objectContaining({
            test_variant: "Academic",
            minimum_reading_band: 8,
          }),
        }),
      ),
    );
  });

  it("starts only the canonical practice activity and has no stale assessment UI", async () => {
    sdk.createPracticeActivity.mockResolvedValue({
      data: {
        outcome: "ASSIGNED",
        activity: {
          practice_activity_id: "activity_1",
          content_revision_id: "reading-bootstrap-classification-001-r1",
          practice_mode_id: "PM-R03",
          practice_type_ids: ["PT-13", "PT-16"],
          canonical_target_ids: ["R-QT-02", "R-QT-03"],
          test_variant: { state: "PRESENT", value: "Academic" },
          content_context_ids: {
            state: "PRESENT",
            values: ["CTX-READING-ACADEMIC"],
          },
          official_family_ids: {
            state: "PRESENT",
            values: ["IELTS-R-QF-02", "IELTS-R-QF-03"],
          },
          presentation_class_ids: { state: "NOT_APPLICABLE", reason: "none" },
          delivery_mode: { state: "NOT_APPLICABLE", reason: "none" },
          primary_activity_purpose: "TRAINING",
          evidence_candidacy: "NOT_EVIDENCE_CANDIDATE",
          assistance_conditions: [],
          exposure_conditions: [],
          material: {
            stimuli: [
              {
                stimulus_id: "s1",
                kind: "TEXT",
                title: "Passage",
                text: "Text",
              },
            ],
            tasks: [],
          },
          assigned_at: "2026-08-29T00:00:00Z",
        },
      },
    });
    renderSubject();
    const start = screen.getByRole("button", { name: "startActivity" });
    await waitFor(() => expect(start).toBeEnabled());
    fireEvent.click(start);
    await waitFor(() =>
      expect(sdk.createPracticeActivity).toHaveBeenCalledWith(
        expect.objectContaining({
          client: apiClient,
          body: { practice_mode_id: "PM-R03" },
          headers: expect.objectContaining({
            "Idempotency-Key": expect.any(String),
          }),
        }),
      ),
    );
    expect(screen.queryByText("startAssessment")).not.toBeInTheDocument();
    expect(screen.queryByText("assessmentHeading")).not.toBeInTheDocument();
  });
});
