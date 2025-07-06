# コーディングルール

- 下記のコマンドでエラーが発生しないように実装すること
    - `make test`
    - `make check`
    - `make build`
    - `make test-e2e`
- 下記コマンドでファイルが更新される場合は必ずコミットすること
    - `make generate`
- コード中のコメントは日本語で記載すること
- 自動テストで固定時間のタイムアウトを使用することは禁止

## フロントエンド

- **言語**: TypeScriptを使用すること。
- **フレームワーク/ライブラリ**: React, React Router DOM, Tailwind CSS, DaisyUIを使用すること。
- **コンポーネント命名**: コンポーネントのファイル名とコンポーネント名はPascalCaseを使用すること (例: `Layout.tsx`, `TopPage.tsx`)。
- **カスタムフック命名**: カスタムフックは`use`プレフィックスを使用すること (例: `useLoading.ts`)。
- **APIクライアント命名**: APIクライアントクラスは`Client`サフィックスを使用すること (例: `PlanningPokerClient.ts`)。
- **エラーハンドリング**: 非同期処理には`try...catch`ブロックを使用し、エラーはログに記録し、ユーザーに表示すること。
- **スタイリング**: Tailwind CSSクラスをJSX内で直接使用し、DaisyUIコンポーネントを活用すること。
- **インポート**: `src`内のモジュールをインポートする際は、絶対パスを使用すること。ただし、同じディレクトリ内のファイルをインポートする場合に限り、相対パスを使用してもよい。
- **コンポーネント形式**: Reactの関数コンポーネントを使用すること。
- **状態管理**: `useState`および`useContext`フックを使用してローカルおよびグローバルな状態管理を行うこと。
- **イベントハンドラ命名**: イベントハンドラ関数は`handle`プレフィックスを使用すること (例: `handleUserNameChange`, `handleSubmit`)。
- **アクセシビリティ**: アクセシビリティのために`aria-label`属性を使用すること。
- **定数/列挙型**: 事前定義された値には型または列挙型を使用すること (例: `ScaleType.ts`)。
- **ファイル構造**: 
    - `api/`: APIクライアント定義と生成された型。
    - `common/`: 再利用可能なコンポーネント、フック、プロバイダー。
    - `planning-poker/`: 機能固有のコンポーネント、ページ、型、ユーティリティ。
    - `talk-roulette/`: 機能固有のコンポーネント、ページ。
    - `src`のルートレベルには、主要なアプリケーションのエントリポイントとレイアウトを配置すること。

## バックエンド

- **言語**: Goを使用すること。
- **フレームワーク/ライブラリ**: Echo (Webフレームワーク), go-redis/redis (Redisクライアント), stretchr/testify (テスト), golang/mock (モック生成) を使用すること。
- **プロジェクト構造**:
    - `cmd/server/`: アプリケーションのエントリポイント (`main.go`)。
    - `internal/`: 内部パッケージ。
        - `api/`: OpenAPIによって生成されたAPI定義とハンドラインターフェース。
        - `planningpoker/`: プランニングポーカー機能のビジネスロジック、ハンドラ、ドメイン、インフラ層。
            - `domain/`: ユースケースとビジネスロジック。
            - `infra/`: 永続化層（Redis）とWebSocketハブの実装。
            - `model/`: ドメインモデルとリポジトリインターフェース。
        - `talkroulette/`: トークルーレット機能のビジネスロジック。
        - `web/`: 静的ファイルハンドリング。
- **APIハンドリング**:
    - `echo.Context` を使用してリクエストとレスポンスを処理すること。
    - リクエストボディのバインディングには `ctx.Bind(&req)` を使用すること。
    - レスポンスの返却には `ctx.JSON()` を使用し、適切なHTTPステータスコードを設定すること。
    - エラーレスポンスには `api.ErrorResponse` 構造体を使用すること。
- **エラーハンドリング**:
    - エラーは `fmt.Errorf` を使用してラップし、元のエラー情報を含めること。
    - ログには `log.Printf` や `log.Fatalf` を使用すること。
    - HTTPハンドラでは、エラーが発生した場合に `http.StatusInternalServerError` や `http.StatusBadRequest` などの適切なHTTPステータスコードを返すこと。
