import { createClient } from "@/generated/public/client";

const baseUrl = process.env.NEXT_PUBLIC_CORE_API_URL ?? "http://127.0.0.1:8080";

export type GetClerkToken = () => Promise<string | null>;

export function createPublicApi(
  getToken: GetClerkToken,
  fetchImpl?: typeof fetch,
) {
  return createClient({
    baseUrl,
    fetch: fetchImpl,
    auth: async () => (await getToken()) ?? undefined,
  });
}
