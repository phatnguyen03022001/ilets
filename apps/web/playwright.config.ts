import { defineConfig, devices } from "@playwright/test";

const host = "127.0.0.1";
const corePort = process.env.ILETS_E2E_CORE_PORT ?? "18080";
const webPort = process.env.ILETS_E2E_WEB_PORT ?? "13000";
const coreOrigin = `http://${host}:${corePort}`;
const webOrigin = `http://${host}:${webPort}`;

export default defineConfig({
  testDir: "./e2e",
  timeout: 30_000,
  retries: 0,
  use: {
    baseURL: webOrigin,
    trace: "retain-on-failure",
  },
  webServer: [
    {
      command: `cd ../../services/core-api && CORE_ADDR=${host}:${corePort} WEB_ORIGINS=${webOrigin} go run ./cmd/core-api`,
      url: `${coreOrigin}/healthz`,
      reuseExistingServer: false,
      timeout: 120_000,
    },
    {
      command: `NEXT_PUBLIC_CORE_API_URL=${coreOrigin} pnpm dev --hostname ${host} --port ${webPort}`,
      url: webOrigin,
      reuseExistingServer: false,
      timeout: 120_000,
    },
  ],
  projects: [{ name: "chromium", use: { ...devices["Desktop Chrome"] } }],
});
