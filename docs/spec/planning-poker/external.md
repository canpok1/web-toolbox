# プランニングポーカー Web アプリケーション - 外部仕様

## 1. はじめに

本ドキュメントは、プランニングポーカーセッションを促進するために設計された Web アプリケーションの外部仕様を概説するものです。プランニングポーカーは、ソフトウェア開発やその他のプロジェクト管理の文脈において、労力や相対的な規模を見積もるための、合意形成に基づいたゲーム化された手法です。

## 2. 目標

*   リモートでプランニングポーカーセッションを実施するための、ユーザーフレンドリーなプラットフォームを提供すること。
*   チームがタスクを効率的かつ協調的に見積もることができるようにすること。
*   見積もりスケールとセッション管理において柔軟性を提供すること。
*   データのプライバシーとセキュリティを確保すること。

## 3. 対象ユーザー

*   アジャイル開発チーム
*   プロジェクトマネージャー
*   スクラムマスター
*   プロダクトオーナー
*   共同で見積もりを行うすべての人

## 4. 機能要件

### 4.1. ユーザーロール

*   **ホスト:**
    *   新しいプランニングポーカーセッションを作成する。
    *   セッションにメンバーを招待する。
    *   セッションを管理する（例：ラウンドの開始/終了、投票の公開、セッションの終了）。
    *   見積もりスケールを定義する。
    *   利用可能なスケールから見積もりを選択する。
    *   投票を送信する。
    *   参加者をキックする。
*   **参加者:**
    *   リンクまたはコードを使用して既存のプランニングポーカーセッションに参加する。
    *   利用可能なスケールから見積もりを選択する。
    *   投票を送信する。
    *   ホストが公開した後、公開された投票を見る。

### 4.2. セッション管理

*   **セッション作成:**
    *   誰でも新しいセッションを作成でき、作成するとセッションURLが自動発行される。
    *   セッションを作成したユーザーがホストとなる。
*   **セッション参加:**
    *   参加者は、ホストから提供されたセッションURLでセッションに参加できる。
    *   参加者は、自分を識別するために名前（またはニックネーム）を入力する必要がある。
    *   参加者は、セッション参加後の次のラウンドから投票に参加できる。
*   **ラウンド管理:**
    *   ホストは、新しいラウンドを開始できる。
    *   ホストと参加者は、ラウンド開始後に投票を送信できる。
    *   ホストは、ラウンドを終了し投票を公開できる。
    *   ホストは、ラウンド終了後に新しいラウンドを開始できる。
*   **セッション終了:**
    * ホストはセッションを終了することができる。
    * セッション終了後、それ以上の投票はできない。
*   **セッション有効期限:**
    * セッションの有効期限は24時間。
    * 有効期限がくるとすべての参加者は退出され、そのセッションの情報はすべて削除される。

### 4.3. 見積もり

*   **見積もりスケール:**
    *   ホストは、定義済みのスケール（フィボナッチ、Tシャツサイズ、2の累乗）から選択できる。
    *   ホストは、カスタムスケールを定義できる。
    *   カスタムスケールはカンマ区切りまたはスペース区切りで指定する。
    *   カスタムスケールに仕様できる値は数字のみで、最大11パターンである。
    *   各スケールには不明（？）が含まれる。
*   **投票:**
    *   ホストと参加者は、利用可能なスケールから 1 つの見積もりを選択できる。
    *   投票で選択された見積もりは、ホストが公開するまで非表示になる。
    *   投票済みかそうでないかは、すべての参加者に表示される。
*   **投票の公開:**
    *   ホストは、ラウンドの終了時にすべての投票を公開できる。
    *   公開された投票は、すべての参加者に表示される。
* **投票の表示:**
    * 投票は明確に表示される。
    * スケールがフィボナッチの場合、投票の平均値と中央値、不明の数が表示される。
    * スケールがTシャツサイズの場合、各サイズの投票数が表示される。
    * 2の累乗、カスタムスケールの場合に表示される内容はフィボナッチと同様。
    * 平均値は、不明以外の合計値 / 不明を除く投票数 で算出され、小数第2位まで表示される。
    * 中央値は不明を除く投票結果から選定され、投票数が偶数の場合は中央の2件が表示される。

