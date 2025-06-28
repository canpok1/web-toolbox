### sessionIdの抽出方法を改善

**対象ファイル:** `e2e/tests/planning-poker/SessionPage.spec.ts`

**問題点:**
招待URLから `sessionId` を抽出する際に、`split("id=")[1]` を使用しているため、URLの形式変更に弱い。

**タスク:**
正規表現などを用いて、より堅牢な `sessionId` の抽出方法に改善する。