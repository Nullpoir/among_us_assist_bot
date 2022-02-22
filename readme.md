# 環境構築
## .envのコピー
```
cp .example_env .env
```
適切な数値に置き換えください。
## dockerによる環境構築
初回のみ
```
docker compose build
```

```
docker compose up -d
```

## 開発時の操作
```
docker compose exec go bash
```

上記コマンド実行後app以下をロードしたコンテナ内でgoコマンドなどが使用可能になります。

### 新規パッケージのインストール
```
go get <package_name>
```
