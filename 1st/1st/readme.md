# 第一回
## 概要
- 開発に必要な各種ツールを導入してもらう。

## 手順
### Golang導入
世界線の移動があります。
#### macOS
1. command + Space
2. terminalと入力
3. `$ /usr/bin/ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"`
4. `$ brew install go`

#### Windows
1. [ここ](https://golang.org/dl/)から自分の環境に合うものをDownload
2. 良しなにinstall(環境を合わせるため、C\に入れてもらう)
3. PATH(環境変数)設定
4. システムのプロパティ[詳細設定]
5. 環境変数(N)...
   1. GOPATHの登録
      - %USERPROFILE%\goで良いんじゃないですかね。
   2. PATHの登録 くれぐれも慎重に！！
      - ;C:\Users\YOUR_ACCOUNT_NAME\go\binを追記
   3. GOROOTの登録
      - C:\Go

### gitの導入
#### macOS
1. terminalを起動
2. `$ git --version`
3. 何か出た貴方は良い感じにやってください
4. `git version 2.5.●`とか出た選ばれし貴方は最新のバージョンを入れていただきます
   - [ここ](https://git-scm.com/download/mac)からダウンロードしてね。
   - その後いい感じにインストールしてください。

#### Windows
1. [ここ](https://git-scm.com/download/win)からダウンロードしてね
2. いい感じにインストールする。（スライド参照）

### 導入確認
- terminal / GitBashでgo --version を確認。

### VSCodeの導入
[ここ](https://azure.microsoft.com/ja-jp/products/visual-studio-code/)からダウンロード & インストール
