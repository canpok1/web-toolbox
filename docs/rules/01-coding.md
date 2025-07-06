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
- **エラーハンドリング**: APIリクエストなどの非同期処理では、`try...catch` ブロックを用いてエラーを捕捉します。捕捉したエラーは、開発者向けにコンソールへログ出力すると同時に、ユーザーには汎用的なエラーメッセージを表示するか、具体的なエラー内容に応じて適切なUIフィードバックを提供してください。これにより、デバッグの容易性とユーザー体験の向上を両立します。
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
- **主な使用ライブラリ**: Echo (Webフレームワーク), go-redis/redis (Redisクライアント), stretchr/testify (テスト), golang/mock (モック生成) を使用しています。
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
    - エラーは `fmt.Errorf` を用いてラップし、スタックトレースにコンテキスト情報を追加していくことを推奨します。これにより、エラーの原因特定が容易になります。 (例: `fmt.Errorf("failed to get session: %w", err)`)
    - ログには `log.Printf` や `log.Fatalf` を使用します。特に、リクエスト処理中のエラーは `log.Printf` で記録し、サーバーの起動などクリティカルな処理の失敗時には `log.Fatalf` を使用してプロセスを終了させます。
    - HTTPハンドラでは、エラーの種類に応じて `http.StatusInternalServerError` や `http.StatusBadRequest` など、適切なHTTPステータスコードを返却してください。クライアントに不要な詳細情報を漏らさないよう、エラーメッセージは汎用的なものに留めるのが基本です。
- **依存性注入**: サーバー、ユースケース、リポジトリなどのコンポーネントは、コンストラクタ（`New...` 関数）を通じて依存関係を受け取ります。これにより、コンポーネントの結合度を下げ、テスト容易性を高めています。例えば、`NewCreateSessionUsecase` は `SessionRepository` インターフェースに依存し、具象的な実装（`RedisRepository` など）は main 関数で注入されます。このパターンに従い、コンポーネントの独立性を保ってください。
- **Redisの使用**:
    - Redisクライアントには `github.com/redis/go-redis/v9` を使用すること。
    - Redisキーの命名規則は `web-toolbox:<feature>:<entity>:<id>` の形式に従うこと (例: `web-toolbox:planning-poker:session:<sessionId>`)。
    - Redisに保存するデータは `json.Marshal` でJSONにシリアライズすること。
    - Redisから取得したデータは `json.Unmarshal` でデシリアライズすること。
    - `context.Context` をRedis操作に渡すこと。
    - `redis.Nil` を使用してキーが存在しない場合をチェックすること。
- **WebSocket**: WebSocket通信にはカスタムの `WebSocketHub` を使用すること。
- **設定**: 環境変数は `os.Getenv` を使用して読み込み、デフォルト値を設定すること。
- **テスト**:
    - テストには `testing` パッケージを使用し、`httptest` でHTTPリクエストをモックすること。
    - アサーションには `github.com/stretchr/testify/assert` を使用すること。
    - モックの生成には `github.com/golang/mock/gomock` を使用すること。
    - モックの期待値は、呼び出し回数に応じて `Times(n)` や `AnyTimes()` を適切に使い分けること。
- **コード生成**: `openapi/generate.go` にある `//go:generate oapi-codegen -config config.yml ../../docs/spec/openapi.yml` コメントから、OpenAPI仕様からコードを自動生成していることがわかる。`make generate` コマンドで実行される。
- **命名規則**:
    - パッケージ名は小文字で単一の単語を使用すること (例: `planningpoker`, `infra`, `domain`)。
    - エクスポートされる識別子（関数、変数、型など）はPascalCase、アンエクスポートな識別子はcamelCaseを使用すること。
    - インターフェース名には `I` プレフィックスを付けず、可能であれば `er` サフィックスを使用すること (例: `Reader`, `Writer`)。振る舞いを表す動詞から命名できない場合は、`Repository` や `Client` のような型を表す名詞を使用してもよい (例: `SessionRepository`, `RedisClient`)。
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
    - ページオブジェクトのクラス名はPascalCaseを使用し、`PagePom` サフィックスを付けること (例: `CreateSessionPagePom`)。ファイル名とクラス名を一致させ、1ファイル1クラスとすること。
    - ページオブジェクトのメソッドは、ユーザーのアクションやページの要素へのアクセスを抽象化すること (例: `inputUserName`, `clickCreateButton`)。
    - `Locator` の取得は getter メソッドで行うこと。
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
