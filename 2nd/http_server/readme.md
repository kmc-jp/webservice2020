# 付録
## 概要
シンプルな掲示板サービスの見本

## 準備
1. ビルドして実行
```sh
$ go build
$ ./appendix
```

## 仕組み
### 前提知識
1. init関数について

Golangではinit関数→main関数の順に最初に実行される。

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

2. http パッケージについて

http package：ウェブサーバを立てるためのパッケージ。

追記：以下の前にISUCON向けに書いた次の資料を読むとわかりやすそうなの載せておきます。<br>
[資料](mux.md)

Handle関数で指定されたパスに対する動作を指定して、ListenAndServe関数でアドレスを指定した上でサーバを起動する。

例：

```go
package main

import (
	"fmt"
	"net/http"
)

//Sample Sample struct
type Sample struct{}

func main() {
	http.Handle("/", &Sample{})
	if err := http.ListenAndServe("0.0.0.0:8080", nil); err != nil {
		log.Fatalln("Error")
	}
}

func (s *Sample) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
	return
}
```

```go
func main() {
	http.Handle("/", &Sample{})
	if err := http.ListenAndServe("0.0.0.0:8080", nil); err != nil {
		log.Fatalln("Error")
	}
}
```
```go
	http.Handle("/", &Sample{})
```

　ここでは
`http://.../`
にリクエストがあったときのハンドラを指定している。第一引数はURLのPath部分が`/`であるものであることを指定していて、第二引数ではServeHTTPメソッド[^1]を実装している集合であるハンドラを指定している。今回ではSample構造体を代入している。

[^1]: `func ServeHTTP(w http.ResponseWriter, r * http.Request)`

```go 
	if err := http.ListenAndServe("0.0.0.0:8080", nil); err != nil {
		log.Fatalln("Error")
	}
```
`ListenAndServe`:httpのリクエストを受付始めるための関数。第一引数にサーバアドレス、第二引数にルーティング設定を入れる。

☆ルーティング設定は省略でき、ない場合はHandle関数で追加した内容を用いる。

```go
//Sample Sample struct
type Sample struct{}
...
func (s *Sample) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
	return
}
```
```go
ServeHTTP(w http.ResponseWriter, r *http.Request)
```
`http.ResponseWriter` : これにクライアント側に返す内容を記述していく。
`http.Request` : これからクライアントのリクエストの内容を参照できる。

```go
fmt.Fprintf(w, "Hello World!")
```
`Fprintf`関数は第一引数にWriterを入れることで、以後の内容をwに出力してくれる。

### init関数
```go
func init() {
	//message.jsonのファイルパスを指定（環境依存）
	MessageFilePath = filepath.Join("data", "message.json")
	//メッセージデータを読み込み
    ReadMessageData()
}
```

```go
MessageFilePath = filepath.Join("data", "message.json")
```
ここでは`message.json`（テキストデータの保存場所）の保存場所を指定しています。

注：`MessageFilePath`というのはグローバル領域に宣言されています。
```go
//MessageFilePath Put a path of message.json
var MessageFilePath string
```

- **filepath package**

ファイルパスの記述方法はOSに依存している。

例：
  - UNIX系
    `/home/~tkmax777/Documents`
  - Windows
    `C:\Users\tkmax777\Documents`

　このような差を自動的にOSに合わせて補完してくれるのがfilepathパッケージ。今回の使い方ではJoin関数で`data`ディレクトリと`message.json`ファイルのパスを結合している。よって、Unix系では`data/message.json`になるし、Windows系では`data\message.json`になる。

　この他では相対パスを絶対パスに変換してくれるAbs関数や、相対パスを最適化した書き方にしてくれるClean関数などがある。

```go
    ReadMessageData()
```
この関数は、main.go内に記述されている。
```go
//ReadMessageData Get messages previously
func ReadMessageData() {
	data, err := ReadFile(MessageFilePath)
	if err != nil {
		fmt.Printf("Read message file error\n%v", err)
		return
	}

	err = json.Unmarshal(data, &Messages)
	if err != nil {
		fmt.Printf("Json unmarshal error\n%v", err)
		return
	}
}
```

`ReadFile`関数はlib.go内で宣言されていて、ファイルパスをしていすると、そのファイルのバイトデータを吸い出して、dataに格納してくれる関数。

その後、そのバイトデータをjsonパッケージのUnmarshal関数でMessages変数内に格納している。

注：Messagesはグローバル領域で宣言している。

```go
//Messages Put all Message data
var Messages []Message

...

//Message Put message data
type Message struct {
	Name string
	Text string
	Time time.Time
}
```

### main関数
```go
func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.Handle("/", &Top)
	http.Handle("/form", http.HandlerFunc(AddMessage))

	if err := http.ListenAndServe("0.0.0.0:8080", nil); err != nil {
		log.Fatalln("Error")
	}
}
```
少しずつ見ていきましょう。

```go
fs := http.FileServer(http.Dir("static"))
```

```go
http.Dir("static")
```

これは関数ではありません。http package内で定義されたstring型の独自定義型です。

以下http package該当箇所
```go
// A Dir implements FileSystem using the native file system restricted to a
// specific directory tree.
//
// While the FileSystem.Open method takes '/'-separated paths, a Dir's string
// value is a filename on the native file system, not a URL, so it is separated
// by filepath.Separator, which isn't necessarily '/'.
//
// Note that Dir will allow access to files and directories starting with a
// period, which could expose sensitive directories like a .git directory or
// sensitive files like .htpasswd. To exclude files with a leading period,
// remove the files/directories from the server or create a custom FileSystem
// implementation.
//
// An empty Dir is treated as ".".
type Dir string
```

