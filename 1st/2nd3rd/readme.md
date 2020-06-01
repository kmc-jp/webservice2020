# 第二／三回

## 概要
- プログラミングの基礎／Golangの基礎を学んでもらう

## 目次

<!-- TOC -->

- [第二／三回](#第二／三回)
    - [概要](#概要)
    - [目次](#目次)
    - [プログラミング始める前に...](#プログラミング始める前に)
    - [Golangの基礎の基礎](#golangの基礎の基礎)
        - [Golangの世界でHello World!](#golangの世界でhello-world)
        - [計算をしてみる](#計算をしてみる)
        - [変数](#変数)
        - [関数を作ってみよう](#関数を作ってみよう)
        - [packageを理解しよう](#packageを理解しよう)
        - [その他](#その他)
            - [public／private 関数・変数について](#public／private-関数・変数について)
            - [functionについて](#functionについて)
            - [変数の宣言の簡略化](#変数の宣言の簡略化)
            - [グローバル変数](#グローバル変数)

<!-- /TOC -->

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

## Golangの基礎の基礎

### Golangの世界でHello World!
実際にプログラムを動かしてみる。
Visual Studio Codeで新たなファイルを作り、以下のコードを写す。

```go
package main

func main(){
	print("Hello World!\n")
	return
}

```

ターミナル → 新しいターミナルを開き、出てきたコンソールに`$ go run main.go`と入力してみる。

**Terminal**

計算機をコマンドを入力することで操作するインターフェース。

慣例的にコマンドラインを入力内容を記述するとき、`$`の後にコマンドを記述します。

実際に書くときは`$`は書かずに入力してください。

```sh
$ command
```
→commandとだけ入力


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
- `func`	functionの略です。先ほどの話で出てきましたね。これから関数を定義することを表しています。
- `main()`	mainはこの関数の名前です。
  - cf)関数f(x)
  - ()はこの関数には何も値を入れないということ。
- `return`	この関数の最終的な結果を返すことを示しています。

**スコープ**

変数が有効な範囲を**スコープ(scope)**と呼ぶ。

Golangでは`{}`で囲まれた範囲（**ブロック**）がこれに相当する。

```go
print("Hello World!\n")
```
- `print()`	標準出力に出力する関数
- `""`		文字列であることを明示する括弧
- `\n`		改行コード。改行する旨を指示。

☆ 同様に`print("文字列\n")`を増やしていくことで文字列をどんどん出力できる。

**コメント機能**

プログラムにどういう作業をしているかなどの注釈を残したいとき、コメントを残す機能があります。
1. `//+文字列`
2. `/* 文章（改行あり）*/`

```go
  printf("文字列\n") //文字列を表示
  /*この後には
    複数行に跨がる
    コメントを残せます。*/
```

☆以降、出力を以下の形式で記述します。

```go
  fmt.Printf("出力") // => 出力
```

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

**演算子一覧**
|演算子|内容|
| ---- | ---- |
|+|加算|
|-|減算|
|*|乗算|
|/|除算|
|%|余りの計算|

### 変数

```go
package main

func main(){
	var x int
	x = 1
	
	print(x) // => 1
	return
}
```

コマンドラインで`$ go run main.go`と入力してみる。

- `var x int`
  - `var`	変数を宣言することを示唆
  - `x`		変数の名前
  - `int`	型

☆このプログラムではxという変数を整数型で宣言し、その変数に1を代入し、その内容を出力している。

**型とは？**
- 計算機にとって、変数の中にどんなものが入るのかはとても重要。
  - 文字列、数字、真偽など。入るものによって扱いが違う
    - 例：`1 + "Hello!"` : 明らかにおかしい。
- そこで型という概念を導入することで、入るものの種類を決めておく

**初期値**
- 変数は宣言した時点で初期化が行われる。
  - 数値型(int,floatなど): 0
  - bool型: false
  - string型: "" (空文字列( empty string ))

```go
var x int
var t bool
var str	string
fmt.Printf("x:%d t:%t str:%q", x, t, str)// => x:0 t:false str:""
```
参考：`fmt.Printf(...)`は後述で詳しくやります。今はとりあえず`%...`という書式を使うことでのちの変数の中身を表示できる物と捉えてください。

**型変換**
- int型とfloat型のような数値同士の計算が必要になった場合、型の変換を行うことができる。
```go
var i int = 42
var f float64 = float64(i)

fmt.Printf("i:%d f:%f\n", i, f)// => i:42 f:42.000000
```
以上のように、`型(変数)`で型を変化させることができる。

**参考**

> C系の言語に触れたことがある人は違和感を持つでしょう。何故`int x`ではなく`x int`なのか...。

> 興味があれば[公式ページに解説](https://blog.golang.org/gos-declaration-syntax)があります。

**代表的な型**

|型|内容|
| ---- | ---- |
|int|符号付整数|
|float32|浮動小数点|
|string|文字列|
|bool|真理値|

他にもcomplex64(複素数)，※byte(文字列)など沢山ある。

※byteとスライスは本質的には同じ物。ただし文字列を扱うだけなら基本的にstringで扱う。

**変数の有効範囲**

変数は宣言したスコープ（`{}`）内のみで有効です。

### 関数を作ってみよう

```go
package	main

func main() {
	var ans int
	ans = compute(1, 1)
	print(ans, "\n") // => 2
}

func compute(x int, y int) int {
	return x + y
}
```

- `func compute(x int, y int) int {`	computeという関数を宣言(cf:f(x, y) = x + y, x, y∈**Q**)
  - `compute(x int, y int)`    	  	整数x, yを与えている。
  - `int`				整数が出力されることを示唆。
- `return x + y`			x + yの結果を返す。

- `ans = compute(1, 1)`			上記のx, yに各々1, 1を代入した答えを変数ansに代入している。

☆ここでx, yのことを**引数（ひきすう）**，`x + y`のことを**返り値（戻り値）**と呼ぶ。

### packageを理解しよう

```go
package main

import "fmt"

func main() {
     var year int = 2020
     fmt.Printf("今年は%d年です。\n", year) // => 今年は2020年です。
     return
}

```

- `import "fmt"`	fmtというpackageを読み込むことを示唆。

**package**

→関数などの集合。先人の知恵ですね。

各packageの中の関数を使うときは`packageNAME.functionNAME`の形で実行する。

☆複数packageをimportするときは、二種類の書き方がある。

1. import を列挙
```go
import "fmt"
import "net/http"
```
2. factored import（集合化）
```go
import(
	"fmt"
	"net/http"
)
```

どちらでもいいですが、後者の方が基本的に書きやすいし、みやすいでしょう。

- `fmt.Printf(...)`		fmt package中のPrintf関数を呼び出している。
  - `print`よりも多くのことができる。上位互換。
- `("...は%d年...", year)`
  - `%d`はのちの引数(=ここでいう`year`)を整数として表示することを意味する。

**書式指定子**

→ `%...`のこと。
ものすごく沢山の種類がある。

代表的なもの

|書式指定子|種類|
| ---- | ---- |
|`%s`|文字列|
|`%q`|エスケープした文字列|
|`%v`|標準書式|
|`%f`|実数|
|`%t`|論理値|

[ここ](https://qiita.com/rock619/items/14eb2b32f189514b5c3c)をみるとめちゃくちゃ書いてあるよ。興味あれば読むといいかもね。

何が嬉しいかというと、例えば数値なら桁数指定とかできる。

```go
fmt.Printf("%02d:%02d\n",12,5) // => 12:05
```

### その他

#### public／private 関数・変数について
   - packageを作る際、外部から参照できるか否かを指定できる。
   - 最初の文字が大文字：Public、package外部から参照できる。
   - 最初の文字が小文字：Private、package外部からは参照されない。

```go
package tkmax

import "fmt"

//DisplayTKMAX Print "TKMAX"
// This function can be used by external functions
func DisplayTKMAX() {
     fmt.Printf("TKMAX\n")
     return
}

//displayTKMINIMUM Print"TKMINIMUM"
// This function can be used only by internal functions
func displayTKMINIMUM() {
     fmt.Println("TKMINIMUM\n")
     return
}

```

#### functionについて
- 幾らか便利な機能がある。
1. returnの省略
  - 何も値を返さない関数(`func () return`型)に関してはreturnを書かなくても良い。
  - あえて明示的に書いて、そこで関数を終わらせることもできる。

例：
```go
func main() {
     print("Hey YO!") // => Hey YO!
     return
     print("Bye!")//呼ばれない。
}
```

2. returnする変数の事前宣言
  - 関数を宣言するときに、返す変数を予め決められる。

例：
```go
func AddInt(x, y int) (z int) {
     z = x + y
     return
}
```
☆この書き方ではzをreturnすることは自明なのでreturn zと書く必要はない。
- ただしreturnは書いてね。

#### 変数の宣言の簡略化

標準の書き方
```go
var x int = 1
```
xに1を代入するため、int型であることが推定できる。

そのため次のようにも書ける。
```go
var x = 1
```
さらにいちいちvarを書くのが面倒なので次のようにも書ける。
```go
x := 1
```
とても柔軟ですね！

**注意**
1. Golangでは使わない変数は宣言することが許されていない
2. 省略は関数内でのみ行うことができる。

- ちなみにこんな感じにfactored宣言もできる。

```go
var (
    x int = 2
    y int = 3
)
```

#### グローバル変数
- 関数の外で変数／定数を宣言すると任意の関数から値を参照できる。
```go
package main

import "fmt"

var x int = 1
const c int = 2

func main() {
     fmt.Printf("x:%d c:%d", x, c)// => x:1 c:2
}
```
**注意**

1. 宣言の省略型を用いることはできない
2. 大文字から始まる変数はpackage外からも参照できる。

`const`:定数

定数はビルドするときに一緒に記録されるのでより正確な値を保持することができる。

その他、この値は他の関数によって書き換わらないことが保証されるので安心できる。

さらに詳しくGoを見ていきましょう！
[続き](./basic.md)
