# Discord: Wikipedia Random Bot
## 概要
　Wikipediaからランダムに記事をとってきて、それをDiscordに送信する。
## 使い方
### 準備
　discordgo の導入
```sh
$ go get github.com/bwmarrin/discordgo
```
　credit.jsonの準備
credit.jsonを作成し、
```json
{
    "discord": "Bot TOKEN"
}
```
### 動作
- こんな感じの挙動
![こんな感じ](/1st/discord_standard/resources/program.png)