# 非同期処理について
[前にもどる](arrow.md)

## 目次
<!-- TOC -->

- [非同期処理について](#非同期処理について)
    - [目次](#目次)
    - [概要](#概要)
    - [非同期処理とは](#非同期処理とは)
    - [Promise](#promise)
        - [コンストラクタ](#コンストラクタ)
        - [thenメソッド](#thenメソッド)
        - [catchメソッド](#catchメソッド)
    - [async await](#async-await)
        - [async](#async)
        - [await](#await)
        - [参考 - 即時関数 -](#参考---即時関数--)
    - [おわり](#おわり)

<!-- /TOC -->

## 概要
前回はにゅっとスルーしましたが、JSを書くうちにこれは重要だな、と思ったので、非同期処理についてまとめることにします。

## 非同期処理とは
まずは次の例をみてください。

```js
setTimeout(
    function() {
        console.log("遅延");
    },
    1000
);
console.log("次");
```

ここで、`setTimeout`というのは、第２引数に待ち時間、第１引数にその後に実行する関数を入力できる関数です。

これを実行したら、どのような出力を得られるでしょうか。実はJavaScriptの世界では次のようになります。

```
次
遅延
```

何故普通に考えると、`遅延`という文字が先にきそうなもの。しかし現実はそうはならないのです。これは、setTimeoutという関数が非同期で行われる関数なことに由来します。

非同期処理とはこのように、処理の度にブロック(=一時停止)することなく、次の処理へ移行する処理方法のことです。この場合、setTimeout関数の処理は待たれず、すぐに`次`という文字が出力されたわけです。

JavaScriptではこのような、非同期関数がその後次に呼び出す関数(**コールバック関数**)を引数として入力することで、非同期関数の処理の連鎖を記述する方式をとります。しかし、この方式はコールバックの回数が多くなればなるほど見にくくなることも予想されますね。例えば次の例をみてください。

```js
setTimeout(
    function() {
        console.log("1秒後に実行");
        setTimeout(
            function() {
                console.log("2秒後に実行");
            },
            1000
        )
    },
    1000
);
```

どうでしょうか。たった一回増やしただけでこのようになります。これでは連鎖を増やせば増やすほど見にくくなっていき、手をつけられなくなることがわかりますね。そこで、`Promise`という、便利なオブジェクトが用意されています。

## Promise
このオブジェクトを用いると、次のように先の処理を記述できます。

```js
function setTimeoutPromise() {
    return new Promise(
        function(resolve, _) {
            setTimeout(resolve, 1000);
            return;
        },
    )
}

setTimeoutPromise()
    .then(
        function() {
            console.log("1秒後に実行");
        },
    ).then(
        function() {
            return setTimeoutPromise()
        }
    ).then(
        function () {
            console.log("2秒後に実行");
        },
    )
    
```

このように縦に処理を書き続けられるので、見やすくなります。とはいえ、これでは恐らく何が起こっているのかわからないと思うので、もうちょっと解説します。

### コンストラクタ

```js
new Promise(
    function (successCallback, failedCallback) {
        処理
        if (error) {
            failedCallback;
            return;
        }
        succecssCallback;
        return;
    },
)
```

まず、Promiseのコンストラクタは第一引数に処理が成功した場合に呼び出す関数、第二引数に失敗した場合に呼び出す関数を入力します。

### thenメソッド

通常はsuccessCallbackの実態を記述します。

```js
let promise = new Promise(
    function (successCallback, failedCallback) {
        処理
        if (error) {
            failedCallback(arg1, arg2, ...);
            return;
        }
        succecssCallback(arg1, arg2, ...);
        return;
    },
)

let promise_2 = promise.then(
    function(arg1, arg2, ...) {
        let res;
        処理
        return res
    }
)

promise_2.then(
    function(res) {
        処理
        return;
    }
)

```

最初の例では直接メソッドを呼び出しましたが、わかりやすく記述すればこうなります。このように、処理後に続けて呼び出す関数を入力できます。このとき、thenメソッドもまた、入力した関数の返り値を、新たに入力する関数の引数とするような、Promise型のオブジェクトを返り値として持ちます。これにより、例のようにさらに続けて処理を記述することも可能です。

### catchメソッド
failedCallbackとして呼び出される関数を入力します。

```js
let promise = new Promise(
    function (successCallback, failedCallback) {
        処理
        if (error) {
            failedCallback(arg1, arg2, ...);
            return;
        }
        succecssCallback(arg1, arg2, ...);
        return;
    },
)

...

promise.catch(
    function (arg1, arg2, ...) {
        処理
        console.log("Error")
    }
)

```

基本的にはthenと同じです。Promiseを返り値として持つため、同様にその後の処理を記述することも可能です。

catchメソッドには次のような使い方もできます。

```js

new Promise(
    function () {
        処理
        if (error) {
            // エラーを投げる
            throw arg
            return;
        }
        return;
    },
).catch(
    function(arg) {
        // エラー時に呼び出される
        処理
        return;
    }
)

```

このように、try-catch文のような働きもします。便利ですね。

## async await
### async
自分で非同期処理の関数を作成する場合、次のように`async`を付けて関数宣言をすることで、その関数内の処理を全て非同期で実行することができます。

- 例

```js
async function awesome_func() {
    ...
}
```

アロー関数の場合は次のように記述します。

```js
let awesome_func = async () => {
    ...   
}
```

このようにして宣言した関数の内部の処理は、呼び出されたときブロックされることなく実行されます。その為、その後の処理を記述しやすいように、暗黙的に返り値はPromise型のオブジェクトになります。

- 例

```js

let awesome_func => async function(){
    // fetchについては後に触れます
    // 時間がかかる処理です
    fetch("https://hoge.example", {}).then(...)
    return "YayYay!"
}

awesome_func().then(
    function(arg) {
        console.log(arg);
    }
)
console.log("Hoge")

```

このようにすると、

```
Hoge
YayYay!
```
という結果が得られます。

### await

**async関数の内部に於いて**、**非同期処理の関数でPromise型の返り値を持つものを呼び出すときに**、その関数の前に`await`を付けることによって、Promiseの解消(=Promiseで管理された、並行処理関数の終了)を待つことができます。

- 例

```js
let awesome_func = async function() {
    処理
    // 終了がまたれる。
    await new Promise(function(resolve) { setTimeout(resolve, 1000) })
    console.log("一秒後")

    return
}

```

これをもちいて、一番最初の例を記述するとこうなります。

```js
function setTimeoutPromise() {
    return new Promise(
        function(resolve, _) {
            setTimeout(resolve, 1000);
            return;
        },
    )
}

(
    async function() {
        await setTimeoutPromise();
        console.log("1秒後に実行");
        await setTimeoutPromise();
        console.log("2秒後に実行");
    }
)();

```

さらにスッキリしましたね。

### 参考 - 即時関数 -
上記の例で、突然グローバル領域で`(function(){ ... })()`という記法を用いました。これは即時関数という記法で、内部で定義した関数を即時に実行することができる記法です。これによって、一回だけ実行したい関数など、グローバル領域の名前空間の圧迫を防ぐことができます。

## おわり
非同期処理については以上です。それでは今度こそWebAPIについて学んでいきましょう。
