package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"math"
	"net/http"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/google"
)

//DisplayNum Set the number of messages displayed at top page
// If you put negative value or zero, all messages will displayed
const DisplayNum int = 0

//MessageFilePath Put a path of message.json
var MessageFilePath string

//Top Put all of messages sent to server
var Top TopPage

//Messages Put all Message data
var Messages []Message

//TopPage Put data used at top page
type TopPage struct {
	Message  []Message
	template *template.Template
	once     sync.Once
}

//Page Handle normal page
type Page struct {
	template *template.Template
	once     sync.Once
	filename string
}

//Message Put message data
type Message struct {
	Name string
	Text string
	Time time.Time
}

func init() {
	//message.jsonのファイルパスを指定（環境依存）
	MessageFilePath = filepath.Join("data", "message.json")
	//メッセージデータを読み込み
	ReadMessageData()

	//認証情報を読み込み
	getKeys()

	//OAuthのSetup
	gomniauth.SetSecurityKey(credit.Secret)
	gomniauth.WithProviders(
		google.New(credit.ClientID, "HZw?uikhiAol3$XWeiA+XMi&I%daJ?RANDOM", "http://localhost:8080/auth/callback/google"),
	)
}

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/auth/", loginHandler)
	http.Handle("/", &Top)
	http.Handle("/login", &Page{filename: "login.html"})
	http.Handle("/form", MustAuth(http.HandlerFunc(AddMessage)))

	if err := http.ListenAndServe("0.0.0.0:8080", nil); err != nil {
		log.Fatalln("Error")
	}
}

//ReadMessageData Get messages previously
func ReadMessageData() {
	data, err := ReadFile(MessageFilePath)
	if err != nil {
		fmt.Printf("Read message file error\n%v", err)
		return
	}

	err = json.Unmarshal(data, &Messages)
	if err != nil {
		fmt.Printf("Json unmarshal error\n%v", err)
		return
	}
}

//AddMessage Add message to the message data
func AddMessage(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Printf("Parse form error\n%v", err)
		http.NotFound(w, r)
		return
	}
	name := r.Form.Get("name")
	message := r.Form.Get("message")

	var newMessage Message = Message{
		Text: message,
		Name: name,
		Time: time.Now(),
	}

	Messages = append(Messages, newMessage)

	data, err := json.MarshalIndent(Messages, "", "	  ")
	if err != nil {
		fmt.Printf("Json marshal error\n%v", err)
		http.Error(w, "500 Internal server error", http.StatusInternalServerError)
		return
	}
	ExportFile(MessageFilePath, data)
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	return
}

func (top *TopPage) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//templateのcompile作業は一回だけで良い
	top.once.Do(func() {
		//templateを定義
		t := template.New("index.html")

		//index.htmlの内容を読み込み（filepath.Joinは複数の環境に対応できるようにするため。）
		//template.MustはtemplateにErrorがあった場合にpanicを起こすための関数
		top.template = template.Must(t.ParseFiles(filepath.Join("resources", "index.html")))
	})

	switch {
	case DisplayNum > 0:
		top.Message = Messages[int(math.Max(0, float64(DisplayNum-len(Messages)))):]
	case DisplayNum <= 0:
		top.Message = Messages
	}

	//昇順降順の入れ替え
	for left, right := 0, len(top.Message)-1; left < right; left, right = left+1, right-1 {
		top.Message[left], top.Message[right] = top.Message[right], top.Message[left]
	}

	if err := top.template.Execute(w, top); err != nil {
		fmt.Fprintf(w, "%v\n", err)
	}

	return
}

func (p *Page) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p.once.Do(func() {
		//templateを定義
		t := template.New(p.filename)

		//index.htmlの内容を読み込み（filepath.Joinは複数の環境に対応できるようにするため。）
		//template.MustはtemplateにErrorがあった場合にpanicを起こすための関数
		p.template = template.Must(t.ParseFiles(filepath.Join("resources", p.filename)))
	})
	if err := p.template.Execute(w, p); err != nil {
		fmt.Fprintf(w, "%v\n", err)
	}

}
