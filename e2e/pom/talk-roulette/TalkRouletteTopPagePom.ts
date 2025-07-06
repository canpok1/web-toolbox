import type { Locator, Page } from "@playwright/test";

export class TalkRouletteTopPagePom {
  constructor(public readonly page: Page) {}

  public get talkTheme(): Locator {
    return this.page.getByTestId("talk-theme");
  }

  public get newThemeButton(): Locator {
    return this.page.getByRole("button", {
      name: "別のテーマを引く",
    });
  }

  public get goodThemeButton(): Locator {
    return this.page.getByRole("button", { name: "良いテーマ" });
  }

  public get badThemeButton(): Locator {
    return this.page.getByRole("button", { name: "悪いテーマ" });
  }

  public get feedbackMessage(): Locator {
    return this.page.getByTestId("feedback-message");
  }

  public get genreCombobox(): Locator {
    return this.page.getByRole("combobox");
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
