import { type Locator, type Page, expect } from "@playwright/test";
import { JoinSessionPagePom } from "./JoinSessionPage";

export class SessionPagePom {
  private page: Page;

  constructor(page: Page) {
    this.page = page;
  }

  public getNameElement(name: string): Locator {
    return this.page.getByText(`あなたの名前: ${name}`);
  }

  public getParticipantNameElement(count: number): Locator {
    return this.page.getByText(`参加者 (${count}名):`);
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

  public async getInviteLink(): Promise<string | null> {
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

  public async joinAsParticipant(participantName: string): Promise<Page> {
    // 招待URLを取得
    await this.clickInviteUrlButton();
    const inviteLink = await this.getInviteLink();
    expect(inviteLink).not.toBeNull();

    // 新しいページで参加者として参加
    const participantPage = await this.page.context().newPage();
    await participantPage.goto(inviteLink as string);
    const participantPom = new JoinSessionPagePom(participantPage);
    await participantPom.fillYourName(participantName);
    await participantPom.clickJoinSessionButton();
    await participantPage.waitForEvent("websocket");

    return participantPage;
  }
}
