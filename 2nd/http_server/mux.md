# 簡単なサーバ構築
## 概観
Go言語での簡単なWebアプリの開発の仕方を理解する。

## 目次
<!-- TOC -->

- [簡単なサーバ構築](#簡単なサーバ構築)
    - [概観](#概観)
    - [目次](#目次)
    - [リクエストの仕組み](#リクエストの仕組み)
    - [マルチプレクサ](#マルチプレクサ)
        - [Hello World](#hello-world)
        - [ファイルサーバを追加する](#ファイルサーバを追加する)
        - [DefaltServeMuxの省略](#defaltservemuxの省略)
        - [DefaultServeMux以外のマルチプレクサ](#defaultservemux以外のマルチプレクサ)
    - [参考](#参考)

<!-- /TOC -->

## リクエストの仕組み

http通信に於いて、クライアントからのリクエストに応じてレスポンスを返送するという流れは、今までに何度も書いてきたとおりです。では具体的にその過程がどのように実行されるのでしょうか。

実際には、リクエストに応じて、ハンドラとなる関数が呼びだされ、そのハンドラがレスポンスを作成します。次の例を見てみましょう。

```go
package main

import (
	"fmt"
	"net/http"
)

type hello struct{}

// ハンドラ
func (h hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
	return
}

func main() {
    // server設定の作成
	var server = &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: &hello{},
	}

    // 起動
	server.ListenAndServe()
}

```


main関数内に於いて、まずhttpサーバの設定を作成し、その中にハンドラ、hello構造体のポインタを渡しています。ここでハンドラとなるものは、ServeHTTPというメソッドを実装している必要があります。このようにすることで、httpリクエストが来たとき、Goはhello構造体のServeHTTPメソッドを実行して、レスポンスを生成します。

## マルチプレクサ

しかし、以上のプログラムはおおきな問題を抱えています。

そう。このままでは全てのパスに対するリクエストに一つのハンドラしか指定できないのです。そこでほとんどの場合、
**マルチプレクサ**
と呼ばれる、通信の受け付けセンター的なものを使います。ここではリクエストの要求するURIを解析し、それに該当するハンドラを選択、実行します。ハンドラは、要求された処理に従い、適切なレスポンスを生成し、マルチプレクサに返却します。これによって得られたものをマルチプレクサはクライアントへ返送します。

これだけでは何が何なのか、という感じなので、具体例を眺めてみましょう。

### Hello World

Hello Worldとだけ返すWebアプリを作成します。

```go
package main

import (
	"fmt"
	"net/http"
)

func main() {
    // マルチプレクサを作成
	var mux = http.NewServeMux()

    // ハンドラを指定
	mux.HandleFunc("/", hello)

    // サーバ設定を記述
	var server = &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}

    // 起動
	server.ListenAndServe()
}

// ハンドラ
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
	return
}

```

以上のプログラムを実行し、
[localhost:8080](http://localhost:8080)
にアクセスすると、`Hello World`という文字列がブラウザ上に表示されます。

このプログラムは第一にマルチプレクサ(`mux`)を作成し、ハンドラ(`hello`)とURL(`/`)を結びつけています。

ハンドラ`hello`はHello Worldという文字列をResponse Writerに書き込んでいます。<a href="#1">^1</a>

以上をもってこのプログラムは、URI(`/`)にリクエストが来たとき、`Hello World`という文字列を返す、というフローができています。

<a id="1">^1</a>: 先程のようにServeHTTPを実装したhello構造体を使い、

```go
mux.Handle("/", &hello{})

```

とすることももちろん可能です。HandleFunc関数は、関数を入れることで、そのようなインターフェースを満たす変数を自動的に背後で作成してくれます。


### ファイルサーバを追加する

では更に、`/static/`以下をプログラムのディレクトリの`public`ディレクトリ以下のファイルサーバとしてみましょう。

```go
package main

import (
	"fmt"
	"net/http"
)

func main() {
    // マルチプレクサを作成
	var mux = http.NewServeMux()

    // ファイルサーバの作成
    // 参照先としてpublic以下を指定
	var fileServe = http.FileServer(http.Dir("public"))
    
    // ハンドラを指定
    
    // / で始まるURIのルーティングを指定
    mux.HandleFunc("/", hello)
    
    // /static/ で始まるURIのルーティングを指定
    mux.Handle("/static/",
        // 接頭辞 /static/ の文字列を除いたパスを参照
        // 例: /static/index.html → index.htmlを参照
        // 今回は参照先がpublic以下なので、public/index.htmlが参照される。
        http.StripPrefix("/static/", fileServe),
    )

    // サーバ設定を記述
	var server = &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
    // 起動
	server.ListenAndServe()
}

// ハンドラ
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
	return
}

```

このプログラムを実行すると
`/static/`
にアクセスすると、public以下のファイルが送信され、それ以外ではHello Worldが出るようになります。

各々の関数の役割はコメントを参照してください。

### DefaltServeMuxの省略
以上のように、毎回`http.NewServeMux()`でマルチプレクサを作成しても良いのですが、httpパッケージではもう少し楽ができるようになっています。

次の例を見てください。


```go
package main

import (
	"fmt"
	"net/http"
)

func main() {
    // ハンドラを指定
	http.HandleFunc("/", hello)

    // 起動
    // 第２引数はマルチプレクサを入れることができる。
    // 入れなければ自動的にDefaultServeMuxになる。
    http.ListenAndServe("0.0.0.0:8080", nil))

}

// ハンドラ
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
	return
}

```

このようにDefaltServeMuxを利用する場合は明示的に宣言せずとも利用することができます。この記法では、背後でDefaltServeMuxが作成され、実行されます。


### DefaultServeMux以外のマルチプレクサ

今まで使ってきたマルチプレクサは、`net/http`が提供する標準のマルチプレクサでした。しかし、このマルチプレクサ以外にも、サードパーティ製のマルチプレクサが多く存在します。isucon7で使われている`echo`はまさにこの、マルチプレクサの一つです。

本番どのようなマルチプレクサが使われるのかは未知数ですが、所詮このような処理をするものにすぎないということを覚えておいて損はないでしょう。

## 参考

[Goプログラミング実践入門 ―標準ライブラリでゼロからWebアプリを作る―](https://elib.maruzen.co.jp/elib/html/BookDetail/Id/3000049505?0)