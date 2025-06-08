import { expect, test } from "@playwright/test";

test.describe("トップページ", () => {
  test.beforeEach(async ({ page }) => {
    await page.goto("/planning-poker");
  });

  test("タイトルがあること", async ({ page }) => {
    await expect(page).toHaveTitle(/プランニングポーカー/);
    await expect(page).toHaveTitle(/Web Toolbox/);
  });

  test("「セッションを作成」ボタンのリンク先が/planning-poker/sessions/createであること", async ({
    page,
  }) => {
    await expect(
      page.getByRole("link", { name: "セッションを作成" }),
    ).toHaveAttribute("href", "/planning-poker/sessions/create");
  });

  test("「セッションに参加」ボタンのリンク先が/planning-poker/sessions/joinであること", async ({
    page,
  }) => {
    await expect(
      page.getByRole("link", { name: "セッションに参加" }),
    ).toHaveAttribute("href", "/planning-poker/sessions/join");
  });
});
