import { describe, expect, it } from "vitest";
import { newIdempotencyKey } from "./idempotency";

describe("newIdempotencyKey", () => {
  it("matches the public contract character envelope", () => {
    const original = globalThis.crypto;
    Object.defineProperty(globalThis, "crypto", {
      value: { randomUUID: () => "123e4567-e89b-12d3-a456-426614174000" },
      configurable: true,
    });
    expect(newIdempotencyKey("submit")).toMatch(/^[A-Za-z0-9._:-]{8,128}$/);
    Object.defineProperty(globalThis, "crypto", {
      value: original,
      configurable: true,
    });
  });
});
