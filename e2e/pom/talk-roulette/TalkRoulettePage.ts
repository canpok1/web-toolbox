import type { Page } from "@playwright/test";

export class TalkRoulettePagePom {
  private page: Page;

  constructor(page: Page) {
    this.page = page;
  }

  public async startRoulette(): Promise<void> {
    // Example: Locating a button by its accessible name 'Start' and clicking it.
    await this.page.getByRole("button", { name: "Start" }).click();
    // Example: Wait for some network activity to settle after the action.
    await this.page.waitForLoadState("networkidle");
  }

  public async stopRoulette(): Promise<void> {
    // Example: Locating a button by its accessible name 'Stop' and clicking it.
    await this.page.getByRole("button", { name: "Stop" }).click();
    // Example: Wait for navigation or a specific element to ensure the page has updated.
    // await this.page.waitForNavigation({ waitUntil: 'domcontentloaded' });
  }

  public async selectRandomTalk(): Promise<void> {
    // Example: Locating a button by its accessible name 'Next Random Talk' and clicking it.
    await this.page.getByRole("button", { name: "Next Random Talk" }).click();
    // Example: Wait for a specific response from the backend API.
    // await this.page.waitForResponse(resp => resp.url().includes('/api/talks/random') && resp.status() === 200);
    // Example: Or wait for a specific element that displays talk content to be visible.
    await this.page
      .locator(".talk-content-wrapper")
      .waitFor({ state: "visible", timeout: 5000 });
  }

  public async getCurrentTalkTitle(): Promise<string | null> {
    // Example: Locating an element assumed to contain the talk title by a test ID.
    const titleElement = this.page.locator('[data-testid="talk-title"]');
    return titleElement.textContent();
  }

  public async getCurrentTalkSpeaker(): Promise<string | null> {
    // Example: Locating an element for speaker name
    const speakerElement = this.page.locator('[data-testid="talk-speaker"]');
    return speakerElement.textContent();
  }

  public async isTalkVisible(): Promise<boolean> {
    // Example: Check if a talk container element is visible
    const talkContainer = this.page.locator('[data-testid="talk-container"]');
    return talkContainer.isVisible();
  }
}
