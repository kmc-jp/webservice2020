package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"math"
	"net/http"
	"path/filepath"
	"sync"
	"time"

	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/google"
	"github.com/stretchr/objx"
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
	UserData objx.Map
	Basic    HTMLbasic
}

//Page Handle static page
type Page struct {
	template *template.Template
	filename string
	Basic    HTMLbasic
}

//HTMLbasic Put header and footer data
type HTMLbasic struct {
	Title string
	Login int
	Name  string
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
	gomniauth.SetSecurityKey("HZw?uikhiAol3$XWeiA+XMi&I%daJ?RANDOM")
	gomniauth.WithProviders(
		google.New(credit.ClientID, credit.Secret, "http://localhost:8080/auth/callback/google"),
	)
}

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/auth/", loginHandler)

	http.Handle("/", &Top)
	http.Handle("/login", &Page{filename: "login.html", Basic: HTMLbasic{Title: "入室 -Webservice2020掲示板-"}})
	http.Handle("/form", MustAuth(http.HandlerFunc(AddMessage)))
	http.Handle("/create", MustAuth(&Page{filename: "new_account.html", Basic: HTMLbasic{Title: "会員登録 -Webservice2020掲示板-"}}))

	http.HandleFunc("/createnew", CreateNewAccount)

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
	message := r.Form.Get("message")
	if message == "" {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	Cookie, err := GetCookieValue(r)
	if err != nil {
		http.Error(w, "Cookieの解析に失敗しました。", http.StatusBadRequest)
		log.Printf("Cookieの解析に失敗しました。\n%v\n", err)
		return
	}

	Accounts, err := ReadAccounts()
	if err != nil {
		http.Error(w, "アカウントデータの読み込みに失敗しました。", http.StatusBadRequest)
		log.Printf("アカウントデータの読み込みに失敗しました。\n%v\n", err)
		return
	}

	name := Accounts[Cookie["id"]].Name

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
		top.template = template.Must(t.ParseFiles(filepath.Join("resources", "index.html"), filepath.Join("resources", "header.html"), filepath.Join("resources", "footer.html")))
	})

	top.Basic = HTMLbasic{Title: "Webservice2020掲示板", Login: AuthCheck(r)}
	if top.Basic.Login == 1 {
		Accounts, err := ReadAccounts()
		if err != nil {
			http.Error(w, "データの読み出しに失敗しました。", http.StatusInternalServerError)
			fmt.Printf("データの読み出しに失敗しました。\n%v", err)
			return
		}
		Cookie, err := GetCookieValue(r)
		if err != nil {
			http.Error(w, "Cookieの読み込みに失敗しました。", http.StatusBadRequest)
			fmt.Printf("Cookieの読み込みに失敗しました。\n%v\n", err)
			return
		}
		top.Basic.Name = Accounts[Cookie["id"]].Name
	}

	var messages []Message = make([]Message, cap(Messages))
	copy(messages, Messages)

	switch {
	case DisplayNum > 0:
		top.Message = messages[int(math.Max(0, float64(DisplayNum-len(Messages)))):]
	case DisplayNum <= 0:
		top.Message = messages
	}

	if authCookie, err := r.Cookie("auth"); err == nil {
		top.UserData = objx.MustFromBase64(authCookie.Value)
	}

	//昇順降順の入れ替え
	for left, right := 0, len(top.Message)-1; left < right; left, right = left+1, right-1 {
		top.Message[left], top.Message[right] = top.Message[right], top.Message[left]
	}

	if err := top.template.ExecuteTemplate(w, "content", top); err != nil {
		fmt.Fprintf(w, "%v\n", err)
	}

	return
}

func (p *Page) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//templateを定義
	t := template.New(p.filename)

	//index.htmlなどの内容を読み込み（filepath.Joinは複数の環境に対応できるようにするため。）
	//template.MustはtemplateにErrorがあった場合にpanicを起こすための関数
	p.template = template.Must(t.ParseFiles(filepath.Join("resources", p.filename), filepath.Join("resources", "header.html"), filepath.Join("resources", "footer.html")))

	p.Basic.Login = AuthCheck(r)
	if p.Basic.Login == 1 {
		Accounts, err := ReadAccounts()
		if err != nil {
			http.Error(w, "データの読み出しに失敗しました。", http.StatusInternalServerError)
			fmt.Printf("データの読み出しに失敗しました。\n%v", err)
			return
		}
		Cookie, err := GetCookieValue(r)
		if err != nil {
			http.Error(w, "Cookieの読み込みに失敗しました。", http.StatusBadRequest)
			fmt.Printf("Cookieの読み込みに失敗しました。\n%v\n", err)
			return
		}
		p.Basic.Name = Accounts[Cookie["id"]].Name
	}
	if err := p.template.ExecuteTemplate(w, "content", p); err != nil {
		fmt.Fprintf(w, "%v\n", err)
	}

}
