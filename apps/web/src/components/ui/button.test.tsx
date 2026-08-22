import { render, screen } from "@testing-library/react";
import userEvent from "@testing-library/user-event";
import { describe, expect, it, vi } from "vitest";
import { Button } from "./button";

describe("Button", () => {
  it("uses native disabled behavior", async () => {
    const user = userEvent.setup();
    const onClick = vi.fn();

    render(
      <Button disabled onClick={onClick}>
        Save target
      </Button>,
    );

    await user.click(screen.getByRole("button", { name: "Save target" }));
    expect(onClick).not.toHaveBeenCalled();
  });
});
