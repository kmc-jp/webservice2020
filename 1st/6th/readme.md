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

    6. `go get`をすればうまくいくはず。

2. credential.jsonの準備
   1. `webservice2020/1st/6th/`に`credential.json`を用意
   2. [GoogleDeceloperConsole](https://console.developers.google.com/apis/credentials)にアクセスしてプロジェクト作成
      - プロジェクト名は`webservice2020-yourname-にしておきましょう。
   3. 再び[GoogleDeceloperConsole](https://console.developers.google.com/apis/credentials)にアクセスし、`+認証情報を作成`をクリック
      - 認証画面の作成が出たら適宜いい感じに入力しておきましょう。
   4. `承認済みの JavaScript 生成元`には`http://localhost:8080`、`承認済みのリダイレクト URI`には`http://localhost:8080/auth/callback/google`と入力しましょう。
   5. 出てきたクライアントIDとシークレットを先ほどの`credential.json`中に次のように入力
    ```json
    {
        "clientID":"クライアントID",
        "secret":"シークレット"
    }
    ```
   6. 保存したら完了