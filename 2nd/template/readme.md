# テンプレートエンジン

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

    var text string = "\{\{ .String \}\}"

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


