# テンプレートエンジン

## 概要
前期には、HTMLの具体的な書き方を学びました。この基礎知識を踏まえて今度は動的なHTMLの生成をしてみます。

## 目次
<!-- TOC -->

- [テンプレートエンジン](#テンプレートエンジン)
    - [概要](#概要)
    - [目次](#目次)
    - [text/template](#texttemplate)
        - [基本例](#基本例)

<!-- /TOC -->

## text/template 

Go言語には、虫食い状態の文書ファイル(テンプレート)の穴を埋めることで動的に変化する文書を作成する機能を実装したパッケージとして、`text/template`パッケージが用意されています。まずはこの使いかたを見てみましょう。

### 基本例

```go
package main

import (
    "text/template"
    "fmt"
)

func main() {
    var text string = `
    `


    // 新たなテンプート、templateを作成
    var t = template.New("template")

    //
}