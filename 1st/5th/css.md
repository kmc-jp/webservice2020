[HTML](./html.md)
[JavaScript](./javascript.md)
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

<body>
    <h1>KMCでの暮らし</h1>
    <p>KMCにはmatuという優秀な会長が居ます。わからないことがあれば、すべて会長に投げましょう。</p>
</body>
</html>
```
**実行例**

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

<table>
    <tr>
        <th>ジャンル</th><th>プロパティ</th><th>値</th><th>効果</th>
    </tr>
    <tr>
        <td rowspan="11">色・画像</td><td>color</td><td>代表値, #RGB記法</td><td>文字色変更</td>
    <tr>
    <tr>
        <td>background-color</td><td>代表値, #RGB記法</td><td>背景色を指定する</td>
    <tr>
    <tr>
        <td>background-image</td><td>url(URL)</td><td>背景色を指定する</td>
    <tr>
    <tr>
        <td>background-attachment</td><td>fixed, scroll</td><td>背景画像の固定・非固定を指定</td>
    <tr>
    <tr>
        <td>background-position</td><td><a href="#position">位置の名前</a>, <br>左上からの%, <br>ピクセル</td><td>位置指定</td>
    <tr>
    <tr>
        <td>background-repeat</td><td>repeat repeat-x repeat-y no-repeat</td><td>背景の繰り返し</td>
    <tr>
    <tr>
        <td rowspan="10">フォント</td><td>font-style</td><td>normal italic oblique</td><td>標準・イタリック体・斜体を切りかえる</td>
    <tr>
    <tr>
        <td>font-variant</td><td>normal small-caps</td><td>小文字大文字を切り替える</td>
    <tr>
    <tr>
        <td>font-size</td><td>ピクセル, %<br>xx-small, x-small, small<br>medium, large, x-large, xx-large</td><td>文字の太さ指定</td>
    <tr>
    <tr>
        <td>font-size-adjust</td><td>none, 数値(相対比率)<br>inherit(上層に合わせる)</td><td>フォント間のサイズ差を自動調整する</td>
    <tr>
    <tr>
        <td>font-variant</td><td>normal small-caps</td><td>小文字大文字を切り替える</td>
    <tr>
    <tr>
        <td rowspan="10">テキスト整形</td><td>line-height</td><td>normal, 数値(比率 or 単位付き), %</td><td>行幅を指定</td>
    <tr>
    <tr>
        <td>text-align</td><td>start, end, left, right<br>center, justify, match-parent</td><td>枠内での文字列の振り分け<br>justify選択時にはtext-justifyプロパティで更に詳しく設定できる。</td>
    <tr>
    <tr>
        <td>text-justify</td><td>auto none inter-word inter-character</td><td>自動 / 無効 / 単語による調整 / 文字による調整</td>
    <tr>
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