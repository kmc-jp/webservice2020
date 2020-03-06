package main

import (
	"net/http"
)

type authHandler struct {
	next http.Handler
}

func (h *authHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if _, err := r.Cookie("auth"); err == http.ErrNoCookie {
		w.Header().Set("Location", "/login")
		w.WriteHeader(http.StatusTemporaryRedirect)
	} else if err != nil {
		http.Error(w, "Internal Server Error!", http.StatusInternalServerError)
		return
	} else {
		h.next.ServeHTTP(w, r)
	}
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

	default:
		http.Error(w, fmt.Sprintf("操作：%sには対応していません。\n", action), http.StatusNotFound)
		return
	}
	return
}
