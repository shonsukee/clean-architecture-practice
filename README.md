# Clean Architecture in Go (Practice)

このプロジェクトは、Go言語でClean Architecture（クリーンアーキテクチャ）を学ぶためのシンプルなAPIサーバーのサンプルです。
外部ライブラリへの依存を最小限に抑え、標準ライブラリを中心に構成することで、アーキテクチャの構造自体を理解しやすくしています。

## ディレクトリ構成

Clean Architectureの原則に基づき、依存関係は **外側から内側** に向かってのみ許可されます。

- **`internal/domain/`** (Enterprise Business Rules)
    - 最も内側の層。外部への依存を一切持ちません。
    - エンティティ（`User`構造体など）や、リポジトリのインターフェース定義が含まれます。
- **`internal/usecase/`** (Application Business Rules)
    - ドメイン層を操作するビジネスロジックが含まれます。
    - 具体的なDB実装などは知らず、ドメイン層で定義されたインターフェース（Repository）を通じてデータを操作します。
- **`internal/interfaces/handler/`** (Interface Adapters)
    - 外部からの入力（HTTPリクエスト）をユースケースが理解できる形に変換し、結果を外部への出力（HTTPレスポンス）に変換します。
- **`internal/infrastructure/`** (Frameworks & Drivers)
    - 最も外側の層。データベース接続や具体的な技術詳細が含まれます。
    - ここではメモリ上のマップを使った簡易リポジトリ (`user_memory.go`) を実装しています。
- **`cmd/api/`**
    - アプリケーションのエントリーポイント。
    - 各層のコンポーネントを初期化し、依存関係を注入（DI）してサーバーを起動します。

## 起動方法

```bash
go run cmd/api/main.go
```

サーバーが `:8080` で起動します。

## 動作確認

### ユーザー作成 (POST)

```bash
curl -X POST http://localhost:8080/users \
  -d '{"id": "1", "name": "Alice", "email": "alice@example.com"}'
```

### ユーザー取得 (GET)

```bash
curl http://localhost:8080/users/1
```

## 練習手順 (Step by Step)

このプロジェクトを使ってClean Architectureを学ぶためのステップです。

### Step 1: コードを読んで依存関係を理解する
各ファイル (`import`文) を見て、どのパッケージがどのパッケージに依存しているか確認してください。
`infrastructure` -> `usecase` (DI時) -> `domain` のように、内側の層が外側の層（DBやHTTP）を知らないことを確認しましょう。

### Step 2: Userエンティティにフィールドを追加する
1. `internal/domain/user/entity.go` の `User` 構造体に `Age int` を追加してください。
2. それに合わせて、Handler (JSONデコード部分) と Usecase (Createメソッドの引数) を修正してください。
   - ※ Clean Architectureでは、変更が内側（Domain）から外側へ波及していく様子を体感してください。

### Step 3: ビジネスロジックを追加する
1. `internal/usecase/user.go` の `CreateUser` メソッドにバリデーションを追加してください。
   - 例: Emailが空の場合はエラーにする、IDが既に存在するかチェックするなど。

### Step 4: 新しいユースケースを作る
1. 全ユーザーのリストを取得する機能を追加してみましょう。
   - `Repository` インターフェースに `FinalAll` メソッドを追加。
   - `UserUsecase` インターフェースと実装に `ListUsers` メソッドを追加。
   - `Handler` に `ListUsers` メソッドを追加し、ルーティングを設定。
   - `Infrastructure` のメモリリポジトリに `FindAll` の実装を追加。

### Step 5 (Advanced): リポジトリを差し替える
現在はメモリ上にデータを保存していますが、これをファイルやDBに保存するように変更しても、**UsecaseやDomainのコードは一切変更する必要がない** ことを確認するのがこのアーキテクチャの醍醐味です。

1. `internal/infrastructure/repository/` に `user_file.go` を作成し、JSONファイルにデータを読み書きする新しいリポジトリ実装を作ってみてください。
2. `cmd/api/main.go` で `NewUserMemoryRepository` の代わりに新しいリポジトリを使うように書き換えてください。
   - サーバーを再起動してもデータが永続化されるようになります。

---
Happy Coding!
