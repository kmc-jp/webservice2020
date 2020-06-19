# デレマスBot
## もくじ
<!-- TOC -->

- [デレマスBot](#デレマスbot)
    - [もくじ](#もくじ)
    - [概要](#概要)
    - [準備](#準備)
    - [httpについて](#httpについて)
    - [jsonについて](#jsonについて)
    - [go routineについて](#go-routineについて)
    - [仕組み](#仕組み)
        - [main package](#main-package)
            - [init関数](#init関数)
            - [main関数](#main関数)
            - [ListenTo関数](#listento関数)
            - [NameCheck関数](#namecheck関数)
            - [Random関数](#random関数)
        - [idol package](#idol-package)
            - [Get関数](#get関数)
            - [MakeDict関数](#makedict関数)
            - [Read関数](#read関数)

<!-- /TOC -->

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
2. main.goのYOUR_TOKENにさっきのを記述

## httpについて

[学んで行きましょう！](./http.md)

## jsonについて

[学んで行きましょう！](./json.md)

## go routineについて

[学んで行きましょう！](./go.md)

## 仕組み
### main package
#### init関数
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

```go
func init() {
	DictPath = filepath.Join("..", "data.json")//1
	var err error = idol.MakeDict(DictPath)//2
	if err != nil {
		panic(err.Error())
	}
}
```

1. アイドル情報を格納するファイルを指定している。
2. ここで時デレマスのデータを取得している。

#### main関数
```go
func main(){
	var api *slack.Client = slack.New(BotToken)//1

	RTM = api.NewRTM()//2

	go RTM.ManageConnection()//3

	for msg := range RTM.IncomingEvents {
		switch ev := msg.Data.(type) {
		case *slack.ConnectedEvent:
			fmt.Printf("Start connection with Slack\n")//4
		case *slack.MessageEvent:
			EV = ev
			ListenTo()//5
		}
    }
}
```

1. Botトークンをもとに、Slackとの通信をするための初期化をする
2. Real Time Messangerを作成。これで各種Slackとの通信を行う。
3. 通信を開始する。
4. Slackとの通信を開始したらメッセージを表示
5. メッセージが送信されたらListenTo関数を実行する。

#### ListenTo関数

```go
//ListenTo Manage incomming messages
func ListenTo() {
	switch {
	case strings.Contains(EV.Text, "ちゃん"): // 1
		NameCheck()
	case strings.Contains(EV.Text, "デレマスランダム"):// 2
		Random()
	}
}

```

1. メッセージに"ちゃん"という文字が含まれていれば`NameCheck`関数を実行する
2. メッセージにデレマスランダムという文字が含まれていれば`NameCheck`関数を実行する

#### NameCheck関数

```go
//NameCheck find the idol from dict
func NameCheck() {
	var Info []idol.Idol = idol.Read() //1
	var NAME string = strings.Split(EV.Text, "ちゃん")[0]//2
	var FoundNUM int

	if NAME == "" {//3
		return
	}
	
	for i, info := range Info {
		if info.Name == NAME {//4
			FoundNUM = i
			goto Found//5
		}
	}
	return

Found:
	var SendText string = MakeText(Info[FoundNUM])//6

	RTM.SendMessage(RTM.NewOutgoingMessage(SendText, EV.Channel))//7
}
```

1. アイドル情報を取得
2. "ちゃん"という文字の前にある文字列(=名前)を取得
3. 名前が空白のときは検索しない
4. 一人ずつ一致する人を探す
5. 見つかればFoundに飛ぶ
6. アイドル情報を表示する文を作成
7. Slackに送信

#### Random関数

```go
//Random chose one Idol
func Random() {
	var Info []idol.Idol = idol.Read()
	rand.Seed(time.Now().UnixNano())//1

	var SendText string = MakeText(Info[int(rand.Float64()*float64(len(Info)))])//2
	RTM.SendMessage(RTM.NewOutgoingMessage(SendText, EV.Channel))
}
```

1. 乱数の初期化
2. ランダムに一人アイドルを選択し、文章を生成

### idol package
#### Get関数

```go
//Get gets idol info from web
func Get() []Idol {
	resp, err := http.Get("https://imascg-slstage-wiki.gamerch.com/%E3%82%A2%E3%82%A4%E3%83%89%E3%83%AB%E4%B8%80%E8%A6%A7")//1
	if err != nil {
		return []Idol{}
	}
	defer resp.Body.Close()

	var buf *bytes.Buffer = new(bytes.Buffer)//2
	io.Copy(buf, resp.Body)//3
	var ret []byte = buf.Bytes()//4

	var top []byte = bytes.Split(ret, []byte("id=\"content_1_2\""))[1]//5
	var content string = string(bytes.Split(bytes.Split(bytes.Split(
		top,
		[]byte("id=\"content_1_2\""))[0],
		[]byte("t-line-img word-keep-all"))[1],
		[]byte("<tbody>"))[1],
	)//6

	var table []Idol

	for i, x := range strings.Split(content, "<tr>") {
		if i == 0 {
			continue
		}

		var idol Idol

		for _, y := range strings.Split(x, "<td") {
			switch {
			case strings.Contains(y, "data-col=\"1\">"):
				...
			case strings.Contains(y, "data-col=\"2\">")://7
				idol.Age = strings.Split(strings.Split(y, "data-col=\"2\">")[1], "</")[0]//8
			case strings.Contains(y, "data-col=\"3\">"):
				...
			}
		}
		table = append(table, idol)

	}
	fmt.Printf("%#v\n", table)
	return table
}
```

1. 
[ここ](https://imascg-slstage-wiki.gamerch.com/%E3%82%A2%E3%82%A4%E3%83%89%E3%83%AB%E4%B8%80%E8%A6%A7)
にGetリクエストを送る(htmlが返ってくる)

2. bytesパッケージのBuffer型(ポインタ)を用意する。
3. レスポンスの内容をここにコピーする。
4. byteとして出力
5. (6も)表の部分を抽出
7. 表の要素ごとに分岐
8. 要素ごとに構造体に格納

#### MakeDict関数

```go
//MakeDict make dictionary file
func MakeDict(Path string) error {
	var x []Idol = Get() //1
	b, err := json.MarshalIndent(x, "", "    ") //2
	if err != nil {
		return fmt.Errorf("Marshal Error")
	}
	err = ioutil.WriteFile(Path, b, os.ModePerm)//3
	if err != nil {
		return fmt.Errorf("Write Error")
	}
	return nil
}
```

1. xにアイドル情報のスライスを格納
2. json形式に変換(byteで出力)
3. この内容をそのままPathに場所に出力

#### Read関数

```go
//Read read idol data
func Read(Path string) []Idol {
	b, err := ioutil.ReadFile(Path))//1
	if err != nil {
		fmt.Printf("Error: Cannot Read file\n")
		return []Idol{}
	}

	var info []Idol

	json.Unmarshal(b, &info)//2
	return info
}
```

1. Pathにあるファイルを読みこむ
2. jsonから構造体に入れ込む

