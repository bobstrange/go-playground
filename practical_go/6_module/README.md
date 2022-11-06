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
