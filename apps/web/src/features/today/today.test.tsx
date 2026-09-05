import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import {
  cleanup,
  fireEvent,
  render,
  screen,
  waitFor,
} from "@testing-library/react";
import { afterEach, beforeEach, describe, expect, it, vi } from "vitest";

vi.stubGlobal(
  "ResizeObserver",
  class ResizeObserver {
    observe() {}
    unobserve() {}
    disconnect() {}
  },
);

const sdk = vi.hoisted(() => ({
  getDailyPlan: vi.fn(),
  listPracticeModes: vi.fn(),
  putTargetProfile: vi.fn(),
  createPracticeActivity: vi.fn(),
  getPracticeActivityMedia: vi.fn(),
  createAttempt: vi.fn(),
  submitAttempt: vi.fn(),
}));
const getToken = vi.hoisted(() => vi.fn(async () => "clerk-token"));
const apiClient = vi.hoisted(() => ({ marker: "generated-client" }));
const authState = vi.hoisted(() => ({ isLoaded: true, isSignedIn: true }));
const mediaUrl = vi.hoisted(() => ({
  createObjectURL: vi.fn(() => "blob:ilets-test"),
  revokeObjectURL: vi.fn(),
}));

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

import Today from "./today";

function renderSubject() {
  const queryClient = new QueryClient({
    defaultOptions: { queries: { retry: false }, mutations: { retry: false } },
  });
  return render(
    <QueryClientProvider client={queryClient}>
      <Today />
    </QueryClientProvider>,
  );
}

function profile(
  variant: "Academic" | "General Training" = "Academic",
  revision = 1,
) {
  return {
    state: "CONFIGURED" as const,
    profile: {
      test_variant: { state: "PRESENT" as const, value: variant },
      delivery_mode: { state: "UNKNOWN" as const },
      purpose_or_receiving_rule: { state: "UNKNOWN" as const },
      selected_skill_retake: { state: "UNKNOWN" as const },
      target_overall_band: 7,
      minimum_listening_band: 6.5,
      minimum_reading_band: 7.5,
      minimum_writing_band: 6,
      minimum_speaking_band: 6.5,
      resolution: { state: "RESOLVED" as const, unresolved_conditions: [] },
      resource_revision: revision,
      updated_at: "2026-08-29T00:00:00Z",
    },
  };
}

function plan(overrides: Record<string, unknown> = {}) {
  return {
    daily_plan_id: "plan_1",
    generated_at: "2026-08-29T00:00:00Z",
    target_context: profile(),
    unresolved_target_conditions: [],
    coverage_gaps: [],
    items: [
      {
        plan_item_id: "plan_item_1",
        practice_mode_id: "PM-R03",
        canonical_target_ids: ["R-QT-02", "R-QT-03"],
        reason_codes: ["INSUFFICIENT_EVIDENCE"],
        primary_activity_purpose: "ASSESSMENT",
        evidence_candidacy: "ASSESSMENT_MAY_ADMIT",
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
      },
    ],
    ...overrides,
  };
}

function assessmentActivity() {
  return {
    practice_activity_id: "activity_assessment",
    content_revision_id: "reading-sampled-at02-001-r1",
    practice_mode_id: "PM-R03",
    practice_type_ids: ["PT-13"],
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
    delivery_mode: { state: "PRESENT", value: "Test-centre computer" },
    primary_activity_purpose: "ASSESSMENT",
    evidence_candidacy: "ASSESSMENT_MAY_ADMIT",
    assistance_conditions: [
      { condition_id: "hints", state: "PRESENT", value: false },
    ],
    exposure_conditions: [
      {
        condition_id: "item_revision_seen_before",
        state: "PRESENT",
        value: false,
      },
    ],
    material: {
      stimuli: [
        {
          stimulus_id: "stimulus_1",
          kind: "TEXT",
          title: "Rooftop garden pilot",
          text: "Learner-safe passage",
        },
      ],
      tasks: [
        {
          task_id: "task_1",
          prompt: "The pilot used the roof.",
          response_contract: {
            kind: "SINGLE_SELECTION",
            options: [
              { value: "TRUE", label: "True" },
              { value: "FALSE", label: "False" },
            ],
          },
        },
      ],
    },
    assigned_at: "2026-08-29T00:00:00Z",
  };
}

