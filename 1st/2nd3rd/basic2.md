<!-- TOC -->

- [Pointers](#pointers)
    - [ポインタとは？](#ポインタとは)
    - [使い方](#使い方)
    - [実用例](#実用例)
- [Type節](#type節)
- [Array（配列）](#array配列)
- [Slice（スライス）](#sliceスライス)

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
p値： 0xc0000a82a0//ここの値は毎回変わる。
i値： 1
*p値： 2
i値： 2
```

1. 番地を格納する変数は`var 変数名 *型名`の形で宣言する。
2. 番地に格納されている値を読み出すときは`*変数名`と記述する。
3. 一般の変数の番地は`&変数名`という形で読み出すことができる。

☆上記の例では変数iの番地をpに格納し、pを通してiを編集しています。
  - 0xhogehogeはメモリ番地を１６進数で表した物です。

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
     fmt.Printf("%s：%t\n", user, verify(user))//
     
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
     
     fmt.Printf("%s：%t\n", user, varify(club))//コンパイルエラー
```
- `verify()`の引数はName型でないといけないため、エラーが発生する。

このように型を作っておくことで、より信頼性の高いプログラムを作成することができる。

## Array（配列）
```go
func main(){
     var user [2]string
     user[0] = "tkmax777"
     user[1] = "matu"

     fmt.Printf("user1:%s\nuser2:%s")//user1:tkmax777 user2:matu
     return
}
```
- `[配列の長さ]型名`という型を持つ。`型名`のところには任意の型が入る。
- 参照するときは`変数名[参照番号]`の形で参照できる。
  - 参照番号は0, 1, 2...と数えていくことに注意。

また、次のように初期化することもできる。
```go
     user = [2]string{"tkmax777", "matu"}
     fmt.Printf("user1:%s\nuser2:%s\n")//user1:tkmax777 user2:matu
```

**Point**

Goにおいて配列は、その長さを型に含むため、後から長さを変更することができない。（固定長）
- 無理矢理長さ変えるなら、その度に新しい配列を宣言することになる。

☆配列は次に挙げるスライス（Slice）があるので基本的にそんなに使わない。

## Slice（スライス）
```go
func main(){
     var userSlice []string = []string{"tkmax777", "matu"}

     fmt.Printf("%v", userSlice)//[tkmax777 matu]
     fmt.Printf("このスライスの容量は%dです。\n", cap(userSlice))//このスライスの長さは2です。
}
```
- `[]型名`という型を持つ
- 配列の参照型
  - 元となる配列が存在し、その配列を参照する。
    - その元の配列の容量を`cap()`関数で知ることができる。
    - 今回ならその配列は`[2]int`の型を持っている。

☆これにより、同じ型でありながら配列と同様の操作ができ、かつ長さを変えることができる。

- 新たな要素を追加するときは`append()`関数を用いる。

例：
```go
     //main関数の続き
     userSlice = append(userSlice, "strelka")//append(追加する前のSlice, 追加する要素)
     fmt.Printf("%v", userSlice)//[tkmax777 matu, strelka]
     fmt.Printf("このスライスの容量は%dです。\n", cap(userSlice))//このスライスの容量は3です。
```
- append関数によって背後で新たな長さの配列が宣言されている。

　次のようにすでに存在する配列を参照先に指定することもできる。

```go
func main(){
     var userArray [4]string = [4]string{"tkmax777", "matu", "strelka", "zeke"}
     var userSlice []string = userArray[0:3]

     fmt.Printf("%v\n",userSlice)//[tkmax777 matu strelka]
     fmt.Printf("容量：%d, 長さ：%d\n", cap(userSlice), len(userSlice))//容量：4, 長さ：3
}
```

- `Array[n:m]`(n, m ∈ **Q**≧0)と記述することで、配列のn番からm番までを取り出すことができる。
  - nは省略すると0が自動的に代入され、mを省略すると配列の末尾番号が自動的に代入される。
- 配列の一部を参照するようにしているため、この例ではスライスの容量(Capacity), 長さ(Length)が異なる。
  - 長さは`len()`関数で調べることができる。

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
     fmt.Printf("容量：%d, 長さ：%d\n", cap(userSlice), len(userSlice))//容量：4, 長さ：3
}
```

- `make([]型名, 長さ, 容量)`の形で書くことで自動的に指定の容量を持った配列でかつ、指定の長さを持ったSliceが用意される。
  - 容量を省略すると長さと容量が等しいスライスが用意される。

この操作のことを**実体化**という。

**注意**
1. Sliceは配列の参照型（≒配列の番地リスト）なため、編集すると元の配列も編集される。

```go
func main(){
     var userArray [4]string = [4]string{"tkmax777", "matu", "strelka", "zeke"}
     var userSlice []string = userArray[0:3]

     fmt.Printf("%v\n",userSlice)//[tkmax777 matu strelka]
     fmt.Printf("容量：%d, 長さ：%d\n", cap(userSlice), len(userSlice))//容量：4, 長さ：3
	userSlice[2] = "dog"
     fmt.Printf("%v\n",userSlice)//[tkmax777 matu dog]
     fmt.Printf("%v\n",userArray)//[tkmax777 matu dog zeke]
}
```
2. 同様に関数に渡したさきの編集がプログラム全体で反映される。
```go
func main(){
     var sweet []string = make([]string, 2)
	sweet = []string{"KINOKO", "Saikou"}
     modifer(sweet)
     fmt.Printf("%v", sweet)//[TAKENOKO Saikou]
}

func modifer(sweet []string) {
     sweet[0] = "TAKENOKO"
     return
}
```