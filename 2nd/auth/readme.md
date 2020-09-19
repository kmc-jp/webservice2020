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
   3. 再び[GoogleDeveloperConsole](https://console.developers.google.com/apis/credentials)にアクセスし、`+認証情報を作成`をクリック
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

## プログラム概要
### init関数
Golangではinit関数→main関数の順に実行される。

**補足-import-**

importするときにはこんな書き方もできる。
```go
import (
   f "fmt"
)
```
これをすると、パッケージ関数を実行するとき、
```go
fmt.Printf("Hello World!\n")
```
とかく所を
```go
f.Printf("Hello World!\n")
```
と書くことができる。これにより、書く手間を省くことができる他、`fmt`という名前のパッケージ名が被った場合、名前を変更できるというメリットもある。

またその他、
```go
import (
   _ "fmt"
)
```
と書くことで、importしても名前を与えないということができる。ここでポイントとなるのがinit関数なのです。

packageは引用されたとき、init関数が最初に実行されるため、名前が与えられていないpackageであってもinit関数だけは実行されます。

今回のプログラムでは、init関数で変数の初期化の記述などを行っています。

## main関数
```go
func main(){
   ...
   if err := http.ListenAndServe("0.0.0.0:8080", nil); err != nil {
		log.Fatalln("Error")
	}
}
```
この部分でhttpサーバを起動している。
```go
http.ListenAndServe(アドレス, ルーティング設定)
```
という形の関数で、第二引数で自分でカスタムルーティング設定を指定することができる。

今回（そして、大体の場合）は次に説明するようにパッケージに入っているルーティング設定を用いて設定する。

```go
   http.Handle(論理アドレス, ServeHTTPメソッドを実装している型の変数のポインタ)
	http.HandleFunc("/logout", logoutHandle)
```