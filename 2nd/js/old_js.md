# 発展 - 旧式JSについて -
[前にもどる](class.md)

## 概要
ここには今までに記さなかった旧式のJSの一部の説明を載せておきます。ここにある記法は現代においては時代遅れのため、使わないほうが良いですが、理解しておくとJSの世界の成りたちがすこし見えます。興味があれば見てください。

## 目次
<!-- TOC -->

- [発展 - 旧式JSについて -](#発展---旧式jsについて--)
    - [概要](#概要)
    - [目次](#目次)
    - [基本事項](#基本事項)
        - [+演算子をもちいた 文字列 => 数値 変換](#演算子をもちいた-文字列--数値-変換)
        - [変数の巻き上げ](#変数の巻き上げ)
        - [for ... in Object](#for--in-object)
    - [クラス以外のインスタンス化](#クラス以外のインスタンス化)
        - [コンストラクタ](#コンストラクタ)
        - [new演算子](#new演算子)
        - [Prototype](#prototype)
        - [継承](#継承)
    - [終わり](#終わり)
    - [参考](#参考)

<!-- /TOC -->

## 基本事項

### +演算子をもちいた 文字列 => 数値 変換

+演算子は、文字列を数値に変換する性質も持っています。次の例を見てください。

```js
console.log("3" + 3);  // => 33
console.log(+"3" + 3); // => 6
```

このように、文字列の前に`+`演算子を付けることで、簡単に数値へと変換することができます。

先程の例ではうまく変換できましたが、ひねくれたユーザが文字列を入力するかもしれません。その場合は次のようになります。

```js
console.log(+"「さん」のつもり"); // => NaN
```

### 変数の巻き上げ
JavaScriptに於いて、実のところは巻き上げルールは関数の限りではありません。次の例を見てください。

```js
function awesome_func_1(){
    var x = "Hello";
    console.log(x);

    var y = function (){
        console.log(x);
        var x = "NULL";
    }

    y();

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

```
Hello
undefined
Hello
```

これは先程の関数の巻き上げと同様の操作が変数に於いても行われていることによります。if文内で記述されているlet節により、関数の宣言だけ先行しているのです。注意しなければいけないのは、次のような記述をしても初期化はなされません。

```js
console.log(x); // -> undefined
var x = 5;
```

この、宣言だけされて、初期化されていない空白のことを、
**TDZ(Temporal Dead Zone)**
と呼びます。

このような記法はバグの原因になりますし、そもそも見にくいです。できるだけ宣言は関数の先頭で行うように心がけましょう。

### for ... in Object
`for ... of Object.keys(obj)`は次のようにも記せます。

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


## クラス以外のインスタンス化

### コンストラクタ
本編ではクラスをもちいた、極めて現代風なクラスの記法を記しました。ここからは、旧式の記法を記します。クラスの内部を直接記述していくイメージになるので、クラスの構造の理解の助けになるかもしれません。

ところで、本編で`Function`ですら、これはオブジェクトなのだ、ということを書きました。このことを利用して、同じオブジェクトのコピー(
**インスタンス**
)を大量に作成することを考えてみましょう。次の例を見てください。

```js
function User(name, age, club) {
    this.Name = name;
    this.Age = age;
    this.Club = club;
    
    return;
}
```

`this`は自分自身を参照する識別子でしたね。このように、実行することで、自身の構造を定義してくれる関数のことを
**コンストラクタ**
と呼びます。ここで、コンストラクタの関数名は最初大文字から始めるのが通例となっていることを覚えておきましょう。では、早速試しに次のように実行してみましょう。

- コード

```js
let user_1 = User("tkmax777", 20, "KMC");
let user_2 = User("morisawa", 20, "KMC");

console.log("Name: %s Age: %d Org: %s", user_1.Name, user_1.Age, user_1.Club);
console.log("Name: %s Age: %d Org: %s", user_2.Name, user_2.Age, user_2.Club);
```

- 出力

```
Uncaught TypeError: Cannot read property 'Name' of undefined at <anonymous>:4:52
```

悲しいことに、このままでは期待していたようにはうまくいきません。何故でしょうか。このプログラムには大きく２つの問題が潜んでいます。１つ目は、そうです。`user_1`や`user_2`には、`User`関数のオブジェクトではなく、実行した結果(`undefined`)が直接代入されてしまったのです。先のエラーはその結果、`undefined`くんには`Name`なんという名前のプロパティは無いと怒られてしまったのですね。

しかしこの関数には、関数がオブジェクトだとということを理解していれば気がつく、もう一つの大きな問題があります。先程、オブジェクトのコピーを大量に生成する、といいましたが、このコードでは、Userオブジェクトを直接編集しているため、もし仮にうまく行っていても`user_1`の情報は`user_2`の処理のときに上書きされてしまい、やはり期待通りには動かないことが予想されますね。

どうにかして、この２つの問題を回避し、この関数のオブジェクトを複製し、さらにその関数(コンストラクタ)の生成するオブジェクトを代入させるため方法は無いでしょうか。

### new演算子
もちろんあります。それが、`new`演算子です。先程のコンストラクタ、`User`を例に見てみましょう。

```js
let user_1 = new User("tkmax777", 20, "KMC");
let user_2 = new User("morisawa", 20, "KMC");

console.log("Name: %s Age: %d Org: %s", user_1.Name, user_1.Age, user_1.Club);
console.log("Name: %s Age: %d Org: %s", user_2.Name, user_2.Age, user_2.Club);

```

```
Name: tkmax777 Age: 20 Org: KMC
Name: morisawa Age: 20 Org: KMC
```

うまくいきましたね。`new`演算子では、先程挙げたように、指定したオブジェクトを複製し、コンストラクタを実行することで初期化してくれる働きがあります。こうして得られた複製オブジェクトをカスタマイズし、新しいオブジェクトを生成していくことでJavaScriptの世界は広がっていきます。重要なので絶対に覚えておきましょう。

### Prototype
さて、ここまでは、基本的に一つのオブジェクトに着目して話をしてきました。しかし、実際にはオブジェクトは複雑な聯関の中で存在しています。その聯関を支える概念、それがPrototypeです。

JavaScriptの世界に於いて、全てのオブジェクトは`prototype`プロパティを保持しています。また、同時にすべてのオブジェクトは`__proto__`プロパティも保持します。(ただし、ブラウザによっては直接参照ができなかったりします。)この二つの対応がとても肝です。

各々の`prototype`プロパティは、任意のオブジェクトを参照先として取ることができます。対して、`__proto__`プロパティの参照先は、そのオブジェクトの複製元となるオブジェクトの`prototype`オブジェクトとなっています。

例えば、次の例を見てください。

```js
function Person(name, age) {
    this.Name = name;
    this.Age = age;

    return;
}

let user_1 = new Person("tkmax777", 20);
let user_2 = new Person("defined", 20);

```

この例に於いて、`user_1`オブジェクトや、`user_2`オブジェクトの`__proto__`オブジェクトの参照先は、複製元である、`Person`オブジェクトの`prototype`オブジェクトになっています。ついでに、`Person`オブジェクトは`Function`オブジェクトを複製することで作成されたオブジェクトなことに注意します。そうすることでわかるように、`__proto__`オブジェクトはこのコード上にはないものの、組み込みオブジェクトである、`Function`オブジェクトの`prototype`オブジェクトとなっています。


さて、ここまで読んだだけでは、なんか複雑なことをしているわりに、何が嬉しいのかわからない、となりますね。実は`prototype`には一つ、有意義な大きな性質があるのです。

それは、自分自身には定義されていないプロパティやメソッドが呼びだされたとき、`__proto__`を遡り、上位にその定義がないかを検索して、あればそれを実行してくれることです。

……わかんねえよ！ってなりますね。具体例で見てみましょう。

```js
function Person(name, age) {
    this.Name = name;
    this.Age = age;
    this.Print = function() {
        console.log("Name:%s Age: %d", this.Name, this.Age);
        return;
    }
    return;
}

let user_1 = new Person("tkmax777", 20);
let user_2 = new Person("defined", 20);

user_1.Print()
user_2.Print()
```

皆さんはこの例を見てどう思うでしょうか。実行してみると、このプログラムは正しく実行され、次のような出力をします。

```
Name:tkmax777 Age: 20
Name:defined Age: 20
```

しかし、一つ問題があります。良くみてみると、このプログラムに於いてPersonコンストラクタはメソッドまでオブジェクトに取り込んでいます。そのようなものを、new演算子で複製していますね。この記法は今後、メソッドが増えていく上でメモリを大量に消費する元となります。

このままでは大きなプログラムを作ることができません。こんなとこに、先程の`prototype`の性質が真価を発揮します。このコードは次のように最適化されます。


```js
function Person(name, age) {
    this.Name = name;
    this.Age = age;

    return;
}

Person.prototype.Print = function() {
    console.log("Name:%s Age: %d", this.Name, this.Age);
    return;
}

let user_1 = new Person("tkmax777", 20);
let user_2 = new Person("defined", 20);

user_1.Print()
user_2.Print()
```

これでも全く同じ出力を得ることができます。これは、`user_1`、`user_2`がともに`Person`オブジェクトから生成することを利用し、Personオブジェクトの`prototype`オブジェクトに登録したのです。

### 継承
さて、ここまででようやく、オブジェクトの複製を気兼ねなく行えるようになってきたのではないでしょうか。では今度は、あるオブジェクトのメソッドやプロパティを保持しつつ、さらに別の機能やプロパティをも持つようなオブジェクト、いわばあるオブジェクトに対する上位互換のオブジェクトの作成方法を考えてみましょう。

先ほどのセクションのプログラムをまた使います。

```js
function Person(name, age) {
    this.Name = name;
    this.Age = age;

    return;
}

Person.prototype.Print = function() {
    console.log("Name:%s Age: %d", this.Name, this.Age);
    return;
}

```

この`Person`のプロパティを満たした上で新たな要素も保持するようなオブジェクト`Student`を作成します。このとき、次のような記法が使えます。

```js
function Person(name, age) {
    this.Name = name;
    this.Age = age;

    return;
}

Person.prototype.Print = function() {
    console.log("Name:%s Age: %d", this.Name, this.Age);
    return;
}

function Student(department, grade, name, age) {
    this.Department = department;
    this.Grade = grade;
    
    Person.call(this, name, age);
}

Student.prototype = new Person();

let user_1 = new Student("工学部", 2,"tkmax777", 20);

user_1.Print()

```

このプログラムで新しい記法が出てきたのはここの部分ですね。

```js
function Student(department, grade, name, age) {
    this.Department = department;
    this.Grade = grade;
    
    Person.call(this, name, age);
}

Student.prototype = new Person();
```

`call`メソッドは、対応するコンストラクタを、第一引数に入れたオブジェクト空間に対して実行してくれます。今回はこれによって、Studentオブジェクトに対して、Personオブジェクトのコンストラクタが設定してくれるプロパティの設定を行った、ということです。

```js
Student.prototype = new Person();
```
ここは、StudentオブジェクトのprototypeオブジェクトをPersonオブジェクトのインスタンスに設定することで、Personの持つ、メソッドをそのまま参照先に登録しています。

JavaScriptに於ける継承の方法はこれだけではありません。気になる場合は次のページを参考にすると良いでしょう。

[【JavaScript入門】初心者でも分かる継承の作り方（ES2015対応）](https://www.sejuku.net/blog/28158)

## 終わり
このページのことを理解する必要は一切ないですが、それでも興味があってわからないところがあれば、#tk-memo_777とかで聞いてもらえれば、と思います。

[前にもどる](class.md)

## 参考
[【JavaScript入門】プロトタイプ(prototype)の使い方と継承まとめ！ | 侍エンジニア塾ブログ（Samurai Blog） - プログラミング入門者向けサイト](https://www.sejuku.net/blog/47722)<br>
[や...やっと理解できた！JavaScriptのプロトタイプチェーン - maeharinの日記](https://maeharin.hatenablog.com/entry/20130215/javascript_prototype_chain)<br>
[Object のプロトタイプ - ウェブ開発を学ぶ ](https://developer.mozilla.org/ja/docs/Learn/JavaScript/Objects/Object_prototypes)<br>