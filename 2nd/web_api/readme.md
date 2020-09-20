# Web APIを使う
## 概要
前回まで学んできたJavaScriptを使って、いよいよ実践的な知識を学んでいきます。

## Web APIを使ったHello World!

では早速、Web APIを使った、簡単なプログラムを見てみましょう。

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

このように、ブラウザに対して直接働きかけることができるのです。少しずつ、使いかたを見ていきましょう。

### HTMLへの埋め込みの仕方

`<script>`
タグの要素として埋め込むことで、
**上から順に**
読み込まれます。

HTML内には複数のscriptタグを入れることができます。しかし、このとき、一つ気をつけなければいけないことがあります。それは、読み込みの順です。

今後、読み込まれた直後の処理などを書いていくとき、複数のタグに分けられた状態にしておくと、

### alert関数

```js
alert("Hello World!")
```

alert関数は、引数にわたされた要素を先程確認したように、ポップアップで表示する、組み込み関数です。引数には任意の型のデータを渡すことができます。引数にはstring文字列を取ることができます。

[alert](https://developer.mozilla.org/ja/docs/XPInstall_API_Reference/Install_Object/Methods/alerts)

## DOM APIを使う
では、ここから実際に要素の移動を実行していみましょう。

次のHTMLを用意します。

```js
<!DOCTYPE html>
<html lang="ja">
    <head>
        <meta charset="UTF-8">
        <script>
            function Ex(){
                document.getElementById("text").textContent = "Excuted!";
            }
        </script>
    </head>
    <body>
        <div id="text"></div>
        <a href="javascript:Ex();">
            <button type="button">実行</button>
        </a>
    </body>
</html>

```

- 実行結果

<iframe src="html/edit_text_content.html" name="sample" width="90%" height="70">
    <a href="html/edit_text_content.html"></a>
</iframe>