- **依存性注入**: 依存関係はコンストラクタを通じて注入すること (例: `NewServer`, `NewRedisClient`, `NewCreateSessionUsecase`)。
- **Redisの使用**:
    - Redisクライアントには `github.com/redis/go-redis/v9` を使用すること。
    - Redisキーの命名規則は `web-toolbox:<feature>:<entity>:<id>` の形式に従うこと (例: `web-toolbox:planning-poker:session:<sessionId>`)。
    - Redisに保存するデータは `json.Marshal` でJSONにシリアライズすること。
    - Redisから取得したデータは `json.Unmarshal` でデシリアライズすること。
    - `context.Context` をRedis操作に渡すこと。
    - `redislib.Nil` を使用してキーが存在しない場合をチェックすること。
- **WebSocket**: WebSocket通信にはカスタムの `WebSocketHub` を使用すること。
- **設定**: 環境変数は `os.Getenv` を使用して読み込み、デフォルト値を設定すること。
- **テスト**:
    - テストには `testing` パッケージを使用し、`httptest` でHTTPリクエストをモックすること。
    - アサーションには `github.com/stretchr/testify/assert` を使用すること。
    - モックの生成には `github.com/golang/mock/gomock` を使用すること。
    - モックの期待値は `EXPECT().AnyTimes()` を使用して設定すること。
- **コード生成**: `openapi/generate.go` にある `//go:generate oapi-codegen -config config.yml ../../docs/spec/openapi.yml` コメントから、OpenAPI仕様からコードを自動生成していることがわかる。`make generate` コマンドで実行される。
- **命名規則**:
    - パッケージ名は小文字で単一の単語を使用すること (例: `planningpoker`, `infra`, `domain`)。
    - 関数名、変数名、フィールド名はcamelCaseを使用すること。
    - インターフェース名は `er` で終わるか、`I` プレフィックスを付けないこと (例: `RedisClient`, `SessionRepository`)。
    - 構造体名はPascalCaseを使用すること。
    - `New` プレフィックスを使用してコンストラクタ関数を定義すること (例: `NewServer`, `NewRedisClient`)。
    - `Handle` プレフィックスを使用してHTTPハンドラ関数を定義すること (例: `HandlePostSessions`)。
    - `Validate` プレフィックスを使用してリクエストバリデーション関数を定義すること (例: `ValidatePostSessions`)。
- **時間処理**: `time.Now()` を使用して現在時刻を取得し、`time.Duration` で期間を扱うこと。
- **UUID生成**: `github.com/google/uuid` を使用してUUIDを生成すること。

## E2Eテスト

- **言語**: TypeScriptを使用すること。
- **フレームワーク/ライブラリ**: Playwrightを使用すること。
- **テストファイル命名**: テストファイルは `*.spec.ts` の形式で命名すること (例: `CreateSessionPage.spec.ts`)。
- **テスト構造**: `test.describe` でテストスイートを定義し、`test` で個別のテストケースを定義すること。
- **Page Object Model (POM)**:
    - POMパターンを使用し、`pom/` ディレクトリにページオブジェクトを配置すること。
    - ページオブジェクトのクラス名はPascalCaseを使用し、`PagePom` サフィックスを付けること (例: `CreateSessionPagePom`)。
    - ページオブジェクトのメソッドは、ユーザーのアクションやページの要素へのアクセスを抽象化すること (例: `inputUserName`, `clickCreateButton`)。
    - ページオブジェクトのコンストラクタは `Page` オブジェクトを受け取ること。
- **要素の特定**: `page.getByRole`, `page.getByLabel` などのPlaywrightのロケータを使用し、堅牢なセレクタを記述すること。
- **アサーション**: `expect` を使用してアサーションを行うこと (例: `expect(page).toHaveTitle()`, `expect(element).toBeVisible()`, `expect(element).toBeDisabled()`)。
- **ナビゲーション**: `page.goto()` を使用してページに遷移すること。
- **待機**: 必要に応じて `page.waitForEvent()` などの待機メカニズムを使用すること。
- **テストデータ**: テストデータはテストファイル内で定義するか、必要に応じて外部からインポートすること。
- **ファイル構造**:
    - `tests/`: テストケースを配置するディレクトリ。
    - `pom/`: Page Object Modelのクラスを配置するディレクトリ。
        - 機能ごとにサブディレクトリを作成すること (例: `planning-poker/`)。
- **テストの実行**: `make test-e2e` コマンドで実行される。
