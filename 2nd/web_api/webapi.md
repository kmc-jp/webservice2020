# Web API

## 目次
<!-- TOC -->

- [Web API](#web-api)
    - [目次](#目次)
    - [HTTPRequestの送信](#httprequestの送信)
    - [XMLHttpRequest オブジェクト](#xmlhttprequest-オブジェクト)
        - [new](#new)

<!-- /TOC -->

## HTTPRequestの送信

JavaScriptによってできることは、もちろんDOM操作だけではありません。例えば、ユーザのコメント投稿をサーバに送信するのに、JavaScriptをもちいてPOSTを送信することも可能です。次の例を見てください。

```js
<!DOCTYPE html>
<html lang="ja">

<head>
    <meta charset="UTF-8">
    <script>
        function Send() {
        var request = new XMLHttpRequest();
        request.open('POST', "https://script.google.com/macros/s/AKfycbyEccZlPjl7GOcdK3SXFhE4al1rwhEu6N1pmBhJPKoUCDym77g/exec");
        request.responseType = "json";
        request.onreadystatechange = function () {
            let state = document.getElementById("state");

            if (request.readyState != 4) {
                // 送信中
            } else if (request.status != 200) {
                state.textContent = "失敗しました。"
            } else {
                let result = request.response;
                state.textContent = result.time 
            }
        };
        request.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded');
        request.send();
    }
    </script>
</head>

<body>
    <div id="state">
        
    </div>
    <a href="javascript:Send()">
        <button type="button">実行</button>
    </a>
</body>

</html>
```
- <a href="html/request.html">実行結果</a>

<iframe src="html/request.html" name="sample" width="90%" height="70">
    <a href="html/request.html"></a>
</iframe>

## XMLHttpRequest オブジェクト
このオブジェクトを用いることで、HTTPリクエストを自在に送受信することができます。

### new

```js
var request = new XMLHttpRequest();
```

このnewを前回説明しませんでした。これは、コンストラクタをもったオブジェクトのインスタンスを作成する演算子です。
