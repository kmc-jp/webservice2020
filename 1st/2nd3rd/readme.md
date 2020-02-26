# 第二回
## 概要
- プログラミングの基礎／Golangの基礎を学んでもらう

## プログラミング始める前に...
そもそもプログラムってなんなのかを考えてみる。
- コンピュータ：電子計算機
  - 計算をさせる。
- プログラミングとは？
  - コンピュータープログラミング（英語: Computer Programming）とは、ある特定のコンピューティングの結果を得ることを目的として、実行可能なコンピュータープログラムを設計・構築するプロセスのことである。(Wikipedia)
  - 計算の"させ方"を記す。
    - "関数"を書いていく。 
- 関数
  - ある入力に対して何かしらの出力を成すもの。
  - 数値とは限らない。
    - 例：文字列、真偽などなど

**僕の見解**
- プログラムとは、ある問題に対して電子計算機が解を導くための関数のこと。

　これから僕たちが作ろうとしているのは、ユーザに対してどのようなWebPageを表示するのかという問題に対して適切なhtmlファイルなどを出力する関数ですね。

## Golangの基礎
### Golangの世界でHello World!
実際にプログラムを動かしてみる。

```go
package main

func main(){
	print("Hello World!\n")
	return
}

```

コマンドラインで`$ go run main.go`と入力してみる。

```sh
$ go run main.go
Hello World!
```
このように出たら成功！


```go
package main
```
ここはおまじないです。一体なんなのか。近いうちにわかります。

```go
func main() {
     ...
     return
}
```
- `func`	functionです。先ほどの話で出てきましたね。関数です。
- `main()`	mainはこの関数の名前です。
  - cf)関数y = f(x)
  - ()はこの関数には何も値を入れていない状態
- `return`	この関数の最終的な結果を返すことを示しています。

**参考**

main関数では最後のreturnでOSに対して終了したという旨(0)を返している。

→自明なことなためにGolangでは省略している。

```go
print("Hello World!\n")
```
- `print()`	標準出力に出力する関数
- `""`		文字列であることを明示する括弧
- `\n`		改行コード。改行する旨を指示。

☆ 同様に`print("hogehoge\n")`を増やしていくことで文字列をどんどん出力できる。

### 計算をしてみる
```go
package main

func main(){
	println(1 + 1)
	return
}
```

- `println`	`print`と違い、末尾に自動的に改行が挿入される。(print line)
このように計算をすることもできる。


### 変数

```go
package main

func main(){
	var x int
	x = 1
	
	print(x) //1
	return
}
```

コマンドラインで`$ go run main.go`と入力してみる。

- `var x int`
  - `var`	変数を宣言することを示唆
  - `x`		変数の名前
  - `int`	型

☆このプログラムではxという変数を整数型で宣言し、その変数にxを代入し、その内容を出力している。

**型とは？**
- 計算機にとって、変数の中にどんなものが入るのかはとても重要。
  - 文字列、数字、真偽など。入るものによって扱いが違う
    - 例：1 + "Hello!" : 明らかにおかしい。
- そこで型という概念を導入することで、入るものの種類を決めておく

**参考**

C系の言語に触れたことがある人は違和感を持つでしょう。何故`int x`ではなく`x int`なのか...。

興味があれば[公式ページに解説](https://blog.golang.org/gos-declaration-syntax)があります。

**代表的な型**

|型|内容|
| ---- | ---- |
|int|符号付整数|
|float32|浮動小数点|
|string|文字列|
|bool|真理値|

他にもcomplex64(複素数)，※byte(文字列)など沢山ある。

※byteはstringとは保存される文字列形式が違う。(UTF-8/Unicode)

### 関数を作ってみよう

```go
package	main

func main() {
	var ans int
	ans = compute(1, 1)
	print(ans, "\n") //2
}

func compute(x int, y int) int {
	return x + y
}
```

- `func compute(x int, y int) int {`	computeという関数を宣言(cf:f(x, y) = z)
  - `compute(x int, y int)`    	  	整数x,yを与えている。
  - `int`				整数が出力されることを示唆。
- `return x + y`			x + yの結果を返す。

- `ans = compute(1, 1)`			上記のx, yに各々1, 1を代入した答えを変数ansに代入している。

### packageを理解しよう

```go
package main

import "fmt"

func main() {
     var year int = 2020
     fmt.Printf("今年は%d年です。\n", year) //今年は2020年です。
     return
}

```

- `import "fmt"`	fmtというpackageを読み込むことを示唆。

**package**

→関数の集合。先人の知恵ですね。

- fmt.Printf(...)		fmt package中のPrintf関数を呼び出している。
  - printよりも多くのことができる。上位互換。
- ("...は%d年...", year)
  - %dはのちの引数(=ここでいうyear)を整数として表示することを意味する。

**書式指定子**

→ `%...`のこと。
ものすごく沢山の種類がある。
代表的なものには、`%s`:文字列`%v`:標準書式`%g`:浮動小数点数がある。

[ここ](https://qiita.com/rock619/items/14eb2b32f189514b5c3c)をみるとめちゃくちゃ書いてあるよ。興味あれば読むといいかもね。

何が嬉しいかというと、例えば数値なら桁数指定とかできる。

`fmt.Printf("%02d:%02d\n",12,5) //12:05`

ここまでわかれば、[A Tour of Go](/1st/2nd3rd/go_tour.md)
も始められます！