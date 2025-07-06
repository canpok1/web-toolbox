import { expect, test } from "@playwright/test";
import { TopPagePom } from "../pom/TopPage";

let topPage: TopPagePom;

test.beforeEach(async ({ page }) => {
  topPage = new TopPagePom(page);
  await topPage.goto();
});

test("タイトルが設定されていること", async () => {
  await topPage.expectTitleToContain("Web Toolbox");
});

test("プランニングポーカーへのリンクが存在すること", async () => {
  await expect(topPage.planningPokerLink).toBeVisible();
});

test("トークルーレットへのリンクが存在すること", async () => {
  await expect(topPage.talkRouletteLink).toBeVisible();
});

test("プランニングポーカーへのリンクをクリックできること", async () => {
  await topPage.clickPlanningPokerLink();
  await topPage.waitForURL(/planning-poker/);
});

test("トークルーレットへのリンクをクリックできること", async () => {
  await topPage.clickTalkRouletteLink();
  await topPage.waitForURL(/talk-roulette/);
  await expect(
    topPage.getByRole("heading", { name: "今日のトークテーマ" }),
  ).toBeVisible();
});
