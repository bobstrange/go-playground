# module

```bash
# モジュールの初期化
go mod init

# モジュールへのファイルの追加
# -u オプションで依存ライブラリの更新も行われる
go get -u <module>

# 依存の更新/削除
go mod tidy
```
