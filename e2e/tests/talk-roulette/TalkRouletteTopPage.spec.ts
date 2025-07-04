import { test, expect } from "@playwright/test";

test.describe("トークルーレット画面", () => {
  test.beforeEach(async ({ page }) => {
    await page.goto("http://localhost:3000/talk-roulette");
  });

  // TODO: タイトル「今日のトークテーマ」が表示されていること
  // TODO: トークテーマが表示されていること
  // TODO: 「いいね」ボタンが表示されていること
  // TODO: 「うーん」ボタンが表示されていること
  // TODO: ジャンルセレクターが表示されていること
  // TODO: 「次のテーマ」ボタンが表示されていること
  // TODO: 「新しいテーマを投稿」リンクが表示されていること

  test.describe("「いいね」ボタンの機能", () => {
    // TODO: 「いいね」ボタンをクリックすると「良いテーマですね！」と表示されること
    // TODO: 「いいね」ボタンを再度クリックするとメッセージが消えること
  });

  test.describe("「うーん」ボタンの機能", () => {
    // TODO: 「うーん」ボタンをクリックすると「テーマを変更しますね。」と表示されること
    // TODO: 「うーん」ボタンを再度クリックするとメッセージが消えること
  });

  test.describe("ジャンルセレクターの機能", () => {
    // TODO: ジャンルを変更するとトークテーマが更新されること
  });

  test.describe("「次のテーマ」ボタンの機能", () => {
    // TODO: 「次のテーマ」ボタンをクリックするとトークテーマが変更されること
  });

  test.describe("「新しいテーマを投稿」リンクのナビゲーション", () => {
    // TODO: 「新しいテーマを投稿」リンクをクリックすると正しいページに遷移すること
  });
});
