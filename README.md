# dynamic-ssl-proxy-sample

動的にssl証明を読み込むリバースプロキシのプログラムのサンプル(証明書はファイルで保存されている想定)

`shared/ssl/`下に`{ドメイン名}.pem`という名前で証明書ファイル、`{ドメイン名}-key.pem`という名前で秘密鍵を保存する。

例)
- ドメイン名: example.com
  - 証明書ファイル名 `shared/ssl/example.com.pem`
  - 秘密鍵ファイル名 `shared/ssl/example.com-key.pem`