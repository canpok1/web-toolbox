import type { Page } from "@playwright/test";

export class TalkRouletteTopPagePom {
  private readonly page: Page;

  constructor(page: Page) {
    this.page = page;
  }

  // TODO: トークテーマのテキストを取得するメソッド
  // 例: async getTalkTheme(): Promise<string | null>

  // TODO: 「いいね」ボタンをクリックするメソッド
  // 例: async clickLikeButton(): Promise<void>

  // TODO: 「うーん」ボタンをクリックするメソッド
  // 例: async clickDislikeButton(): Promise<void>

  // TODO: フィードバックメッセージのテキストを取得するメソッド
  // 例: async getFeedbackMessage(): Promise<string | null>

  // TODO: ジャンルを選択するメソッド
  // 例: async selectGenre(genre: string): Promise<void>

  // TODO: 「次のテーマ」ボタンをクリックするメソッド
  // 例: async clickNextThemeButton(): Promise<void>

  // TODO: 「新しいテーマを投稿」リンクをクリックするメソッド
  // 例: async clickSubmitNewThemeLink(): Promise<void>
}
