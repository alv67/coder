import { type FC, type MouseEventHandler } from "react";
import { type ClickableAriaRole, useClickable } from "./useClickable";
import { render, screen } from "@testing-library/react";
import userEvent from "@testing-library/user-event";

/**
 * Since the point of the hook is to take a traditionally non-interactive HTML
 * element and make it interactive, it made the most sense to test against an
 * element directly.
 */
type NonNativeButtonProps<TElement extends HTMLElement = HTMLElement> =
  Readonly<{
    role?: ClickableAriaRole;
    onInteraction: MouseEventHandler<TElement>;
  }>;

const NonNativeButton: FC<NonNativeButtonProps<HTMLDivElement>> = ({
  onInteraction,
  role = "button",
}) => {
  const clickableProps = useClickable(onInteraction, role);
  return <div {...clickableProps} />;
};

describe(useClickable.name, () => {
  it("Always defaults to role 'button'", () => {
    render(<NonNativeButton onInteraction={jest.fn()} />);
    expect(() => screen.getByRole("button")).not.toThrow();
  });

  it("Always returns out the same role override received via arguments", () => {
    const placeholderCallback = jest.fn();
    const roles = [
      "button",
      "switch",
    ] as const satisfies readonly ClickableAriaRole[];

    for (const role of roles) {
      const { unmount } = render(
        <NonNativeButton role={role} onInteraction={placeholderCallback} />,
      );

      expect(() => screen.getByRole(role)).not.toThrow();
      unmount();
    }
  });

  it("Allows an element to receive keyboard focus", async () => {
    const user = userEvent.setup();
    const mockCallback = jest.fn();

    render(<NonNativeButton role="button" onInteraction={mockCallback} />, {
      wrapper: ({ children }) => (
        <>
          <button>Native button</button>
          {children}
        </>
      ),
    });

    await user.keyboard("[Tab][Tab]");
    await user.keyboard(" ");
    expect(mockCallback).toBeCalledTimes(1);
  });

  it("Allows an element to respond to clicks and Space/Enter, following all rules for native Button element interactions", async () => {
    const mockCallback = jest.fn();
    const user = userEvent.setup();
    render(<NonNativeButton role="button" onInteraction={mockCallback} />);

    await user.click(document.body);
    await user.keyboard(" ");
    await user.keyboard("[Enter]");
    expect(mockCallback).not.toBeCalled();

    const button = screen.getByRole("button");
    await user.click(button);
    await user.keyboard(" ");
    await user.keyboard("[Enter]");
    expect(mockCallback).toBeCalledTimes(3);
  });

  it("Will keep firing events if the Enter key is held down", async () => {
    const mockCallback = jest.fn();
    const user = userEvent.setup();
    render(<NonNativeButton role="button" onInteraction={mockCallback} />);

    // Focus over to element, hold down Enter for specified keydown period
    // (count determined by browser/library), and then release the Enter key
    const keydownCount = 5;
    await user.keyboard("[Tab]");
    await user.keyboard(`{Enter>${keydownCount}}`);
    await user.keyboard("{/Enter}");
    expect(mockCallback).toBeCalledTimes(keydownCount);
  });

  it("Will NOT keep firing events if the Space key is held down", async () => {
    const mockCallback = jest.fn();
    const user = userEvent.setup();
    render(<NonNativeButton role="button" onInteraction={mockCallback} />);

    // Focus over to element, hold down Space for specified keydown period
    // (count determined by browser/library)
    const keydownCount = 5;
    await user.keyboard("[Tab]");
    await user.keyboard(`{ >${keydownCount}}`);
    expect(mockCallback).not.toBeCalled();

    // Then release the space bar
    await user.keyboard(`{/ }`);
    expect(mockCallback).toBeCalledTimes(1);
  });

  test("If focus is lost while Space is held down, then releasing the key will do nothing", async () => {
    const mockCallback = jest.fn();
    const user = userEvent.setup();

    render(<NonNativeButton role="button" onInteraction={mockCallback} />, {
      wrapper: ({ children }) => (
        <>
          {children}
          <button>Native button</button>
        </>
      ),
    });

    // Focus over to element, hold down Space for an indefinite amount of time,
    // move focus away from element, and then release Space
    await user.keyboard("[Tab]");
    await user.keyboard("{ >}");
    await user.keyboard("[Tab]");
    await user.keyboard("{/ }");
    expect(mockCallback).not.toBeCalled();
  });
});
