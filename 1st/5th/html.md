
## HTML
Webページを構成するのに最低限必要不可欠なもの、それがHTMLです。まずはこれがないと何もはじまりません。まずは簡単な例を見てみましょう。

### 基本

```html
<!DOCTYPE html>
<html lang="ja">
    <head>
        <meta charset="UTF-8">
        <link rel="author" href="mailto:someone@example.jp">
        <title lang="jp">京大マイコンクラブ(KMC)</title> 
    </head>
    <body>
        <article>
            <h1 lang="ja">KMCとは？</h1>
            <p>
                京大マイコンクラブは京都大学を中心に活動している京都大学全学公認のコンピュータサークルです。
                <b>マイコン</b>
                という言葉は今では死語ですが、1977年の設立以来その名を引き継いで現在に至ります。
                KMCの歴史については
                <a href="https://kmc.gr.jp/guidance/history.html">KMCの歴史</a>
                をご覧下さい。
            </p>
        </article>
    </body>
</html>

```
**実行例**<br>
<article>
    <h1 lang="ja">KMCとは？</h1>
    <p>
        京大マイコンクラブは京都大学を中心に活動している京都大学全学公認のコンピュータサークルです。
        <b>マイコン</b>
        という言葉は今では死語ですが、1977年の設立以来その名を引き継いで現在に至ります。
        KMCの歴史については
        <a href="https://kmc.gr.jp/guidance/history.html">KMCの歴史</a>
        をご覧下さい。
    </p>
</article>


順番に要素をみていきましょう。

```html
<!DOCTYPE html>
<html lang="ja">
...
</html>
```

htmlでは、`<...>`をタグといい、基本的に一対の`<...>`と`</...>`で囲われた部分(要素)を入れ子構造で記述していくことで構成されます。

 一行目はこのスクリプトがHTML5で記述されていることを示すDocumentType宣言(Document Type Definition, DTD, 文書型宣言)▽です。 この情報をもとに、ブラウザはここからファイルを解釈していきます。 

 二行目と最終行のhtmlタグで囲われた範囲は、ここがhtml文書であることを示しています。二行目の`lang="ja"`は、この文章の言語が日本語であることを指定しています。このように、タグに属性(ここでいう`lang`)を指定することで、タグにさらに細かいオプションを指定することが出来ます。


```html
    <head>
        <meta charset="UTF-8">
        <link rel="author" href="mailto:someone@example.jp">
        <title lang="jp">京大マイコンクラブ(KMC)</title> 
    </head>
```

head要素には、このhtml全体のメタデータ、つまり文章の情報を指定します。今回の例では、様々なメタデータを格納することができるmetaタグで文字コードを指定したり、この文章に関連する外部リンクを指定するlinkタグで筆者の連絡先を指定したり、このページのタイトルをtitleタグで指定したりしています。

```html
    <body>
        <article>
            <h1 lang="ja">KMCとは？</h1>
            <p>
                京大マイコンクラブは京都大学を中心に活動している京都大学全学公認のコンピュータサークルです。
                <b>マイコン</b>
                という言葉は今では死語ですが、1977年の設立以来その名を引き継いで現在に至ります。
                KMCの歴史については
                <a href="https://kmc.gr.jp/guidance/history.html">KMCの歴史</a>
                をご覧下さい。
            </p>
        </article>
    </body>
```

body要素には実際にユーザが見る要素を記述していきます。

```html
<article>
...
</article>
```

この要素のなかには、この文章のなかで、記事として内容が完結しているものを記述します。

