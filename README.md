# CASystem
CASystemの個人開発用リポジトリ

## システム構成
参考文献(https://qiita.com/tozastation/items/a69a102fdc3f62d566b4)

```
├── docker: 仮想環境設定
├── public: 実行ファイル
├── test:   実行確認
├── pkg:
│   ├── application: ビジネスロジック
│   ├── di:          依存性の注入
│   ├── infra:       外部との通信
│   ├── DB:          データベースのマイグレーション
│   └── server
│       ├── handler:    エンドポイント
│       ├── middleware: 認証
│       ├── response
│       └── server.go:  ルーティング
├── main.go
└── Makefile: ビルド設定
```

## 技術選定

- 技術選定
    - アプリケーション: Go
        - アーキテクチャ: Layered Architecture
        - ログ: zap
        - 認証・認可: Firebase
        - ルーティング: net
        - マイグレーション: goose
        - データベースアクセス: gorm
        - テスト: Go test
        - モック・スタブ: Go mock
    - データベース: MySQL

### SQLの構築

#### Docker
```
docker run --name コンテナ名 -p 3333:3306 -e MYSQL_ROOT_PASSWORD=password -d mysql:latest
docker exec -it コンテナ名 bash
mysql -u root -p -P 3333
    password
```

#### ホスト
MySQL Workbenchでmwbファイルを開く

Database -> ForEngineer
Stored Connection -> Manage Stored Connections
ポートフォワーディングの設定を入力
Continueを押してDocker上のMySQLにDBを作成