## 5. 非機能要件

### 5.1. ユーザビリティ

*   アプリケーションは直感的で使いやすいものであること。
*   ユーザーインターフェースは、すっきりとしていて、整理されていること。
*   アプリケーションはレスポンシブであり、さまざまなデバイス（デスクトップ、タブレット、モバイル）で適切に動作すること。

### 5.2. パフォーマンス

*   アプリケーションは迅速にロードされること。
*   投票と投票の公開は、参加者にほぼリアルタイム（5秒以内）で反映されること。
*   参加者数は最大20名であること。

### 5.3. セキュリティ

*   セッションは不正アクセスから保護されること。
*   ユーザーデータは安全に処理されること。
*   アプリケーションは、一般的な Web の脆弱性から保護されること。

### 5.4. アクセシビリティ

*   アプリケーションは、WCAG ガイドラインに準拠し、障害を持つユーザーにもアクセス可能であること。

### 5.5. 信頼性

* アプリケーションは安定しており、信頼できること。
* アプリケーションはエラーを適切に処理すること。

## 6. 画面設計とURL

*   **トップ画面:** `/`
    *   メインのランディングページ。ユーザーがセッションを作成したり、参加したりするための入り口です。
    *   UI
        *   セッション作成画面とセッション参加画面へのリンク。
*   **セッション作成画面:** `/sessions/create`
    *   ホストが新しいセッションを作成するためのページ。フォームが表示されます。
    *   作成後はセッションロビーにリダイレクトされます。
    *   UI
        *   見積もりスケールを選択するドロップダウン。
        *   ホストの名前（またはニックネーム）を入力するフィールド。
        *   セッションを作成するボタン。
*   **セッション参加画面:** `/sessions/join?sessionId={sessionId}`
    *   参加者がセッションに参加するためのページ。フォームが表示されます。
    *   参加後はセッションロビーにリダイレクトされます。
    *   クエリパラメータでセッションIDを指定するとフォームにセッションIDが自動入力されます。
    *   クエリパラメータでセッションIDが未指定のときはセッションIDフォームは空で表示される。
    *   UI
        *   セッションIDを入力するフィールド。
        *   参加者の名前（またはニックネーム）を入力するフィールド。
        *   セッションに参加するボタン。
*   **セッションロビー:** `/sessions/{sessionId}`
    *   このURLは、コンテキストに応じて以下の機能を果たします。
        *   ロビー（ラウンド開始前）: 参加者リストと「ラウンド開始」ボタン（ホスト用）を表示します。
        *   投票（ラウンド中）: 投票オプションと現在のラウンドのステータスを表示します。
        *   結果（ラウンド後）: 公開された投票、平均値、中央値を表示します。
    *   UI（ラウンド開始前）
        *   セッションIDを表示する。
        *   セッション参加画面のURLとコピーボタンを表示する。
        *   セッションの有効期限を表示する。
        *   セッション参加者の一覧。
        *   （ホスト）ラウンドを開始するボタン。
    *   UI（ラウンド中）
        *   セッションIDを表示する。
        *   セッション参加画面のURLとコピーボタンを表示する。
        *   セッションの有効期限を表示する。
        *   セッション参加者の一覧。
        *   ラウンド参加者の投票状況（投票済、未投票）を表示する。
        *   （ラウンド参加者とホスト）各見積もりのボタンまたはカード。
        *   （ホスト）投票を公開するボタン。
    *   UI（ラウンド後）
        *   セッションIDを表示する。
        *   セッション参加画面のURLとコピーボタンを表示する。
        *   セッションの有効期限を表示する。
        *   セッション参加者の一覧。
        *   ラウンド参加者の投票内容を表示する。
        *   投票の平均値と中央値を表示する。
        *   （ホスト）次のラウンドを開始するボタン。
