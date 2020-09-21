# Web API

## 目次
<!-- TOC -->

- [Web API](#web-api)
    - [目次](#目次)
    - [XMLHttpRequestオブジェクト](#xmlhttprequestオブジェクト)
        - [new](#new)

<!-- /TOC -->

## XMLHttpRequestオブジェクト

JavaScriptによってできることは、もちろんDOM操作だけではありません。例えば、ユーザのコメント投稿をサーバに送信するのに、JavaScriptをもちいてHTTP Requestを送信することも可能です。次の例を見てください。

```js
<!DOCTYPE html>
<html lang="ja">

<head>
    <meta charset="UTF-8">
    <script>
        function Send() {
            // 入力内容を取得
            let text = document.getElementById("text").value;

            console.log(text);

            // POSTするリクエストボディを生成
            let data = "text=" + text;

            let state = document.getElementById("state");

            let request = new XMLHttpRequest();
            request.open('POST', "https://script.google.com/macros/s/AKfycbyEccZlPjl7GOcdK3SXFhE4al1rwhEu6N1pmBhJPKoUCDym77g/exec");
            request.responseType = "json";
            request.onreadystatechange = function () {
                if (request.readyState !== XMLHttpRequest.DONE) {
                    // 送信中の処理
                    state.textContent = "Now Sending..."
                    return;
                }

                if (request.status != 200) {
                    // 送信に失敗
                    state.textContent = "失敗しました。"
                    return;
                }

                // 送信に成功
                let result = request.response;
                state.textContent = result.text;

                return;
            };

            // HTTPヘッダをセット
            request.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded');

            // 送信
            request.send(data);
        }
    </script>
</head>

<body>
    <div id="state">

    </div>
    <form onsubmit="Send();return false;">
        <input type="text" id="text"></input>
        <button type="submit">実行</button>
    </form>
</body>

</html>
```
- <a href="html/request.html">実行結果</a>

<iframe src="html/request.html" name="sample" width="90%" height="70">
    <a href="html/request.html"></a>
</iframe>

### new

```js
var request = new XMLHttpRequest();
```

このnewを前回説明しませんでした。これは、コンストラクタをもったオブジェクトのインスタンスを作成する演算子です。
