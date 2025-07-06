import { expect, test } from "@playwright/test";
import { PlanningPokerTopPagePom } from "../../pom/planning-poker/PlanningPokerTopPagePom";

test.describe("トップページ", () => {
  let planningPokerTopPage: PlanningPokerTopPagePom;
  test.beforeEach(async ({ page }) => {
    planningPokerTopPage = new PlanningPokerTopPagePom(page);
    await planningPokerTopPage.goto();
  });

  test("タイトルがあること", async () => {
    await planningPokerTopPage.expectTitleToContain("プランニングポーカー");
    await planningPokerTopPage.expectTitleToContain("Web Toolbox");
  });

  test("「セッションを作成」ボタンのリンク先が/planning-poker/sessions/createであること", async () => {
    await expect(planningPokerTopPage.createSessionLink).toHaveAttribute(
      "href",
      "/planning-poker/sessions/create",
    );
  });

  test("「セッションに参加」ボタンのリンク先が/planning-poker/sessions/joinであること", async () => {
    await expect(planningPokerTopPage.joinSessionLink).toHaveAttribute(
      "href",
      "/planning-poker/sessions/join",
    );
  });
});
