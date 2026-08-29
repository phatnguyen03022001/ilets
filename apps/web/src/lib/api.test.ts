import { beforeEach, describe, expect, it, vi } from "vitest";
import { getMe } from "@/generated/public";
import { createPublicApi } from "./api";

describe("generated public client auth", () => {
  beforeEach(() => vi.restoreAllMocks());

  it("attaches the current Clerk token as Bearer without localStorage persistence", async () => {
    const getToken = vi.fn(async () => "signed-clerk-session-jwt");
    const storageGet = vi.spyOn(Storage.prototype, "getItem");
    const storageSet = vi.spyOn(Storage.prototype, "setItem");
    const fetchImpl = vi.fn(
      async (request: RequestInfo | URL, init?: RequestInit) => {
        const req =
          request instanceof Request ? request : new Request(request, init);
        expect(req.headers.get("Authorization")).toBe(
          "Bearer signed-clerk-session-jwt",
        );
        return new Response(
          JSON.stringify({ actor_id: "actor_1", learner_id: "learner_1" }),
          {
            status: 200,
            headers: { "Content-Type": "application/json" },
          },
        );
      },
    );

    const client = createPublicApi(getToken, fetchImpl as typeof fetch);
    const response = await getMe({ client });

    expect(response.data).toEqual({
      actor_id: "actor_1",
      learner_id: "learner_1",
    });
    expect(getToken).toHaveBeenCalledTimes(1);
    expect(storageGet).not.toHaveBeenCalled();
    expect(storageSet).not.toHaveBeenCalled();
  });
});
