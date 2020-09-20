# その他の型

## 目次
<!-- TOC -->

- [その他の型](#その他の型)
    - [目次](#目次)
    - [配列](#配列)
        - [配列の長さ](#配列の長さ)
        - [配列の要素の付加](#配列の要素の付加)
        - [要素の削除](#要素の削除)
        - [要素の検索](#要素の検索)
        - [配列の複製](#配列の複製)
        - [注意事項](#注意事項)
        - [参考 - for ... of [Array] -](#参考---for--of-array--)
        - [発展 - 多次元配列 -](#発展---多次元配列--)
    - [オブジェクト](#オブジェクト)
        - [プロパティ](#プロパティ)
            - [リテラル](#リテラル)
        - [参考 - 関連事項 -](#参考---関連事項--)
            - [個別登録](#個別登録)
            - [参考 連想配列的な参照](#参考-連想配列的な参照)
            - [参考 - for ... in 〜 -](#参考---for--in-〜--)
        - [メソッド](#メソッド)
            - [リテラル](#リテラル-1)
            - [個別登録](#個別登録-1)
            - [自身を参照する](#自身を参照する)
        - [組み込みオブジェクト](#組み込みオブジェクト)
    - [終わりに](#終わりに)
    - [参考](#参考)

<!-- /TOC -->

## 配列

JavaScriptでは次のように配列を宣言することができます。

```js
let ExArray = ["First", "Second"];
console.log(ExArray[0]); // => First
```

一応リテラルは次の通りです。

```js
let Array_Name = [
    Element_1,
    Element_2,
    ...
    Element_n
]
```

### 配列の長さ

```js
let ExArray = ["First", "Second"];
console.log(ExArray.length); // => 2
```

### 配列の要素の付加
次の操作で、末尾にデータを追加することができます。

```js
let ExArray = ["First", "Second"];
console.log(ExArray.push("Third")); // => 3
```
このとき返されるのは、要素の追加された配列の長さです。

また、次の操作で先頭に要素を追加することもできます。

```js
let ExArray = ["First", "Second"];
console.length(ExArray.unshift("Zero")); // => 3
```

### 要素の削除
次の操作で、末尾のデータが削除されます。

```js
let ExArray = ["First", "Second"];
console.log(ExArray.pop()) // Second
```

また、次の操作で先頭の要素を削除することができます。

```js
let ExArray = ["First", "Second"];
console.log(ExArray.shift()); // First
```

### 要素の検索
JavaScriptでは、次のメソッドで簡単に要素をの添え字を調べることができます。

```js
let ExArray = ["First", "Second"];
console.log(ExArray.indexOf("First")); // 0
```

このとき、見つからない場合は`-1`を返します。

### 配列の複製
配列を実態ごと複製したい場合は、次のようにすることでできます。

```js
let ExArray = ["First", "Second"];
let ExArrayCopy = ExArray.slice();
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
var ExArray []string = []string{"First", "Second"}
fmt.Printf(ExArray[2])
```

このプログラムを実行するとどうなるでしょうか。答えは簡単──Panicが発生し、その時点でプログラムが停止します。では、同様のプログラムをJavaScriptで動かしてみましょう。

```js
let ExArray = ["First", "Second"];
console.log(ExArray[2]);
```

このプログラムはなんと正しく実行されます。その結果は次のとおりです。

```
undefined
```

undefined「呼んだ？」

というわけで、配列外参照によって停止しない可能性があるので、十分に気をつけておきましょう。

### 参考 - for ... of [Array] -
for文に於いて、配列の要素一つ一つに関して処理を実行したいとき、次のように記述することで、これを達成することができます。

- 例

```js
let ExArray = ["First", "Second", "Third"];
for (let a of ExArray) {
    console.log(a);
}
```

- 出力

```
First
Second
Third
```

### 発展 - 多次元配列 -
次のように記述することで、多次元配列を形成することができます。

```js
let ExArray = [
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
おまたせしました。JavaScriptの基礎、といって必須なものを挙げたとしたら、確実に入るであろう、オブジェクトという概念を学びましょう。まずは次の例をみてください。

```js
let Member_1 = {
    ID: "tkmax777",
    Club: "KMC",
    Projects: [
        "WebService2020",
        "DTM練習会2020",
        "お絵かき練習会2020"
    ]
};

console.log(Member_1.ID); // => tkmax777
```

この例では、Member_1というオブジェクトを宣言し、その中にIDや、Clubといった
**プロパティ**
を登録しています。

### プロパティ
#### リテラル
プロパティとは、先のような、オブジェクトに登録された、データのことです。リテラルを次に示します。

```js
let Object_Name = {
    property_1 : Data_1,
    property_2 : Data_2,
    ...
    property_n : Data_n
}
```

### 参考 - 関連事項 -
さて、皆さんはこの記法、今までに見たことありますよね。え？わからない？

では、今までに
**JS**
っていう名前を含んだものを扱ったこと、ありましたよね？

そうです！
**json**
ですね。

実はjsonは、
**JavaScript Object Notation**
の略であり、JavaScriptの表現方法を摸して作られたものなのです。しかし、全く同じなわけではありません。違いを以下に載せておきます。

> - JSON は、"property": value 構文を使用するプロパティ定義のみ許可します。プロパティ名称は二重引用符で囲まなければなりません。そして、その定義は簡略にすることはできません。
> - JSON ではその値は strings、numbers、arrays、true、false、null、別の (JSON) オブジェクトのみです。
> - 関数の値 (上記"メソッド"を参照) は JSON では値を割り当てることができません。
> - Date のようなオブジェクトは JSON.parse() の後で string になります。
> - JSON.parse() は計算されたプロパティ名を拒否し、エラーがスローされます。<br>

[引用元](https://developer.mozilla.org/ja/docs/Web/JavaScript/Reference/Operators/Object_initializer)

注意しましょう。

#### 個別登録
次のように記述することで、個別に登録も可能です。

```js
let Member_1 = {};
Member_1.Name = "tkmax777";

console.log(Member_1.Name) // => tkmax777
```

このとき、既に存在しているプロパティに対して登録しようとした場合、以前のデータは上書きれて、消えます。

#### 参考 連想配列的な参照
オブジェクトの各々の名称は、識別子の名称規則に則っている必要があります。しかし、時としてそれに反した名前を付けたいことがあるかもしれません。そのときは次のように連想配列的な参照方法がつかえることを覚えておくと良いかもしれません。

```js
let Member_1 = {};
Member_1.Name = "tkmax777";

console.log(Member_1["Name"]) // => tkmax777
```

#### 参考 - for ... in 〜 - 
配列の時みたく、オブジェクトの全プロパティに対してある処理をしたい場合、次のように記述できます。

- 例

```js
let ExObject = {
    Proparty_1: "First",
    Proparty_2: "Second",
    Proparty_3: "Third"
}

for (proparty in ExObject) {
    console.log("Proparty Name: " + proparty + " Proparty: " + ExObject[proparty]);
}

```

- 出力

```
Proparty Name: Proparty_1 Proparty: First
Proparty Name: Proparty_2 Proparty: Second
Proparty Name: Proparty_3 Proparty: Third
```


### メソッド
#### リテラル
オブジェクトに登録できるのは、データだけではありません。
**メソッド**
と呼ばれる関数も登録することができます。

```js
let MethodObject = {
    Print : function(data) {
        console.log(data);
        return undefined;
    }
}

MethodObject.Print("YEAH"); // => YEAH
```

一応リテラルも示しておきます。

```js
let Object = {
    Method_1: function(parameter) {
        処理
        return 
    },
    Method_2: function(parameter) {
        ...
    },
    ...
}
```

#### 個別登録
メソッドに関しても、同様に後付けできます。

```js
let MethodObject;
MethodObject.Print = function(data) {
    console.log(data);
    return undefined;
};
```

#### 自身を参照する
メソッドの定義に於いて、そのオブジェクトの持つプロパティなどを参照したい場合があるかもしれません。その場合は、自身を`this`として表すことで、これを達成することができます。

```js
let Member = {
    Print : function() {
        console.log(this.Name);
        return undefined;
    },
    Name: "tkmax777"
}

Member.Print(); // => tkmax777
```

### 組み込みオブジェクト
今までに既に多くのオブジェクトが登場してきたことに、皆さんは気付いているかもしれません。

例えば`console`。ブラウザのコンソール周りのメソッドやプロパティを持った組み込みオブジェクトです。そのなかの`log`メソッドを利用して、ここまで出力を書いてきました。

しかし、それだけでもないです。今までに型として登場した、`Array`、`Number`、`String`など、さらには関数`Function`まで、JavaScriptの世界の全てがオブジェクトから成り立っています。

その全貌を説明することはここではできませんが、興味があるのなら、各々調べてみると良いでしょう。

[標準組み込みオブジェクト](https://developer.mozilla.org/ja/docs/Web/JavaScript/Reference/Global_Objects)<br>
[Web API](https://developer.mozilla.org/ja/docs/Web/API)<br>

## 終わりに
さて、ここまでJavaScriptの基本文法をずっとお話してきました。しかし、残念ながらここまでではまだDOMの操作は一切できませんね。これを使うためにはWeb APIのDOMインターフェースについて、もうすこし知る必要があります。次回はこれについて、学んでいこうと思います。お疲れさまでした。

## 参考
[Array - JavaScript](https://developer.mozilla.org/ja/docs/Web/JavaScript/Reference/Global_Objects/Array)<br>
[JavaScript 型付き配列 - JavaScript](https://developer.mozilla.org/ja/docs/Web/JavaScript/Typed_arrays)<br>
[オブジェクト初期化子](https://developer.mozilla.org/ja/docs/Web/JavaScript/Reference/Operators/Object_initializer)<br>
[オブジェクトでの作業 - JavaScript](https://developer.mozilla.org/ja/docs/Web/JavaScript/Guide/Working_with_Objects)<br>
[JavaScriptでメソッドを使う方法【初心者向け】](https://techacademy.jp/magazine/5562#:~:text=%E3%83%A1%E3%82%BD%E3%83%83%E3%83%89%E3%81%A8%E3%81%AF-,%E3%83%A1%E3%82%BD%E3%83%83%E3%83%89%E3%81%A8%E3%81%AF%E3%80%81%E3%82%AA%E3%83%96%E3%82%B8%E3%82%A7%E3%82%AF%E3%83%88%E3%81%AE%E3%83%97%E3%83%AD%E3%83%91%E3%83%86%E3%82%A3%E3%81%AB%E4%BB%A3%E5%85%A5%E3%81%95%E3%82%8C%E3%81%9F,%E3%82%92%E3%81%A4%E3%81%91%E3%81%9F%E3%82%82%E3%81%AE%E3%81%A7%E3%81%99%E3%80%82)<br>
[JavaScript Object Notation - Wikipedia](https://ja.wikipedia.org/wiki/JavaScript_Object_Notation)<br>
[【JavaScript入門】「for – of文」の使い方と間違いやすいポイントを徹底解説！](https://www.sejuku.net/blog/20228)<br>
