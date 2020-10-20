# 便利なAPI集
## 概要
WebAPIには、DOMの操作をするDOM APIの他にも沢山の便利なAPIが含まれています。そのうち、特に便利なものの利用例を少しだけ紹介します。

## 目次
<!-- TOC -->

- [便利なAPI集](#便利なapi集)
    - [概要](#概要)
    - [目次](#目次)
    - [Media Capture and Streams API](#media-capture-and-streams-api)
        - [概要](#概要-1)
        - [簡単な実装例](#簡単な実装例)
    - [その他の便利なWebAPI](#その他の便利なwebapi)
    - [おわり](#おわり)

<!-- /TOC -->

## Media Capture and Streams API 
### 概要
皆さんが今つかっているパソコンやスマートフォンなどには、大抵カメラやマイクなどの外部リソースとなる装置が付いているのでは無いでしょうか。そのような装置からの情報を操作できるAPIが用意されています。

### 簡単な実装例

```js
<!DOCTYPE html>
<html lang="ja">

<head>
    <meta charset="UTF-8">
    <script>
        function startVideo() {
            // 動画を表示する先を指定
            let videoElement = document.getElementById("video")

            // 取ってくる資源の種類(制限)を指定
            // getUserMediaメソッドはPromise型によってメディアストリームを返す
            navigator.mediaDevices.getUserMedia(
                {
                    video: true,
                    audio: true,
                }
            ).then(
                // この関数は、getUserMediaの後に実行され、streamにはメディアストーリームが実行される
                function (stream) {
                    // それぞれ、端末のメディア資源の一覧が代入される
                    let audioTracks = stream.getAudioTracks();
                    let videoTracks = stream.getVideoTracks();

                    console.log(audioTracks);

                    // video要素の資源にストリームを指定
                    videoElement.srcObject = stream;

                    // 端末のメディア資源の数が0以外(truthy)のときに実行する
                    if (audioTracks.length) {
                        audioTrack = audioTracks[0];
                    }
                    if (videoTracks.length) {
                        videoTrack = videoTracks[0];
                    }
                    
                    return;
                }
            ).catch(
                // エラー処理
                function (error) {
                    console.log(error);
                    return
                }
            );
        }
    </script>
</head>

<body>
    <video id="video" autoplay></video><br>
    <a href="javascript:startVideo();">
        <button type="button">実行</button>
    </a>
</body>

</html>
```

[実行結果](html/media_stream.html)

このように動画のストリームを取得することができます。便利ですね。

この他、取得する画像サイズやフレームレートの指定などの詳しい設定もできます。詳しい利用方法は次の参照すると良いでしょう。

[Media Capture and Streams API](https://developer.mozilla.org/ja/docs/Web/API/Media_Streams_API)

## その他の便利なWebAPI

今回は時間がなくて用意できませんでしたが、次のような便利なWebAPIが他にもあります

- [CanvasAPI](https://developer.mozilla.org/ja/docs/Web/API/Canvas_API)
HTMLのキャンバスタグ内に、自由な図形を描画できるAPIです。

- [Server-sent events](https://developer.mozilla.org/ja/docs/Web/API/Server-sent_events)
サーバからテキストデータをプッシュ送信することが可能になるAPIです。

- [WebSocketAPI](https://developer.mozilla.org/ja/docs/Web/API/WebSockets_API)
サーバと双方向通信が可能になるAPIです。

## おわり
今回は代表的なWebAPIの使いかたについて見てみました。しかし、もちろんブラウザに実装されているWebAPIはこれだけではありません。今後、自分の知らない動作をしているWebページを見かけたときは、是非、それを支えているWebAPIを調べてみると良いかもしれません。

ではでは。