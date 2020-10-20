# Web API
[前にもどる](readme.md)
[次にすすむ](media_api.md)

## 目次
<!-- TOC -->

- [Web API](#web-api)
    - [目次](#目次)
    - [fetch API](#fetch-api)

<!-- /TOC -->

## fetch API
JavaScriptによってできることは、もちろんDOM操作だけではありません。例えば次のように、httpリクエストを投げることでWeb上からJSONを取ってくるといったことも可能です。


・簡単な例

```js
function GetDate() {
    // fetch先
    let url = "https://script.googleusercontent.com/macros/echo?user_content_key=88KKZ1jRZCH7RZvsu8h68R23CstM3-PtBYTbHEZTUlvAFaT0qK3EyBjJ2qwlRIfO4tUKxwWUzFz0ltDsIVZ6JjwsI5gjoKLMm5_BxDlH2jW0nuo2oDemN9CCS2h10ox_1xSncGQajx_ryfhECjZEnLFAC2QYl-PaklxCbk9-bQcQmzv57AI_5Sp-MhCci-zJVa4Q2qeT-HoJIkrsr1xRObxy_u9Yn2NR&lib=MK_r-2dv9bXi6nfG3QlPNaLGHHm8mtwzb"
    
    fetch(url,
        // 各種オプションの入力
        {
            method: "GET",
        }
    ).then(
        function (response) {
            if (!response.ok) {
                console.log("Error: " + response.status);
            }
            return response.json();
        }
    ).then(
        function (res) {
            console.log(res.Date);
        }
    ).catch(
        function (error) {
            console.log("ClientError: " + error);
        }
    )
}
```

・出力

```
2020-10-18T13:34:24.436Z
```

極めて簡単にかけますね。詳しい使いかたは次のページを見ると良いでしょう。

[Fetch API - Web API](https://developer.mozilla.org/ja/docs/Web/API/Fetch_API)

[次にすすむ](media_api.md)
