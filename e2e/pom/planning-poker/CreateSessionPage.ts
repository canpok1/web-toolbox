import type { Page } from "@playwright/test";

export type Scale = "fibonacci" | "t-shirt" | "power-of-two";

export class CreateSessionPagePom {
  private page: Page;

  constructor(page: Page) {
    this.page = page;
  }

  public async goto(): Promise<void> {
    await this.page.goto("/planning-poker/sessions/create");
  }

  public async inputUserName(userName: string): Promise<void> {
    await this.page.getByLabel("あなたの名前").fill(userName);
  }

  public async selectScale(scale: Scale): Promise<void> {
    await this.page.getByLabel("スケール").selectOption(scale);
  }

  public async clickCreateButton(): Promise<void> {
    await this.page
      .getByRole("button", { name: "セッションを作成", exact: true })
      .click();
    await this.page.waitForEvent("websocket");
  }

  public async createSession(userName: string, scale: Scale): Promise<void> {
    await this.inputUserName(userName);
    await this.selectScale(scale);
    await this.clickCreateButton();
  }
}
