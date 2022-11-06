# module

## commands

```bash
# モジュールの初期化
go mod init

# モジュールへのファイルの追加
# -u オプションで依存ライブラリの更新も行われる
go get -u <module>

# 依存の更新/削除
go mod tidy
```

## package structure

- 同一フォルダ内のファイルは同一パッケージ
- 実行ファイルのエントリーポイントは main パッケージ

実行ファイルを含む moudle のフォルダ構成例
cmd フォルダ -> コマンド名のフォルダ

```bash
.
  ./cmd
    ./commandA
      ./main.go
    ./commandB
      ./main.go
  ./go.mod
  ./lib.go
```

循環参照を避けるための方法

1. 共通要素のパッケージを作り、依存先をそこに集中する
  - util, common などの名前になりがち
2. ルートパッケージのロジックを全て移動し、ルートパッケージを共通要素の置き場とする
  - 1 の util や common の代わりにルートのパッケージを使う
3. 共通部分を持たない末端のロジックを子パッケージとして切り出す
