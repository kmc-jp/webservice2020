# その他の型

## 配列

JavaScriptでは次のように配列を宣言することができます。

```js
let Array = ["First", "Second"];
console.log(Array[0]); // => First
```

### 配列の長さ

```js
let Array = ["First", "Second"];
console.log(Array.length); // => 2
```

### 配列の要素の付加
次の操作で、末尾にデータを追加することができます。

```js
let Array = ["First", "Second"];
console.log(Array.push("Third")); // => 3
```
このとき返されるのは、要素の追加された配列の長さです。

また、次の操作で先頭に要素を追加することもできます。

```js
let Array = ["First", "Second"];
console.length(Array.unshift("Zero")); // => 3
```

### 要素の削除
次の操作で、末尾のデータが削除されます。

```js
let Array = ["First", "Second"];
console.log(Array.pop()) // Second
```

また、次の操作で先頭の要素を削除することができます。

```js
let Array = ["First", "Second"];
console.log(Array.shift()); // First
```

### 要素の検索
JavaScriptでは、次のメソッドで簡単に要素をの添え字を調べることができます。

```js
let Array = ["First", "Second"];
console.log(Array.indexOf("First")); // 0
```

このとき、見つからない場合は`-1`を返します。

### 配列の複製
配列を実態ごと複製したい場合は、次のようにすることでできます。

```js
let Array = ["First", "Second"];
let ArrayCopy = Array.slice();
```

### 注意事項
JavaScriptの配列はGoのSliceとは性質が異なる点があります。それについて、対比しつつ、軽く確認しておきましょう。

- 配列の要素が全て同じ型とは限らない

GoのSliceでは、そもそもそのSlice自体の型名に要素の型が入っていました。そのため、他の要素の型が入ることはまずありえませんでした。しかし、JSではその限りではありません。0番目はNumber、1番目は文字列、2番目は真理値…などということがありうるのです。気をつけましょう。

- 配列が密とは限らない

GoのSliceで穴が発生することは、あり得ません。すなわち、初期化されていない要素は存在せず、メモリを確保した時点で、その分の要素は初期値(0や"")で初期化されるため、書き手は安心して利用することができます。しかし、JSでは先程の条件から、`undefined`要素が入り得るのです。十分に気をつけましょう。

- 配列外参照が可能

Go言語を振り返ってみましょう。

```go
var Array []string = []string{"First", "Second"}
fmt.Printf(Array[2])
```

このプログラムを実行するとどうなるでしょうか。答えは簡単──Panicが発生し、その時点でプログラムが停止します。では、同様のプログラムをJavaScriptで動かしてみましょう。

```js
let Array = ["First", "Second"];
console.log(Array[2]);
```

このプログラムはなんと正しく実行されます。その結果は次のとおりです。

```
undefined
```

undefined「呼んだ？」

というわけで、配列外参照によって停止しない可能性があるので、十分に気をつけておきましょう。

### 発展 - 多次元配列 -
次のように記述することで、多次元配列を形成することができます。

```js
let Array = [
    // 一次元
    [
        // 二次元
        [
            // 三次元
            ...
        ],
        [
            // 三次元
            ...
        ]
    ],
    [
        ...
    ],
    ...
]

```

## オブジェクト

おまたせしました。JavaScriptの基礎、といって切っても切り離せないものにオブジェクトは絶対入ります。まずは次の例をみてください。

```js
let Object = {
    
}
```

## 参考
[Array - JavaScript](https://developer.mozilla.org/ja/docs/Web/JavaScript/Reference/Global_Objects/Array)<br>
[JavaScript 型付き配列 - JavaScript](https://developer.mozilla.org/ja/docs/Web/JavaScript/Typed_arrays)<br>
[JavaScriptでメソッドを使う方法【初心者向け】](https://techacademy.jp/magazine/5562#:~:text=%E3%83%A1%E3%82%BD%E3%83%83%E3%83%89%E3%81%A8%E3%81%AF-,%E3%83%A1%E3%82%BD%E3%83%83%E3%83%89%E3%81%A8%E3%81%AF%E3%80%81%E3%82%AA%E3%83%96%E3%82%B8%E3%82%A7%E3%82%AF%E3%83%88%E3%81%AE%E3%83%97%E3%83%AD%E3%83%91%E3%83%86%E3%82%A3%E3%81%AB%E4%BB%A3%E5%85%A5%E3%81%95%E3%82%8C%E3%81%9F,%E3%82%92%E3%81%A4%E3%81%91%E3%81%9F%E3%82%82%E3%81%AE%E3%81%A7%E3%81%99%E3%80%82)<br>
[JavaScript Object Notation - Wikipedia](https://ja.wikipedia.org/wiki/JavaScript_Object_Notation)<br>
