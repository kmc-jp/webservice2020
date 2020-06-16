[前に戻る](./basic.md)
## 目次
<!-- TOC -->

- [目次](#目次)
- [Pointers](#pointers)
    - [ポインタとは？](#ポインタとは)
    - [使い方](#使い方)
    - [実用例](#実用例)
- [Type節](#type節)
- [method](#method)
- [Array（配列）](#array配列)
    - [基本形](#基本形)
- [Slice（スライス）](#sliceスライス)
    - [基本形](#基本形-1)
    - [新たな要素を追加するとき](#新たな要素を追加するとき)
    - [配列やスライスからの切抜](#配列やスライスからの切抜)
    - [make関数](#make関数)
    - [for range](#for-range)
    - [注意](#注意)
    - [発展 入れ子](#発展-入れ子)
- [map](#map)
    - [記法](#記法)
        - [stringとintを紐付けた例](#stringとintを紐付けた例)
    - [要素の追加](#要素の追加)
    - [要素が存在するか判別](#要素が存在するか判別)
    - [要素の削除](#要素の削除)
    - [参考 for range](#参考-for-range)
    - [発展 入れ子](#発展-入れ子-1)
- [Struct（構造体）](#struct構造体)
    - [記法](#記法-1)
    - [参照方法](#参照方法)
    - [初期化](#初期化)
- [TypeとStruct](#typeとstruct)
- [interface(インターフェース)](#interfaceインターフェース)
- [最後に...](#最後に)

<!-- /TOC -->

## Pointers

### ポインタとは？
変数の値などを一時的に記録するのに、計算機ではメモリという装置を持ちいる。
- いわゆるRAMってやつですね。
- この装置は高速にデータの書き込み／読み出しができる
  - ただし電源を落とすと同時にデータは消えてしまう。

☆全ての変数には値を格納する前に値の保存先となるメモリの住所が割り当てられる。
- この住所、通称"番地"をポインタという。

### 使い方
```go
func main() {
     var p *int
     var i int

     i = 1 //変数iに1を代入
     p = &i//ポインタ変数pにiのメモリ番地を代入
     fmt.Println("p値：", p)
     fmt.Println("i値：", i)

     *p = 2//pの番地に2を代入
     fmt.Println("*p値：", *p)//pに格納されている値を参照
     fmt.Println("i値：", i)//元データであるiを参照

     return
}

```
実行結果
```sh
p値： 0xc0000a82a0
i値： 1
*p値： 2
i値： 2
```
（p値の`0x...`は実行する度に変化します。）

1. 番地を格納する変数は`var 変数名 *型名`の形で宣言する。
2. 番地に格納されている値を読み出すときは`*変数名`と記述する。
3. 一般の変数の番地は`&変数名`という形で読み出すことができる。

☆上記の例では変数iの番地をpに格納し、pを通してiを編集しています。
  - 0x...はメモリ番地を１６進数で表した物です。

### 実用例
```go
func main() {
     var a int
     var b int

     a = 0
     b = 0

     EditVar(a, &b)

     fmt.Printf("a:%d\nb:%d", a, b)
     return
}


func EditVar(m int, n *int) {
     m = m + 1
     *n = *n + 1
     return
}

```

実行結果
```sh
a:0
b:1
```

　この例において、aは値をそのままEditVar関数に、bはそのメモリ番地をEditVar関数に渡しています。このとき、それぞれEditVar関数ではaはmという名前の整数型変数に、bはnという名前の整数型ポインタ変数に各々内容が格納されています。そのため、mの値を変化させてもaには一切変化が生じず、その反面、メモリ番地で指定したnの値を変化させると、それと同時にbの値も変化していることが観察されました。

## Type節
```go
type Name string

func main(){
     var user Name

     user = "matu"
     fmt.Printf("%s：%t\n", user, verify(user))
     
     user = "tkmax777"
     fmt.Printf("%s：%t\n", user, verify(user))
}

func verify(user Name) bool {
     return user == "tkmax777"
}
```

実行結果
```sh
matu：false
tkmax777：true
```
- `type 名前 型名`記述することで、新たな型を定義することができる。

☆専用の型名を用意しておくことで、コンパイル時に記述ミスに気づくことができる。

例：
```go
//先ほどの例のmain関数に追記
     var club string
     club = "KMC"
     
     fmt.Printf("%s：%t\n", user, verify(club))//コンパイルエラー
```
- `verify()`の引数はName型でないといけないため、エラーが発生する。

このように型を作っておくことで、より信頼性の高いプログラムを作成することができる。

## method
Golangでは自分で定義した型に対し、メソッドと呼ばれる操作関数を定義できる。

例：
```go
package main

import "fmt"

type Akari string

func(a Akari) String() string{
     return fmt.Sprintf("わぁい%s あかり%s大好き\n", string(a), string(a))
}

func main() {
     var food Akari = "うすしお"
     fmt.Printf(food.String())// => わぁいうすしお あかりうすしお大好き
     return
}

```

`food.String()`：変数`food`に対して型`Akari`に定義されたメソッド`String`を呼び出している。

```go
func(a Akari) String() string{...}
```
ここでメソッドを定義している。

・記法
```go
func(変数名 自分で定義した型) メソッド名(引数) 返り値の型 {...}
```
で記述する。ここで、最初の括弧内の自分で定義した型に対してメソッドが定義される。

この、最初の括弧で宣言された変数（これを**レシーバ**と呼ぶ）に、メソッドを呼び出した変数の値（先例では`food`）が代入される。

☆それ以外の動作は関数と同じ。

**補足**

1. fmtパッケージで`func(...)String() string{...}`メソッドを実装している物をPrintしようとした時、自動でこのメソッドが実行される。
```go
//先ほどのmain関数を変える。
func main() {
     var food Akari = "うすしお"
     print(food)// => うすしお
     print("\n")
     fmt.Printf(food)// => わぁいうすしお あかりうすしお大好き
     return
}
```

2. レシーバにポインタを指定することもできる。

例：
```go
package main

import (
     "math"
     "fmt"
)

type Sqrt float64

func(s *Sqrt) Sqrt(){
     *s = math.Sqrt(*s)
     return
}

func main() {
     var x Sqrt = 2

     fmt.Printf("x:%g", float64(x))// => 2.0000000000000000
     x.Sqrt()
     fmt.Printf("x:%g", float64(x))// => 1.4142135623730951
     return
}
```
この例ではSqrtメソッド内でレシーバにポインタを指定することで、その実体を上書きしています。

## Array（配列）
### 基本形
```go
func main(){
     var user [2]string
     user[0] = "tkmax777"
     user[1] = "matu"
     
     fmt.Printf("user1:%s user2:%s", )// => user1:tkmax777 user2:matu
     return
}
```
- `[配列の長さ]型名`という型を持つ。`型名`のところには任意の型が入る。
- 参照するときは`変数名[参照番号]`の形で参照できる。
  - **参照番号は0, 1, 2...と数えていくことに注意。**

また、次のように初期化することもできる。
```go
     var user [2]string
     user = [2]string{"tkmax777", "matu"}
     fmt.Printf("user1:%s user2:%s\n", user[0], user[1])// => user1:tkmax777 user2:matu
```

**Point**

Goにおいて配列は、その長さを型に含むため、後から長さを変更することができない。（固定長）

☆配列は次に挙げるスライス（Slice）があるので基本的にそんなに使わない。

## Slice（スライス）
### 基本形
```go
func main(){
     var userSlice []string = []string{"tkmax777", "matu"}

     fmt.Printf("%v", userSlice)// => [tkmax777 matu]
     fmt.Printf("このスライスの容量は%dです。\n", cap(userSlice))// => このスライスの容量は2です。
}
```
- `[]型名`という型を持つ
- 配列の参照型
  - 元となる配列が存在し、その配列を参照する。
    - その元の配列の容量を`cap()`関数で知ることができる。
    - 今回ならその配列は`[2]string`の型を持っている。

☆これにより、同じ型でありながら配列と同様の操作ができ、かつ長さを変えることができる。

### 新たな要素を追加するとき
`append()`関数を用いる。

例：
```go
     //main関数の続き
     userSlice = append(userSlice, "strelka")// => append(追加する前のSlice, 追加する要素)
     fmt.Printf("%v", userSlice)// => [tkmax777 matu, strelka]
     fmt.Printf("このスライスの容量は%dです。\n", cap(userSlice))// => このスライスの容量は3です。
```
- append関数によって背後で新たな長さのスライスが宣言されている。

### 配列やスライスからの切抜
　次のようにすでに存在する配列を参照先に指定することもできる。

```go
func main(){
     var userArray [4]string = [4]string{"tkmax777", "matu", "strelka", "zeke"}
     var userSlice []string = userArray[0:2]

     fmt.Printf("%v\n",userSlice)// => [tkmax777 matu strelka]
     fmt.Printf("容量：%d, 長さ：%d\n", cap(userSlice), len(userSlice))// => 容量：4, 長さ：3
}
```

- `Array[n:m]`(n, m ∈ **Q**≧0)と記述することで、配列のn番からm番までを取り出すことができる。
  - nは省略すると0が自動的に代入され、mを省略すると配列の末尾番号が自動的に代入される。
- 配列の一部を参照するようにしているため、この例ではスライスの容量(Capacity), 長さ(Length)が異なる。
  - 長さは`len()`関数で調べることができる。
- 同様な操作がSlice同士でも出来る。

### make関数
　また次のように記述することで、既に初期化された任意長のスライスを用意できる。
```go
//Name Put user name
type Name string

func main(){
     var userSlice []Name
     userSlice = make([]Name, 2, 4)//長さ2, 容量4のスライスを用意する
     userSlice[0] = "tkmax777"
     userSlice[1] = "matu"
     
     fmt.Printf("%v\n", userSlice)//[tkmax777 matu]
     fmt.Printf("容量：%d, 長さ：%d\n", cap(userSlice), len(userSlice))// => 容量：4, 長さ：2
}
```

- `make([]型名, 長さ, 容量)`の形で書くことで自動的に指定の容量を持った配列でかつ、指定の長さを持ったSliceが用意される。
  - 容量を省略すると長さと容量が等しいスライスが用意される。

この操作のことを**実体化**という。

### for range
for文で要素を一つずつ見ていきたいとき、次のような記法が可能

```go
var x []string = []string{"tkmax777", "matu", "defined", "murayama"}

for i := range x{
     fmt.Printf("%d:%s\n", i+1, x[i])
}

```

出力
```
1:tkmax777
2:matu
3:defined
4:murayama
```
さらに、要素を一つずつ取ってくることが出来る。

```go
var x []string = []string{"tkmax777", "matu", "defined", "murayama"}

for i, name := range x{
     fmt.Printf("%d:%s\n", i+1, name)
}
```

出力
```
1:tkmax777
2:matu
3:defined
4:murayama
```


### 注意
<ol>
<li> Sliceは配列の参照型（≒配列の番地リスト）なため、編集すると元の配列も編集される。</li> 
</ol>

```go

func main(){
     var userArray [4]string = [4]string{"tkmax777", "matu", "strelka", "zeke"}
     var userSlice []string = userArray[0:2]

     fmt.Printf("%v\n",userSlice)// => [tkmax777 matu strelka]
     fmt.Printf("容量：%d, 長さ：%d\n", cap(userSlice), len(userSlice))// => 容量：4, 長さ：3

     userSlice[2] = "dog"
     
     fmt.Printf("%v\n",userSlice)// => [tkmax777 matu dog]
     fmt.Printf("%v\n",userArray)// => [tkmax777 matu dog zeke]
}
```

<ol>
<li value="2"> 同様に関数に渡したさきの編集がプログラム全体で反映される。</li> 
</ol>

```go
func main(){
     var sweet []string = make([]string, 2)
	sweet = []string{"KINOKO", "Saikou"}
     modifier(sweet)
     fmt.Printf("%v", sweet)// => [TAKENOKO Saikou]
}

func modifier(sweet []string) {
     sweet[0] = "TAKENOKO"
     return
}
```

<ol>
<li value="3">  実体化していないスライス変数は`nil`であるため、当然参照できない。</li> 
</ol>

### 発展 入れ子
入れ子もできる。

```go
var x [][]string = [][]string{
     {"ミミ・バルゲリー", "竹中s", "袴田かげつ", "藤井まさみ"},
     {"湊智花", "三沢真帆", "袴田ひなた", "永塚紗季", "香椎愛莉"},
     {"荻山葵", "上原一成","柿園さつき","御庄寺多恵","長谷川昴","香椎万里"},
}

fmt.Printf("%s",x[1][0])//=>湊智花

```

## map
Pythonでいう辞書型のデータ。
ある型の入力(キー)に対して、出力を対応づけられる。

### 記法

型は、
```go
map[キーの型]出力の型
```
の形で定義される。

初期化時の書式は次の通り
```go
変数 = map[キーの型]出力の型{
     キー1:データ1,
     キー2:データ2,
     ...
}
```

要素を追加せずに用意したい場合はmake関数を用いる。
```go
var x map[キーの型]データ型 = make(map[キーの型]データ型)
```
どちらかをしないとメモリの実体が割りあてられないので注意。

#### stringとintを紐付けた例
```go
func main() {
	var age map[string]int = map[string]int{
          "湊智花": 12,
          "三沢真帆": 12,
          "袴田ひなた": 11,
          "永塚紗季": 12,
          "香椎愛莉": 12,
          "荻山葵": 16,
	}

	fmt.Printf("%s。%d歳\n", "湊智花", age["湊智花"])// => 湊智花。12歳    
	fmt.Printf("%s。%d歳\n", "袴田ひなた", age["袴田ひなた"])// => 袴田ひなた。11歳

	return
}

```
### 要素の追加

単純に
```go
変数[新しいキー] = 新しいデータ
```
だけでなされる。便利。

### 要素が存在するか判別

次の記法ができる。
```go
_, ok := MAP[Key]
if !ok {
     初期化していないときの処理
}
```

### 要素の削除
次の関数で削除できる。
```go
delete(MAP, KEY)
```

例

```go
---先程のmain関数---
delete(age, "荻山葵")
_, ok := age["荻山葵"]
if !ok {
     fmt.Printf("存在しません。\n") // =>存在しません。
}
```

### 参考 for range
・Sliceと同様、for文をrangeで回すことができる。、
```go
--先程のmain関数--
for name, num := range age {
     fmt.Printf("%s。%d歳\n", name, num)
}
```

出力
```
湊智花。12歳
三沢真帆。12歳
袴田ひなた。11歳
永塚紗季。12歳
香椎愛莉。12歳
```
注意:mapのアクセスにおける順番はランダム。

### 発展 入れ子
これもSliceと同様に入れ子ができる。

```go
func main() {
	var x map[string]map[string]int = map[string]map[string]int{
		"湊智花": {
			"年齢": 12,
			"学年": 6,
		},
		"三沢真帆": {
			"年齢": 12,
			"学年": 6,
		},
		"袴田ひなた": {
			"年齢": 11,
			"学年": 6,
		},
		"永塚紗季": {
			"年齢": 12,
			"学年": 6,
		},
		"香椎愛莉": {
			"年齢": 12,
			"学年": 6,
		},
	}

	fmt.Printf("%s。年齢は%d歳、%d年生\n", "湊智花", x["湊智花"]["年齢"], x["湊智花"]["学年"])// => 湊智花。年齢は12歳、6年生

	return
}
```
## Struct（構造体）
Struct（構造体）：変数の集合
### 記法

`struct{...}`という形の型として定義される。
```go
var 変数名 struct{
     要素名1   型名
     要素名2   型名
     ...
}
```
### 参照方法
```go
変数名.要素名
```
という形でかかれる。

例：
```go
func main() {
     var user struct{
          Name string
          Age  int
     }

     user.Name = "tkmax777"
     user.Age = 20

     fmt.Printf("Name:%s Age:%d", user.Name, user.Age)// => Name:tkmax777 Age:20
}
```

### 初期化
```go
var 変数名 struct{
          要素名1   型名
          要素名2   型名
          ...
     } = struct{
          要素名1   型名
          要素名2   型名
          ...
     }{
          要素名1:要素名1の内容
          要素名2:要素名2の内容
          ...
     }
```
例：
```go
func main() {
	var user struct {
		Name string
		Age  int
	} = struct {
		Name string
		Age  int
	}{
		Name: "tkmax777",
		Age:  20,
	}

	fmt.Printf("Name:%s Age:%d", user.Name, user.Age)// => Name:tkmax777 Age:20
}
```
見た瞬間は「は？」ってなりますが、これは配列やスライスの初期化方法とよく似ています。

・復習-Slice-
```go
var userSlice []string = []string{"tkmax777", "matu"}
```

このように、
```go
var 変数名 型名 = 型名{要素の内容}
```
という初期化方法をちゃんと満たしています。

また、この構造体の型のうちで、`Name`という要素が先で、`Age`という要素が後にきているということに注意すると、次のようにも初期化できます。

```go
var user struct {
          Name string
          Age  int
     } = struct {
          Name string
          Age  int
     }{
          "tkmax777",
          20,
     }
```
これは次のように一般化できます。

```go
var 変数名 struct{
          要素名1   型名
          要素名2   型名
          ...
     }=struct{
          要素名1   型名
          要素名2   型名
          ...
     }{
          要素名1の内容,
          要素名2の内容,
          ...
     }
```

このほか、今まで通り変数宣言の略記法もできます。
```go
変数名 := struct{
          要素名1   型名
          要素名2   型名
          ...
     }{
          要素名1の内容
          要素名2の内容
          ...
     }
```

例：
```go
func main() {
	user := struct {
		Name string
		Age  int
	}{
		"tkmax777",
		20,
	}

	fmt.Printf("Name:%s Age:%d\n", user.Name, user.Age)// => Name:tkmax777 Age:20
}
```

以上にStructの使い方の基本を記しました。しかし、一般的にはstructをそのまま型として変数を宣言するケースはあまりありません。

## TypeとStruct
では実際どのようにStruct使われるかというと、type節と組み合わせて使います。

例：
```go
package main

import "fmt"

//User Put user data
type User struct{
     Name      string
     Age       int
     Locate    string
     Group     string
}

func main() {
     var user User
     user.init()//initメソッドを実行
     user.Name = "tkmax777"
     user.Age = 20

     fmt.Printf("%v", user)
     check(user)
     fmt.Println()

     user.init()
     user.Name = "matu"
     user.Age = 19

     fmt.Printf("%v", user)
     check(user)
}

func(u User) String() string{
     return fmt.Sprintf("ユーザ名：%s 年齢：%d歳\n場所：%s 所属：%s\n", u.Name, u.Age, u.Locate, u.Group)
}

func(u *User) init() {
     (*u).Locate = "京都"
     (*u).Group = "KMC"
     (*u).Name = ""
     (*u).Age = 0
     return
}

func check(user User){
     if user.Age < 20{
          fmt.Printf("お酒とタバコは禁止です！\n")
     } else {
          fmt.Printf("飲み過ぎ注意！\n")
     }
     return
}
```
実行結果
```sh
ユーザ名：tkmax777 年齢：20歳
場所：京都 所属：KMC
飲み過ぎ注意！

ユーザ名：matu 年齢：19歳
場所：京都 所属：KMC
お酒とタバコは禁止です！
```

このように、グローバル領域においてtype節であらたな型としてstructを定義すると、メソッドが使えるようになったり、関数の引数に使えたりと嬉しいことがおおいのです。

ちなみに、構造体をポインタで表した時、わざわざ`(*変数名).要素`と書かず、`変数名.要素`と書いて参照することが認められています。そのため、以上のinitメソッドは次のように書くことができます。
```go
func(u *User) init() {
     u.Locate = "京都"
     u.Group = "KMC"
     u.Name = ""
     u.Age = 0
     return
}
```

## interface(インターフェース)
　今までにType節と、それに付随するメソッドについて学びました。ここで注意したいのは、Type節で定義する型の名前はpackage内で固有なものである必要がありますが、その型に対して定義されたメソッドの名前については一意でなくて良い、ということです。例えば、次のようなプログラムを書くことができます。

```go
package main

import "fmt"

type Human struct {
	Name   string
	Age    int
	Locate string
	Job    string
}

type Cat struct {
	Name   string
	Age    int
	Color  string
	Locate string
}

func main() {
	var human Human
	var cat Cat

	human.Input()
	fmt.Println()
	cat.Input()

	fmt.Printf("%#v\n", human)
	fmt.Printf("%#v\n", cat)

	return
}

func (h *Human) Input() {
	fmt.Printf("人間の情報を入力します。\n")
Name:
	fmt.Printf("名前を入力してください。\n")
	fmt.Scanf("%s", &h.Name)
	if h.Name == "" {
		fmt.Printf("無効な入力です。\n")
		goto Name
	}
Age:
	fmt.Printf("年齢を入力してください。\n")
	fmt.Scanf("%d", &h.Age)
	if h.Age == 0 {
		fmt.Printf("年齢を半角算用数字で入力してください。\n")
		goto Age
	}
Locate:
	fmt.Printf("所在地を入力してください。\n")
	fmt.Scanf("%s", &h.Locate)
	if h.Locate == "" {
		fmt.Printf("無効な入力です。\n")
		goto Locate
	}
Job:
	fmt.Printf("職業を入力してください。")
	fmt.Scanf("%s", &h.Job)
	if h.Job == "" {
		fmt.Printf("無効な入力です。\n")
		goto Job
	}
}

func (c *Cat) Input() {
	fmt.Printf("猫の情報を入力します。\n")
Name:
	fmt.Printf("名前を入力してください。\n")
	fmt.Scanf("%s", &c.Name)
	if c.Name == "" {
		fmt.Printf("無効な入力です。\n")
		goto Name
	}
Age:
	fmt.Printf("年齢を入力してください。\n")
	fmt.Scanf("%d", &c.Age)
	if c.Age == 0 {
		fmt.Printf("年齢を半角算用数字で入力してください。\n")
		goto Age
	}
Locate:
	fmt.Printf("所在地を入力してください。\n")
	fmt.Scanf("%s", &c.Locate)
	if c.Locate == "" {
		fmt.Printf("無効な入力です。\n")
		goto Locate
	}
Color:
	fmt.Printf("色を入力してください。\n")
	fmt.Scanf("%s", &c.Color)
	if c.Color == "" {
		fmt.Printf("無効な入力です。\n")
		goto Color
	}
}

```

このプログラムでは、`Human`と`Cat`各々に対して`Input`メソッドが定義されています。いま、対象となる型が違えど、これらの同じ名前のメソッドは同じ働きをしていますね。Golangではこのような同じメソッドを持つ型を関数の引数などに用いたい時に使える手法として`interface`というものが用意されています。

・記法
```go
interface{
     メソッド1
     メソッド2
     ...
}
```
例：
```go
func input(i interface{ Input() }) {
	i.Input()
}
```
この関数では`Input`というメソッドを実装している集合（最初の例では`Human`と`Cat`がこれに相当します。）を引数としてとることができます。

さらに、このinterfaceに対しても、今までのようにtype節で名前をつけることができます。

例：
```go
type Inputter interface{
     Input()
}
func input(i Inputter) {
	i.Input()
}
```

さて、今までにこのような、あるメソッドを実装している集合に対して特殊な操作をするpackageが一度だけ登場していることにお気づきですか？

そうです。`fmt`packageです。

このpackageでは、`String()`メソッドを実装している集合に対しては自動的にそのメソッドを実行していました。この集合を示すinterfaceは今までと同様に次のように表すことができます。

```go
type Stringer interface{
     String()
}
```

interfaceは使いこなせればとても柔軟にプログラムを書くことができます。是非うまく使えるように、意識してプログラムを書いてみてください。

## 最後に...
ここまでこれば、あなたは一通りGolangを書くための基礎を抑えることができたはずです。

しかし今回は時間の都合もあり省略してしまいましたが、他にもgo routineという並列処理をするための優秀な機能や、それに付随する便利な機能がまだまだあります。この言語のさらなる理解を深め、使いこなせるようになるためにも、是非[A Tour of Go](https://go-tour-jp.appspot.com/welcome/1)をやってみてください。

さて、次回からは実際に動くプログラムを作っていきます。お楽しみに。