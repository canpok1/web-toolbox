import { type Locator, type Page, expect } from "@playwright/test";
import { JoinSessionPagePom } from "./JoinSessionPagePom";

export class SessionPagePom {
  private page: Page;

  constructor(page: Page) {
    this.page = page;
  }

  getNameElement(name: string): Locator {
    return this.page.getByText(`あなたの名前: ${name}`);
  }

  getParticipantNameElement(count: number): Locator {
    return this.page.getByText(`参加者 (${count}名):`);
  }

  getParticipantByName(name: string): Locator {
    return this.page.getByText(name, { exact: true });
  }

  getInviteUrlButton(): Locator {
    return this.page.getByRole("button", {
      name: "招待URL/QRコード",
      exact: true,
    });
  }

  getInviteUrlCopyButton(): Locator {
    return this.page.getByRole("button", {
      name: "参加ページのURLをコピー",
      exact: true,
    });
  }

  async clickInviteUrlButton(): Promise<void> {
    await this.getInviteUrlButton().click();
  }

  get inviteLink(): Locator {
    return this.page.locator('a[href*="/planning-poker/sessions/join?id="]');
  }

  get startVoteButton(): Locator {
    return this.page.getByRole("button", { name: "投票を開始", exact: true });
  }

  get openVoteButton(): Locator {
    return this.page.getByRole("button", { name: "投票を公開", exact: true });
  }

  getVoteButton(value: string): Locator {
    return this.page.getByRole("button", { name: value, exact: true });
  }

  async clickStartVoteButton(): Promise<void> {
    await this.startVoteButton.click();
  }

  async clickVoteButton(value: string): Promise<void> {
    await this.getVoteButton(value).click();
  }

  async clickOpenVoteButton(): Promise<void> {
    await this.openVoteButton.click();
  }

  getVotedIndicator(): Locator {
    return this.page.locator('[data-tip="投票済み"]');
  }

  async joinAsParticipant(participantName: string): Promise<Page> {
    // 招待URLを取得
    await this.clickInviteUrlButton();
    const inviteLink = await this.inviteLink.getAttribute("href");
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
