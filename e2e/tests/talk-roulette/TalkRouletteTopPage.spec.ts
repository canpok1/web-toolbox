import { test, expect } from "@playwright/test";
import { TalkRouletteTopPagePom } from "../../pom/talk-roulette/TalkRouletteTopPage.pom.ts";

test.describe("トークルーレット画面", () => {
  let talkRouletteTopPagePom: TalkRouletteTopPagePom;

  test.beforeEach(async ({ page }) => {
    await page.goto("http://localhost:3000/talk-roulette");
    talkRouletteTopPagePom = new TalkRouletteTopPagePom(page);
  });

  // TODO: POMを利用してタイトル「今日のトークテーマ」が表示されていることを確認する
  // TODO: POMを利用してトークテーマが表示されていることを確認する
  // TODO: POMを利用して「いいね」ボタンが表示されていることを確認する
  // TODO: POMを利用して「うーん」ボタンが表示されていることを確認する
  // TODO: POMを利用してジャンルセレクターが表示されていることを確認する
  // TODO: POMを利用して「次のテーマ」ボタンが表示されていることを確認する
  // TODO: POMを利用して「新しいテーマを投稿」リンクが表示されていることを確認する

  test.describe("「いいね」ボタンの機能", () => {
    // TODO: POMのメソッドを使って「いいね」ボタンをクリックすると「良いテーマですね！」と表示されること
    // TODO: POMのメソッドを使って「いいね」ボタンを再度クリックするとメッセージが消えること
  });

  test.describe("「うーん」ボタンの機能", () => {
    // TODO: POMのメソッドを使って「うーん」ボタンをクリックすると「テーマを変更しますね。」と表示されること
    // TODO: POMのメソッドを使って「うーん」ボタンを再度クリックするとメッセージが消えること
  });

  test.describe("ジャンルセレクターの機能", () => {
    // TODO: POMのメソッドを使ってジャンルを変更するとトークテーマが更新されること
  });

  test.describe("「次のテーマ」ボタンの機能", () => {
    // TODO: POMのメソッドを使って「次のテーマ」ボタンをクリックするとトークテーマが変更されること
  });

  test.describe("「新しいテーマを投稿」リンクのナビゲーション", () => {
    // TODO: POMのメソッドを使って「新しいテーマを投稿」リンクをクリックすると正しいページに遷移すること
  });
});
