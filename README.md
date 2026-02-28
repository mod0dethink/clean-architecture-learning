# clean-architecture-learning

フルスタックタスク管理アプリを題材に、クリーンアーキテクチャについて学んでいるリポジトリです。

## アーキテクチャ

```
backend/internal/
├── domain/          # エンティティ・ドメインルール（DBやHTTPに依存しない）
├── usecase/         # ユースケース（操作の手順を管理）
├── interface/
│   ├── repository/  # リポジトリインターフェース
│   └── handler/     # Echo HTTPハンドラー
└── infrastructure/
    └── sqlite/      # SQLite具体実装
```
