import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import {
  cleanup,
  fireEvent,
  render,
  screen,
  waitFor,
} from "@testing-library/react";
import { afterEach, beforeEach, describe, expect, it, vi } from "vitest";

const apiMocks = vi.hoisted(() => ({
  get: vi.fn(),
  post: vi.fn(),
  put: vi.fn(),
}));

vi.mock("next-intl", () => ({
  useTranslations: () => (key: string) => key,
}));

vi.mock("@/lib/api", () => ({
  api: {
    GET: apiMocks.get,
    POST: apiMocks.post,
    PUT: apiMocks.put,
  },
}));

import ReadingPractice from "./reading-practice";

function renderReadingPractice() {
  const queryClient = new QueryClient({
    defaultOptions: { queries: { retry: false } },
  });

  return render(
    <QueryClientProvider client={queryClient}>
      <ReadingPractice />
    </QueryClientProvider>,
  );
}

describe("ReadingPractice target profile", () => {
  afterEach(() => cleanup());

  beforeEach(() => {
    apiMocks.get.mockReset();
    apiMocks.post.mockReset();
    apiMocks.put.mockReset();
    apiMocks.post.mockImplementation(async (path: string) => {
      if (path === "/v1/session") {
        return { data: { learner_id: "learner_test", human_actor: "Learner" } };
      }
      throw new Error(`unexpected POST ${path}`);
    });
  });

  it("hydrates all persisted TargetProfile Band constraints", async () => {
    apiMocks.get.mockResolvedValue({
      data: {
        test_variant: "GENERAL_TRAINING",
        target_overall_band: 7,
        minimum_listening_band: 6.5,
        minimum_reading_band: 7.5,
        minimum_writing_band: 6,
        minimum_speaking_band: 6.5,
        resource_revision: 3,
        updated_at: "2026-08-28T00:00:00Z",
      },
      response: { status: 200 },
    });

    renderReadingPractice();

    await waitFor(() => {
      expect(screen.getByLabelText("variant")).toHaveValue("GENERAL_TRAINING");
      expect(screen.getByLabelText("targetOverallBand")).toHaveValue(7);
      expect(screen.getByLabelText("minimumListeningBand")).toHaveValue(6.5);
      expect(screen.getByLabelText("minimumReadingBand")).toHaveValue(7.5);
      expect(screen.getByLabelText("minimumWritingBand")).toHaveValue(6);
      expect(screen.getByLabelText("minimumSpeakingBand")).toHaveValue(6.5);
    });
  });

  it("blocks the Academic-only practice when the persisted target is General Training", async () => {
    apiMocks.get.mockResolvedValue({
      data: {
        test_variant: "GENERAL_TRAINING",
        minimum_reading_band: 7,
        resource_revision: 2,
        updated_at: "2026-08-28T00:00:00Z",
      },
      response: { status: 200 },
    });

    renderReadingPractice();

    const button = screen.getByRole("button", { name: "startActivity" });
    await waitFor(() =>
      expect(screen.getByText("academicPracticeOnly")).toBeVisible(),
    );
    expect(button).toBeDisabled();
  });

  it("preserves existing target constraints when Reading Band changes", async () => {
    const existingTarget = {
      test_variant: "GENERAL_TRAINING" as const,
      target_overall_band: 7,
      minimum_listening_band: 6.5,
      minimum_reading_band: 7.5,
      minimum_writing_band: 6,
      minimum_speaking_band: 6.5,
      resource_revision: 4,
      updated_at: "2026-08-28T00:00:00Z",
    };
    apiMocks.get.mockResolvedValue({
      data: existingTarget,
      response: { status: 200 },
    });
    apiMocks.put.mockResolvedValue({
      data: {
        ...existingTarget,
        minimum_reading_band: 8,
        resource_revision: 5,
      },
      response: { status: 200 },
    });

    renderReadingPractice();

    const input = screen.getByLabelText("minimumReadingBand");
    await waitFor(() => expect(input).toHaveValue(7.5));
    fireEvent.change(input, { target: { value: "8" } });
    fireEvent.click(screen.getByRole("button", { name: "saveTarget" }));

    await waitFor(() =>
      expect(apiMocks.put).toHaveBeenCalledWith("/v1/target-profile", {
        body: {
          test_variant: "GENERAL_TRAINING",
          target_overall_band: 7,
          minimum_listening_band: 6.5,
          minimum_reading_band: 8,
          minimum_writing_band: 6,
          minimum_speaking_band: 6.5,
          expected_resource_revision: 4,
        },
      }),
    );
  });

  it("allows editing Listening, Writing, and Speaking minima while preserving the rest", async () => {
    const existingTarget = {
      test_variant: "ACADEMIC" as const,
      target_overall_band: 7,
      minimum_listening_band: 6,
      minimum_reading_band: 7.5,
      minimum_writing_band: 6,
      minimum_speaking_band: 6,
      resource_revision: 7,
      updated_at: "2026-08-29T00:00:00Z",
    };
    apiMocks.get.mockResolvedValue({
      data: existingTarget,
      response: { status: 200 },
    });
    apiMocks.put.mockResolvedValue({
      data: {
        ...existingTarget,
        minimum_listening_band: 6.5,
        minimum_writing_band: 6.5,
        minimum_speaking_band: 7,
        resource_revision: 8,
      },
      response: { status: 200 },
    });

    renderReadingPractice();

    const listening = screen.getByLabelText("minimumListeningBand");
    const writing = screen.getByLabelText("minimumWritingBand");
    const speaking = screen.getByLabelText("minimumSpeakingBand");
    await waitFor(() => {
      expect(listening).toHaveValue(6);
      expect(writing).toHaveValue(6);
      expect(speaking).toHaveValue(6);
    });

    fireEvent.change(listening, { target: { value: "6.5" } });
    fireEvent.change(writing, { target: { value: "6.5" } });
    fireEvent.change(speaking, { target: { value: "7" } });
    fireEvent.click(screen.getByRole("button", { name: "saveTarget" }));

    await waitFor(() =>
      expect(apiMocks.put).toHaveBeenCalledWith("/v1/target-profile", {
        body: {
          test_variant: "ACADEMIC",
          target_overall_band: 7,
          minimum_listening_band: 6.5,
          minimum_reading_band: 7.5,
          minimum_writing_band: 6.5,
          minimum_speaking_band: 7,
          expected_resource_revision: 7,
        },
      }),
    );
  });

  it("allows a learner to save an overall-only Band target", async () => {
    apiMocks.get.mockResolvedValue({
      error: { error: { message: "resource not found" } },
      response: { status: 404 },
    });
    apiMocks.put.mockResolvedValue({
      data: {
        test_variant: "ACADEMIC",
        target_overall_band: 7,
        resource_revision: 1,
        updated_at: "2026-08-29T00:00:00Z",
      },
      response: { status: 201 },
    });

    renderReadingPractice();

    const save = screen.getByRole("button", { name: "saveTarget" });
    await waitFor(() => expect(save).toBeEnabled());
    fireEvent.change(screen.getByLabelText("variant"), {
      target: { value: "ACADEMIC" },
    });
    fireEvent.change(screen.getByLabelText("targetOverallBand"), {
      target: { value: "7" },
    });
    fireEvent.click(save);

    await waitFor(() =>
      expect(apiMocks.put).toHaveBeenCalledWith("/v1/target-profile", {
        body: {
          test_variant: "ACADEMIC",
          target_overall_band: 7,
          minimum_listening_band: undefined,
          minimum_reading_band: undefined,
          minimum_writing_band: undefined,
          minimum_speaking_band: undefined,
          expected_resource_revision: 0,
        },
      }),
    );
  });

  it("keeps an unknown Band constraint blank for a new learner", async () => {
    apiMocks.get.mockResolvedValue({
      error: { error: { message: "resource not found" } },
      response: { status: 404 },
    });

    renderReadingPractice();

    await waitFor(() => {
      expect(screen.getByLabelText("variant")).toHaveValue("");
      expect(screen.getByLabelText("targetOverallBand")).toHaveValue(null);
      expect(screen.getByLabelText("minimumListeningBand")).toHaveValue(null);
      expect(screen.getByLabelText("minimumReadingBand")).toHaveValue(null);
      expect(screen.getByLabelText("minimumWritingBand")).toHaveValue(null);
      expect(screen.getByLabelText("minimumSpeakingBand")).toHaveValue(null);
    });
  });
});
