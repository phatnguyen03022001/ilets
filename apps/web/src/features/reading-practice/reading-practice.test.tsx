import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { cleanup, render, screen, waitFor } from "@testing-library/react";
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

  it("hydrates the persisted minimum Reading Band", async () => {
    apiMocks.get.mockResolvedValue({
      data: {
        test_variant: "ACADEMIC",
        minimum_reading_band: 7.5,
        resource_revision: 3,
        updated_at: "2026-08-28T00:00:00Z",
      },
      response: { status: 200 },
    });

    renderReadingPractice();

    const input = screen.getByLabelText("minimumReadingBand");
    await waitFor(() => expect(input).toHaveValue(7.5));
  });

  it("keeps an unknown Band constraint blank for a new learner", async () => {
    apiMocks.get.mockResolvedValue({
      error: { error: { message: "resource not found" } },
      response: { status: 404 },
    });

    renderReadingPractice();

    const input = screen.getByLabelText("minimumReadingBand");
    await waitFor(() => expect(input).toHaveValue(null));
  });
});
