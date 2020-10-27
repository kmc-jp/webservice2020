# テンプレートエンジン
[前にもどる](readme.md)
## 概要
いままではtext/templateをつかって虫食い状態のテンプレートを埋める方法をまなんできました。今度はhtmlを埋めていきましょう。

## 目次
<!-- TOC -->

- [テンプレートエンジン](#テンプレートエンジン)
    - [概要](#概要)
    - [目次](#目次)
    - [html/template](#htmltemplate)
        - [injection](#injection)
        - [html型](#html型)
    - [おわり](#おわり)

<!-- /TOC -->

## html/template
このパッケージの基本的な使い方はtext/templateと相違ありません。しかし、機能に差があります。大きな違いは、htmlのタグ情報などを入力しても、自動的にエスケープしてくれることです。これによって、injectionと呼ばれる攻撃に備えた    プログラムを作成することが可能です。

### injection
例えば次の例を見てみましょう。

```go
package main

import (
    "html/template"
    "os"
    "log"
)

func main() {
    var text string = `{{.}}`
    
    var t = template.New("template")

    t, err := t.Parse(text);

    err = t.Execute(os.Stdout, "<a href='#'></a>")
    if err != nil {
        log.Println("executing template:", err)
    }

    return
}
```

- <a href="https://play.golang.org/p/Bc71pzeTI_X">実行結果</a>

```
&lt;a href=&#39;#&#39;&gt;&lt;/a&gt;
```

ぐちゃぐちゃしていますね。これはhtml内に使われると表示する以外の用途にかかわる文字、例えば`<`や`>`などがエスケープされている結果です。このように自動的に処理してくれるため、プログラマは安心して実行することができるのです。

### html型
しかし、場合によってはあえてタグなどの情報を入れ込みたいかもしれません。その場合はhtml関数に渡すことでこれを実行できます。

```go
package main

import (
    "html/template"
    "os"
    "log"
)

func main() {
    var text string = `{{.}}`
    
    var t = template.New("template")

    t, err := t.Parse(text);

    err = t.Execute(os.Stdout, template.HTML("<a href='#'></a>"))
    if err != nil {
        log.Println("executing template:", err)
    }

    return
}
```

- <a href="https://play.golang.org/p/6cZJcLIgv9q">実行結果</a>

```
<a href='#'></a>
```

## おわり

html/templateの具体的な使用例は次回みる予定です。おつかれさまでした。