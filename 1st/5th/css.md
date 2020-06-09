[HTML](./html.md)
[javaScript](./javascript.md)
# CSS
とりあえずHTMLの作成はできるようになりましたが、このままではデザインが素朴(ってかダサい)。というわけで、最低限、色付けたり出来るようになりましょう！

## 目次
<!-- TOC -->

- [CSS](#css)
    - [目次](#目次)
        - [基本](#基本)
    - [各要素の名前](#各要素の名前)
    - [セレクタ](#セレクタ)
    - [プロパティ](#プロパティ)
    - [色々な指定の仕方](#色々な指定の仕方)
        - [タグに属性付加](#タグに属性付加)
        - [別ファイルでの指定](#別ファイルでの指定)
    - [外部CSSの利用](#外部cssの利用)

<!-- /TOC -->
### 基本
```html
<!DOCTYPE html>
<html lang="ja">

<head>
    <meta charset="UTF-8">
    <link rel="author" href="mailto:someone@example.jp">
    <title lang="jp">WebService2020</title>
        <style>
            h1 {
                font-family: "ＭＳ 明朝", serif;
                color: #616161;
            }
        </style>
</head>
<body>
    <h1>KMCでの暮らし</h1>
    <p>KMCにはmatuという優秀な会長が居ます。わからないことがあれば、すべて会長に投げましょう。</p>
</body>
</html>
```
・<a href="./cssSample/css1.html">
実行例
</a>

<iframe src="./cssSample/css1.html" name="sample" width="90%" height="200">
    <a href="./cssSample/css1.html"></a>
</iframe>


CSSはこのようにhead内にStyleタグの要素として記述していくことで記述することが出来ます。では、ひとつずつ要素を眺めていきましょう。

## 各要素の名前
CSSの文法はおおまかに次の塊の集合であらわされる。
```css
セレクタ {
    プロパティ: 値;
}
```

セレクタで選択された対象のデザイン設定をプロパティとよばれる設定項目に値を付けていくことで、上書きしていきます。

## セレクタ
次のようなセレクタがあります。
[参考](http://www.htmq.com/csskihon/005.shtml)

<table>
    <tr>
        <th>種別</th><th>セレクタ</th><th>選択される要素</th>
    </tr>
    <tr>
        <td>要素型セレクタ</td><td>タグ</td><td>タグの要素</td>
    </tr>
    <tr>
        <td>全称セレクタ</td><td>*</td><td>すべての要素</td>
    </tr>
    <tr>
        <td>classセレクタ</td><td>要素名.クラス名</td><td>該当するクラスに属する要素</td>
    </tr>
    <tr>
        <td>idセレクタ</td><td>要素名#id名</td><td>該当する要素</td>
    </tr>
    <tr>
        <td rowspan="7">疑似クラス</td><td>要素名:link</td><td>未訪問のリンク</td>
    </tr>
    <tr>
        <td>要素名:visited</td><td>訪問済のリンク</td>
    </tr>
    <tr>
        <td>要素名:hover</td><td>カーソルが乗っている要素</td>
    </tr>
    <tr>
        <td>要素名:active</td><td>クリック中の要素</td>
    </tr>
    <tr>
        <td>要素名:focus</td><td>フォーカスされている要素</td>
    </tr>
    <tr>
        <td>要素名:lang</td><td>特定の言語を指定した要素</td>
    </tr>
    <tr>
        <td>要素名:first-child</td><td>要素内の最初の子要素</td>
    </tr>
    <tr>
        <td rowspan="4">擬似要素</td><td>要素名:first-line</td><td>要素の最初の一行</td>
    </tr>
    <tr>
        <td>要素名:first-letter</td><td>要素の最初の一文字</td>
    </tr>
    <tr>
        <td>要素名:before</td><td>要素の直前</td>
    </tr>
    <tr>
        <td>素名:after</td><td>要素の直後</td>
    </tr>
</table>

無理に覚える必要はありません。設定したくなったらこの表を見ましょう。

## プロパティ

主要なもの。全部は<a href="http://www.htmq.com/style/">このページ</a>が詳しくかいてあります。これもまた、無理に覚える必要はありません。必要になったら調べましょう。

<table>
    <tr>
        <th>ジャンル</th><th>プロパティ</th><th>値</th><th>効果</th>
    </tr>
    <tr>
        <td rowspan="6">色・画像</td><td>color</td><td>代表値, #RGB記法</td><td>文字色変更</td>
    </tr>
    <tr>
        <td>background-color</td><td>代表値, #RGB記法</td><td>背景色を指定する</td>
    </tr>
    <tr>
        <td>background-image</td><td>url(URL)</td><td>背景色を指定する</td>
    </tr>
    <tr>
        <td>background-attachment</td><td>fixed, scroll</td><td>背景画像の固定・非固定を指定</td>
    </tr>
    <tr>
        <td>background-position</td><td><a href="#position">位置の名前</a>, <br>左上からの%, <br>ピクセル</td><td>位置指定</td>
    </tr>
    <tr>
        <td>background-repeat</td><td>repeat repeat-x repeat-y no-repeat</td><td>背景の繰り返し</td>
    </tr>
    <tr>
        <td rowspan="5">フォント</td><td>font-style</td><td>normal italic oblique</td><td>標準・イタリック体・斜体を切りかえる</td>
    </tr>
    <tr>
        <td>font-variant</td><td>normal small-caps</td><td>小文字大文字を切り替える</td>
    </tr>
    <tr>
        <td>font-size</td><td>ピクセル, %<br>xx-small, x-small, small<br>medium, large, x-large, xx-large</td><td>文字の太さ指定</td>
    </tr>
    <tr>
        <td>font-size-adjust</td><td>none, 数値(相対比率)<br>inherit(上層に合わせる)</td><td>フォント間のサイズ差を自動調整する</td>
    </tr>
    <tr>
        <td>font-variant</td><td>normal small-caps</td><td>小文字大文字を切り替える</td>
    </tr>
    <tr>
        <td rowspan="6">テキスト整形</td><td>line-height</td><td>normal, 数値(比率 or 単位付き), %</td><td>行幅を指定</td>
    </tr>
    <tr>
        <td>text-align</td><td>start, end, left, right<br>center, justify, match-parent</td><td>枠内での文字列の振り分け<br>justify選択時にはtext-justifyプロパティで更に詳しく設定できる。</td>
    </tr>
    <tr>
        <td>text-justify</td><td>auto none inter-word inter-character</td><td>自動 / 無効 / 単語による調整 / 文字による調整</td>
    </tr>
    <tr>
        <td>vertical-align</td><td>baseline, top, middle, bottom<br>text-top, text-bottom, super, sub<br>% 数値(単位付き)</td><td>文字列の配置を指定できる。<a href="http://www.htmq.com/style/vertical-align.shtml">詳細はココ参照</a></td>
    </tr>
    <tr>
        <td>text-underline-position</td><td>auto, under, left, right</td><td>傍線の位置を指定。<br>但し、left, rightは縦書き用</td>
    </tr>
    <tr>
        <td>text-indent</td><td>%, 数値, each-line, hanging</td><td>pタグに適応することで、一行目のインデント幅を調整できる。<br>但し、hangingを指定するとこれが反転する。
        </td>
    </tr>
    <tr>
        <td rowspan="6">大きさ</td><td>width</td><td>auto %, 数値(単位付き)</td><td>幅を指定する。</td>
    </tr>
    <tr>
        <td>max-width</td><td>auto %, 数値(単位付き)</td><td>幅の最大値を指定</td>
    </tr>
    <tr>
        <td>min-width</td><td>auto %, 数値(単位付き)</td><td>幅の最小値を指定</td>
    </tr>
    <tr>
        <td>height</td><td>auto %, 数値(単位付き)</td><td>高さを指定する。</td>
    </tr>
    <tr>
        <td>max-height</td><td>auto %, 数値(単位付き)</td><td>高さの最大値を指定</td>
    </tr>
    <tr>
        <td>min-height</td><td>auto %, 数値(単位付き)</td><td>高さの最小値を指定</td>
    </tr>
</table>

<a id="position">・位置の名前</a><br>
<table>
    <tr>
        <th>軸</th><th>名前</th><th>場所</th><th>軸</th><th>名前</th><th>場所</th>
    </tr>
    <tr>
        <td rowspan="3">X</td><td>left</td><td>左</td><td rowspan="3">Y</td><td>top</td><td>上</td>
    </tr>
    <tr>
        <td>center</td><td>中央</td><td>center</td><td>中央</td>
    </tr>
    <tr>
        <td>right</td><td>右</td><td>bottom</td><td>下</td>
    </tr>
</table>

<a id="weight">・太さの名前</a><br>

<table>
    <tr>
        <th>名前</th><th>対応</th>
    </tr>
    <tr>
        <td>normal</td><td>400</td>
    </tr>
    <tr>
        <td>bold </td><td>700</td>
    </tr>
    <tr>
        <td>lighter</td><td>現行より一段階細く</td>
    </tr>
    <tr>
        <td>bolder </td><td>現行より一段階太く</td>
    </tr>
</table>

## 色々な指定の仕方
冒頭のように、HTMLのヘッダ内に記述していってもいいのですが、この方法以外にも次の方法があります。

### タグに属性付加
各タグにたいして、単純にstyleタグを設定するだけでも設定できます。

・HTML

```html

<p><span style="font-weight: bold;">このように</span>ちょっとしたスタイルの挿入に便利です。</p>

```

<p><span style="font-weight: bold;">このように</span>ちょっとしたスタイルの挿入に便利です。</p>

### 別ファイルでの指定
今までの方法の他に、次のように、HTMLのhead要素内にlinkタグで紐付けることも出来ます。この方法で記述されることが１番主流です。

・HTML

```html
<!DOCTYPE html>
<html>

<head>
    <meta charset="UTF-8">
    <link rel="stylesheet" href="style.css" type="text/css">
    <title lang="jp">WebService2020</title>
</head>

<body>
    <div class="text">
        <div class="head">
            <p>実験</p>
        </div>
        <div class="body">
            <p>こんな感じになります。</p>
        </div>
    </div>
</body>

</html>
```
・CSS(style.css)

```css
div.text div.head{
    background-color: rgb(228, 228, 228);
    
    border-radius: 5px;

    font-weight: bold;

    height: 50px;

    padding-left: 2em;
    padding-right: 2em;
    padding-top: 5px;
    padding-bottom: 5px;
}

div.text div.body{
    border: 3px solid ;
    border-color: rgb(228, 228, 228);
    border-radius: 5px;
    border-top: transparent;


    background-color: whitesmoke;
    padding-left: 2em;
    padding-right: 2em;
    padding-top: 1em;
    padding-bottom: 1em;
}
```

・<a href="./cssSample/fileSample/index.html">実行結果<br></a>
<iframe src="./cssSample/fileSample/index.html" name="sample" width="90%" height="200">
    <a href="./cssSample/fileSample/index.html"></a>
</iframe>

良い感じですね！

## 外部CSSの利用
さて、ここまでざっと軽くCSSの記述の仕方を見てきました。


```html
<!DOCTYPE html>
<html lang="ja">
    <head>
        <meta charset="UTF-8">
        <link rel="author" href="mailto:someone@example.jp">
        <title lang="jp">WebService2020</title> 
        <link href="https://fonts.googleapis.com/css2?family=M+PLUS+1p:wght@500;700&display=swap" rel="stylesheet">
        <style>
            .sample {                
                font-family: 'M PLUS 1p', sans-serif;
            }
            h1.sample {
                font-weight: 700;
            }
            p.sample {
                font-weight: 500; 
            }
        </style>
    </head>
    <body>
        <h1 class="sample">KMCでの暮らし</h1>
        <p class="sample">KMCにはmatuという優秀な会長が居ます。わからないことがあれば、すべて会長に投げましょう。</p>
    </body>
</html>
```