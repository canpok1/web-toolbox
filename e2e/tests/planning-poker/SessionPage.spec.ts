import { test, Page, expect } from '@playwright/test';

test.describe('セッション画面', () => {
    const hostUserName = 'ホストユーザー';

    test.beforeEach(async ({ page, context }) => {
        await page.goto('/planning-poker/sessions/create');
        await page.getByLabel('あなたの名前').fill(hostUserName);
        await page.getByRole('button', { name: 'セッションを作成' }).click();
    });

    test.describe('参加者一覧と招待リンク', () => {
        test.describe('ホストのみ', () => {
            test('表示内容が正しいこと', async ({ page: hostPage, context }) => {
                await expect(hostPage.getByText(`あなたの名前: ${hostUserName}`)).toBeVisible();
                await expect(hostPage.getByText(/参加者 \(1名\):\s*/)).toBeVisible();
                await expect(hostPage.getByText(hostUserName, { exact: true })).toBeVisible();
                await expect(hostPage.getByRole('button', { name: '参加ページのURLをコピー' })).toBeVisible();

                await hostPage.getByRole('button', { name: '招待URL/QRコード' }).click();

                const inviteLink = hostPage.locator('a[href*="/planning-poker/sessions/join?id="]');
                await expect(inviteLink).toBeVisible();
                const inviteLinkValue = await inviteLink.getAttribute('href');
                await expect(inviteLinkValue).toMatch(/planning-poker\/sessions\/join\?id=.*/);
            });
        });
        test.describe('参加者が複数', () => {
            test('表示内容が正しいこと', async ({ }) => {
                // TODO 参加者ユーザー画面に自分の名前が表示されるか確認
                // TODO 参加者ユーザー画面に参加者一覧が表示されるか確認
                // TODO 参加者ユーザー画面で招待リンクをコピーできることを確認
                // TODO 参加者ユーザー画面の招待リンクが正しいことを確認
                // TODO ホストユーザー画面に自分の名前が表示されるか確認
                // TODO ホストユーザー画面に参加者一覧が表示されるか確認
                // TODO ホストユーザー画面で招待リンクをコピーできることを確認
                // TODO ホストユーザー画面の招待リンクが正しいことを確認
            });
        });
    });

    test.describe('ホスト用ボタンと投票ボタンと投票結果', () => {
        test.describe('フィボナッチ', () => {
            test('投票開始→投票→投票公開→投票開始', async ({ }) => {
                // 画面表示確認
                // TODO 参加者ユーザー画面に投票開始ボタンが表示されないことを確認
                // TODO 参加者ユーザー画面に投票結果が表示されないことを確認
                // TODO ホストユーザー画面に投票開始ボタンが表示されることを確認
                // TODO ホストユーザー画面に投票結果が表示されないことを確認

                // TODO ホストが投票を開始する

                // 画面表示確認
                // TODO 参加者ユーザー画面に投票ボタンが表示されることを確認
                // TODO 参加者ユーザー画面に投票状況が表示されることを確認
                // TODO ホストユーザー画面に投票ボタンが表示されることを確認
                // TODO ホストユーザー画面に投票状況が表示されることを確認

                // TODO 参加者ユーザーが投票する

                // 画面表示確認
                // TODO 参加者ユーザー画面に投票状況が表示されることを確認
                // TODO ホストユーザー画面に投票状況が表示されることを確認

                // TODO 参加者ユーザーが投票内容を変更する

                // 画面表示確認
                // TODO 参加者ユーザー画面に投票状況が表示されることを確認
                // TODO ホストユーザー画面に投票状況が表示されることを確認

                // TODO ホストユーザーが投票する

                // 画面表示確認
                // TODO 参加者ユーザー画面に投票状況が表示されることを確認
                // TODO ホストユーザー画面に投票状況が表示されることを確認

                // TODO ホストが投票を公開する

                // 画面表示確認
                // TODO 参加者ユーザー画面に投票開始ボタンが表示されないことを確認
                // TODO 参加者ユーザー画面に投票結果が表示されることを確認
                // TODO ホストユーザー画面に投票開始ボタンが表示されないことを確認
                // TODO ホストユーザー画面に投票結果が表示されることを確認

                // TODO ホストが投票を開始する

                // 画面表示確認
                // TODO 参加者ユーザー画面に投票ボタンが表示されることを確認
                // TODO 参加者ユーザー画面に投票状況が表示されることを確認
                // TODO ホストユーザー画面に投票ボタンが表示されることを確認
                // TODO ホストユーザー画面に投票状況が表示されることを確認
            });
        });

        test.describe('Tシャツサイズ', () => {
            test('投票開始→投票→投票公開→投票開始', async ({ }) => {
                // TODO フィボナッチと同様のテストを実施
            });
        });

        test.describe('2の累乗', () => {
            test('投票開始→投票→投票公開→投票開始', async ({ }) => {
                // TODO フィボナッチと同様のテストを実施
            });
        });
    });
});
