# 第六回
## 概要
実際にWeb掲示板サービスを動かしてみる。

## 準備
必要なpackageを入れる。

1. `gomniauth`:OAuthをするためのpackage
```sh
$ go get github.com/stretchr/gomniauth/...
```
- 次のエラーが出た場合
```sh
fatal: could not read Username for 'https://github.com': terminal prompts disabled
```
1. [https://github.com/settings/tokens](https://github.com/settings/tokens)にアクセス
2. `Generate new token`
3. `Note`に適当な名前を入力
4. `repo`，`admin:org`にチェックを入れ`Generate token`をクリック
5. 次のコマンドをtokenを代入して実行
```sh
$ git config --global url."https://YOUR_TOKEN:x-oauth-basic@github.com/".insteadOf "https://github.com/"
```

この後`go get`をすればうまくいくはず。

