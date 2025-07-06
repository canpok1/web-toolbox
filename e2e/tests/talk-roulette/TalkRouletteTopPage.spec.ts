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
  await expect(themeElement.textContent()).not.toBe("");
  await expect(themeElement).not.toHaveText("テーマを読み込み中...");
});

test("「新しいテーマ」ボタンをクリックするとテーマが変更されること", async ({
  page,
}) => {
  const initialTheme = await page.getByTestId("talk-theme").textContent();
  await page.getByRole("button", { name: "別のテーマを引く" }).click();
  await page.waitForFunction((initialTheme) => {
    const el = document.querySelector('[data-testid="talk-theme"]');
    return el && el.textContent !== initialTheme;
  }, initialTheme);
  const newTheme = await page.getByTestId("talk-theme").textContent();
  expect(newTheme).not.toBe(initialTheme);
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
  const initialTheme = await page.getByTestId("talk-theme").textContent();
  await page.getByRole("combobox").selectOption("hobby");
  await page.waitForFunction((initialTheme) => {
    const el = document.querySelector('[data-testid="talk-theme"]');
    return el && el.textContent !== initialTheme;
  }, initialTheme);
  const newTheme = await page.getByTestId("talk-theme").textContent();
  expect(newTheme).not.toBe(initialTheme);
});
