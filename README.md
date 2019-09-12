# CASystem
CASystemの個人開発用リポジトリ

## システム構成
参考文献(https://qiita.com/tozastation/items/a69a102fdc3f62d566b4)

├── domain: ドメイン層
│   ├── repository: 依存性逆転の原則
│   └── service: ロジック
├── implements: アプリケーション層
├── infrastructure: インフラ層
│   └── persistence
│       ├── model: DBモデル
│       └── mysql: mysqlサーバ用Repository
├── migrations: DB初期化
├── interfaces: その他
│   ├── auth: 認証
│   ├── di: 依存性の注入
│   └── handler: ハンドラー
├── main.go: ルーティング
└── vendor: Goのパッケージ

## 技術選定

- 技術選定
    - アプリケーション: Go
        - アーキテクチャ: Layered Architecture
        - フレームワーク: Gin
        - ログ: seelog
        - 認証・認可: Firebase
        - ルーティング: ?
        - マイグレーション: goose
        - データベースアクセス: github.com/go-sql-driver/mysql
        - テスト: Go test
        - モック・スタブ: Go mock
    - データベース: MySQL
