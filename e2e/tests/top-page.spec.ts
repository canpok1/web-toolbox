import { test, expect } from '@playwright/test';

test.beforeEach(async ({ page }) => {
    await page.goto('/');
});

test('タイトルが設定されていること', async ({ page }) => {
    await expect(page).toHaveTitle(/Web Toolbox/);
});

test('プランニングポーカーへのリンクが存在すること', async ({ page }) => {
    await expect(page.getByRole('link', { name: 'プランニングポーカー' })).toBeVisible();
});

test('トークルーレットへのリンクが存在すること', async ({ page }) => {
    await expect(page.getByRole('link', { name: 'トークルーレット' })).toBeVisible();
});

test('プランニングポーカーへのリンクをクリックできること', async ({ page }) => {
    await page.getByRole('link', { name: 'プランニングポーカー' }).click();
    await expect(page).toHaveURL(/planning-poker/);
});

test('トークルーレットへのリンクをクリックできること', async ({ page }) => {
    await page.getByRole('link', { name: 'トークルーレット' }).click();
    await expect(page).toHaveURL(/talk-roulette/);
});
