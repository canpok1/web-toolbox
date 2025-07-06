import type { Page } from "@playwright/test";

export type Scale = "fibonacci" | "t-shirt" | "power-of-two";

export class CreateSessionPagePom {
  private page: Page;

  constructor(page: Page) {
    this.page = page;
  }

  async goto(): Promise<void> {
    await this.page.goto("/planning-poker/sessions/create");
  }

  get scaleSelect() {
    return this.page.getByLabel("スケール");
  }

  get yourNameInput() {
    return this.page.getByLabel("あなたの名前");
  }

  get createSessionButton() {
    return this.page.getByRole("button", { name: "セッションを作成" });
  }

  get backLink() {
    return this.page.getByRole("link", { name: "戻る" });
  }

  async fillYourName(userName: string): Promise<void> {
    await this.yourNameInput.fill(userName);
  }

  async selectScale(scale: Scale): Promise<void> {
    await this.scaleSelect.selectOption(scale);
  }

  async clickCreateSessionButton(): Promise<void> {
    await this.createSessionButton.click();
    await this.page.waitForEvent("websocket");
  }

  async createSession(userName: string, scale: Scale): Promise<void> {
    await this.fillYourName(userName);
    await this.selectScale(scale);
    await this.clickCreateSessionButton();
  }
}
