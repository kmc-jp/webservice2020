# Web API
[前にもどる](readme.md)
## 目次
<!-- TOC -->

- [Web API](#web-api)
    - [目次](#目次)
    - [XMLHttpRequestオブジェクト](#xmlhttprequestオブジェクト)

<!-- /TOC -->

## XMLHttpRequestオブジェクト

JavaScriptによってできることは、もちろんDOM操作だけではありません。例えば、ユーザのコメント投稿をサーバに送信するのに、JavaScriptをもちいてHTTP Requestを送信することも可能です。次の例を見てください。

```html
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
            let data = "text=" + encodeURIComponent(text);

            let state = document.getElementById("state");

            // XMLHttpRequestのインスタンスを生成
            let request = new XMLHttpRequest();

            // 通信方式 ／ 接続先を指定 
            request.open('POST', "https://script.google.com/macros/s/AKfycbyEccZlPjl7GOcdK3SXFhE4al1rwhEu6N1pmBhJPKoUCDym77g/exec");
            
            // レスポンス内容がJSONで送られてくることを登録
            request.responseType = "json";

            // 通信状態が変化するたびに呼ばれる関数を定義
            request.onreadystatechange = function () {
                // 完了を確認
                if (request.readyState !== XMLHttpRequest.DONE) {
                    // 送信中の処理
                    state.textContent = "Now Sending..."
                    return;
                }

                // HTTPステータスコードを取得
                if (request.status != 200) {
                    // 送信に失敗
                    // (ここのコードはサーバの実装によっては200以外を取り得ることに気をつける)
                    state.textContent = "失敗しました。"
                    return;
                }

                // 送信に成功
                // レスポンス内容を取得
                // JSONレスポンスであることを指定していたので、自動的にパースされる。
                let result = request.response;

                // 反映
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

大よその使いかたは恐らく例を見ればわかるとおもうので、詳しくは説明しません。XMLHttpRequestオブジェクトについて、さらに詳しい説明が読みたい場合は次のページを参照すると良いでしょう。

[XMLHttpRequest](https://developer.mozilla.org/ja/docs/Web/API/XMLHttpRequest)
