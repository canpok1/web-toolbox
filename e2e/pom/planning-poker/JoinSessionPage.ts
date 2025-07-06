import type { Locator, Page } from "@playwright/test";

export class JoinSessionPagePom {
  private page: Page;

  readonly sessionIdInput: Locator;
  readonly yourNameInput: Locator;
  readonly joinSessionButton: Locator;
  readonly backLink: Locator;
  readonly alertMessage: Locator;

  constructor(page: Page) {
    this.page = page;
    this.sessionIdInput = page.getByLabel("セッションID");
    this.yourNameInput = page.getByLabel("あなたの名前");
    this.joinSessionButton = page.getByRole("button", {
      name: "セッションに参加",
    });
    this.backLink = page.getByRole("link", { name: "戻る" });
    this.alertMessage = page.locator(".alert");
  }

  public async goto(): Promise<void> {
    await this.page.goto("/planning-poker/sessions/join");
  }

  public async fillSessionId(sessionId: string): Promise<void> {
    await this.sessionIdInput.fill(sessionId);
  }

  public async fillYourName(yourName: string): Promise<void> {
    await this.yourNameInput.fill(yourName);
  }

  public async clickJoinSessionButton(): Promise<void> {
    await this.joinSessionButton.click();
  }
}
