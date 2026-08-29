import { generateKeyPairSync, sign } from "node:crypto";
import { defineConfig, devices } from "@playwright/test";

const host = "127.0.0.1";
const corePort = process.env.ILETS_E2E_CORE_PORT ?? "18080";
const webPort = process.env.ILETS_E2E_WEB_PORT ?? "13000";
const coreOrigin = `http://${host}:${corePort}`;
const webOrigin = `http://${host}:${webPort}`;
const clerkIssuer = "https://e2e.clerk.accounts.dev";
const clerkAudience = "ilets-core";

const { publicKey, privateKey } = generateKeyPairSync("rsa", {
  modulusLength: 2048,
  publicKeyEncoding: { type: "spki", format: "pem" },
  privateKeyEncoding: { type: "pkcs8", format: "pem" },
});

function encode(value: object) {
  return Buffer.from(JSON.stringify(value)).toString("base64url");
}

const now = Math.floor(Date.now() / 1000);
const protectedHeader = encode({ alg: "RS256", typ: "JWT", kid: "e2e-key" });
const payload = encode({
  iss: clerkIssuer,
  sub: "user_playwright",
  aud: clerkAudience,
  azp: webOrigin,
  sid: "sess_playwright",
  exp: now + 3600,
  nbf: now - 60,
  iat: now - 60,
});
const signingInput = `${protectedHeader}.${payload}`;
const bearerToken = `${signingInput}.${sign(
  "RSA-SHA256",
  Buffer.from(signingInput),
  privateKey,
).toString("base64url")}`;

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
      command:
        "cd ../../services/core-api && go test ./internal/httpapi -run '^TestPlaywrightServer$' -count=1 -timeout=0",
      env: {
        ...process.env,
        CORE_ADDR: `${host}:${corePort}`,
        WEB_ORIGINS: webOrigin,
        CLERK_ISSUER: clerkIssuer,
        CLERK_AUDIENCE: clerkAudience,
        CLERK_AUTHORIZED_PARTIES: webOrigin,
        ILETS_E2E_PUBLIC_KEY_PEM: publicKey,
        ILETS_E2E_SERVER: "1",
      },
      url: `${coreOrigin}/healthz`,
      reuseExistingServer: false,
      timeout: 120_000,
    },
    {
      command: `corepack pnpm dev --hostname ${host} --port ${webPort}`,
      env: {
        ...process.env,
        ILETS_E2E: "1",
        NEXT_PUBLIC_CORE_API_URL: coreOrigin,
        NEXT_PUBLIC_ILETS_E2E_BEARER_TOKEN: bearerToken,
      },
      url: webOrigin,
      reuseExistingServer: false,
      timeout: 120_000,
    },
  ],
  projects: [{ name: "chromium", use: { ...devices["Desktop Chrome"] } }],
});
