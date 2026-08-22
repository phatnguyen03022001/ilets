import createClient from "openapi-fetch";
import type { paths } from "@/generated/public-v1";

const baseUrl = process.env.NEXT_PUBLIC_CORE_API_URL ?? "http://127.0.0.1:8080";

export const api = createClient<paths>({
  baseUrl,
  credentials: "include",
});
