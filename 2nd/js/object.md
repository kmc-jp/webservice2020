# その他の型
[前にもどる](flow.md)
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
            - [参考 - for ... of Object.keys(obj) 〜 -](#参考---for--of-objectkeysobj-〜--)
        - [メソッド](#メソッド)
            - [リテラル](#リテラル-1)
            - [短縮表記](#短縮表記)
            - [個別登録](#個別登録-1)
            - [自身を参照する](#自身を参照する)
        - [組み込みオブジェクト](#組み込みオブジェクト)
        - [コンストラクタ](#コンストラクタ)
        - [new演算子](#new演算子)
        - [Global Object](#global-object)
            - [参考 - グローバル変数 -](#参考---グローバル変数--)
        - [Prototype](#prototype)
        - [継承](#継承)
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

また、findメソッドをつかって、次のように検索することもできます。

```js
const inventory = [
  {name: 'apples', quantity: 2},
  {name: 'bananas', quantity: 0},
  {name: 'cherries', quantity: 5}
];

function isCherries(fruit) { 
  return fruit.name === 'cherries';
}

console.log(inventory.find(isCherries)); 
// { name: 'cherries', quantity: 5 }
```
[引用元](https://developer.mozilla.org/ja/docs/Web/JavaScript/Reference/Global_Objects/Array/find)

このメソッドでは、入力された関数を各要素について実行します。その結果、truthyが返された時点で、その値を返します。要素が見つからなかった場合、undefinedを返します。

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

#### 参考 - for ... of Object.keys(obj) 〜 - 
配列の時みたく、オブジェクトの全プロパティに対してある処理をしたい場合、次のように記述できます。

- 例

```js
let ExObject = {
    Proparty_1: "First",
    Proparty_2: "Second",
    Proparty_3: "Third"
}

for (proparty of Object.keys(ExObject)) {
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

#### 短縮表記
メソッドの登録には、次のようにも記述することが可能です。

```js
let Object = {
    Method_1(parameter) {
        処理
        return 
    },
    Method_2(parameter) {
        ...
    },
    ...
}
```

どちらで記述しても全く同じなので、好みな方で記述しましょう。

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
これからオブジェクトを扱っていく上で、自身の持つ他のプロパティなどを参照したい場合があるかもしれません。その場合は、自身を`this`として表すことで、これを達成することができます。

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

しかし、それだけではありせん。今までに型として登場した、`Array`、`Number`、`String`など、さらには関数`Function`まで、JavaScriptの世界の全てがオブジェクトから成り立っています。

その全貌を説明することはここではできませんが、興味があるのなら、各々調べてみると良いでしょう。

[標準組み込みオブジェクト](https://developer.mozilla.org/ja/docs/Web/JavaScript/Reference/Global_Objects)<br>
[Web API](https://developer.mozilla.org/ja/docs/Web/API)<br>

### コンストラクタ
今までは毎回オブジェクトをリテラルをもちいて一つ一つ生成してきました。では、同じプロパティ、メソッドを持ったものを複数生成したい場合はどうするのが良いでしょうか。もちろん、毎回書くとか、そうではなくても初期化する機構をfor文でまわすとか、やりようはありますが、面倒ですね。

ところで、先程`Function`ですら、これはオブジェクトなのだ、ということを書きました。このことを利用して、同じオブジェクトのコピー(
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

### Global Object

ところで、JavaScriptの世界は、全てのモノがオブジェクトなのだ、ということを以前に軽く触れました。実はそんなオブジェクトですが、その全てのオブジェクトは
**Global Object**
と呼ばれる、唯一のオブジェクトのプロパティとして保持されています。そのプロパティに、Functionオブジェクトが組み込みで存在しています。それを複製し、カスタマイズすることで、プログラマは関数オブジェクトを構築でき、さらにその関数オブジェクトを複製し、カスタマイズすることで多くのオブジェクトをどんどん作成できるのです。

ブラウザでは、このGlobal Objectは`window`という名前のオブジェクトで存在します。実際に次のような操作が可能です。<br>ref)
[グローバル関数へのアクセス](https://developer.mozilla.org/ja/docs/Glossary/Global_object#Access_Global_Functions)

```js
function greeting() {
   console.log("Hi!");
}

window.greeting(); // ふつうに greeting(); を呼び出すのと同じ
```

ついでに一応、`Function`もオブジェクトだということを示すために記しますが、次のように関数をfunction宣言を使わずに宣言することも可能です。(好ましくはありません。)<br>ref) 
[Funciton コンストラクターと関数宣言の違い](https://developer.mozilla.org/ja/docs/Web/JavaScript/Reference/Global_Objects/Function#Difference_between_Function_constructor_and_function_declaration)

```js
let func = new Function ([arg1[, arg2[, ...argN]],] functionBody)
```

このことは次のPrototypeを理解する上でとても大事なことなので、絶対に覚えておきましょう。

#### 参考 - グローバル変数 -
このことを活かして、ブラウザに於けるJavaScriptではグローバル変数を次のように定義することができます。

```js
window.VAR_NAME = "値"
```

また、ブラウザ以外ではGlobalObjectの名前は違うことに注意します。このことから、他の環境のJSと互換性を持たせるために
**globalThis**
という名前のオブジェクトが用意されています。これを次のように用いても、同様の操作が可能です。

```js
globalThis.VAR_NAME = "値"
```

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

## 終わりに
さて、ここまでJavaScriptの基本文法をずっとお話してきました。しかし残念ながら、ここまでではまだDOMの操作は一切できませんね。これを使うためにはWeb APIのDOMインターフェースについて、もうすこし知る必要があります。次回はこれについて、学んでいこうと思います。お疲れさまでした。

[前にもどる](flow.md)

## 参考
[Array - JavaScript](https://developer.mozilla.org/ja/docs/Web/JavaScript/Reference/Global_Objects/Array)<br>
[JavaScript 型付き配列 - JavaScript](https://developer.mozilla.org/ja/docs/Web/JavaScript/Typed_arrays)<br>
[オブジェクト初期化子](https://developer.mozilla.org/ja/docs/Web/JavaScript/Reference/Operators/Object_initializer)<br>
[オブジェクトでの作業 - JavaScript](https://developer.mozilla.org/ja/docs/Web/JavaScript/Guide/Working_with_Objects)<br>
[Object のプロトタイプ - ウェブ開発を学ぶ ](https://developer.mozilla.org/ja/docs/Learn/JavaScript/Objects/Object_prototypes)<br>
[JavaScriptでメソッドを使う方法【初心者向け】](https://techacademy.jp/magazine/5562#:~:text=%E3%83%A1%E3%82%BD%E3%83%83%E3%83%89%E3%81%A8%E3%81%AF-,%E3%83%A1%E3%82%BD%E3%83%83%E3%83%89%E3%81%A8%E3%81%AF%E3%80%81%E3%82%AA%E3%83%96%E3%82%B8%E3%82%A7%E3%82%AF%E3%83%88%E3%81%AE%E3%83%97%E3%83%AD%E3%83%91%E3%83%86%E3%82%A3%E3%81%AB%E4%BB%A3%E5%85%A5%E3%81%95%E3%82%8C%E3%81%9F,%E3%82%92%E3%81%A4%E3%81%91%E3%81%9F%E3%82%82%E3%81%AE%E3%81%A7%E3%81%99%E3%80%82)<br>
[JavaScript Object Notation - Wikipedia](https://ja.wikipedia.org/wiki/JavaScript_Object_Notation)<br>
[【JavaScript入門】「for – of文」の使い方と間違いやすいポイントを徹底解説！](https://www.sejuku.net/blog/20228)<br>
[【JavaScript入門】プロトタイプ(prototype)の使い方と継承まとめ！ | 侍エンジニア塾ブログ（Samurai Blog） - プログラミング入門者向けサイト](https://www.sejuku.net/blog/47722)<br>
[や...やっと理解できた！JavaScriptのプロトタイプチェーン - maeharinの日記](https://maeharin.hatenablog.com/entry/20130215/javascript_prototype_chain)<br>