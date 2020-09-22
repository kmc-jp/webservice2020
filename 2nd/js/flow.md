# フロー制御

[前にもどる](readme.md)
[次にすすむ](object.md)

## 概要
ここまではJavaScriptの基本文法を軽くみてきました。ここからは条件分岐などをさらに細かく見ていきましょう。

## 目次

<!-- TOC -->

- [フロー制御](#フロー制御)
    - [概要](#概要)
    - [目次](#目次)
    - [条件分岐](#条件分岐)
        - [基本](#基本)
        - [condition](#condition)
            - [falsy](#falsy)
            - [truthy](#truthy)
        - [if...else](#ifelse)
        - [else if](#else-if)
        - [switch文](#switch文)
    - [Loop](#loop)
        - [For文](#for文)
        - [while文](#while文)
        - [do...while文](#dowhile文)
        - [補足 - continue -](#補足---continue--)
    - [参考](#参考)

<!-- /TOC -->

## 条件分岐
まずはif文の記法を見ましょう。

### 基本

```js
if (condition) {
    実行内容
} 
```

基本的にはGo言語と大差ありません。conditionには同じく論理式そ入れることができます。詳しくは
[ここ](https://developer.mozilla.org/ja/docs/Web/JavaScript/Reference/Operators/Logical_Operators)
を見ると良いでしょう。

ただし、JavaScriptのconditionに入るのは`true`、`false`だけではありません。偽と見なされるのは、
**falsy**
なもの、対して真として扱われるのは
**truthy**
なものです。

### condition

#### falsy

falseの他、次の8つは全てfalsyなものとして認識されます。

| false | false キーワード |
| --- | --- |
|0|数値ゼロ|
|-0|数値マイナスゼロ|
|0n|BigInt で、論理値として使用された場合、 Number と同じ規則に従います。 0n は偽値です。|
|""|空文字列の値|
|null|null - 何も値が存在しないこと|
|undefined|undefined - プリミティブ値|
|NaN|NaN - 非数|

[引用元](https://developer.mozilla.org/ja/docs/Glossary/Falsy)

#### truthy

falsy以外のものは何を入れても全てtruthyとして認識されます。

### if...else
Goと同じく、else文を用いることができます。

```js
if (false) {
    console.log("実行されない");
} else {
    console.log("実行される");
}

```

### else if
これまたGoと同じくelse ifでさらに分岐できます。

```js
if (false) {
    console.log("実行されない");
} else if (0){
    console.log("実行されない");
} else {
    console.log("実行される");
}

```

### switch文

```js
switch (expression) {
  case value1:
    // 式の結果が value1 に一致する場合に実行する文
    [break;]
  case value2:
    // 式の結果が value2 に一致する場合に実行する文
    [break;]
  ...
  case valueN:
    // 式の結果が valueN に一致する場合に実行する文
    [break;]
  [default:
    // 式の値に一致するものが存在しない場合に実行する文
    [break;]]
}
```

記法は基本的に変わらないのですが、一つおおきな違いがあります。次の例を見てください。

```js
let i = 1
switch (i){
case 0:
    console.log("zero");
case 1:
    console.log("one");
case 2:
    console.log("two");
default:
    console.log("not matched");
}
```

さて、この場合、どのように処理されるでしょうか。正解はこうです。

- 出力

```
one
two
```

Goで考えた場合は、oneが実行された後、switchからその時点で抜け出しそうなものですが、他の言語ではそうはいきません。各caseは該当するものがあったとき、そのcaseのみならず、その後のcaseも一致するかにかかわらず全て実行します。これを防ぐためには、各caseの最後にbreakを記述する必要があります。よって、次のように修正することで、想定した挙動を得ることができます。


```js
let i = 1
switch (i){
case 0:
    console.log("zero");
    break;
case 1:
    console.log("one");
    break;
case 2:
    console.log("two");
    break;
default:
    console.log("not matched");
}

```

まちがえないように気を付けましょう。

## Loop
### For文

```js
for ([initialization]; [condition]; [final-expression]) {
   statement
}
```

この記法は、Golangに於ける

```go
// 初期化ステートメント; 継続ステートメント; 後処理ステートメント
for [InitStatement]; [Condition]; [PostStatement] {
    statement
}
```

と全く同じです。

- 例

```js
for (let i = 0; i < 3; i++) {
    console.log(i);
}
```

- 出力

```
0
1
2
```

### while文
これはGo言語でいうところの、`for condition {}`節に相当します。

```js
while (condition) {
  statement
}
```

- 例

```js
let i = 0;
while (i < 3>) {
    console.log(i);
}
```

- 出力

```
0
1
2
```

無限ループにしたいときは、次のように記述するのが一般的です。

```js
while true {
    処理
}
```

### do...while文
これはGoに対応する文はありません。

```js
do {
   statement
} while (condition);
```
この文のwhile文との違いは、必ずdloの中身を一度は実行することです。それ以外はwhileとかわりません。

### 補足 - continue -

以上の繰り返し操作に於いて、`continue`文を用いると、その時点で次のループへと移行させることができます。

```js
for (let i = 0; i < 10; i++) {
    continue
    console.log(i); // 実行されない。
}
```

それでは、JavaScriptの基本の最後として、他の型を少しみてみましょう。

[次にすすむ](object.md)

## 参考

[for - JavaScript](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Statements/for)<br>
[while - JavaScript](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Statements/while)<br>
[do...while - JavaScript](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Statements/do...while)<br>