import { Locator, Page } from "@playwright/test";

export class TalkRouletteTopPagePom {
  readonly page: Page;
  readonly talkTheme: Locator;
  readonly newThemeButton: Locator;
  readonly goodThemeButton: Locator;
  readonly badThemeButton: Locator;
  readonly feedbackMessage: Locator;
  readonly genreCombobox: Locator;

  constructor(page: Page) {
    this.page = page;
    this.talkTheme = page.getByTestId("talk-theme");
    this.newThemeButton = page.getByRole("button", { name: "別のテーマを引く" });
    this.goodThemeButton = page.getByRole("button", { name: "良いテーマ" });
    this.badThemeButton = page.getByRole("button", { name: "悪いテーマ" });
    this.feedbackMessage = page.getByTestId("feedback-message");
    this.genreCombobox = page.getByRole("combobox");
  }

  async goto() {
    await this.page.goto("/talk-roulette");
  }

  async waitForTalkThemeVisible() {
    await this.talkTheme.waitFor({ state: "visible" });
  }

  async getTalkThemeText() {
    return this.talkTheme.textContent();
  }

  async clickNewThemeButton() {
    await this.newThemeButton.click();
  }

  async clickGoodThemeButton() {
    await this.goodThemeButton.click();
  }

  async clickBadThemeButton() {
    await this.badThemeButton.click();
  }

  async selectGenre(genre: string) {
    await this.genreCombobox.selectOption(genre);
  }
}