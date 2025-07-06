import crypto from "node:crypto";
import { expect, test } from "@playwright/test";
import { JoinSessionPagePom } from "../../pom/planning-poker/JoinSessionPagePom";

test.describe("セッション参加画面", () => {
  let joinSessionPage: JoinSessionPagePom;

  test.beforeEach(async ({ page }) => {
    joinSessionPage = new JoinSessionPagePom(page);
    await joinSessionPage.goto();
  });

  test("タイトルがあること", async ({ page }) => {
    await expect(page).toHaveTitle(/プランニングポーカー/);
    await expect(page).toHaveTitle(/Web Toolbox/);
  });

  test("セッションIDを入力できること", async () => {
    await expect(joinSessionPage.sessionIdInput).toBeVisible();
  });

  test("参加者名を入力できること", async () => {
    await expect(joinSessionPage.yourNameInput).toBeVisible();
  });

  test("「セッションに参加」ボタンがあること", async () => {
    await expect(joinSessionPage.joinSessionButton).toBeVisible();
  });

  test("戻るボタンがあること", async () => {
    await expect(joinSessionPage.backLink).toBeVisible();
    await expect(joinSessionPage.backLink).toHaveAttribute(
      "href",
      "/planning-poker",
    );
  });

  test("セッションIDと名前が入力されていないときはセッション参加ボタンが押せないこと", async () => {
    await expect(joinSessionPage.joinSessionButton).toBeDisabled();
  });

  test("セッションIDと名前が入力されたときはセッション参加ボタンが押せること", async () => {
    await joinSessionPage.fillSessionId("test-session-id");
    await joinSessionPage.fillYourName("テストユーザー");
    await expect(joinSessionPage.joinSessionButton).toBeEnabled();
  });

  test("存在しないセッションIDを入力してセッション参加ボタンを押すと画面遷移せずにエラーが表示されること", async () => {
    const invalidSessionId = crypto.randomUUID();
    await joinSessionPage.fillSessionId(invalidSessionId);
    await joinSessionPage.fillYourName("テストユーザー");
    await joinSessionPage.clickJoinSessionButton();
    await expect(joinSessionPage.alertMessage).toBeVisible();
  });
});