function trainingActivity() {
  return {
    ...assessmentActivity(),
    practice_activity_id: "activity_training",
    content_revision_id: "reading-bootstrap-classification-001-r1",
    primary_activity_purpose: "TRAINING",
    evidence_candidacy: "NOT_EVIDENCE_CANDIDATE",
    assistance_conditions: [],
    exposure_conditions: [],
  };
}

function gistActivity(variant: "Academic" | "General Training" = "Academic") {
  return {
    ...trainingActivity(),
    practice_activity_id: `activity_gist_${variant}`,
    content_revision_id: "listening-bootstrap-gist-001-r1",
    practice_mode_id: "PM-L02",
    practice_type_ids: ["PT-12"],
    canonical_target_ids: ["L-COMP-01"],
    test_variant: { state: "PRESENT" as const, value: variant },
    content_context_ids: {
      state: "PRESENT" as const,
      values: ["CTX-LISTENING-SHARED"],
    },
    official_family_ids: {
      state: "PRESENT" as const,
      values: ["IELTS-L-QF-01"],
    },
    delivery_mode: {
      state: "NOT_APPLICABLE" as const,
      reason:
        "This bounded Listening activity has no delivery-mode-specific interaction.",
    },
    material: {
      stimuli: [
        {
          stimulus_id: "gist-stimulus",
          kind: "MEDIA_REFERENCE" as const,
          title: "Marsha introduction",
          media_reference: "hello-this-is-marsha",
        },
      ],
      tasks: [
        {
          task_id: "listening_gist_001",
          prompt: "What is the recording mainly about?",
          response_contract: {
            kind: "SINGLE_SELECTION" as const,
            options: [
              { value: "Introducing Marsha", label: "Introducing Marsha" },
              { value: "Ordering a meal", label: "Ordering a meal" },
              { value: "Making a travel plan", label: "Making a travel plan" },
            ],
          },
        },
      ],
    },
  };
}

function draftAttempt(
  activity: Pick<
    ReturnType<typeof assessmentActivity>,
    "practice_activity_id" | "content_revision_id"
  > = assessmentActivity(),
) {
  return {
    attempt_id: "attempt_1",
    practice_activity_id: activity.practice_activity_id,
    content_revision_id: activity.content_revision_id,
    status: "draft",
    resource_revision: 1,
    evaluation_ids: [],
    started_at: "2026-08-29T00:00:00Z",
  };
}

function submittedResult(
  activity: Pick<
    ReturnType<typeof assessmentActivity>,
    "practice_activity_id" | "content_revision_id"
  > = assessmentActivity(),
) {
  return {
    attempt: {
      ...draftAttempt(activity),
      status: "evaluated",
      resource_revision: 2,
      submitted_at: "2026-08-29T00:01:00Z",
      evaluated_at: "2026-08-29T00:01:00Z",
    },
    evaluation_state: { state: "NOT_REQUIRED" },
  };
}

