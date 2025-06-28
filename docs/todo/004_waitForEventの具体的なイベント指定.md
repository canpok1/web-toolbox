### waitForEvent("websocket")の具体的なイベント指定

**対象ファイル:** `e2e/tests/planning-poker/SessionPage.spec.ts`

**問題点:**
`page.waitForEvent("websocket")` が汎用的すぎるため、意図しないWebSocketイベントを待ってしまう可能性がある。これにより、テストが不安定になる可能性がある。

**タスク:**
WebSocketイベントをより具体的に指定する（例: 特定のURLへの接続、特定のメッセージの受信など）ことで、テストの信頼性を向上させる。