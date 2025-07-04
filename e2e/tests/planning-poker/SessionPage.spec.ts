import { type Page, expect, test } from "@playwright/test";
import { CreateSessionPagePom } from "../../pom/planning-poker/CreateSessionPage";
import { JoinSessionPagePom } from "../../pom/planning-poker/JoinSessionPage";
import { SessionPagePom } from "../../pom/planning-poker/SessionPage";

async function checkVoteButtonsVisibility(page: Page, buttons: string[]) {
  for (const voteButton of buttons) {
    await expect(
      page.getByRole("button", {
        name: voteButton,
        exact: true,
      }),
    ).toBeVisible();
  }
}

async function performVotingFlow(
  hostPage: Page,
  participantPage: Page,
  hostPom: SessionPagePom,
  participantPom: SessionPagePom,
  voteButtons: string[],
  hostVote: string,
  participantVote: string,
) {
  // 参加者ユーザー画面に投票開始ボタンが表示されないことを確認
  await expect(
    participantPage.getByRole("button", {
      name: "投票を開始",
      exact: true,
    }),
  ).not.toBeVisible();
  // ホストユーザー画面に投票開始ボタンが表示されることを確認
  await expect(
    hostPage.getByRole("button", { name: "投票を開始", exact: true }),
  ).toBeVisible();

  // ホストが投票を開始する
  await hostPom.clickStartVoteButton();

  // 画面表示確認
  // 参加者ユーザー画面に投票ボタンが表示されることを確認
  await checkVoteButtonsVisibility(participantPage, voteButtons);
  await checkVoteButtonsVisibility(hostPage, voteButtons);

  // 参加者ユーザーが投票する
  await participantPom.clickVoteButton(participantVote);

  // 画面表示確認
  await expect(hostPom.getVotedIndicator()).toHaveCount(1);
  await expect(participantPom.getVotedIndicator()).toHaveCount(1);

  // ホストユーザーが投票する
  await hostPom.clickVoteButton(hostVote);

  // 画面表示確認
  await expect(hostPom.getVotedIndicator()).toHaveCount(2);
  await expect(participantPom.getVotedIndicator()).toHaveCount(2);

  // ホストが投票を公開する
  await hostPom.clickOpenVoteButton();

  // 画面表示確認
  // 参加者ユーザー画面に投票開始ボタンが表示されないことを確認
  await expect(
    participantPage.getByRole("button", {
      name: "投票を開始",
      exact: true,
    }),
  ).not.toBeVisible();
  // ホストユーザー画面に投票開始ボタンが表示されることを確認
  await expect(
    hostPage.getByRole("button", { name: "投票を開始", exact: true }),
  ).toBeVisible();

  // ホストが投票を開始する
  await hostPom.clickStartVoteButton();

  // 画面表示確認
  // 参加者ユーザー画面に投票ボタンが表示されることを確認
  await checkVoteButtonsVisibility(participantPage, voteButtons);
  await checkVoteButtonsVisibility(hostPage, voteButtons);
}

