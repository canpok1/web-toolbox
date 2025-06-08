
import { test, expect } from '@playwright/test';

test('has title', async ({ page }) => {
    await page.goto('/planning-poker');

    await expect(page).toHaveTitle(/プランニングポーカー/);
    await expect(page).toHaveTitle(/Web Toolbox/);
});
