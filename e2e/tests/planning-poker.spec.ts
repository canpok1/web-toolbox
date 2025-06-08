import { test, expect } from '@playwright/test';

test.describe('トップページ', () => {
    test.beforeEach(async ({ page }) => {
        await page.goto('/planning-poker');
    });

    test('タイトルがあること', async ({ page }) => {
        await expect(page).toHaveTitle(/プランニングポーカー/);
        await expect(page).toHaveTitle(/Web Toolbox/);
    });

    test('「セッションを作成」ボタンのリンク先が/planning-poker/sessions/createであること', async ({ page }) => {
        await expect(page.getByRole('link', { name: 'セッションを作成' })).toHaveAttribute('href', '/planning-poker/sessions/create');
    });

    test('「セッションに参加」ボタンのリンク先が/planning-poker/sessions/joinであること', async ({ page }) => {
        await expect(page.getByRole('link', { name: 'セッションに参加' })).toHaveAttribute('href', '/planning-poker/sessions/join');
    });
});

test.describe('セッション作成画面', () => {
    test.beforeEach(async ({ page }) => {
        await page.goto('/planning-poker/sessions/create');
    });

    test('タイトルがあること', async ({ page }) => {
        await expect(page).toHaveTitle(/プランニングポーカー/);
        await expect(page).toHaveTitle(/Web Toolbox/);
    });

    test('スケールを選択できること', async ({ page }) => {
        await expect(page.getByLabel('スケール')).toBeVisible();
    });

    test('スケールの選択肢がフィボナッチ・Tシャツサイズ・2の累乗であること', async ({ page }) => {
        const scaleSelect = page.getByLabel('スケール');
        await expect(scaleSelect).toContainText('フィボナッチ');
        await expect(scaleSelect).toContainText('Tシャツサイズ');
        await expect(scaleSelect).toContainText('2の累乗');
    });

    test('参加者名を入力できること', async ({ page }) => {
        await expect(page.getByLabel('あなたの名前')).toBeVisible();
    });

    test('「セッションを作成」ボタンがあること', async ({ page }) => {
        await expect(page.getByRole('button', { name: 'セッションを作成' })).toBeVisible();
    });

    test('戻るボタンがあること', async ({ page }) => {
        await expect(page.getByRole('link', { name: '戻る' })).toBeVisible();
        await expect(page.getByRole('link', { name: '戻る' })).toHaveAttribute('href', '/planning-poker');
    });

    test('名前が入力されていないときはセッション作成ボタンが押せないこと', async ({ page }) => {
        const createSessionButton = page.getByRole('button', { name: 'セッションを作成' });
        await expect(createSessionButton).toBeDisabled();
    });

    test('名前が入力されたときはセッション作成ボタンが押せること', async ({ page }) => {
        await page.getByLabel('あなたの名前').fill('テストユーザー');
        const createSessionButton = page.getByRole('button', { name: 'セッションを作成' });
        await expect(createSessionButton).toBeEnabled();
    });

    test('セッション作成ボタンを押すとセッションページに遷移すること', async ({ page }) => {
        await page.getByLabel('あなたの名前').fill('テストユーザー');
        await page.getByRole('button', { name: 'セッションを作成' }).click();
        await expect(page).toHaveURL(/planning-poker\/sessions\/.*/);
    });
});
