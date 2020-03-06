package main

import (
	"fmt"
	"net/http"
)

type authHandler struct {
	next http.Handler
}

func (h *authHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if _, err := r, Cookie("auth"); err == http.ErrNoCookie {
		w.Header().Set("Location", "/login")
		w.WriteHeader(http.StatusTemporaryRedirect)
	} else if err != nil {
		fmt.Fprintf(http.Error(w, "Internal Server Error!", http.StatusInternalServerError)
		return
	} else {
		h.next.ServeHTTP(w, r)
	}
}

func MustAuth(handler http.Handler) http.Handler{
	return &authHandler{next: handler}
}
