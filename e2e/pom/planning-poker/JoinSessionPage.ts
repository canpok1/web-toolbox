import type { Page } from "@playwright/test";

export class JoinSessionPagePom {
  private page: Page;

  constructor(page: Page) {
    this.page = page;
  }

  public async inputUserName(userName: string): Promise<void> {
    await this.page.getByLabel("あなたの名前").fill(userName);
  }

  public async clickJoinButton(): Promise<void> {
    await this.page
      .getByRole("button", { name: "セッションに参加", exact: true })
      .click();
    await this.page.waitForEvent("websocket");
  }

  public async joinSession(userName: string): Promise<void> {
    await this.inputUserName(userName);
    await this.clickJoinButton();
  }
}
