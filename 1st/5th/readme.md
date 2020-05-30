# HTMLやCSS, JavaScriptについて

## 目次
<!-- TOC -->

- [HTMLやCSS, JavaScriptについて](#html%E3%82%84css-javascript%E3%81%AB%E3%81%A4%E3%81%84%E3%81%A6)
    - [目次](#%E7%9B%AE%E6%AC%A1)
    - [概要](#%E6%A6%82%E8%A6%81)
    - [俯瞰](#%E4%BF%AF%E7%9E%B0)
    - [HTML](#html)
        - [基本](#%E5%9F%BA%E6%9C%AC)
        - [表](#%E8%A1%A8)
        - [メディアの表示](#%E3%83%A1%E3%83%87%E3%82%A3%E3%82%A2%E3%81%AE%E8%A1%A8%E7%A4%BA)
            - [画像・imgタグ](#%E7%94%BB%E5%83%8F%E3%83%BBimg%E3%82%BF%E3%82%B0)
            - [動画・videoタグ](#%E5%8B%95%E7%94%BB%E3%83%BBvideo%E3%82%BF%E3%82%B0)

<!-- /TOC -->

## 概要
今回はWebページを作成するための、最低限度の知識を身に付けます。

## 俯瞰
今日、ブラウザを使えば、インターネットを通して様々な画像や、テキストを用いたページを見ることができますね。では、このようなページはどのような形式サーバから送られてくるのでしょうか。

その答えはブラウザの開発者ツールを用いれば見ることができます。次の画像はKMCの外部ページをChromeの開発者ツールを用いて眺めてみた例です。

![開発者ツール](/1st/5th/resources/browserDev.png)

なにやら沢山の記号を用いたスクリプトがでてきました。これは
**HTML(HyperTextMarkupLanguage)**
という基本的に
**文章の構造を指定するため**
の言語で書かれたスクリプトです。ブラウザはこの文章を骨格として、ページの文章や画像を表示していきます。

しかし、これだけでは構造を指定するだけで、要素の配置を細かく指定することはできません[^1]。そこで登場するのが
**CSS(Cascading Style Sheets, 段階スタイルシート)**
です。これを用いることで各々の要素の色、配置などの
**Webページのデザインを指定する**
ことができます。

また、その他、ページの要素を動かしたり、サーバに送るデータの整形をしたりするために、
**JavaScript**
という、プログラミング言語を用いることができます。

以上の三本柱を元に全てのWebページは構成されています。このことを念頭におきつつ、簡単にそれぞれ見ていきましょう。

 
[^1]:正確にはできますが、望ましくありません

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

```html
<tr>...<tr>
```
このタグで表の行を一行挿入できます。

```html
<th>...<th>
```
このタグで表の見出しを定義できます。

```html
<td>...<td>
```
このタグで表の要素を作成できます。

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

### メディアの表示
HTMLの中には音楽、画像、動画を挿入することが出来ます。

#### 画像・imgタグ
```html
<img src="https://www.kmc.gr.jp/~tkmax777/data/webservice/html/sample.jpg">
```
**実行例**<br>
<img src="https://www.kmc.gr.jp/~tkmax777/data/webservice/html/sample.jpg">

imgタグのsrc(source)属性に該当する画像ファイルのURLを指定することで、その画像を挿入できる。

更に次のように、画像のタイトルや、代替テキストを指定しておくと良いでしょう。ついでに、このように大きさの指定もできます。

```html
<img src="https://www.kmc.gr.jp/~tkmax777/data/webservice/html/sample.jpg" alt="がっこうぐらし!" title="がっこうぐらし!" width="100px">
```
**実行例**<br>
<img src="https://www.kmc.gr.jp/~tkmax777/data/webservice/html/sample.jpg" alt="がっこうぐらし!" title="がっこうぐらし!" width="100px">

**参考**<br>
画像にキャプションなど付けたいときはこのようにする。

```html
<figure>
<legend>がっこうぐらし!</legend>
<img src="https://www.kmc.gr.jp/~tkmax777/data/webservice/html/sample.jpg" alt="がっこうぐらし!" title="がっこうぐらし!" width="100px">
</figure>
```

<figure>
<legend>がっこうぐらし!</legend>
<img src="https://www.kmc.gr.jp/~tkmax777/data/webservice/html/sample.jpg" alt="がっこうぐらし!" title="がっこうぐらし!" width="100px">
</figure>

もしくは

```html
<figure>
<img src="https://www.kmc.gr.jp/~tkmax777/data/webservice/html/sample.jpg" alt="がっこうぐらし!" title="がっこうぐらし!" width="100px">
<figcaption>がっこうぐらし!</figcaption>
</figure>
```

<figure>
<img src="https://www.kmc.gr.jp/~tkmax777/data/webservice/html/sample.jpg" alt="がっこうぐらし!" title="がっこうぐらし!" width="100px">
<figcaption>がっこうぐらし!</figcaption>
</figure>

#### 動画・videoタグ