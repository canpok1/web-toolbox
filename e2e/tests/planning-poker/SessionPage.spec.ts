import { type Page, expect, test } from "@playwright/test";
import { CreateSessionPagePom } from "../../pom/planning-poker/CreateSessionPagePom";
import { SessionPagePom } from "../../pom/planning-poker/SessionPagePom";

const fibonacciVoteButtons = [
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

const tShirtVoteButtons = ["XS", "S", "M", "L", "XL", "?"];

const powerOfTwoVoteButtons = [
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

async function checkVoteButtonsVisibility(
  pom: SessionPagePom, // PageからSessionPagePomに変更
  buttons: string[],
) {
  for (const voteButton of buttons) {
    await expect(pom.getVoteButton(voteButton)).toBeVisible(); // pom.getVoteButtonを使用
  }
}

async function performVotingFlow(
  _hostPage: Page,
  _participantPage: Page,
  hostPom: SessionPagePom,
  participantPom: SessionPagePom,
  voteButtons: string[],
  hostVote: string,
  participantVote: string,
) {
  // 参加者ユーザー画面に投票開始ボタンが表示されないことを確認
  await expect(participantPom.startVoteButton).not.toBeVisible(); // pom.startVoteButtonを使用
  // ホストユーザー画面に投票開始ボタンが表示されることを確認
  await expect(hostPom.startVoteButton).toBeVisible(); // pom.startVoteButtonを使用

  // ホストが投票を開始する
  await hostPom.clickStartVoteButton();

  // 画面表示確認
  // 参加者ユーザー画面に投票ボタンが表示されることを確認
  await checkVoteButtonsVisibility(participantPom, voteButtons); // pomを渡す
  await checkVoteButtonsVisibility(hostPom, voteButtons); // pomを渡す

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
  await expect(participantPom.startVoteButton).not.toBeVisible(); // pom.startVoteButtonを使用
  // ホストユーザー画面に投票開始ボタンが表示されることを確認
  await expect(hostPom.startVoteButton).toBeVisible(); // pom.startVoteButtonを使用

  // ホストが投票を開始する
  await hostPom.clickStartVoteButton();

  // 画面表示確認
  // 参加者ユーザー画面に投票ボタンが表示されることを確認
  await checkVoteButtonsVisibility(participantPom, voteButtons); // pomを渡す
  await checkVoteButtonsVisibility(hostPom, voteButtons); // pomを渡す
}

function createVotingFlowTests(
  scaleName: string,
  scaleType: "fibonacci" | "t-shirt" | "power-of-two",
  voteButtons: string[],
  hostVote: string,
  participantVote: string,
  hostUserName: string,
  joinAsParticipantAndGetPom: (
    hostPom: SessionPagePom,
    participantUserName: string,
  ) => Promise<{ participantPage: Page; participantPom: SessionPagePom }>,
) {
  test.describe(scaleName, () => {
    test.beforeEach(async ({ page }) => {
      const pom = new CreateSessionPagePom(page);
      await pom.goto();
      await pom.createSession(hostUserName, scaleType);
    });

    test("投票フロー", async ({ page: hostPage }) => {
      const hostPom = new SessionPagePom(hostPage);
      const participantUserName = "参加者ユーザー";
      const { participantPage, participantPom } =
        await joinAsParticipantAndGetPom(hostPom, participantUserName);

      await performVotingFlow(
        hostPage,
        participantPage,
        hostPom,
        participantPom,
        voteButtons,
        hostVote,
        participantVote,
      );
    });
  });
}

test.describe("セッション画面", () => {
  async function joinAsParticipantAndGetPom(
    hostPom: SessionPagePom,
    participantUserName: string,
  ): Promise<{ participantPage: Page; participantPom: SessionPagePom }> {
    const participantPage =
      await hostPom.joinAsParticipant(participantUserName);
    const participantPom = new SessionPagePom(participantPage);
    return { participantPage, participantPom };
  }
  const hostUserName = "ホストユーザー";

  test.describe("参加者一覧と招待リンク", () => {
    test.describe("ホストのみ", () => {
      test.beforeEach(async ({ page }) => {
        const pom = new CreateSessionPagePom(page);
        await pom.goto();
        await pom.createSession(hostUserName, "fibonacci");
      });
      test("表示内容が正しいこと", async ({ page: hostPage }) => {
        const hostPom = new SessionPagePom(hostPage);

        // ホストユーザー画面に自分の名前が表示されるか確認
        await expect(hostPom.getNameElement(hostUserName)).toBeVisible();
        await expect(hostPom.getParticipantNameElement(1)).toBeVisible();
        await expect(hostPom.getNameElement(hostUserName)).toBeVisible();
        await expect(hostPom.getInviteUrlCopyButton()).toBeVisible();

        const pom = new SessionPagePom(hostPage);
        await pom.clickInviteUrlButton();
        const inviteLinkValue = await pom.inviteLink.getAttribute("href");
        await expect(inviteLinkValue).toMatch(
          /planning-poker\/sessions\/join\?id=.*/,
        );
      });
    });

    test.describe("参加者が複数", () => {
      const participantUserName = "参加者ユーザー";
      test.beforeEach(async ({ page }) => {
        const pom = new CreateSessionPagePom(page);
        await pom.goto();
        await pom.createSession(hostUserName, "fibonacci");
      });
      test("表示内容が正しいこと", async ({ page: hostPage }) => {
        const hostPom = new SessionPagePom(hostPage);
        const { participantPage, participantPom } =
          await joinAsParticipantAndGetPom(hostPom, participantUserName);

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
        const participantInviteLinkValue =
          await participantPom.inviteLink.getAttribute("href");

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
    createVotingFlowTests(
      "フィボナッチ",
      "fibonacci",
      fibonacciVoteButtons,
      "13",
      "5",
      hostUserName,
      joinAsParticipantAndGetPom,
    );
    createVotingFlowTests(
      "Tシャツサイズ",
      "t-shirt",
      tShirtVoteButtons,
      "L",
      "M",
      hostUserName,
      joinAsParticipantAndGetPom,
    );
    createVotingFlowTests(
      "2の累乗",
      "power-of-two",
      powerOfTwoVoteButtons,
      "32",
      "8",
      hostUserName,
      joinAsParticipantAndGetPom,
    );
  });
});
