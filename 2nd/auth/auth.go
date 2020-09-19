package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strings"
	"sync"

	"github.com/stretchr/gomniauth"
	"github.com/stretchr/objx"
)

//NewAccountPage Handle new account page
type NewAccountPage struct {
	template *template.Template
	once     sync.Once
	Data     OAuthData
	Basic    HTMLbasic
}

//OAuthData Put data got from OAuth
type OAuthData struct {
	nickname string
}

type authHandler struct {
	next http.Handler
}

//AccountList put accounts data
type AccountList struct {
	Name string `json:"name"`
}

func (h *authHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if _, err := r.Cookie("auth"); err == http.ErrNoCookie {
		w.Header().Set("Location", "/login")
		w.WriteHeader(http.StatusTemporaryRedirect)
	} else if err != nil {
		http.Error(w, "Internal Server Error!", http.StatusInternalServerError)
		return
	} else {
		Cookie, err := GetCookieValue(r)
		if err != nil {
			http.Error(w, "Cookieの取得に失敗しました。", http.StatusBadRequest)
			return
		}
		Accounts, err := ReadAccounts()
		if err != nil {
			http.Error(w, "データベースの読み出しに失敗しました。", http.StatusInternalServerError)
			fmt.Printf("データベースの読み出しに失敗しました。\n")
			return
		}
		if _, ok := Accounts[Cookie["id"]]; !ok {
			w.Header().Set("Location", "/login")
			w.WriteHeader(http.StatusTemporaryRedirect)
			return
		}
		h.next.ServeHTTP(w, r)
		return
	}
}

//AuthCheck checks the user's status of loging
//0:not login
//1:login
func AuthCheck(r *http.Request) (status int) {
	if _, err := r.Cookie("auth"); err == nil {
		Cookie, err := GetCookieValue(r)
		if err != nil {
			return 0
		}
		Accounts, err := ReadAccounts()
		if err != nil {
			return 0
		}
		if _, ok := Accounts[Cookie["id"]]; !ok {
			return 0
		}
		return 1
	}
	return 0
}

//MustAuth Add handler to cheking prcess of authorizing
func MustAuth(handler http.Handler) http.Handler {
	return &authHandler{next: handler}
}

