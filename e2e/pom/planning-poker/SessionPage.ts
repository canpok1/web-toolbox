import { type Page } from "@playwright/test";

export class SessionPagePom {
  private page: Page;

  constructor(page: Page) {
    this.page = page;
  }

  public async clickInviteUrlButton(): Promise<void> {
    await this.page
      .getByRole("button", { name: "招待URL/QRコード", exact: true })
      .click();
  }

  public async copyInviteUrl(): Promise<string | null> {
    const inviteLink = this.page.locator(
      'a[href*="/planning-poker/sessions/join?id="]',
    );
    return await inviteLink.getAttribute("href");
  }

  public async clickStartVoteButton(): Promise<void> {
    await this.page
      .getByRole("button", { name: "投票を開始", exact: true })
      .click();
  }

  public async clickVoteButton(value: string): Promise<void> {
    await this.page
      .getByRole("button", { name: value, exact: true })
      .click();
  }

  public async clickOpenVoteButton(): Promise<void> {
    await this.page
      .getByRole("button", { name: "投票を公開", exact: true })
      .click();
  }
}
