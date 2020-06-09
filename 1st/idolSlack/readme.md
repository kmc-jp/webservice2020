# デレマスBot

## 概要
デレマスのキャラクターの詳細を表示するBot。ランダム or 指名によってキャラを選択する。

## 準備

1. Slack用の外部パッケージを入れる。

```sh
$ go get -u github.com/nlopes/slack
```

2. SlackAppの作成
   1. [Slack Appの作成](https://api.slack.com/apps?new_classic_app=1)に飛ぶ
   2. 次の図を参考に入力
   ![CreateNewSlackApp](resources/CreateNewSlackApp.png)
   3. Bots選択
   ![Bots](resources/Bots.png)
   4. Add Legacy Bot Userを選択
   ![AddLegacyBotUser](resources/AddLegacyBotUser.png)
   5. 必要事項を記入してAdd
   - 今回Display Nameは`Webservice2020_YOUR_NAME`にしましょう。
   6. `OAuth & Permissions`から`Install App to Workspace`選択
   ![InstallBot](resources/InstallBot.png)
   7. 内容を確認したら`許可する`押しましょう。
   8. `OAuth & Permissions`に戻ると表示される、`xoxb-`から始まる`Bot User OAuth Access Token`を後に使います。

## 仕組み
### init関数
Golangではinit関数→main関数の順に実行される。

**補足-import-**

importするときにはこんな書き方もできる。

```go
import (
   f "fmt"
)
```
これをすると、パッケージ関数を実行するとき、

```go
fmt.Printf("Hello World!\n")
```
とかく所を

```go
f.Printf("Hello World!\n")
```
と書くことができる。これにより、書く手間を省くことができる他、`fmt`という名前のパッケージ名が被った場合、名前を変更できるというメリットもある。

またその他、

```go
import (
   _ "fmt"
)
```
と書くことで、importしても名前を与えないということができる。ここでポイントとなるのがinit関数なのです。

packageは引用されたとき、init関数が最初に実行されるため、名前が与えられていないpackageであってもinit関数だけは実行されます。

今回のプログラムでは、init関数で
[このページ](https://imascg-slstage-wiki.gamerch.com/%E3%82%A2%E3%82%A4%E3%83%89%E3%83%AB%E4%B8%80%E8%A6%A7)
からデレマスキャラの詳細を取得しています。