test.describe("セッション画面", () => {
  const hostUserName = "ホストユーザー";

  test.beforeEach(async ({ page }) => {
    const pom = new CreateSessionPagePom(page);
    await pom.goto();
    await pom.createSession(hostUserName, "fibonacci");
  });

  test.describe("参加者一覧と招待リンク", () => {
    test.describe("ホストのみ", () => {
      test("表示内容が正しいこと", async ({ page: hostPage }) => {
        const hostPom = new SessionPagePom(hostPage);

        // ホストユーザー画面に自分の名前が表示されるか確認
        await expect(hostPom.getNameElement(hostUserName)).toBeVisible();
        await expect(hostPom.getParticipantNameElement(1)).toBeVisible();
        await expect(
          hostPage.getByText(hostUserName, { exact: true }),
        ).toBeVisible();
        await expect(hostPom.getInviteUrlCopyButton()).toBeVisible();

        const pom = new SessionPagePom(hostPage);
        await pom.clickInviteUrlButton();
        const inviteLinkValue = await pom.copyInviteUrl();
        await expect(inviteLinkValue).toMatch(
          /planning-poker\/sessions\/join\?id=.*/,
        );
      });
    });

    test.describe("参加者が複数", () => {
      const participantUserName = "参加者ユーザー";
      test("表示内容が正しいこと", async ({ page: hostPage }) => {
        const hostPom = new SessionPagePom(hostPage);
        const participantPage = await hostPom.joinAsParticipant(
          participantUserName,
        );
        const participantPom = new SessionPagePom(participantPage);

        // 参加者ユーザー画面に自分の名前が表示されるか確認
        await expect(
          participantPom.getNameElement(participantUserName),
        ).toBeVisible();

        // 参加者ユーザー画面に参加者一覧が表示されるか確認
        await expect(participantPom.getParticipantNameElement(2)).toBeVisible();
        await expect(
          participantPage.getByText(participantUserName, { exact: true }),
        ).toBeVisible();
        await expect(
          participantPage.getByText(hostUserName, { exact: true }),
        ).toBeVisible();
        // 参加者ユーザー画面で招待リンクをコピーできることを確認
        await expect(participantPom.getInviteUrlCopyButton()).toBeVisible();

        // 参加者ユーザー画面の招待リンクが正しいことを確認
        await participantPom.clickInviteUrlButton();
        const participantInviteLinkValue = await participantPom.copyInviteUrl();

        await expect(participantInviteLinkValue).toMatch(
          /planning-poker\/sessions\/join\?id=.*/,
        );

        await hostPage.bringToFront();

        // ホストユーザー画面に自分の名前が表示されるか確認
        await expect(hostPom.getNameElement(hostUserName)).toBeVisible();
        // ホストユーザー画面に参加者一覧が表示されるか確認
        await expect(hostPom.getParticipantNameElement(2)).toBeVisible();
        await expect(
          hostPage.getByText(hostUserName, { exact: true }),
        ).toBeVisible();
        await expect(
          hostPage.getByText(participantUserName, { exact: true }),
        ).toBeVisible();
      });
    });
  });

  test.describe("ホスト用ボタンと投票ボタンと投票結果", () => {
    test.beforeEach(async ({ page }, testInfo) => {
      const pom = new CreateSessionPagePom(page);
      await pom.goto();
      let scaleType: "fibonacci" | "t-shirt" | "power-of-two";
      if (testInfo.titlePath.includes("フィボナッチ")) {
        scaleType = "fibonacci";
      } else if (testInfo.titlePath.includes("Tシャツサイズ")) {
        scaleType = "t-shirt";
      } else if (testInfo.titlePath.includes("2の累乗")) {
        scaleType = "power-of-two";
      } else {
        throw new Error("Unknown scale type");
      }
      await pom.createSession(hostUserName, scaleType);
    });
    test.describe("フィボナッチ", () => {
      test("投票フロー", async ({ page: hostPage }) => {
        const hostPom = new SessionPagePom(hostPage);
        const participantUserName = "参加者ユーザー";
        const participantPage = await hostPom.joinAsParticipant(
          participantUserName,
        );
        const participantPom = new SessionPagePom(participantPage);

        const voteButtons = [
          "0",
          "1",
          "2",
          "3",
          "5",
          "8",
          "13",
          "21",
          "34",
          "55",
          "89",
          "?",
        ];
        await performVotingFlow(
          hostPage,
          participantPage,
          hostPom,
          participantPom,
          voteButtons,
          "13",
          "5",
        );
      });
    });

    test.describe("Tシャツサイズ", () => {

      test("投票フロー", async ({ page: hostPage }) => {
        const hostPom = new SessionPagePom(hostPage);
        const participantUserName = "参加者ユーザー";
        const participantPage = await hostPom.joinAsParticipant(
          participantUserName,
        );
        const participantPom = new SessionPagePom(participantPage);

        const voteButtons = ["XS", "S", "M", "L", "XL", "?"];
        await performVotingFlow(
          hostPage,
          participantPage,
          hostPom,
          participantPom,
          voteButtons,
          "L",
          "M",
        );
      });
    });

    test.describe("2の累乗", () => {

      test("投票フロー", async ({ page: hostPage }) => {
        const hostPom = new SessionPagePom(hostPage);
        const participantUserName = "参加者ユーザー";
        const participantPage = await hostPom.joinAsParticipant(
          participantUserName,
        );
        const participantPom = new SessionPagePom(participantPage);

        const voteButtons = [
          "1",
          "2",
          "4",
          "8",
          "16",
          "32",
          "64",
          "128",
          "256",
          "512",
          "1024",
          "?",
        ];
        await performVotingFlow(
          hostPage,
          participantPage,
          hostPom,
          participantPom,
          voteButtons,
          "32",
          "8",
        );
      });
    });
  });
});