//loginHandler Handle login page
// Set cookies to client and complete authorizing process
func loginHandler(w http.ResponseWriter, r *http.Request) {
	//path: /auth/{action}/{provider}

	//Split the string requested path to []string augments by "/"
	var segs []string = strings.Split(r.URL.Path, "/")

	//Check the number of the url augments
	if len(segs) != 4 {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	//Set variables
	var action string = segs[2]
	var provider string = segs[3]

	switch action {
	case "login":
		var loginURL string

		provider, err := gomniauth.Provider(provider)
		if err != nil {
			log.Printf("認証サーバの取得に失敗しました。\n%s\n%v\n", provider, err)
			http.Error(w, "認証サーバの取得に失敗しました。", http.StatusInternalServerError)
			return
		}

		//(nil, nil)の前者はサーバ側（こちら側）の内部状態を示すハッシュ
		//後者は追加設定（追加の権限要求など）を記述するハッシュ
		//今回はどちらも必要ない。
		loginURL, err = provider.GetBeginAuthURL(nil, nil)
		if err != nil {
			log.Printf("GetBeginAuthURLの呼び出しに失敗しました。\n")
			http.Error(w, "認証エラー", http.StatusInternalServerError)
		}
		w.Header().Set("Location", loginURL)
		w.WriteHeader(http.StatusTemporaryRedirect)
	case "callback":
		provider, err := gomniauth.Provider(provider)
		if err != nil {
			log.Printf("認証サーバの取得に失敗しました。\n%s\n%v\n", provider, err)
			http.Error(w, "認証サーバの取得に失敗しました。", http.StatusInternalServerError)
			return
		}
		creds, err := provider.CompleteAuth(objx.MustFromURLQuery(r.URL.RawQuery))
		if err != nil {
			log.Printf("認証の完了に失敗しました。\n%s\n%v\n", provider, err)
			http.Error(w, "認証の完了に失敗しました。", http.StatusInternalServerError)
			return
		}
		user, err := provider.GetUser(creds)
		if err != nil {
			log.Printf("ユーザデータの取得に失敗しました。\n%s\n%v\n", provider, err)
			http.Error(w, "ユーザデータの取得に失敗しました。", http.StatusInternalServerError)
			return
		}
		authCookieValue := objx.New(map[string]interface{}{
			"id": user.Data()["id"].(string),
		}).MustBase64()

		Accounts, err := ReadAccounts()
		if err != nil {
			log.Printf("データベースの読み取りに失敗しました。\n%s\n%v\n", provider, err)
			http.Error(w, "データベースの読み取りに失敗しました。", http.StatusInternalServerError)
			return
		}

		if _, ok := Accounts[user.Data()["id"].(string)]; ok {
			http.SetCookie(w, &http.Cookie{
				Name:  "auth",
				Value: authCookieValue,
				Path:  "/",
			})
			w.Header()["Location"] = []string{"/"}
			w.WriteHeader(http.StatusTemporaryRedirect)
		} else {
			http.SetCookie(w, &http.Cookie{
				Name:  "auth",
				Value: authCookieValue,
				Path:  "/",
			})

			Accounts[user.Data()["id"].(string)] = AccountList{Name: user.Data()["name"].(string)}

			err := ExportAccounts(Accounts)
			if err != nil {
				log.Printf("データベースの書き出しに失敗しました。\n%s\n", err)
				http.Error(w, "データベースの書き出しに失敗しました。", http.StatusInternalServerError)
				return
			}

			w.Header().Set("Location", "/create")
			w.WriteHeader(http.StatusTemporaryRedirect)
		}
	default:
		http.Error(w, fmt.Sprintf("操作：%sには対応していません。\n", action), http.StatusNotFound)
		return
	}
	return
}

func logoutHandle(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("auth") // クライアントから送信されたCookieを格納
	if err == http.ErrNoCookie {
		w.Header().Set("Location", "/create")
		w.WriteHeader(http.StatusTemporaryRedirect)
	}
	c.MaxAge = -1        // 格納した変数cのMaxAgeフィールドに-1を指定
	http.SetCookie(w, c) // 変更を反映するためにcをCookieにセット
	w.Header().Set("Location", "/")
	w.WriteHeader(http.StatusTemporaryRedirect)
}

//CreateNewAccount Create new account
func CreateNewAccount(w http.ResponseWriter, r *http.Request) {
	Accounts, err := ReadAccounts()
	if err != nil {
		http.Error(w, "データベースの参照に失敗しました。\n", http.StatusInternalServerError)
		log.Printf("データベースの参照に失敗しました。\n%v\n", err)
		return
	}

	CookieData, err := GetCookieValue(r)
	if err != nil {
		http.Error(w, "Cookieの取得に失敗しました。", http.StatusBadRequest)
		return
	}

	err = r.ParseForm()
	if err != nil {
		fmt.Printf("Parse form error\n%v\n", err)
		http.NotFound(w, r)
		return
	}

	nickname := r.Form.Get("nickname")
	if nickname == "" {
		nickname = Accounts[CookieData["id"]].Name
	}

	Accounts[CookieData["id"]] = AccountList{Name: nickname}

	if err = ExportAccounts(Accounts); err != nil {
		http.Error(w, "データの書き出しに失敗しました。\n", http.StatusInternalServerError)
		log.Printf("データの書き出しに失敗しました。\n%v\n", err)
		return
	}

	w.Header().Set("Location", "/")
	w.WriteHeader(http.StatusTemporaryRedirect)
}

//ReadAccounts Read accounts.json file
func ReadAccounts() (Accounts map[string]AccountList, err error) {
	bytes, err := ReadFile(filepath.Join("data", "accounts.json"))
	if err != nil {
		return nil, err
	}
	json.Unmarshal(bytes, &Accounts)
	return Accounts, nil
}

//ExportAccounts export accounts.json file
func ExportAccounts(Accounts map[string]AccountList) error {
	bytes, err := json.MarshalIndent(Accounts, "", "    ")
	if err != nil {
		return err
	}

	ExportFile(filepath.Join("data", "accounts.json"), bytes)
	return nil
}

//GetCookieValue Get cookie values from *http.Request
func GetCookieValue(r *http.Request) (map[string]string, error) {
	CookieData := make(map[string]string)

	cookie, err := r.Cookie("auth")
	if err == http.ErrNoCookie {
		log.Printf("必要なデータが見つかりませんでした。\nCookieが無効な可能性があります。\n%v\n", err)
		return nil, err
	}

	data, err := objx.FromBase64(cookie.Value)
	if err != nil {
		log.Printf("データ参照に失敗しました。\n%v\n", err)
		return nil, err
	}

	if _, ok := data["id"].(string); !ok {
		log.Printf("データの読み出しに失敗しました。\n%v\n", err)
		return nil, err
	}

	CookieData["id"] = data["id"].(string)

	return CookieData, nil
}