```go
fs := http.FileServer(http.Dir("static"))
```
これは`static`ディレクトリの中身をhttp上に流すハンドラを作成する関数。

```go
http.Handle("/static/", http.StripPrefix("/static/", fs))
```
これもさらに分解して見てみると

```go
http.Handle("/static/", ...)
```
この部分は、URL:`http://ドメインなどホスト部/static/`にアクセスがあった場合、第二引数のハンドラを呼ぶ、という関数。いわゆるルーティング設定をしています。
```go
http.StripPrefix("/static/", fs)
```
この部分はURLのPathの部分のうち、最初の`/static/`という文字列を削除した情報を第二引数のハンドラ、ここでいう`fs`に渡す、という物。

```go
	http.Handle("/", &Top)
	http.Handle("/form", http.HandlerFunc(AddMessage))
```
この部分はどちらも先ほどのものと同じくルーティング設定をしています。
```go
http.HandlerFunc(AddMessage)
```
は、いちいちServeHTTPメソッドを実装している型を定義する必要がない場合に
```go
func (w http.ResponseWriter, r *http.Request)
```
というインターフェースを実装している関数を直接ハンドラに変えてしまうための関数です。
今回は上記のメソッドを実装している`AddHandler`関数を代入しています。

```go
	if err := http.ListenAndServe("0.0.0.0:8080", nil); err != nil {
		log.Fatalln("Error")
	}
```
　ここで先ほどまでに設定したルーティング設定を実際に有効かしています。

### 動作
　それでは、実際に
`/`へリクエストがきた場合の挙動を確認してみましょう。

　まず、main関数のルーティング設定にしたがって、変数Top(TopPage型)の中の、ServeHTTPメソッドが実行されます。
```go
//TopPage Put data used at top page
type TopPage struct {
	Message    []Message
	template   *template.Template
	once       sync.Once
	Basic      HTMLbasic
	PresetName string
}

//HTMLbasic Put header and footer data
type HTMLbasic struct {
	Title string
	Name  string
}

//Message Put message data
type Message struct {
	Name string
	Text string
	Time time.Time
}

...

func (top *TopPage) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//templateのcompile作業は一回だけで良い
	top.once.Do(func() {
		//templateを定義
		t := template.New("index.html")

		//index.htmlの内容を読み込み（filepath.Joinは複数の環境に対応できるようにするため。）
		//template.MustはtemplateにErrorがあった場合にpanicを起こすための関数
		top.template = template.Must(t.ParseFiles(filepath.Join("resources", "index.html"), filepath.Join("resources", "header.html"), filepath.Join("resources", "footer.html")))
	})

	top.PresetName = r.URL.Query().Get("user")

	top.Basic = HTMLbasic{Title: "Webservice2020掲示板"}

	var messages []Message = make([]Message, len(Messages))
	copy(messages, Messages)

	switch {
	case DisplayNum > 0:
		top.Message = messages[int(math.Max(0, float64(DisplayNum-len(Messages)))):]
	case DisplayNum <= 0:
		top.Message = messages
	}

	//昇順降順の入れ替え
	for left, right := 0, len(top.Message)-1; left < right; left, right = left+1, right-1 {
		top.Message[left], top.Message[right] = top.Message[right], top.Message[left]
	}

	if err := top.template.ExecuteTemplate(w, "content", top); err != nil {
		fmt.Fprintf(w, "%v\n", err)
	}

	return
}
```

少しずつ見ていきましょう。

```go
	top.once.Do(func() {
		//templateを定義
		t := template.New("index.html")

		//index.html／header.html／footer.htmlの内容を読み込み（filepath.Joinは複数の環境に対応できるようにするため。）
		//template.MustはtemplateにErrorがあった場合にpanicを起こすための関数
		top.template = template.Must(t.ParseFiles(filepath.Join("resources", "index.html"), filepath.Join("resources", "header.html"), filepath.Join("resources", "footer.html")))
	})
```

　ここでは最初にresourcesフォルダに入っているindex.html／header.html／footer.htmlという名前のテンプレートデータを処理しています。この作業は一度だけ行えば良いため、syncパッケージの中のOnce関数を用い、複数回呼ばれても一回しか処理しない設定をしています。別に無理に使う必要はありませんが、大きなサービスを作るならば一回で済む処理は二度以上させないことで処理を軽くすることができるでしょう。

```go

	top.PresetName = r.URL.Query().Get("user")

	top.Basic = HTMLbasic{Title: "Webservice2020掲示板"}

	var messages []Message = make([]Message, len(Messages))
	copy(messages, Messages)

	switch {
	case DisplayNum > 0:
		top.Message = messages[int(math.Max(0, float64(DisplayNum-len(Messages)))):]
	case DisplayNum <= 0:
		top.Message = messages
	}

	//昇順降順の入れ替え
	for left, right := 0, len(top.Message)-1; left < right; left, right = left+1, right-1 {
		top.Message[left], top.Message[right] = top.Message[right], top.Message[left]
	}
```

　templateに埋め込むための変数データを各々代入しています。switch文の部分は
`DisplayNum`に設定された個数の投稿内容を表示するようにしています。

```go
	if err := top.template.ExecuteTemplate(w, "content", top); err != nil {
		fmt.Fprintf(w, "%v\n", err)
	}
	return
```

この部分で先ほどまでに設定した内容のtemplateを実際に展開して、できたデータをクライアントマシンに送信しています。

## 終わりに

[終わりに](/2nd/final.md)