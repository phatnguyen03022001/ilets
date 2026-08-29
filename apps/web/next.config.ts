import type { NextConfig } from "next";
import createNextIntlPlugin from "next-intl/plugin";

const withNextIntl = createNextIntlPlugin("./src/i18n/request.ts");

const e2eClerkStub = process.env.ILETS_E2E === "1";

const nextConfig: NextConfig = {
  poweredByHeader: false,
  ...(e2eClerkStub
    ? {
        turbopack: {
          resolveAlias: {
            "@clerk/nextjs": "./e2e/clerk-nextjs.e2e.tsx",
          },
        },
      }
    : {}),
};

export default withNextIntl(nextConfig);
