# 第一回

## 概要
後期初回。今回は前期に時間がなくて触れずに終わったJavaScriptについて、軽く使いかたを学んでいきます。

## 目次
<!-- TOC -->

- [第一回](#第一回)
    - [概要](#概要)
    - [目次](#目次)
    - [基本知識](#基本知識)
        - [DOM](#dom)
        - [JavaScript](#javascript)
    - [JavaScriptの世界でHello World!](#javascriptの世界でhello-world)
        - [HTMLへの埋め込みの仕方](#htmlへの埋め込みの仕方)
        - [コメントの記法](#コメントの記法)
        - [function節](#function節)
        - [alert関数](#alert関数)
        - [console.log](#consolelog)
    - [宣言](#宣言)
    - [型](#型)
    - [型変換](#型変換)
    - [巻き上げ](#巻き上げ)
    - [参考](#参考)

<!-- /TOC -->

## 基本知識
世のなかのウェブページを眺めてみると、今まで学んできたような静的な描画のページだけではないことがわかります。例えば、Twitterでは、タイムラインが自動的に更新されていったり、ユーザアイコンにポインタを重ねるだけで、その人の詳細が出てきたりしますね。このように、ブラウザでは再読み込みをせずに動的にHTMLの要素を動かすことができるようになっています。この機能をささえている仕組み、それがDOMです。

### DOM
DOMとは
**Document Object Model**の略であり、HTMLなどの文章をプログラムから操作するためのインターフェースです。

### JavaScript
JavaScript(JS)は先のDOMを操作するために最もよく用いられるプログラミング言語です。その他、Node.jsやGASなどでサーバ側のプログラムを記述するのにも用いられます。

## JavaScriptの世界でHello World!
実際にプログラムを動かしてみましょう。次のようにHTMLを記述することによって、簡単にJSを実行することができます。

- HTML

```html
<!DOCTYPE html>
<html lang="ja">
    <head>
        <meta charset="UTF-8">
        <script>
        // javascript開始
            function Hello(){
                alert("Hello World!")
            }
        </script>
    </head>
    <body>
        <a href="javascript:Hello();">
            <button type="button">実行</button>
        </a>
    </body>
</html>

```

- 実行結果

<iframe src="html/hello_world.html" name="sample" width="90%" height="50">
    <a href="html/hello_world.html"></a>
</iframe>

それでは、各要素について細かくみていきましょう。

### HTMLへの埋め込みの仕方

`<script>`
タグの要素として埋め込むことで、
**上から順に**
読み込まれます。

### コメントの記法

```js
// コメント
/*
複数行に
またがる
コメント
*/
```

上記のように、
`//`
もしくは
`/**/`を用いることで、コメントを挿入できます。Golangと同じですね。


### function節

```js
function Hello(){
    ...
}
```

関数
`Hello`
を宣言しています。

### alert関数

```js
alert("Hello World!")
```

alert関数は、引数にわたされた要素を先程確認したように、ポップアップで表示します。引数には任意の型のデータを渡すことができます。

### console.log

先程の例にはありませんでしたが、次のように記述することで、ブラウザのコンソール(開発者ビューにあるやつ)に出力することができます。

```js

console.log(arg1, arg2, ...)

```

arg: 任意の型

基本的にデバッグなどはこれをもちいておこないます。


## 宣言

JSでは次のように各種宣言することができます。

```js
// 変数
var x;
let y;

// 定数
const z;
```

以上の記述をするだけで、変数の宣言が行われます。型については後述しますが、記述する必要はありません。さて、var と letの違いが気になりますね。これらはスコープが異なります。

次の例を見てくださjい

```js
function Awesome() {
    if (true) {
        var x = 5;
        let y = 5;

        console.log("x = " + x + " y = " + y); // => x = 5 y = 5
    }
    console.log("x = " + x + " y = " + y); // => x = 5 y = undefined
}

```

このように、varでは関数全体がスコープになるのに対し、letではブロックがスコープになります。varを下手に用いると、予期しない動作となる場合があるので、できるだけletを用いると良いでしょう。但し、グローバル変数の宣言にはvarしか用いられません。気をつけましょう。

## 型

それでは、JavaScriptの型を見ていきましょう。

基本的には次の型が用意されています。

| 型 | 説明 |
| --- | --- |
|`Boolean`|真偽値|
|`BigInt`|長整数。精度が自由な整数値|
|`Number`|整数もしくは不動性点数|
|`String`|文字列|
|`Object`|構造体|
|`null`|空集合|
|`NaN`|非数|
|`undefined`|未定義値|

他にもSymbol (シンボル)という型もありますが、とりあえずのところは良いでしょう。JSでは変数の型は、自動的に定義され、かつ動的に変化します、よって次のような操作は全く問題ありません。


```js
let x = "Hello";
console.log(x); // => Hello

x = 5;
console.log(x); // => 5

x = true;
console.log(x); // => true
```

<iframe src="html/vars.html" name="sample" width="90%" height="50">
    <a href="html/vars.html"></a>
</iframe>

## 型変換

型について学んだとき、次に気になるのはやはり型変換の方法でしょう。JavaScriptで問題になる型変換の方法は、数値 <=> 文字列の変換ですね。これはいたって簡単で、次のように行うことができます。

```js
// 数値 => 文字列
let num = 5;

console.log("数値は"+num+"です");

// 文字列 => 数値
let x = "5.1"

console.log("parseInt: "+parseInt(x));  // => 5
console.log("parseFloat: "+parseInt(x));// => 5.1

```

数値=>文字列に関しては、必要になったときにStringとそのまま+演算子にて結合すれば勝手に変換されます。また、文字列=>数値に関しては、parseIntもしくはparseFloatを用いることで、それぞれ整数、浮動小数型に変換できます。

## 巻き上げ

JSでは、変数の宣言順序の規則がゆるく設定されています。まず、次のコードを見てください。

```js
function awesome_func1() {
    awesome_func2(); // => func2 excuted!
}

function awesome_func2() {
    console.log("func2 excuted!");
}

```

みなさんはこの記法をみて、何か違和感を覚えますか。もしあなたがC言語などに触れていれば、awesome_func2はawesome_func1よりあとに宣言されているので、awesome_func1からは参照できない可能性がある、と思うかもしれなせん。しかし、さしあたってGo言語のように、JSでは関数は宣言する順番に関係なく、使用することができます。<a href="#p1">^1</a>このことを、
**関数の巻き上げ**
と呼びます。

しかし、このJavaScriptに於いて、ルールは関数の限りではありません。次の例を見てください。


```js
function awesome_func(){
    let x = "Hello";
    console.log(x);

    if (true){
        console.log(x);
        let x = "NULL";
    }

    console.log(x);
}

```

Go言語などに慣れた人であれば、きっと次のような出力を期待することでしょう。

```
Hello
Hello
NULL
```

しかし、JavaScriptではこうなりません。正しくは次のようになります。

<iframe src="html/vars2.html" name="sample" width="90%" height="50">
    <a href="html/vars2.html"></a>
</iframe>

![実行例](static/var.png)

これは先程の関数の巻き上げと同様の操作が変数に於いても行われていることによります。if文内で記述されているlet節により、関数の宣言だけ先行しているのです。注意しなければいけないのは、次のような記述をしても初期化はなされません。

```js
console.log(x); // ~> undefined
let x = 5;
```

この、宣言だけされて、初期化されていない空白のことを、
**TDZ(Temporal Dead Zone)**
と呼びます。

このような記法はバグの原因になりますし、そもそも見にくいです。できるだけ宣言は関数の先頭で行うように心がけましょう。

<a id="p1">^1</a>: ただし、htmlは上から順に読み込まれていくため、複数のscriptタグをまたいで、読み込み前に利用することはできません。

それではフロー制御を見ていきましょう。

[次にすすむ](flow.md)

## 参考
[文法とデータ型 - JavaScript | MDN](https://developer.mozilla.org/ja/docs/Web/JavaScript/Guide/Grammar_and_types)<br>
[console.log() - Web API | MDN](https://developer.mozilla.org/ja/docs/Web/API/Console/log)<br>
[Symbol - JavaScript | MDN](https://developer.mozilla.org/ja/docs/Web/JavaScript/Reference/Global_Objects/Symbol)<br>
[JavaScript: var/letのスコープとletによる巻き上げ](https://qiita.com/jonakp/items/c943c8d457871fa6ceb6)<br>
[【javascript】関数の宣言はどこに書いてもいいの？（呼び出しと宣言の位置） at softelメモ](https://www.softel.co.jp/blogs/tech/archives/1703)<br>