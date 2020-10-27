# テンプレートエンジン
[次にすすむ](html.md)
## 概要
前期には、HTMLの具体的な書き方を学びました。この基礎知識を踏まえて今度は動的なHTMLの生成をしていくことを考えます。そのために、テンプレートエンジンである、text/templateの使いかたから、最終的に、そのHTMLのテンプレートとして特化された版である、html/templateを見ていきます。

## 目次
<!-- TOC -->

- [テンプレートエンジン](#テンプレートエンジン)
    - [概要](#概要)
    - [目次](#目次)
    - [text/template](#texttemplate)
        - [基本例](#基本例)
        - [構造体を渡す](#構造体を渡す)
        - [メソッド](#メソッド)
        - [テンプレート内関数](#テンプレート内関数)
            - [関数マップの作成](#関数マップの作成)
                - [参考 - ...型名 -](#参考---型名--)
            - [関数の登録](#関数の登録)
            - [関数の呼び出し](#関数の呼び出し)
                - [参考 - Pipe -](#参考---pipe--)
        - [条件分岐](#条件分岐)
            - [if 関数](#if-関数)
            - [比較関数](#比較関数)
        - [繰り返し](#繰り返し)
            - [range関数](#range関数)
                - [参考 - テンプレート内変数 -](#参考---テンプレート内変数--)
        - [ファイルからの読み込み](#ファイルからの読み込み)
            - [単一ファイル](#単一ファイル)
            - [複数ファイル](#複数ファイル)
                - [ExecuteTemplate](#executetemplate)
            - [テンプレートの入れ子](#テンプレートの入れ子)
                - [define](#define)
                - [template](#template)
        - [おわり](#おわり)

<!-- /TOC -->

## text/template 

Go言語には、虫食い状態の文書ファイル(テンプレート)の穴を埋めることで動的に変化する文書を作成する機能を実装したパッケージとして、`text/template`パッケージが用意されています。まずはこの使いかたを見てみましょう。

### 基本例

```go
package main

import (
    "text/template"
    "os"
    "log"
)

func main() {
    var text string = `こんにちは、{{.}}さん`
    
    var t = template.New("template")

    t, err := t.Parse(text);

    err = t.Execute(os.Stdout, "tkmax777")
    if err != nil {
        log.Println("executing template:", err)
    }

    return
}
```

- <a href="https://play.golang.org/p/ZqD8C7U_07c">実行結果</a>

```
こんにちは、tkmax777さん
```

順番に見てみましょう。

```go
var text string = `こんにちは、{{.}}さん`
```

この部分はテンプレートデータを作成しています。結果からわかるように、Goではテンプレートの虫食いの穴を `{{` と `}}` で挟み込むことで表現します。`.`は、渡したデータそのものを指しています。

```go
var t = template.New("template")
```

テンプレート`template`を作成しています。<a href="#note_1">^1</a>

```go
t, err := t.Parse(text);
```

ここで、最初に作ったテンプレートデータの虫食い部分を解析しています。

```go
err = t.Execute(os.Stdout, "tkmax777")
if err != nil {
    log.Println("executing template:", err)
}
```

この部分でテンプレートに"tkmax777"というデータを渡して、虫食いをそれで埋めて、それを標準出力に出力させています。

1. テンプレートの生成(template.New())
2. テンプレートデータの読み込み／解析(template.Parse())
3. テンプレートにデータを渡し、埋めさせる

このような作業の流れは、今後複雑になっていっても基本は変わりません。覚えておきましょう。

<a id="note_1">^1</a>

### 構造体を渡す
先程の例ではたった一つのデータしか渡していませんでした。しかし実際には、もっと多くのデータを渡したいですね。そこで今度は構造体に複数のデータを含め、それを渡そうしていみましょう。

```go
package main

import (
    "text/template"
    "os"
    "log"
)

type Profile struct {
    Name    string
    Age     int
    Org     string
}

func main() {
    var profile Profile = Profile {
        Name:   "tkmax777",
        Age:    20,
        Org:    "KMC",
    }

    var text string = "こんにちは、{{.Name}}さん\n" +
                    "あなたのプロフィール\n" +
                    "名前: {{.Name}}\n" +
                    "年齢: {{.Age}}\n" +
                    "所属: {{.Org}}"

    var t = template.New("template")

    t, err := t.Parse(text);

    err = t.Execute(os.Stdout, profile)
    if err != nil {
        log.Println("executing template:", err)
    }

    return
}
```

- <a href="https://play.golang.org/p/0yZS3M7nBVR">実行結果</a>

```
こんにちは、tkmax777さん
あなたのプロフィール
名前: tkmax777
年齢: 20
所属: KMC
```

このようにテンプレートに構造体を渡した場合、`.`から`.Field1`のようにたどることで、取得することができます。

### メソッド
テンプレートに渡せるのは、構造体のデータそのものだけではありません。次のように、テンプレート上からメソッドを使用することも可能です。

```go
package main

import (
    "text/template"
    "os"
    "fmt"
    "log"
    "strings"
)

type Profile struct {
    Name    string
    Age     int
    Projects []string
}

func (p Profile) String() string {
    return fmt.Sprintf(
        "こんにちは、%[1]sさん\n" +
        "あなたのプロフィール\n" +
        "名前: %[1]s\n" +
        "年齢: %[2]d\n" +
        "プロジェクト:\n  " +
        strings.Join(p.Projects, "\n  "),
        p.Name, p.Age,
    )
}

func main() {
    var profile Profile = Profile {
        Name:   	"tkmax777",
        Age:    	20,
	Projects:	[]string{
		"WebService2020",
		"DTM練習会2020",
		"お絵描き練習会2020",
	},
    }

    var text string = "{{ .String }}"

    var t = template.New("template")

    t, err := t.Parse(text);

    err = t.Execute(os.Stdout, profile)
    if err != nil {
        log.Println("executing template:", err)
    }

    return
}
```

- <a href="https://play.golang.org/p/u-sqImX3tuX">実行結果</a>

```
こんにちは、tkmax777さん
あなたのプロフィール
名前: tkmax777
年齢: 20
プロジェクト:
  WebService2020
  DTM練習会2020
  お絵描き練習会2020
```

~~テンプレート使う意味のないプログラムということは置いておいて、~~
このように、通常メソッドを利用するように、`.Method`の形式で利用できます。もし引数をとるメソッドの場合、次のように空白で区切ることでメソッドに渡すことができます。

```
{{ .Method arg1 arg2 ... argn }}
```

便利ですね。

### テンプレート内関数

さらに、テンプレート内で使える関数を渡すこともできます。次の例を見てください。

```go
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"text/template"
)

var t = template.New("template")

type Profile struct {
	Name     string
	Age      int
	Projects []string
}

func (p Profile) String() string {
	return fmt.Sprintf(
		"こんにちは、%[1]sさん\n"+
			"あなたのプロフィール\n"+
			"名前: %[1]s\n"+
			"年齢: %[2]d\n"+
			"プロジェクト:\n  "+
			strings.Join(p.Projects, "\n  "),
		p.Name, p.Age,
	)
}

func main() {
	var profile Profile = Profile{
		Name: "tkmax777",
		Age:  20,
		Projects: []string{
			"WebService2020",
			"DTM練習会2020",
			"お絵描き練習会2020",
		},
	}

	var funcMap template.FuncMap = template.FuncMap{
		"num": func(args ...interface{}) string {
			var slice = args[0].([]string)
			return strconv.Itoa(len(slice))
		},
	}

	var text string = "{{ .String }}\nプロジェクト数合計: {{ num .Projects }}"

	var t = template.New("template")

	t = t.Funcs(funcMap)

	t, err := t.Parse(text)

	err = t.Execute(os.Stdout, profile)
	if err != nil {
		log.Println("executing template:", err)
	}

	return
}

```

- <a href="https://play.golang.org/p/VmjNSF4wk1R">実行結果</a><br>

```
こんにちは、tkmax777さん
あなたのプロフィール
名前: tkmax777
年齢: 20
プロジェクト:
  WebService2020
  DTM練習会2020
  お絵描き練習会2020
プロジェクト数合計: 3
```

今新しくプロジェクト数の合計を出力する機構を追加しました。変更した点はここです。

```go
    // 1
var funcMap template.FuncMap = template.FuncMap{
		"num": func(args ...interface{}) string {
			var slice = args[0].([]string)
			return strconv.Itoa(len(slice))
		},
	}

	var t = template.New("template")
    // 2
	t = t.Funcs(funcMap)

    // 3
	var text string = "{{ .String }}\nプロジェクト数合計: {{ num .Projects }}"

```

#### 関数マップの作成

`template.FuncMap`というのは、`map[string]interface{}`をtype節を用いて定義した型です。このマップのキーに関数名を入れ、その内容に関数を登録します。今回登録する関数の書式は次の通りです。

```go
func (args ...interface{}) string {}
```

この関数の返り値が、そのまま出力としてtemplate内で処理されます。

※今回は次の可変長列の話をかきたかったのでこの形式にしましたが、別に`[]string`で受けても問題ないです。

##### 参考 - ...型名 -
ここで、`...型名`という書式が登場しました。これは関数が可変長の引数を取るときに指定する書き方です。こうして入力された値はSliceとして指定した変数に格納されます。

例えば`fmt.Println()`という関数がありました。この関数は入力した値を全てまとめて表示してくれる関数でしたね。この関数も同様に`...interface{}`を引数としていて、これによって無数の引数を与えることができたのです。

- 例

```go

func Len(args ...interface) int {
    return len(args)
}

func main() {
    fmt.Printf("%d", Len("Yeah", "Foo", "Hoge", "Fuga")) // => 4
}

```

#### 関数の登録
template内に関数を登録します。

```go
template.Funcs(FuncMap)
```

注意しなくてはならないのは、Parseする前に登録する必要があるということです。注意しましょう。

#### 関数の呼び出し

テンプレート内で次のように記述することで、関数を呼び出すことができます。

```
{{ 関数名 引数1 引数2 ...}}
```

簡単ですね。

##### 参考 - Pipe -
引数の渡す方法として、次のような書き方もあります。

```
{{ 引数 | 関数名 }}
```

`|`(pipe)で区切って関数を書くと、その前までの出力内容ががそのまま関数の入力として渡されます。この記法は端末上でも使えるので覚えておきましょう。



### 条件分岐

次のように条件分岐をさせることができます。(ここではテンプレートだけ記述します。)

```
こんにちは、{{.Name}}さん\n
あなたのプロフィール\n
名前: {{.Name}}\n
年齢: {{.Age}}\n
所属: {{.Org}}\n
状態: {{if ge .Age 20 }} 成人 {{ else }} 未成年 {{ end }}
```

- <a href="https://play.golang.org/p/qZAO6uMr6Z8">実行結果</a>

```
こんにちは、tkmax777さん
あなたのプロフィール
名前: tkmax777
年齢: 20
所属: KMC 
状態:  成人 
```

#### if 関数
if 関数は入力がfalseまたはそれに準じるもの(例えば`0`や`""`など)を入力されたとき直後に記述された`{{end}}`までのテンプレート内容を飛ばします。もしそうではない場合、すなわちtrueないし、`0`や`""`以外が入力された場合は直後の内容を実行します。

```
{{if true}}表示される{{end}}
{{if "ほげほげ"}}表示される{{end}}
{{if false}}飛ばされる{{end}}
{{if 0}}飛ばされる{{end}}
```

また、ifとendの間にelseがあった場合は、falseまたはそれに準じる入力があった場合その箇所が実行されます。

#### 比較関数

次に比較です。比較は次のビルドイン関数を用います。

|関数|役割|
| --- | --- |
|eq A B|A == B|
|ne A B|A != B|
|lt A B|A > B|
|le A B|A => B|
|gt A B|A < B|
|ge A B|A <= B|

これらは状態に応じてtrue / falseを返します。そのため、これの前にif関数を記述し、入力させれば無事に条件分岐をすることができるという寸法です。

### 繰り返し
templateに入れたスライスの各要素について、操作を繰り返すことができます。

```
こんにちは、{{.Name}}さん\n
あなたのプロフィール\n
名前: {{.Name}}\n
年齢: {{.Age}}\n
所属: {{.Org}}\n
状態: {{if ge .Age 20 }} 成人 {{ else }} 未成年 {{ end }}
プロジェクト\n
{{ range $i, $data := .Projects }}
   {{$data}}
{{end}}
```

- <a href="https://play.golang.org/p/u-TVAYIiX_f">実行結果</a>

```
こんにちは、tkmax777さん
あなたのプロフィール
名前: tkmax777
年齢: 20
所属: KMC 
状態:  成人 
プロジェクト:
   WebService2020
   DTM練習会2020
   お絵描き練習会2020
```

#### range関数

```
{{ range $index, $data := .Slice_or_Map }}
```

とてもGo自体と似た記法なので、敢えて説明はそこまでいらないとは思います。ただ、変数名に$がついていますね。

##### 参考 - テンプレート内変数 -
template内で変数を別途定義することが可能です。この変数は`$`から始まる名称をつける必要があります。

```
{{ $text := `ほげ` }}
{{ $text }}
```

こういうことが可能です。

### ファイルからの読み込み

Golangでのテンプレートには、外部のファイルを用いることもできます。

#### 単一ファイル


- 構成

```
┬
├ main.go
└ template.txt
```

- template.txt

```
こんにちは、{{.Name}}さん\n
あなたのプロフィール\n
名前: {{.Name}}\n
年齢: {{.Age}}\n
所属: {{.Org}}\n
状態: {{if ge .Age 20 }} 成人 {{ else }} 未成年 {{ end }}
プロジェクト\n
{{ range $i, $data := .Projects }}
   {{$data}}
{{end}}
```


```go
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"text/template"
)

var t = template.New("template")

type Profile struct {
	Name     string
	Age      int
	Projects []string
}

func main() {
	var profile Profile = Profile{
		Name: "tkmax777",
		Age:  20,
		Projects: []string{
			"WebService2020",
			"DTM練習会2020",
			"お絵描き練習会2020",
		},
    }

	var t = template.New("template")

	t = t.Funcs(funcMap)

	t, err := t.ParseFiles("template.txt")

	err = t.Execute(os.Stdout, profile)
	if err != nil {
		log.Println("executing template:", err)
	}

	return
}
```

結果は変わりません。

何ら難しいことはなく、ParseがParseFilesにかわり、String文字列の代わりにファイル名を入力しただけです。簡単ですね。


#### 複数ファイル


- 構成

```
┬
├ main.go
└ template.txt
```

- japanese.txt

```
こんにちは、{{.Name}}さん\n
あなたのプロフィール\n
名前: {{.Name}}\n
年齢: {{.Age}}\n
所属: {{.Org}}\n
状態: {{if ge .Age 20 }} 成人 {{ else }} 未成年 {{ end }}
プロジェクト: \n
{{ range $i, $data := .Projects }}
   {{$data}}\n
{{end}}
```

- english.txt

```
Hello, {{.Name}}\n
Your Profile\n
Name: {{.Name}}
Age: {{.Age}}
Organization: {{.Org}}\n
Status: {{if ge .Age 20 }} Adult {{ else }} Child {{ end }}\n
Projects: \n
{{ range $i, $data := .Projects }}
   {{$data}}\n
{{end}}
```

- main.go

```go
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"text/template"
)

var t = template.New("template")

type Profile struct {
	Name     string
	Age      int
	Projects []string
}

func main() {
	var profile Profile = Profile{
		Name: "tkmax777",
		Age:  20,
		Projects: []string{
			"WebService2020",
			"DTM練習会2020",
			"お絵描き練習会2020",
		},
    }

	var t = template.New("")

	t, err := t.ParseFiles("japanese.txt", "english.txt")

	err = t.ExecuteTemplate(os.Stdout, "japanese", profile)
	if err != nil {
		log.Println("executing template:", err)
	}

    err = t.ExecuteTemplate(os.Stdout, "english", profile)
	if err != nil {
		log.Println("executing template:", err)
	}

	return
}

```

  ・実行結果

```
こんにちは、tkmax777さん
あなたのプロフィール
名前: tkmax777
年齢: 20
所属: KMC 
状態:  成人 
プロジェクト:
   WebService2020
   DTM練習会2020
   お絵描き練習会2020
Hello, tkmax777さん
Your Profile
Name: tkmax777
Age: 20
Organization: KMC 
Status:  Adult 
Projects:
   WebService2020
   DTM練習会2020
   お絵描き練習会2020

```

##### ExecuteTemplate
今までExecuteだったものがExecuteTamplateになりました。Executeは一番はじめに入力されたテンプレートを自動的に選択し、展開してくれます。対してExecuteTemplate関数は、テンプレートを指定することができます。ここで、テンプレートの名前はファイル名になります。そのため、ここでjapaneseを呼びだせば、日本語のものがでてきて、englishにすると、英語のものがでてきたというわけです。

#### テンプレートの入れ子
次のようにテンプレートを入れ子状に定義することもできます。

- profile.txt

```
{{define "profile"}}

こんにちは、{{.Name}}さん\n
あなたのプロフィール\n
名前: {{.Name}}\n
年齢: {{.Age}}\n
所属: {{.Org}}\n
状態: {{if ge .Age 20 }} 成人 {{ else }} 未成年 {{ end }}
プロジェクト: \n
{{ range $i, $data := .Projects }}
   {{$data}}\n
{{end}}

{{end}}
```

- date.txt

```
{{define "date"}}

今日は{{.Month}} 月 {{.Day}} 日です。

{{end}}
```

- template.txt

```
{{define "index"}}

{{template "profile" .}}

{{template "date" .Day}}

{{end}}
```

- main.go

```go
package main

import (
	"fmt"
	"log"
	"os"
    "strconv"
    "time"
	"strings"
	"text/template"
)

var t = template.New("template")

type Profile struct {
	Name     string
	Age      int
    Projects []string
    Day      Day
}

type Day struct { 
    Month   int
    Day     int
}

func main() {
	var profile Profile = Profile{
		Name: "tkmax777",
		Age:  20,
		Projects: []string{
			"WebService2020",
			"DTM練習会2020",
			"お絵描き練習会2020",
        },
        Day: Day {
            Month: int(time.Now().Month()),
            Day: int(time.Now().Day()),
        },
    }

	var t = template.New("")

	t = t.Funcs(funcMap)

	t, err := t.ParseFiles("profile.txt", "date.txt", "index.txt")

	err = t.ExecuteTemplate(os.Stdout, "index", profile)
	if err != nil {
		log.Println("executing template:", err)
	}

	return
}

```

- 出力

```
こんにちは、tkmax777さん
あなたのプロフィール
名前: tkmax777
年齢: 20
所属: KMC 
状態:  成人 
プロジェクト:
   WebService2020
   DTM練習会2020
   お絵描き練習会2020
今日は 10 月 27 日です。
```

##### define

テンプレートファイル内でテンプレートの名前を指定できます。書式は次の通りです。

```
{{ define "Tenplate_Name" }}
```

##### template

テンプレートを呼び出す関数です。書式は次のとおりです。

```
{{ template "Template_Name" 格納する変数}}
```

### おわり
以上がtext/templateの使いかたです。これでも十分便利ですが、Golangではhtmlの虫食いを埋めるのに特化したhtml/templateというパッケージが用意されています。このパッケージの使いかたを今度は学んでいきましょう。

[html/template](html.md)