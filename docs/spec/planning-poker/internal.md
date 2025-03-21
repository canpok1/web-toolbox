# プランニングポーカーアプリ - 詳細設計書

## 1. 概要

本ドキュメントは、プランニングポーカーアプリの詳細設計を記述するものです。クライアント/サーバー構成とし、REST APIで通信を行います。データストアにはRedisを使用し、サーバーサイドはGo言語、クライアントサイドはTypeScriptとReactで開発します。

## 2. システム構成

### 2.1. アーキテクチャ

*   **クライアント:** TypeScript/Reactで構築されたWebアプリケーション。
*   **サーバー:** Go言語で構築されたREST APIサーバー。
*   **データストア:** Redis。セッション情報、参加者情報、投票情報などを保持。

### 2.2. 通信方式

*   クライアントとサーバー間は、REST API（JSON形式）で通信を行います。
*   リアルタイムな更新が必要な箇所（例：新しい参加者の追加、投票の送信、ラウンドの開始/終了）は、WebSocketを使用します。

## 3. データモデル（Redis）

Redisには以下のデータ構造で情報を格納します。

### 3.1. セッション (Session)

*   **キー:** `web-toolbox:planning-poker:session:{sessionId}`
*   **型:** Hash
*   **フィールド:**
    *   `sessionName` (String): セッション名
    *   `hostId` (String): ホストのparticipantId
    *   `scaleType` (String): 見積もりスケールのタイプ (例: "fibonacci", "tshirt", "custom")
    *   `customScale` (String, JSON Array): カスタムスケールの場合のスケール値の配列 (例: `["1", "2", "3", "5"]`)
    *   `currentRoundId` (String): 現在のラウンドのroundId
    *   `status` (String): セッションのステータス ("lobby", "inProgress", "finished")
    *   `createdAt` (String): 作成日時 (ISO 8601形式)
    *   `updatedAt` (String): 更新日時 (ISO 8601形式)

### 3.2. ラウンド (Round)

*   **キー:** `web-toolbox:planning-poker:round:{roundId}`
*   **型:** Hash
*   **フィールド:**
    *   `sessionId` (String): 属するセッションのsessionId
    *   `status` (String): ラウンドのステータス ("voting", "revealed")
    *   `createdAt` (String): 作成日時 (ISO 8601形式)
    *   `updatedAt` (String): 更新日時 (ISO 8601形式)

### 3.3. 参加者 (Participant)

*   **キー:** `web-toolbox:planning-poker:participant:{participantId}`
*   **型:** Hash
*   **フィールド:**
    *   `sessionId` (String): 属するセッションのsessionId
    *   `name` (String): 参加者の名前
    *   `isHost` (Boolean): ホストかどうか
    *   `createdAt` (String): 作成日時 (ISO 8601形式)
    *   `updatedAt` (String): 更新日時 (ISO 8601形式)

### 3.4. 投票 (Vote)

*   **キー:** `web-toolbox:planning-poker:vote:{voteId}`
*   **型:** Hash
*   **フィールド:**
    *   `roundId` (String): 属するラウンドのroundId
    *   `participantId` (String): 投票者のparticipantId
    *   `value` (String): 投票値
    *   `createdAt` (String): 作成日時 (ISO 8601形式)
    *   `updatedAt` (String): 更新日時 (ISO 8601形式)

### 3.5. セッション参加者リスト (Session Participants)

*   **キー:** `web-toolbox:planning-poker:session:{sessionId}:participants`
*   **型:** Set
*   **値:** participantId のリスト

### 3.6. ラウンド投票リスト (Round Votes)

*   **キー:** `web-toolbox:planning-poker:round:{roundId}:votes`
*   **型:** Set
*   **値:** voteId のリスト

## 4. WebSocket 仕様

### 4.1. イベント

*   **`participantJoined`:** 新しい参加者がセッションに参加したことを通知。
    *   **ペイロード:**
        ```json
        {
          "participantId": "zzzzzzzz-zzzz-zzzz-zzzz-zzzzzzzzzzzz",
          "name": "Jane Smith"
        }
        ```
*   **`roundStarted`:** 新しいラウンドが開始されたことを通知。
    *   **ペイロード:**
        ```json
        {
          "roundId": "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa"
        }
        ```
*   **`voteSubmitted`:** 参加者が投票を送信したことを通知。
    *   **ペイロード:**
        ```json
        {
          "participantId": "zzzzzzzz-zzzz-zzzz-zzzz-zzzzzzzzzzzz"
        }
        ```
*   **`votesRevealed`:** ラウンドが終了し、投票が公開されたことを通知。
    *   **ペイロード:**
        ```json
        {
          "votes": [
            {
              "participantId": "zzzzzzzz-zzzz-zzzz-zzzz-zzzzzzzzzzzz",
              "value": "5"
            },
            {
              "participantId": "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
              "value": "8"
            }
          ],
          "average": 6.5,
          "median": 6.5
        }
        ```
* **`sessionEnded`:** セッションが終了したことを通知。
    * **ペイロード:**
        ```json
        {}
        ```

## 5. クライアントサイド（TypeScript/React）

### 5.1. コンポーネント構成

*   **SessionCreate:** セッション作成画面。
*   **SessionJoin:** セッション参加画面。
*   **SessionLobby:** セッションロビー画面。参加者一覧、ラウンド開始ボタンなどを表示。
*   **Voting:** 投票画面。
*   **Result:** 結果表示画面。

### 5.2. 状態管理

*   React ContextまたはReduxなどの状態管理ライブラリを使用。
*   セッション情報、参加者情報、ラウンド情報、投票情報などを管理。

### 5.3. 通信

*   `fetch` APIまたは`axios`などのHTTPクライアントライブラリを使用。
*   WebSocketクライアントライブラリを使用。

## 6. サーバーサイド（Go言語）

### 6.1. フレームワーク

*   `gin`などのWebフレームワークを使用。

### 6.2. Redisクライアント

*   `go-redis`などのRedisクライアントライブラリを使用。

### 6.3. WebSocket

*   `gorilla/websocket`などのWebSocketライブラリを使用。

### 6.4. 処理の流れ

1.  クライアントからのリクエストを受け付ける。
2.  Redisからデータを取得または更新する。
3.  必要に応じてWebSocketでクライアントに通知する。
4.  クライアントにレスポンスを返す。

## 7. エラー処理

*   **REST API:**
    *   400 Bad Request: リクエストボディの形式が不正な場合。
    *   404 Not Found: リソースが見つからない場合。
    *   500 Internal Server Error: サーバー内部エラー。
*   **WebSocket:**
    *   エラーが発生した場合、クライアントにエラーメッセージを送信。

## 8. セキュリティ

*   **セッションID:** UUIDを使用し、推測困難なIDを生成。
*   **入力値の検証:** クライアントからの入力値を検証し、不正な値を排除。
*   **アクセス制御:** ホストのみがラウンドの開始/終了、セッションの終了を行えるようにする。
* **CORS:** 適切なCORS設定を行う。

## 9. 今後の拡張

*   ユーザー認証機能の追加。
*   投票結果の履歴保存機能。
*   チャット機能の追加。

