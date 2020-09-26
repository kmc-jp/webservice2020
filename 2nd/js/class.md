# Class
[前にもどる](object.md)
[次にすすむ](old_js.md)

## 概要
この章ではJavaScriptのクラスについて説明していきます。

## 目次
<!-- TOC -->

- [Class](#class)
    - [概要](#概要)
    - [目次](#目次)
    - [導入](#導入)
    - [Classとは](#classとは)
    - [class宣言](#class宣言)
    - [インスタンスの作成 - new 演算子 -](#インスタンスの作成---new-演算子--)
    - [プロトタイプメソッド](#プロトタイプメソッド)
    - [コンストラクタ](#コンストラクタ)
        - [補足 - インスタンスへのメソッド付加 -](#補足---インスタンスへのメソッド付加--)
    - [getter / setter](#getter--setter)
    - [継承](#継承)
    - [super](#super)
    - [終わりに](#終わりに)
    - [補足](#補足)

<!-- /TOC -->

## 導入
今までは毎回オブジェクトをリテラルをもちいて一つ一つ生成してきました。では、同じプロパティ、メソッドを持ったものを複数生成したい場合はどうするのが良いでしょうか。もちろん、毎回書くとか、そうではなくても初期化する機構をfor文でまわすとか、やりようはありますが、面倒ですね。そのような場面で活躍するのがClassです。

例えば、次のようなオブジェクトを考えてみましょう。

```js
let user = [];

user[0] = {
    Name: "tkmax777",
    Age: 20,
    Print: function() {
        console.log("Name:%s Age: %d", this.Name, this.Age);
        return;
    },
}
```

もう一人同様なオブジェクトを用意したとします。すると、これまでの知識では、もう一度次のような記述をする必要があります。

```js
user[1] = {
    Name: "defined",
    Age: 20,
    Print: function() {
        console.log("Name:%s Age: %d", this.Name, this.Age);
        return;
    },
}
```

このような記述をすることで確かに目的は達成されるものの、問題もありますね。今仮にこれをfor文を使うなどして簡単に記述できるようにしたとしても、毎回`Print`のようなメソッドごと複製していればメモリを明らかに冗長に消費してしまいます。この方針で大きなオブジェクトを大量に複製したら、とんでもないことになる未来が見えます。これからClassを使うことで、これをもっと簡潔に、そして効率的に記述していきましょう。

## Classとは
先程のような、同様のプロパティ、メソッドを持ったオブジェクトを大量に作成したい場合に用いる、オブジェクトの雛形のことです。この雛形をもちいて作成されたオブジェクトのことを、
**インスタンス**
と呼びます。

## class宣言
まずはクラスの作成をしましょう。次のように記述するだけです。

```js
class ClassName { }
```

これで、複製する雛形の下地が用意されました。

## インスタンスの作成 - new 演算子 -  
次のように記述することで、生成したクラスのインスタンスを作成できます。

```js
let instance_1 = new ClassName(arg1, arg2..., argn)
```

ここで、`argi`には後に出てくるコンストラクタに渡す値を入れます。それではこの下地に更に要素を詰めていきます。

## プロトタイプメソッド
classで作成されるオブジェクトで共有されるメソッドのことをプロトタイプメソッドと呼びます。この登録は次のように行います。

```js
class ClassName {
    Method1() {
        ...
    }
}
```

**注意**<br>
ここでは次のような、オブジェクトのメソッド定義のようなfunctionをつかった定義は使えません。

```js
// 間違い
class ClassName {
    Method1 = function() {
        ...
    }
}
```
注意しましょう。

## コンストラクタ
オブジェクトを作成したとき、まず第一に重要なのは値の初期化ですね。各々のインスタンスを作成したときに自動的に呼ばれるメソッドのことを
**コンストラクタ**
と呼びます。

- 記法

```js
class ClassName {
    constructor(arg1, arg2..., argn) {
        this.Property_1 = arg1;
        this.Property_2 = arg2;
        ...
        this.Property_n = argn;
    }
}
```

以上を踏まえて、先程のuserオブジェクトをclassに落としこむ場合は次のようにします。

```js
class Person { 
    constructor(name, age) {
        this.Name = name;
        this.Age = age;
        return;
    }
    Print() {
        console.log("Name:%s Age: %d", this.Name, this.Age);
        return;
    }
}

// 使ってみる
let user = [];

user[0] = new Person("tkmax777", 20);
user[1] = new Person("defined", 20);
// スッキリ！
```

### 補足 - インスタンスへのメソッド付加 -
上記のような全インスタンスで共有されるようなプロトタイプメソッドが好ましくない場合があるかもしれません。その場合は、constructorで変数を登録するのと同様に、functionを登録すれば問題ありません。

## getter / setter
ここまでで、クラスと、そのインスタンスの作成ができるようになりました。しかし、クラスに登録できるのは単なるプロパティやメソッドだけではありません。次のような、プロパティを呼ばれたときに実行される関数を定義することも可能です。

```js
class ClassName {
    constructor {
        this.Proparty = Data;
    }
    // getter
    get PropartyName() {
        ...
        return PropatyData;
    }
    // setter
    set PropartyName(param) {
        ...
        this.Proparty = param;
    }
}
```

例えば、先程のPersonクラスで、今度は名前の文字数を教えてくれるプロパティ、NameLengthを作ってみましょう。

```js
class Person { 
    constructor(name, age) {
        this.Name = name;
        this.Age = age;
        return;
    }
    Print() {
        console.log("Name:%s Age: %d", this.Name, this.Age);
        return;
    }
    get NameLength() {
        return this.Name.length;
    } 
}

// 使ってみる
let user = [];

user[0] = new Person("tkmax777", 20);

console.log(user[0].NameLength); // => 8
```

アホみたいなプログラムになりますが、setterを使って、登録した文字数のランダム名に変更するプログラムも作れます。

```js
class Person { 
    constructor(name, age) {
        this.Name = name;
        this.Age = age;
        return;
    }
    Print() {
        console.log("Name:%s Age: %d", this.Name, this.Age);
        return;
    }
    get NameLength() {
        return this.Name.length;
    } 
    set NameLength(put_lebgth) {
        this.Name = "";
        
        let src = "abcdefghijklmnopqrstuvwxyz0123456789";
        for (let i=0; i<put_lebgth; i++) {
            this.Name += src[Math.floor(Math.random()*src.length)];
        }
    } 
}

// 使ってみる
let user = [];

user[0] = new Person("tkmax777", 20);
console.log("Name: %s Length: %d", user[0].Name, user[0].NameLength); // => Name: tkmax777 Length: 8

user[0].NameLength = 8;
console.log("Name: %s Length: %d", user[0].Name, user[0].NameLength); // => Name: cy47i4qh Length: 8

```

## 継承
ここまでで、とりあえずクラスの作成が一通りできるようになりました。では次は、既存のクラスの拡張の方法を学んでいきましょう。

例として、次の場合を考えましょう。

```js
class Person { 
    constructor(name, age) {
        this.Name = name;
        this.Age = age;
        return;
    }
    Print() {
        console.log("Name:%s Age: %d", this.Name, this.Age);
        return;
    }
}

class Student(name, age, department, grade) {
    constructor(name, age) {
        this.Name = name;
        this.Age = age;
        this.Department = department;
        this.Grade = grade;
        return;
    }
    Print() {
        console.log(
            "Name:%s Age: %d\nDepartment: %s Grade: %d",
            this.Name, this.Age, this.Department this.Grade,
        );
        return;
    }
}

```

皆さんは、このクラス設計を見てどう思うでしょうか。明らかにStudentクラスはPersonクラスの上位概念として定義されていますよね。今はパラメータが精々名前、年齢、学部、学年で済んでいますが、もしこれが、Personクラスに住所、血液型などの個人情報を詰めこんで、StudentクラスにはPersonクラスの情報に加えて学籍番号、各科目の成績などなどと詰めこんでいったらどうでしょう。これはどう考えてもコード記述的にもメモリ的にも二度手間ですね。

このようなときに使えるのが
**継承**
という操作です。

継承は次のような公式で行います。

```js
class ParentClass { }
class ChildClass extends ParentClass { }
```

これによって定義された`ChildClass`からは、`ParentClass`で定義されたメソッドやプロパティを扱うことができます。

**補足**<br>
この時、`ParentClass`と`ChildClass`、双方で定義された同名のメソッド／プロパティはChildClassが優先されます。

具体例は、次のsuper関数を示したあとに示します。

## super
`ChildClass`から、`ParentClass`のコンストラクタを呼ぶためには`super`という名の関数を呼びだします。

```js
class ParentClass {
    constructor {
        ...
    }
}
class ChildClass extends ParentClass {
    constructor {
        super(arg1, arg2, ...argn)
        ...
    }
}
```
先程の例で記述すると次のようになります。

```js
class Person { 
    constructor(name, age) {
        this.Name = name;
        this.Age = age;
        return;
    }
    Print() {
        console.log("Name:%s Age: %d", this.Name, this.Age);
        return;
    }
}

class Student extends Person {
    constructor(name, age, department, grade) {
        super(name, age);
        this.Department = department;
        this.Grade = grade;
        return;
    }
    Print() {
        console.log(
            "Name: %s Age: %d\nDepartment: %s Grade: %d",
            this.Name, this.Age, this.Department, this.Grade,
        );
        return;
    }
}

let student = new Student("tkmax777", 20, "工学部", 2);
student.Print(); 

```

- 出力

```
Name: tkmax777 Age: 20
Department: 工学部 Grade: 2
```

**注意**<br>
今までやってきたように、変数の内容などは実行された順に上書きされていきます。先程`ParentClass`と`ChildClass`、双方で定義された同名のメソッド／プロパティはChildClassが優先される、ということを記述しました。これを実現するために、`super`関数は他の如何なる`this`の呼び出しよりも先に呼び出さなければならないという規則があります。これを呼び出すまえにthisを呼び出した場合、`ReferenceError`が発生するので気をつけましょう。


## 終わりに

さて、ここまでJavaScriptの基本文法をずっとお話してきました。しかし残念ながら、ここまでではまだDOMの操作は一切できませんね。これを使うためにはWeb APIのDOMインターフェースについて、もうすこし知る必要があります。次回はこれについて、学んでいこうと思います。お疲れさまでした。

## 補足
本編では触れませんが、もし旧式のJSにも興味があれば、次にすすんでも良いです。クラスの内部を直接書く方法など触れるので、ひょっとしたらJSの理解が進むかもしれません。

[前にもどる](object.md)
[次にすすむ](old_js.md)