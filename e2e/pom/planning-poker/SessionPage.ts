import type { Locator, Page } from "@playwright/test";

export class SessionPagePom {
  private page: Page;

  constructor(page: Page) {
    this.page = page;
  }

  public getNameElement(name: string): Locator {
    return this.page.getByText(`あなたの名前: ${name}`);
  }

  public getParticipantNameElement(count: number): Locator {
    return this.page.getByText(new RegExp(`参加者 \\(${count}名\\):\\s*`));
  }

  public getInviteUrlButton(): Locator {
    return this.page.getByRole("button", {
      name: "招待URL/QRコード",
      exact: true,
    });
  }

  public getInviteUrlCopyButton(): Locator {
    return this.page.getByRole("button", {
      name: "参加ページのURLをコピー",
      exact: true,
    });
  }

  public async clickInviteUrlButton(): Promise<void> {
    await this.getInviteUrlButton().click();
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
    await this.page.getByRole("button", { name: value, exact: true }).click();
  }

  public async clickOpenVoteButton(): Promise<void> {
    await this.page
      .getByRole("button", { name: "投票を公開", exact: true })
      .click();
  }

  public getVotedIndicator(): Locator {
    return this.page.locator('[data-tip="投票済み"]');
  }
}
