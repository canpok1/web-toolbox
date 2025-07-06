import { expect, test } from "@playwright/test";
import { CreateSessionPagePom } from "../../pom/planning-poker/CreateSessionPage";

test.describe("セッション作成画面", () => {
  let createSessionPage: CreateSessionPagePom;

  test.beforeEach(async ({ page }) => {
    createSessionPage = new CreateSessionPagePom(page);
    await createSessionPage.goto();
  });

  test("タイトルがあること", async ({ page }) => {
    await expect(page).toHaveTitle(/プランニングポーカー/);
    await expect(page).toHaveTitle(/Web Toolbox/);
  });

  test("スケールを選択できること", async () => {
    await expect(createSessionPage.scaleSelect).toBeVisible();
  });

  test("スケールの選択肢がフィボナッチ・Tシャツサイズ・2の累乗であること", async () => {
    const scaleSelect = createSessionPage.scaleSelect;
    await expect(scaleSelect).toContainText("フィボナッチ");
    await expect(scaleSelect).toContainText("Tシャツサイズ");
    await expect(scaleSelect).toContainText("2の累乗");
  });

  test("参加者名を入力できること", async () => {
    await expect(createSessionPage.yourNameInput).toBeVisible();
  });

  test("「セッションを作成」ボタンがあること", async () => {
    await expect(createSessionPage.createSessionButton).toBeVisible();
  });

  test("戻るボタンがあること", async () => {
    const backLink = createSessionPage.backLink;
    await expect(backLink).toBeVisible();
    await expect(backLink).toHaveAttribute(
      "href",
      "/planning-poker",
    );
  });

  test("名前が入力されていないときはセッション作成ボタンが押せないこと", async () => {
    await expect(createSessionPage.createSessionButton).toBeDisabled();
  });

  test("名前が入力されたときはセッション作成ボタンが押せること", async () => {
    await createSessionPage.fillYourName("テストユーザー");
    await expect(createSessionPage.createSessionButton).toBeEnabled();
  });

  test("セッション作成ボタンを押すとセッションページに遷移すること", async ({
    page,
  }) => {
    await createSessionPage.createSession("テストユーザー", "fibonacci");
    await expect(page).toHaveURL(/planning-poker\/sessions\/.*/);
  });
});