describe("Today canonical consumer", () => {
  afterEach(cleanup);
  beforeEach(() => {
    Object.values(sdk).forEach((mock) => mock.mockReset());
    Object.defineProperty(URL, "createObjectURL", {
      configurable: true,
      value: mediaUrl.createObjectURL,
    });
    Object.defineProperty(URL, "revokeObjectURL", {
      configurable: true,
      value: mediaUrl.revokeObjectURL,
    });
    mediaUrl.createObjectURL.mockClear();
    mediaUrl.revokeObjectURL.mockClear();
    authState.isLoaded = true;
    authState.isSignedIn = true;
    sdk.getDailyPlan.mockResolvedValue({ data: plan() });
    sdk.listPracticeModes.mockResolvedValue({
      data: {
        modes: [
          {
            practice_mode_id: "PM-R03",
            label: "Reading classification",
            practice_type_ids: ["PT-13"],
            duration_label: "10 min",
          },
        ],
      },
    });
  });

  it("does not make durable learner calls while signed out", () => {
    authState.isSignedIn = false;
    renderSubject();
    expect(screen.getByText("signInRequired")).toBeVisible();
    expect(screen.getByRole("button", { name: "signIn" })).toBeVisible();
    expect(sdk.getDailyPlan).not.toHaveBeenCalled();
    expect(sdk.listPracticeModes).not.toHaveBeenCalled();
    expect(sdk.putTargetProfile).not.toHaveBeenCalled();
    expect(sdk.createPracticeActivity).not.toHaveBeenCalled();
  });

  it("uses DailyPlan as the root target and recommendation read authority", async () => {
    renderSubject();
    await waitFor(() =>
      expect(sdk.getDailyPlan).toHaveBeenCalledWith(
        expect.objectContaining({ client: apiClient }),
      ),
    );
    expect(await screen.findByLabelText("variant")).toHaveValue("Academic");
    expect(screen.getByLabelText("minimumReadingBand")).toHaveValue(7.5);
    expect(screen.getByText("Reading classification")).toBeVisible();
    expect(screen.getByText("why.insufficientEvidence")).toBeVisible();
    expect(screen.queryByText("INSUFFICIENT_EVIDENCE")).not.toBeInTheDocument();
    expect(screen.queryByText("PM-R03")).not.toBeInTheDocument();
    expect(screen.queryByText("R-QT-02")).not.toBeInTheDocument();
  });

  it("renders NOT_CONFIGURED as target setup without an Academic default", async () => {
    sdk.getDailyPlan.mockResolvedValue({
      data: plan({
        target_context: { state: "NOT_CONFIGURED" },
        items: [],
      }),
    });
    renderSubject();
    const variant = await screen.findByLabelText("variant");
    expect(variant).toHaveValue("");
    expect(screen.getByText("targetSetupHeading")).toBeVisible();
    expect(
      screen.queryByRole("heading", { name: "recommendationHeading" }),
    ).not.toBeInTheDocument();
    expect(screen.getByRole("button", { name: "startDirect" })).toBeDisabled();
  });

  it("keeps direct PM-R03 non-actionable for configured General Training", async () => {
    sdk.getDailyPlan.mockResolvedValue({
      data: plan({
        target_context: profile("General Training"),
        items: [],
      }),
    });
    renderSubject();
    await waitFor(() =>
      expect(screen.getByLabelText("variant")).toHaveValue("General Training"),
    );
    const direct = screen.getByRole("button", { name: "startDirect" });
    expect(direct).toBeDisabled();
    expect(
      screen.getByRole("button", { name: "startGistSprint" }),
    ).toBeEnabled();
    fireEvent.click(direct);
    expect(sdk.createPracticeActivity).not.toHaveBeenCalled();
  });

  it("keeps Gist Sprint truthful while assignment is pending", async () => {
    let resolveAssignment: (value: unknown) => void = () => {};
    sdk.createPracticeActivity.mockReturnValue(
      new Promise((resolve) => {
        resolveAssignment = resolve;
      }),
    );
    renderSubject();
    const gist = await screen.findByRole("button", {
      name: "startGistSprint",
    });
    await waitFor(() => expect(gist).toBeEnabled());
    fireEvent.click(gist);
    await waitFor(() => expect(gist).toBeDisabled());
    resolveAssignment({
      data: { outcome: "ASSIGNED", activity: gistActivity() },
    });
    expect(await screen.findByText("Marsha introduction")).toBeVisible();
  });

  it("renders unresolved target explanations without exposing condition IDs", async () => {
    sdk.getDailyPlan.mockResolvedValue({
      data: plan({
        items: [],
        unresolved_target_conditions: [
          {
            condition_id: "delivery_mode",
            explanation: "Choose the delivery mode required by your target.",
          },
        ],
      }),
    });
    renderSubject();
    expect(
      await screen.findByText(
        "Choose the delivery mode required by your target.",
      ),
    ).toBeVisible();
    expect(screen.queryByText("delivery_mode")).not.toBeInTheDocument();
  });

  it("saves TargetProfile with Expected-Resource-Revision then refetches DailyPlan", async () => {
    sdk.getDailyPlan.mockResolvedValue({ data: plan() });
    sdk.putTargetProfile.mockResolvedValue({
      data: profile("Academic", 2).profile,
    });
    renderSubject();
    const reading = await screen.findByLabelText("minimumReadingBand");
    await waitFor(() =>
      expect(screen.getByLabelText("variant")).toHaveValue("Academic"),
    );
    await waitFor(() =>
      expect(screen.getByRole("button", { name: "saveTarget" })).toBeEnabled(),
    );
    fireEvent.change(reading, { target: { value: "8" } });
    expect(reading).toHaveValue(8);
    fireEvent.click(screen.getByRole("button", { name: "saveTarget" }));
    await waitFor(() =>
      expect(sdk.putTargetProfile).toHaveBeenCalledWith(
        expect.objectContaining({
          client: apiClient,
          headers: { "Expected-Resource-Revision": 1 },
          body: expect.objectContaining({
            test_variant: "Academic",
            minimum_reading_band: 8,
          }),
        }),
      ),
    );
    await waitFor(() => expect(sdk.getDailyPlan).toHaveBeenCalledTimes(2));
  });

  it("starts a recommendation only through daily_plan_item_id and renders learner-safe AT-02 material", async () => {
    sdk.createPracticeActivity.mockResolvedValue({
      data: { outcome: "ASSIGNED", activity: assessmentActivity() },
    });
    renderSubject();
    const start = await screen.findByRole("button", {
      name: "startRecommendation",
    });
    fireEvent.click(start);
    await waitFor(() =>
      expect(sdk.createPracticeActivity).toHaveBeenCalledWith(
        expect.objectContaining({
          client: apiClient,
          body: { daily_plan_item_id: "plan_item_1" },
        }),
      ),
    );
    expect(await screen.findByText("Rooftop garden pilot")).toBeVisible();
    expect(screen.getByText("The pilot used the roof.")).toBeVisible();
    expect(screen.getByText("assessmentBoundary")).toBeVisible();
    expect(screen.queryByText("correct_choice")).not.toBeInTheDocument();
    expect(
      screen.queryByText("reading-sampled-at02-001-r1"),
    ).not.toBeInTheDocument();
  });

  it("treats UNAVAILABLE as a domain state and refreshes the plan", async () => {
    sdk.createPracticeActivity.mockResolvedValue({
      data: {
        outcome: "UNAVAILABLE",
        unavailability: {
          reason: "CURRENT_ELIGIBILITY_BLOCKED",
          unresolved_target_conditions: [],
          coverage_gaps: [],
          explanation:
            "This recommendation is no longer available. Today was refreshed.",
        },
      },
    });
    renderSubject();
    fireEvent.click(
      await screen.findByRole("button", { name: "startRecommendation" }),
    );
    expect(
      await screen.findByText(
        "This recommendation is no longer available. Today was refreshed.",
      ),
    ).toBeVisible();
    expect(screen.getByTestId("unavailable")).toBeVisible();
    await waitFor(() => expect(sdk.getDailyPlan).toHaveBeenCalledTimes(2));
  });

  it("surfaces a direct Gist Sprint UNAVAILABLE result and refreshes Today", async () => {
    sdk.getDailyPlan.mockResolvedValue({
      data: plan({ items: [] }),
    });
    sdk.createPracticeActivity.mockResolvedValue({
      data: {
        outcome: "UNAVAILABLE",
        unavailability: {
          reason: "CURRENT_ELIGIBILITY_BLOCKED",
          unresolved_target_conditions: [],
          coverage_gaps: [],
          explanation:
            "Gist Sprint is not currently available. Today was refreshed.",
        },
      },
    });
    renderSubject();
    const gist = await screen.findByRole("button", {
      name: "startGistSprint",
    });
    await waitFor(() => expect(gist).toBeEnabled());
    fireEvent.click(gist);
    await waitFor(() =>
      expect(sdk.createPracticeActivity).toHaveBeenCalledWith(
        expect.objectContaining({
          client: apiClient,
          body: { practice_mode_id: "PM-L02" },
        }),
      ),
    );
    expect(
      await screen.findByText(
        "Gist Sprint is not currently available. Today was refreshed.",
      ),
    ).toBeVisible();
    expect(screen.getByTestId("unavailable")).toBeVisible();
    await waitFor(() => expect(sdk.getDailyPlan).toHaveBeenCalledTimes(2));
  });

  it("preserves assigned conditions, uses bounded assessment copy, and refreshes post-submission plan", async () => {
    const activity = assessmentActivity();
    const blocker =
      "A prior assignment exists for the only bounded Reading assessment sample; actual learner exposure is not established, so fresh/unseen eligibility can no longer be proven and no new fresh-independent opportunity is issued.";
    sdk.getDailyPlan.mockResolvedValueOnce({ data: plan() }).mockResolvedValue({
      data: plan({
        items: [],
        coverage_gaps: [
          {
            gap_class: "CONTENT_OR_ASSET",
            scoped_target_ids: ["R-QT-02", "R-QT-03"],
            condition_id: "content_assets",
            condition_status: "BLOCKED",
            blocking_consequence: blocker,
            dependencies: ["fresh eligible sampled Reading assessment content"],
            demand_class: "content/assets/supply route",
            provenance_version: "planner-v1",
          },
        ],
      }),
    });
    sdk.createPracticeActivity.mockResolvedValue({
      data: { outcome: "ASSIGNED", activity },
    });
    sdk.createAttempt.mockResolvedValue({ data: draftAttempt(activity) });
    sdk.submitAttempt.mockResolvedValue({ data: submittedResult(activity) });

    renderSubject();
    fireEvent.click(
      await screen.findByRole("button", { name: "startRecommendation" }),
    );
    await screen.findByText("The pilot used the roof.");
    fireEvent.click(screen.getByRole("radio", { name: "True" }));
    fireEvent.click(screen.getByRole("button", { name: "submitAnswers" }));

    await waitFor(() =>
      expect(sdk.submitAttempt).toHaveBeenCalledWith(
        expect.objectContaining({
          body: expect.objectContaining({
            actual_conditions: {
              delivery: activity.delivery_mode,
              assistance: activity.assistance_conditions,
              exposure: activity.exposure_conditions,
              input: [],
              timing: [],
            },
          }),
        }),
      ),
    );
    expect(await screen.findByText("assessmentCompleted")).toBeVisible();
    expect(screen.queryByText(/evidence admitted/i)).not.toBeInTheDocument();
    expect(screen.queryByText(/Band improved/i)).not.toBeInTheDocument();
    await waitFor(() => expect(sdk.getDailyPlan).toHaveBeenCalledTimes(2));
    expect(await screen.findByText(blocker)).toBeVisible();
    expect(screen.queryByText("content_assets")).not.toBeInTheDocument();
    expect(screen.queryByText("CONTENT_OR_ASSET")).not.toBeInTheDocument();
    expect(screen.queryByText(/readiness/i)).not.toBeInTheDocument();
  });

  it("keeps direct PM-R03 training secondary and non-evidence", async () => {
    sdk.createPracticeActivity.mockResolvedValue({
      data: { outcome: "ASSIGNED", activity: trainingActivity() },
    });
    renderSubject();
    expect(
      await screen.findByRole("heading", { name: "recommendationHeading" }),
    ).toBeVisible();
    expect(
      screen.getByRole("heading", { name: "directHeading" }),
    ).toBeVisible();
    const direct = screen.getByRole("button", { name: "startDirect" });
    expect(direct).toBeEnabled();
    fireEvent.click(direct);
    await waitFor(() =>
      expect(sdk.createPracticeActivity).toHaveBeenCalledWith(
        expect.objectContaining({ body: { practice_mode_id: "PM-R03" } }),
      ),
    );
    expect(await screen.findByText("trainingBoundary")).toBeVisible();
  });

  it.each(["Academic", "General Training"] as const)(
    "starts PM-L02 Gist Sprint for %s and submits observation-only training",
    async (variant) => {
      const activity = gistActivity(variant);
      sdk.getDailyPlan.mockResolvedValue({
        data: plan({ target_context: profile(variant), items: [] }),
      });
      sdk.createPracticeActivity.mockResolvedValue({
        data: { outcome: "ASSIGNED", activity },
      });
      sdk.getPracticeActivityMedia.mockResolvedValue({
        data: new Blob(["audio"], { type: "audio/ogg" }),
      });
      sdk.createAttempt.mockResolvedValue({ data: draftAttempt(activity) });
      sdk.submitAttempt.mockResolvedValue({ data: submittedResult(activity) });

      renderSubject();
      await waitFor(() =>
        expect(screen.getByLabelText("variant")).toHaveValue(variant),
      );
      fireEvent.click(screen.getByRole("button", { name: "startGistSprint" }));

      await waitFor(() =>
        expect(sdk.createPracticeActivity).toHaveBeenCalledWith(
          expect.objectContaining({
            client: apiClient,
            body: { practice_mode_id: "PM-L02" },
          }),
        ),
      );
      expect(await screen.findByTestId("activity-audio")).toBeVisible();
      expect(sdk.getPracticeActivityMedia).toHaveBeenCalledWith({
        client: apiClient,
        path: {
          practice_activity_id: activity.practice_activity_id,
          media_reference: "hello-this-is-marsha",
        },
      });

      fireEvent.click(
        screen.getByRole("radio", { name: "Introducing Marsha" }),
      );
      fireEvent.click(screen.getByRole("button", { name: "submitAnswers" }));
      await waitFor(() => expect(sdk.submitAttempt).toHaveBeenCalled());
      expect(await screen.findByText("trainingCompleted")).toBeVisible();
      expect(screen.queryByText("assessmentCompleted")).not.toBeInTheDocument();
    },
  );

  it("blocks Gist Sprint submission when authorized audio fails and releases audio URLs on unmount", async () => {
    const activity = gistActivity();
    sdk.createPracticeActivity.mockResolvedValue({
      data: { outcome: "ASSIGNED", activity },
    });
    sdk.getPracticeActivityMedia.mockResolvedValue({
      data: new Blob(["audio"], { type: "audio/ogg" }),
    });

    const rendered = renderSubject();
    const startGist = await screen.findByRole("button", {
      name: "startGistSprint",
    });
    await waitFor(() => expect(startGist).toBeEnabled());
    fireEvent.click(startGist);
    expect(await screen.findByTestId("activity-audio")).toBeVisible();
    fireEvent.click(screen.getByRole("radio", { name: "Introducing Marsha" }));
    rendered.unmount();
    expect(mediaUrl.revokeObjectURL).toHaveBeenCalledWith("blob:ilets-test");

    sdk.getPracticeActivityMedia.mockRejectedValueOnce(
      new Error("media unavailable"),
    );
    renderSubject();
    const retryGist = await screen.findByRole("button", {
      name: "startGistSprint",
    });
    await waitFor(() => expect(retryGist).toBeEnabled());
    fireEvent.click(retryGist);
    expect(await screen.findByRole("alert")).toHaveTextContent("mediaError");
    fireEvent.click(screen.getByRole("radio", { name: "Introducing Marsha" }));
    expect(
      screen.getByRole("button", { name: "submitAnswers" }),
    ).toBeDisabled();
  });
});