**参考**<br>
記事ではなく、単なる要素の塊をまとめる際には`section`タグを用います。詳しくは、
[ここ](http://www.htmq.com/html5/)
のページを参考にするとわかり易く解説されています。

```html
<h1>...</h1>
```

一般に、`<h番号>`タグは、見出しを記述するためにもちいます。articleやsectionなどのセクションの先頭にはその性質上、基本的にこのタグを記述することになることを覚えておくといいかもしれません。

```html
<p>...</p>
```
このタグは、パラグラフを指定するときにつかいます。覚えやすい!!

```html
<b>マイコン</b>
```
このタグで囲われた部分は、他の単語と区別すべき単語、例えばキーワードや、固有名詞などをいれることができます。大抵太文字になりますが、強調したい場合は、`<strong>`や、`<em>`タグを一般的に用いることに注意しましょう。

```html
<a href="https://kmc.gr.jp/guidance/history.html">KMCの歴史</a>
```
このタグはリンクを指定するときに用います。`href`属性には相対パスや絶対パスなどを指定することが出来ます。超つかうので絶対覚えておきましょう。


### 表
htmlでは、次のように表を作成することができます。

```html
<table>
    <tr>
        <th>A</th>
        <th>B</th>
        <th>C</th>
    </tr>
    <tr>
        <td>a</td>
        <td>b</td>
        <td>c</td>
    </tr>
    <tr>
        <td>α</td>
        <td>β</td>
        <td>γ</td>
    </tr>
</table>  
```
**実行例**<br>
<table>
    <tr>
        <th>A</th>
        <th>B</th>
        <th>C</th>
    </tr>
    <tr>
        <td>a</td>
        <td>b</td>
        <td>c</td>
    </tr>
    <tr>
        <td>α</td>
        <td>β</td>
        <td>γ</td>
    </tr>
</table>  

#### tr
```html
<tr>...<tr>
```
このタグで表の行を一行挿入できます。

#### th
```html
<th>...<th>
```
このタグで表の見出しを定義できます。

#### td
```html
<td>...<td>
```
このタグで表の要素を作成できます。

#### 結合
更に、次のように記述することで表の要素を結合できたりもします。

```html
<table>
    <tr>
        <th>タイトル1</th>
        <th>タイトル2</th>
        <th>タイトル3</th>
    </tr>
    <tr>
        <td colspan="2">列結合</td>
        <td rowspan="2">行結合</td>
    </tr>
    <tr>
        <td>A</td>
        <td>B</td>
    </tr>
</table>

```
<table>
    <tr>
        <th>タイトル1</th>
        <th>タイトル2</th>
        <th>タイトル3</th>
    </tr>
    <tr>
        <td colspan="2">列結合</td>
        <td rowspan="2">行結合</td>
    </tr>
    <tr>
        <td>A</td>
        <td>B</td>
    </tr>
</table>

### 画像・imgタグ
htmlでは次のように画像を埋めこむことが出来ます。

```html
<img src="https://www.kmc.gr.jp/~tkmax777/data/webservice/html/sample.jpg">
```
**実行例**<br>
<img src="https://www.kmc.gr.jp/~tkmax777/data/webservice/html/sample.jpg">

imgタグのsrc(source)属性に該当する画像ファイルのURLを指定することで、その画像を挿入できます。

更に次のように、title属性で画像のタイトルや、alt属性で代替テキストを指定しておくと良いでしょう。ついでに、このようにwidth属性や、height属性によって大きさの指定もできます。

```html
<img src="https://www.kmc.gr.jp/~tkmax777/data/webservice/html/sample.jpg" alt="がっこうぐらし!" title="がっこうぐらし!" width="150px">
```
**実行例**<br>
<img src="https://www.kmc.gr.jp/~tkmax777/data/webservice/html/sample.jpg" alt="がっこうぐらし!" title="がっこうぐらし!" width="150px">

**参考**<br>
画像にキャプションなど付けたいときはこのようにします。

```html
<figure>
<legend>がっこうぐらし!</legend>
<img src="https://www.kmc.gr.jp/~tkmax777/data/webservice/html/sample.jpg" alt="がっこうぐらし!" title="がっこうぐらし!" width="150px">
</figure>
```

<figure>
<legend>がっこうぐらし!</legend>
<img src="https://www.kmc.gr.jp/~tkmax777/data/webservice/html/sample.jpg" alt="がっこうぐらし!" title="がっこうぐらし!" width="150px">
</figure>

もしくは

```html
<figure>
<img src="https://www.kmc.gr.jp/~tkmax777/data/webservice/html/sample.jpg" alt="がっこうぐらし!" title="がっこうぐらし!" width="150px">
<figcaption>がっこうぐらし!</figcaption>
</figure>
```

<figure>
<img src="https://www.kmc.gr.jp/~tkmax777/data/webservice/html/sample.jpg" alt="がっこうぐらし!" title="がっこうぐらし!" width="150px">
<figcaption>がっこうぐらし!</figcaption>
</figure>

### 音声・audioタグ
#### 基本形
HTML5で次のように、音声の挿入が出来るようになりました。
```html
<audio src="https://static.kmc.gr.jp/projects/webservice2020/5th/sample.mp3" controls>
<p>このブラウザでは音楽再生をサポートしていません。</p>
</audio>

```
**実行例**<br>
<body>
<audio src="https://static.kmc.gr.jp/projects/webservice2020/5th/sample.mp3" controls>
<p>このブラウザでは音楽再生をサポートしていません。</p>
</audio>
</body>

順番に解説していきます。

`audio`タグでは、`img`タグと同様にsrc属性で音源を選択したあと、一般的なプレーヤーを表示するために、`controls`属性を指定する必要があります。

```html
<audio src="https://static.kmc.gr.jp/projects/webservice2020/5th/sample.mp3" controls></audio>
```
**実行例**<br>
<audio src="https://static.kmc.gr.jp/projects/webservice2020/5th/sample.mp3" controls></audio>


これで音声の挿入ができました。しかし、ブラウザのバージョンによっては`audio`タグに対応していない可能性があります。では一般に非対応のタグの内容がどのように処理されるかというと、そのタグの内容は無視され、その要素をそのまま文章として出力されるようになっています。そこで、この性質を利用し、最初のように`audio`要素に未対応の旨を記述しておくとユーザに通知できてよりベターでしょう。


#### 応用

ブラウザによっては、対応していない音声フォーマットがあります。この差埋めるために、`source`タグで複数の形式を用意しておくことが可能です。

```html
<body>
<audio controls>
<source src="https://static.kmc.gr.jp/projects/webservice2020/5th/sample.mp3">
<source src="https://static.kmc.gr.jp/projects/webservice2020/5th/sample.wav">
<p>このブラウザは非対応です。</p>
</audio>
</body>
```
**実行例**<br>
<body>
<audio controls>
<source src="https://static.kmc.gr.jp/projects/webservice2020/5th/sample.wav">
<source src="https://static.kmc.gr.jp/projects/webservice2020/5th/sample.mp3">
<p>このブラウザは非対応です。</p>
</audio>
</body>

また、次の属性をaudioタグに付加することで、細かい設定が出来ます。

| 属性 | 効果 |
| ---- | ---- |
|preload属性|音声の読み込み時期を変更できます。（auto、metadata、none）|
|loop|繰り返し再生を行います。|

この他、autoplay属性も存在しますが、今日多くのブラウザではこの機能が無視されるようになっているので、あまりつかえません。

### 動画・videoタグ
#### 基本形
画像や音声の挿入の仕方を複合した感じで挿入できます。

```html
<body>
<video src="https://static.kmc.gr.jp/projects/webservice2020/5th/sample.mp4" width="300px" controls>
<p>このブラウザでは対応していません。</p>
</video>
</body>
```
**実行例**<br>
<body>
<video src="https://static.kmc.gr.jp/projects/webservice2020/5th/sample.mp4" width="300px" controls>
<p>このブラウザでは対応していません。</p>
</video>
</body>

#### 応用
音声と同じく複数の形式を用意することが可能です。
```html
<body>
<video  width="300px" controls>
<source src="https://static.kmc.gr.jp/projects/webservice2020/5th/sample.mp4">
<source src="https://static.kmc.gr.jp/projects/webservice2020/5th/sample.ogv">
<p>このブラウザでは対応していません。</p>
</video>
</body>
```
**実行例**<br>
<body>
<video  width="300px" controls>
<source src="https://static.kmc.gr.jp/projects/webservice2020/5th/sample.mp4">
<source src="https://static.kmc.gr.jp/projects/webservice2020/5th/sample.ogv">
<p>このブラウザでは対応していません。</p>
</video>
</body>

`audio`タグと同様に、読み込み設定などの細かい設定が出来ます。


| 属性 | 効果 |
| ---- | ---- |
|preload|動画の読み込み時期を変更できます。（auto、metadata、none）|
|loop|繰り返し再生を行います。|
|muted|初期状態を消音します。|
|autoplay|自動再生を有効にします。但し、`muted`属性が必要です。|
|playsinline|スマホなどで全画面表示をオフにします。<br>但し、`autoplay`, `muted`属性が同時に必要です。|
|poster|動画のサムネイルになる画像をURLで指定できます。|

詳細は
[ココ](https://webliker.info/5250/)
のページがわかりやすく解説してくれているので、興味があれば是非一読あれ。

### フォーム
#### 基本形
```html
<form action="https://script.google.com/macros/s/AKfycbzqmk6foQdJ1RicRXr4nRk6Fsk8da34yut_4QRVtRqHiqeG7ZXf/exec"  method="post">
<p>
    <label>
        入力:
        <input type="text" name="text">
    </label>
</p>
<input type="submit" value="送信">
</form>
```
**実行例**<br>
<form action="https://script.google.com/macros/s/AKfycbzqmk6foQdJ1RicRXr4nRk6Fsk8da34yut_4QRVtRqHiqeG7ZXf/exec"  method="post">
<p>
    <label>
        入力:
        <input type="text" name="text">
    </label>
</p>
<input type="submit" value="送信">
</form>