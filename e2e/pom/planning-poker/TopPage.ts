import { Page } from "@playwright/test";

export class PlanningPokerTopPagePom {
  constructor(private page: Page) {}

  async goto() {
    await this.page.goto("/planning-poker");
  }

  get createSessionLink() {
    return this.page.getByRole("link", { name: "セッションを作成" });
  }

  get joinSessionLink() {
    return this.page.getByRole("link", { name: "セッションに参加" });
  }

  async clickCreateSessionLink() {
    await this.createSessionLink.click();
  }

  async clickJoinSessionLink() {
    await this.joinSessionLink.click();
  }
}
