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

  public get scaleSelect() {
    return this.page.getByLabel("スケール");
  }

  public get yourNameInput() {
    return this.page.getByLabel("あなたの名前");
  }

  public get createSessionButton() {
    return this.page.getByRole("button", { name: "セッションを作成" });
  }

  public get backLink() {
    return this.page.getByRole("link", { name: "戻る" });
  }

  public async fillYourName(userName: string): Promise<void> {
    await this.yourNameInput.fill(userName);
  }

  public async selectScale(scale: Scale): Promise<void> {
    await this.scaleSelect.selectOption(scale);
  }

  public async clickCreateSessionButton(): Promise<void> {
    await this.createSessionButton.click();
    await this.page.waitForEvent("websocket");
  }

  public async createSession(userName: string, scale: Scale): Promise<void> {
    await this.fillYourName(userName);
    await this.selectScale(scale);
    await this.clickCreateSessionButton();
  }
}
