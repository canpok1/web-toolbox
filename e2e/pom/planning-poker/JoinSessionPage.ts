import type { Locator, Page } from "@playwright/test";

export class JoinSessionPagePom {
  private page: Page;

  constructor(page: Page) {
    this.page = page;
  }

  public async goto(): Promise<void> {
    await this.page.goto("/planning-poker/sessions/join");
  }

  public get sessionIdInput(): Locator {
    return this.page.getByLabel("セッションID");
  }

  public get yourNameInput(): Locator {
    return this.page.getByLabel("あなたの名前");
  }

  public get joinSessionButton(): Locator {
    return this.page.getByRole("button", {
      name: "セッションに参加",
    });
  }

  public get backLink(): Locator {
    return this.page.getByRole("link", { name: "戻る" });
  }

  public get alertMessage(): Locator {
    return this.page.locator(".alert");
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
