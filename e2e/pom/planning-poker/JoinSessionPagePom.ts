import type { Locator, Page } from "@playwright/test";

export class JoinSessionPagePom {
  private page: Page;

  constructor(page: Page) {
    this.page = page;
  }

  async goto(): Promise<void> {
    await this.page.goto("/planning-poker/sessions/join");
  }

  get sessionIdInput(): Locator {
    return this.page.getByLabel("セッションID");
  }

  get yourNameInput(): Locator {
    return this.page.getByLabel("あなたの名前");
  }

  get joinSessionButton(): Locator {
    return this.page.getByRole("button", {
      name: "セッションに参加",
    });
  }

  get backLink(): Locator {
    return this.page.getByRole("link", { name: "戻る" });
  }

  get alertMessage(): Locator {
    return this.page.locator("[data-testid=\"alert-message\"]");
  }

  async fillSessionId(sessionId: string): Promise<void> {
    await this.sessionIdInput.fill(sessionId);
  }

  async fillYourName(yourName: string): Promise<void> {
    await this.yourNameInput.fill(yourName);
  }

  async clickJoinSessionButton(): Promise<void> {
    await this.joinSessionButton.click();
  }
}
