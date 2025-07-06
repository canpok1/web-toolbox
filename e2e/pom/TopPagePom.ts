import { type Page, expect } from "@playwright/test";

export class TopPagePom {
  constructor(private readonly page: Page) {}

  async goto() {
    await this.page.goto("/");
  }

  async expectTitleToContain(text: string) {
    await expect(this.page).toHaveTitle(new RegExp(text));
  }

  get planningPokerLink() {
    return this.page.getByRole("link", { name: "プランニングポーカー" });
  }

  get talkRouletteLink() {
    return this.page.getByRole("link", { name: "トークルーレット" });
  }

  async clickPlanningPokerLink() {
    await this.planningPokerLink.click();
  }

  async clickTalkRouletteLink() {
    await this.talkRouletteLink.click();
  }

  async waitForURL(url: string | RegExp) {
    await this.page.waitForURL(url);
  }

  getByRole(
    role: Parameters<Page["getByRole"]>[0],
    options: { name?: string | RegExp; level?: number },
  ) {
    return this.page.getByRole(role, options);
  }
}
