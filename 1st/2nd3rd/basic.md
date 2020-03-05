<!-- TOC -->

- [フロー制御](#フロー制御)
    - [If文 -条件分岐-](#if文--条件分岐-)
    - [if...else](#ifelse)
    - [For文 -繰り返し処理-](#for文--繰り返し処理-)
    - [switch文](#switch文)
    - [defer文](#defer文)

<!-- /TOC -->

## フロー制御

### If文 -条件分岐-
状況によって実行するブロックを選ぶことができる。
```go
package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))// => 1.4142135623730951 2i
}
```

if文の構成

`if condition {...}`

- conditoin : 論理値（`true`or`false`）

`true`のとき、指定したブロックが実行される。

**論理演算子**

|演算子|内容|
| ---- | ---- |
|a == b|a = bのとき|
|a != b|a ≠ bのとき|
|a < b|a < bのとき|
|a > b|a > bのとき|
|a <= b|a ≦ bのとき|
|a  >= b|a ≧ bのとき|

**二つ以上の条件を比較する時**

and : `&&`
or : `||`

例:
```go
fmt.Printf("%t %t\n", 1 == 1 || 1 == 2, 1 == 1 && 1 == 2)// => true false
```
(`%t`:論理値を代入する書式指定子)

### if...else
if文の条件に適合しなかった場合についてはelse文で分岐できる。

```go
func main() {
	var a string
	if 1 == 2 {
		a = "実行されない"
	} else {
		a = "実行される"
	}
	fmt.Printf("%s\n", a)
}
```
`else if`でさらに分岐できる。
```go
func main() {
	var a string
        if 1 == 2 {
                a = "実行されない"
        } else if 1 == 1 {
                a = "実行される"
	} else {
		a = "実行されない"
	}
        fmt.Printf("%s\n", a)
}
```

### For文 -繰り返し処理-
ブロックの繰り返し処理を行うときに用いることができる。

```go
package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)//45
}
```

for文の構成要素

`for init; condition; post {...}`
- `init`:初期化ステートメント: 最初のイテレーション(繰り返し)の前に初期化が実行されます
- `condition`:条件式: 継続条件。繰り返し毎に評価されます
- `post`:後処理ステートメント: 繰り返し毎の最後に実行されます

なので今回の例の関数ならば、
1. iという0値変数を宣言
2. iが10未満の間繰り返す
3. 繰り返しごとにiに1を足し合わせる。

という意味になる。

**インクリメントとデクリメント**

1を足したり引いたりする作業は頻繁に行うので次のような略記が許されている。
- インクリメント

int++：1を足す（`i++` ↔︎ `i = i + 1` ↔︎ `i += 1`）
- デクリメント

int--：1を引く（`i--` ↔︎ `i = i - 1` ↔︎ `i -= 1`）

☆for文構成要素`init`, `post`は省略することができる。

例：
```go
package main

import "fmt"

func main() {
	sum := 1
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum) // => 1024
}
```
こうなると、もはやfor分の`;`は意をなしていないので省略できる。
```go
for sum < 1000 {...}
```

**参考**

> この記法はCなどの言語におけるwhile文に相当します。

さらに、`condition`節も省略することができる。この時、この`condition`節には自動的に`true`が代入され、無限ループ処理となる。

`for {...}`↔︎`for true {...}`

**継続／脱出フラグ**

- `continue`と書くことで、loopの最後まで達していなくても次のloopに移行できる。

例：
```go
for i := 0; i < 2; i++ {
    if i == 0 {
       continue
    }

    fmt.Printf("%d\n", %d)// => 1
}
```

- `break`と書くことで、本来forから脱出できない条件下でも脱出できる。

例：
```go
for {
    fmt.Printf("実行\n")
    break
}

fmt.Printf("forからの脱出\n")
```

**参考 - 他段階脱出**

forを重ねがけしているとき、Labelをつけていることで複数個のforを跨いで脱出できる。

例：
```go

First:
	for {
		for {
			break First
		}
	}

	fmt.Printf("多段階脱出完了！")
	
```

ここまで来たら、試しに次の練習問題を解いてみましょう

[A Tour of Go](https://go-tour-jp.appspot.com/flowcontrol/8)

### switch文
`if else`が頻発するような場合わけをするとき、Switch文を使えば可読性をあげることができます。

例：
```go
func main() {
	var i int
	
	fmt.Scanf("%d", &i)

	switch i {
	case 0:
		fmt.Printf("虚無\n")
	case 1:
		fmt.Printf("素数ではない。\n")
	default:
		fmt.Printf("素数かもしれない。\n")
	}
}
```

これは次と同値：
```go
func main() {
        var i int

        fmt.Scanf("%d", &i)

	if i == 0 {
	   fmt.Printf("虚無\n")
	} else if i == 1 {
	   fmt.Printf("素数ではない。\n")
	} else {
	   fmt.Printf("素数かもしれない。\n")
	}
}
```

また、goのSwitch文はさらに次のように、caseに条件式を持ってくる拡張を認めています。

```go
func main() {
        var i int

	fmt.Scanf("%d", &i)
	switch {
	case i == 0:
		fmt.Printf("虚無\n")
	case i == 1:
		fmt.Printf("素数ではない。\n")
	case i < 0:
		fmt.Printf("負の整数\n")
	default:
		fmt.Printf("素数かもしれない。\n")
	}
}

```
とても便利ですね！

**参考**

> Cではcaseごとにいちいちbreakする必要がありますが、Goではその必要はありません。
> 
> あえてCのような挙動にする場合、各case末尾に`fallthrough`を書く必要があります。

### defer文
```go
//OpenFile 指定したfilenameのファイルの内容を出力
func OpenFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("File open error\n%v", err)
	}
	defer file.Close()

	var b []byte = make([]byte, 10)
	for {
		n, err := file.Read(b)
		if err == io.EOF {
			break
		}
		fmt.Printf("%s", b[:n])
	}
	return
}

```

- defer文に書いた関数は関数から抜け出すときに必ず実行される。
  - 上記の例では開いたファイルが関数から抜け出すときに必ず閉じられる。

**Point**

defer文内の関数f(x, y, ...)の評価はdeferが呼ばれた時点で行われる。

例:
```go

func main() {
     var i int = 0
     
     defer fmt.Printf("i = %d\n", i)
     i = 1
     fmt.Printf("iに%dを代入しました。", i)
}

```
実行結果
```sh
iに1を代入しました。
i = 0
```
- 今回ではdeferを実行した時点ではiは0だったため、`i = 0`だったときの内容が出てきたというわけです。

**Stacking defers**

関数内でdeferを複数回読んだとき、関数終了後に呼び出されるdefer節は実行した逆順になる。

例：
```go
func main(){
     defer fmt.Println("First")
     defer fmt.Println("Second")
}
```
実行結果
```sh
Second
First
```

ここまでで、基本的な関数の作成はできるようになりました。ここからは、さらに変数などに関して掘り下げていきます。

[続き](/1st/2nd3rd/basic2.md)


