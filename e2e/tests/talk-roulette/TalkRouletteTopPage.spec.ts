import { expect, test } from "@playwright/test";
import { TalkRouletteTopPagePom } from "../../pom/talk-roulette/TalkRouletteTopPagePom";

let talkRoulettePage: TalkRouletteTopPagePom;

test.beforeEach(async ({ page }) => {
  talkRoulettePage = new TalkRouletteTopPagePom(page);
  await talkRoulettePage.goto();
  await talkRoulettePage.waitForTalkThemeVisible();
});

test("トークルーレットページにアクセスできること", async ({ page }) => {
  await expect(page).toHaveURL(/talk-roulette/);
  await expect(page.getByText("今日のトークテーマ")).toBeVisible();
});

test("初期表示でテーマが表示されていること", async () => {
  await expect(talkRoulettePage.talkTheme).toBeVisible();
  await expect(talkRoulettePage.talkTheme).not.toBeEmpty();
  await expect(talkRoulettePage.talkTheme).not.toHaveText(
    "テーマを読み込み中...",
  );
});

test("「新しいテーマ」ボタンをクリックするとテーマが変更されること", async () => {
  const initialTheme = await talkRoulettePage.getTalkThemeText();
  await talkRoulettePage.clickNewThemeButton();
  await expect(talkRoulettePage.talkTheme).not.toHaveText(
    initialTheme as string,
  );
});

test("「良いね」ボタンをクリックするとフィードバックメッセージが表示されること", async () => {
  await expect(talkRoulettePage.goodThemeButton).toBeVisible();
  await talkRoulettePage.clickGoodThemeButton();
  await expect(talkRoulettePage.feedbackMessage).toHaveText(
    "良いテーマですね！",
  );
});

test("「良くないね」ボタンをクリックするとフィードバックメッセージが表示されること", async () => {
  await expect(talkRoulettePage.badThemeButton).toBeVisible();
  await talkRoulettePage.clickBadThemeButton();
  await expect(talkRoulettePage.feedbackMessage).toHaveText(
    "テーマを変更しますね。",
  );
});

test("ジャンルを選択するとテーマが変更されること", async () => {
  const initialTheme = await talkRoulettePage.getTalkThemeText();
  await talkRoulettePage.selectGenre("hobby");
  await expect(talkRoulettePage.talkTheme).not.toHaveText(
    initialTheme as string,
  );
});
