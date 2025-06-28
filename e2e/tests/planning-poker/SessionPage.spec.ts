import { expect, test } from "@playwright/test";

test.describe("セッション画面", () => {
  const hostUserName = "ホストユーザー";

  test.describe("参加者一覧と招待リンク", () => {
    test.beforeEach(async ({ page }) => {
      await page.goto("/planning-poker/sessions/create");
      await page.getByLabel("あなたの名前").fill(hostUserName);
      await page
        .getByRole("button", { name: "セッションを作成", exact: true })
        .click();
      await page.waitForEvent("websocket");
    });

    test.describe("ホストのみ", () => {
      test("表示内容が正しいこと", async ({ page: hostPage }) => {
        await expect(
          hostPage.getByText(`あなたの名前: ${hostUserName}`),
        ).toBeVisible();
        await expect(hostPage.getByText(/参加者 \(1名\):\s*/)).toBeVisible();
        await expect(
          hostPage.getByText(hostUserName, { exact: true }),
        ).toBeVisible();
        await expect(
          hostPage.getByRole("button", { name: "参加ページのURLをコピー" }),
        ).toBeVisible();

        await hostPage
          .getByRole("button", { name: "招待URL/QRコード", exact: true })
          .click();

        const inviteLink = hostPage.locator(
          'a[href*="/planning-poker/sessions/join?id="]',
        );
        await expect(inviteLink).toBeVisible();
        const inviteLinkValue = await inviteLink.getAttribute("href");
        await expect(inviteLinkValue).toMatch(
          /planning-poker\/sessions\/join\?id=.*/,
        );
      });
    });

    test.describe("参加者が複数", () => {
      const participantUserName = "参加者ユーザー";
      test("表示内容が正しいこと", async ({ page: hostPage }) => {
        // 招待URLを取得
        await hostPage
          .getByRole("button", { name: "招待URL/QRコード" })
          .click();
        const hostInviteLink = hostPage.locator(
          'a[href*="/planning-poker/sessions/join?id="]',
        );
        await expect(hostInviteLink).toBeVisible();
        const hostInviteLinkValue = await hostInviteLink.getAttribute("href");
        let sessionId = "";
        if (hostInviteLinkValue) {
          sessionId = hostInviteLinkValue.split("id=")[1];
        }

        // 参加者としてセッションに参加
        const participantPage = await hostPage.context().newPage();
        await participantPage.goto(
          `/planning-poker/sessions/join?id=${sessionId}`,
        );

        await participantPage
          .getByLabel("あなたの名前")
          .fill(participantUserName);
        await participantPage
          .getByRole("button", { name: "セッションに参加", exact: true })
          .click();
        await participantPage.waitForEvent("websocket");

        // 参加者ユーザー画面に自分の名前が表示されるか確認
        await expect(
          participantPage.getByText(`あなたの名前: ${participantUserName}`),
        ).toBeVisible();
        // 参加者ユーザー画面に参加者一覧が表示されるか確認
        await expect(
          participantPage.getByText(/参加者 \(2名\):\s*/),
        ).toBeVisible();
        await expect(
          participantPage.getByText(participantUserName, { exact: true }),
        ).toBeVisible();
        await expect(
          participantPage.getByText(hostUserName, { exact: true }),
        ).toBeVisible();
        // 参加者ユーザー画面で招待リンクをコピーできることを確認
        await expect(
          participantPage.getByRole("button", {
            name: "参加ページのURLをコピー",
            exact: true,
          }),
        ).toBeVisible();
        // 参加者ユーザー画面の招待リンクが正しいことを確認
        await participantPage
          .getByRole("button", { name: "招待URL/QRコード", exact: true })
          .click();
        const participantInviteLink = participantPage.locator(
          'a[href*="/planning-poker/sessions/join?id="]',
        );
        await expect(participantInviteLink).toBeVisible();
        const participantInviteLinkValue =
          await participantInviteLink.getAttribute("href");
        await expect(participantInviteLinkValue).toMatch(
          /planning-poker\/sessions\/join\?id=.*/,
        );

        await hostPage.bringToFront();

        // ホストユーザー画面に自分の名前が表示されるか確認
        await expect(
          hostPage.getByText(`あなたの名前: ${hostUserName}`),
        ).toBeVisible();
        // ホストユーザー画面に参加者一覧が表示されるか確認
        await expect(hostPage.getByText(/参加者 \(2名\):\s*/)).toBeVisible();
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
    test.describe("フィボナッチ", () => {
      test.beforeEach(async ({ page }) => {
        await page.goto("/planning-poker/sessions/create");
        await page.getByLabel("あなたの名前").fill(hostUserName);
        await page.getByRole("button", { name: "セッションを作成" }).click();
        await page.waitForEvent("websocket");
      });

      test("投票開始→投票→投票公開→投票開始", async ({ page: hostPage }) => {
        // 招待URLを取得
        await hostPage
          .getByRole("button", { name: "招待URL/QRコード" })
          .click();
        const hostInviteLink = hostPage.locator(
          'a[href*="/planning-poker/sessions/join?id="]',
        );
        const hostInviteLinkValue = await hostInviteLink.getAttribute("href");
        let sessionId = "";
        if (hostInviteLinkValue) {
          sessionId = hostInviteLinkValue.split("id=")[1];
        }

        // 画面表示確認
        const participantUserName = "参加者ユーザー";
        const participantPage = await hostPage.context().newPage();
        await participantPage.goto(
          `/planning-poker/sessions/join?id=${sessionId}`,
        );

        await participantPage
          .getByLabel("あなたの名前")
          .fill(participantUserName);
        await participantPage
          .getByRole("button", { name: "セッションに参加" })
          .click();
        await participantPage.waitForEvent("websocket");

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
        await hostPage
          .getByRole("button", { name: "投票を開始", exact: true })
          .click();

        // 画面表示確認
        // 参加者ユーザー画面に投票ボタンが表示されることを確認
        await expect(
          participantPage.getByRole("button", { name: "0", exact: true }),
        ).toBeVisible();
        await expect(
          participantPage.getByRole("button", { name: "1", exact: true }),
        ).toBeVisible();
        await expect(
          participantPage.getByRole("button", { name: "2", exact: true }),
        ).toBeVisible();
        await expect(
          participantPage.getByRole("button", { name: "3", exact: true }),
        ).toBeVisible();
        await expect(
          participantPage.getByRole("button", { name: "5", exact: true }),
        ).toBeVisible();
        await expect(
          participantPage.getByRole("button", { name: "8", exact: true }),
        ).toBeVisible();
        await expect(
          participantPage.getByRole("button", { name: "13", exact: true }),
        ).toBeVisible();
        await expect(
          participantPage.getByRole("button", { name: "21", exact: true }),
        ).toBeVisible();
        await expect(
          participantPage.getByRole("button", { name: "34", exact: true }),
        ).toBeVisible();
        await expect(
          participantPage.getByRole("button", { name: "55", exact: true }),
        ).toBeVisible();
        await expect(
          participantPage.getByRole("button", { name: "89", exact: true }),
        ).toBeVisible();
        await expect(
          participantPage.getByRole("button", { name: "?", exact: true }),
        ).toBeVisible();
        // ホストユーザー画面に投票ボタンが表示されることを確認
        await expect(
          hostPage.getByRole("button", { name: "0", exact: true }),
        ).toBeVisible();
        await expect(
          hostPage.getByRole("button", { name: "1", exact: true }),
        ).toBeVisible();
        await expect(
          hostPage.getByRole("button", { name: "2", exact: true }),
        ).toBeVisible();
        await expect(
          hostPage.getByRole("button", { name: "3", exact: true }),
        ).toBeVisible();
        await expect(
          hostPage.getByRole("button", { name: "5", exact: true }),
        ).toBeVisible();
        await expect(
          hostPage.getByRole("button", { name: "8", exact: true }),
        ).toBeVisible();
        await expect(
          hostPage.getByRole("button", { name: "13", exact: true }),
        ).toBeVisible();
        await expect(
          hostPage.getByRole("button", { name: "21", exact: true }),
        ).toBeVisible();
        await expect(
          hostPage.getByRole("button", { name: "34", exact: true }),
        ).toBeVisible();
        await expect(
          hostPage.getByRole("button", { name: "55", exact: true }),
        ).toBeVisible();
        await expect(
          hostPage.getByRole("button", { name: "89", exact: true }),
        ).toBeVisible();
        await expect(
          hostPage.getByRole("button", { name: "?", exact: true }),
        ).toBeVisible();

        // 参加者ユーザーが投票する
        await participantPage
          .getByRole("button", { name: "5", exact: true })
          .click();

        // 画面表示確認
        await expect(hostPage.locator('[data-tip="投票済み"]')).toHaveCount(1);
        await expect(
          participantPage.locator('[data-tip="投票済み"]'),
        ).toHaveCount(1);

        // ホストユーザーが投票する
        await hostPage.getByRole("button", { name: "13", exact: true }).click();

        // 画面表示確認
        await expect(hostPage.locator('[data-tip="投票済み"]')).toHaveCount(2);
        await expect(
          participantPage.locator('[data-tip="投票済み"]'),
        ).toHaveCount(2);

        // ホストが投票を公開する
        await hostPage
          .getByRole("button", { name: "投票を公開", exact: true })
          .click();

        // 画面表示確認
        // 参加者ユーザー画面に投票開始ボタンが表示されないことを確認
        await expect(
          participantPage.getByRole("button", { name: "投票を開始" }),
        ).not.toBeVisible();
        // ホストユーザー画面に投票開始ボタンが表示されることを確認
        await expect(
          hostPage.getByRole("button", { name: "投票を開始" }),
        ).toBeVisible();

        // ホストが投票を開始する
        await hostPage
          .getByRole("button", { name: "投票を開始", exact: true })
          .click();

        // 画面表示確認
        // 参加者ユーザー画面に投票ボタンが表示されることを確認
        await expect(
          participantPage.getByRole("button", { name: "0", exact: true }),
        ).toBeVisible();
        await expect(
          participantPage.getByRole("button", { name: "1", exact: true }),
        ).toBeVisible();
        await expect(
          participantPage.getByRole("button", { name: "2", exact: true }),
        ).toBeVisible();
        await expect(
          participantPage.getByRole("button", { name: "3", exact: true }),
        ).toBeVisible();
        await expect(
          participantPage.getByRole("button", { name: "5", exact: true }),
        ).toBeVisible();
        await expect(
          participantPage.getByRole("button", { name: "8", exact: true }),
        ).toBeVisible();
        await expect(
          participantPage.getByRole("button", { name: "13", exact: true }),
        ).toBeVisible();
        await expect(
          participantPage.getByRole("button", { name: "21", exact: true }),
        ).toBeVisible();
        await expect(
          participantPage.getByRole("button", { name: "34", exact: true }),
        ).toBeVisible();
        await expect(
          participantPage.getByRole("button", { name: "55", exact: true }),
        ).toBeVisible();
        await expect(
          participantPage.getByRole("button", { name: "89", exact: true }),
        ).toBeVisible();
        await expect(
          participantPage.getByRole("button", { name: "?", exact: true }),
        ).toBeVisible();
        // ホストユーザー画面に投票ボタンが表示されることを確認
        await expect(
          hostPage.getByRole("button", { name: "0", exact: true }),
        ).toBeVisible();
        await expect(
          hostPage.getByRole("button", { name: "1", exact: true }),
        ).toBeVisible();
        await expect(
          hostPage.getByRole("button", { name: "2", exact: true }),
        ).toBeVisible();
        await expect(
          hostPage.getByRole("button", { name: "3", exact: true }),
        ).toBeVisible();
        await expect(
          hostPage.getByRole("button", { name: "5", exact: true }),
        ).toBeVisible();
        await expect(
          hostPage.getByRole("button", { name: "8", exact: true }),
        ).toBeVisible();
        await expect(
          hostPage.getByRole("button", { name: "13", exact: true }),
        ).toBeVisible();
        await expect(
          hostPage.getByRole("button", { name: "21", exact: true }),
        ).toBeVisible();
        await expect(
          hostPage.getByRole("button", { name: "34", exact: true }),
        ).toBeVisible();
        await expect(
          hostPage.getByRole("button", { name: "55", exact: true }),
        ).toBeVisible();
        await expect(
          hostPage.getByRole("button", { name: "89", exact: true }),
        ).toBeVisible();
        await expect(
          hostPage.getByRole("button", { name: "?", exact: true }),
        ).toBeVisible();
      });
    });

    test.describe("Tシャツサイズ", () => {
      // TODO 実装
    });

    test.describe("2の累乗", () => {
      // TODO 実装
    });
  });
});
