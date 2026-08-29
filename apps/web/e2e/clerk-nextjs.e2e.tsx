"use client";

import type { ReactNode } from "react";

const token = process.env.NEXT_PUBLIC_ILETS_E2E_BEARER_TOKEN ?? null;

export function ClerkProvider({ children }: { children: ReactNode }) {
  return <>{children}</>;
}

export function useAuth() {
  return {
    getToken: async () => token,
    isLoaded: true,
    isSignedIn: token !== null,
  };
}

export function SignInButton({ children }: { children: ReactNode }) {
  return <>{children}</>;
}
