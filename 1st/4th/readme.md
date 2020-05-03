# 第四回
## 概要
今までで学習したGolangの使い方を踏まえてSlackBotを作ってみる。

## 準備
1. まずはSackBot用のgoのpackageを入れる。
```sh
$ go get -u github.com/slack-go/slack
```
（`-u`:update、既に入っている場合アップデートする。）

2. もしも入会していなければ
[ここ](https://slack.com/intl/ja-jp/get-started#/create)でワークスペース作ってください...。

**（もしくは入会）**
3. SlackAppの作成
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
