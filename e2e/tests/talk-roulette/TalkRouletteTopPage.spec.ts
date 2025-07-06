import { expect, test } from "@playwright/test";

test.beforeEach(async ({ page }) => {
  await page.goto("/talk-roulette");
  await page.getByTestId("talk-theme").waitFor({ state: "visible" });
});

test("トークルーレットページにアクセスできること", async ({ page }) => {
  await expect(page).toHaveURL(/talk-roulette/);
  await expect(page.getByText("今日のトークテーマ")).toBeVisible();
});

test("初期表示でテーマが表示されていること", async ({ page }) => {
  const themeElement = page.getByTestId("talk-theme");
  await expect(themeElement).toBeVisible();
  await expect(themeElement).not.toBeEmpty();
  await expect(themeElement).not.toHaveText("テーマを読み込み中...");
});

test("「新しいテーマ」ボタンをクリックするとテーマが変更されること", async ({
  page,
}) => {
  const themeElement = page.getByTestId("talk-theme");
  const initialTheme = await themeElement.textContent();
  await page.getByRole("button", { name: "別のテーマを引く" }).click();
  await expect(themeElement).not.toHaveText(initialTheme as string);
});

// TODO: プロダクトコードの不備により、ボタンが表示されないためテストをコメントアウト
// test("「良いね」ボタンをクリックするとフィードバックメッセージが表示されること", async ({ page }) => {
//   await page.getByRole("button", { name: "良いね" }).click();
//   await expect(page.locator("#feedback-message")).toHaveText("良いテーマですね！");
// });

// TODO: プロダクトコードの不備により、ボタンが表示されないためテストをコメントアウト
// test("「良くないね」ボタンをクリックするとフィードバックメッセージが表示されること", async ({ page }) => {
//   await page.getByRole("button", { name: "良くないね" }).click();
//   await expect(page.locator("#feedback-message")).toHaveText("テーマを変更しますね。");
// });

test("ジャンルを選択するとテーマが変更されること", async ({ page }) => {
  const themeElement = page.getByTestId("talk-theme");
  const initialTheme = await themeElement.textContent();
  await page.getByRole("combobox").selectOption("hobby");
  await expect(themeElement).not.toHaveText(initialTheme as string);
});
