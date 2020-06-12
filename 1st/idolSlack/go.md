# go routine
## 概要
Go言語に於いて便利な機能な一つ、go runtine(ゴルーチン)について使いかたを簡単にだけ学ぶ。

## 使い方

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    go func (){
        for i := 0; i < 5; i++{
            fmt.Printf("%d\n", i)
            time.Sleep(time.Second)
        }
    }()

    fmt.Printf("Excuted!\n")

}

```

実行結果
```sh
0
Excuted!
1
2
3
4
```

上記のように、

```go
go 実行する関数
```
と記述することで、この関数を並列に処理することができる。とても簡単ですね。

## まとめ
Golangではとても簡単に並列処理が実装出来るので、気が向けば使ってみましょう。

今回は省略しますが、この並列処理を同期するために`sync`パッケージやChannelが使えるので興味があれば調べてにれば良いです。(要望があればやるかもしれません。)
