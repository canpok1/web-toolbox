import crypto from "node:crypto";
import { expect, test } from "@playwright/test";

test.describe("セッション参加画面", () => {
  test.beforeEach(async ({ page }) => {
    await page.goto("/planning-poker/sessions/join");
  });

  test("タイトルがあること", async ({ page }) => {
    await expect(page).toHaveTitle(/プランニングポーカー/);
    await expect(page).toHaveTitle(/Web Toolbox/);
  });

  test("セッションIDを入力できること", async ({ page }) => {
    await expect(page.getByLabel("セッションID")).toBeVisible();
  });

  test("参加者名を入力できること", async ({ page }) => {
    await expect(page.getByLabel("あなたの名前")).toBeVisible();
  });

  test("「セッションに参加」ボタンがあること", async ({ page }) => {
    await expect(
      page.getByRole("button", { name: "セッションに参加" }),
    ).toBeVisible();
  });

  test("戻るボタンがあること", async ({ page }) => {
    await expect(page.getByRole("link", { name: "戻る" })).toBeVisible();
    await expect(page.getByRole("link", { name: "戻る" })).toHaveAttribute(
      "href",
      "/planning-poker",
    );
  });

  test("セッションIDと名前が入力されていないときはセッション参加ボタンが押せないこと", async ({
    page,
  }) => {
    const joinSessionButton = page.getByRole("button", {
      name: "セッションに参加",
    });
    await expect(joinSessionButton).toBeDisabled();
  });

  test("セッションIDと名前が入力されたときはセッション参加ボタンが押せること", async ({
    page,
  }) => {
    await page.getByLabel("セッションID").fill("test-session-id");
    await page.getByLabel("名前").fill("テストユーザー");
    const joinSessionButton = page.getByRole("button", {
      name: "セッションに参加",
    });
    await expect(joinSessionButton).toBeEnabled();
  });

  test("存在しないセッションIDを入力してセッション参加ボタンを押すと画面遷移せずにエラーが表示されること", async ({
    page,
  }) => {
    const invalidSessionId = crypto.randomUUID();
    await page.getByLabel("セッションID").fill(invalidSessionId);
    await page.getByLabel("名前").fill("テストユーザー");
    await page.getByRole("button", { name: "セッションに参加" }).click();
    await expect(page.locator(".alert")).toBeVisible();
  });
});
