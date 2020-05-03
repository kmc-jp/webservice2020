package main

import (
	"fmt"
	"log"
	"net/http"
)

//Page show main page
type Page struct{}

func main() {
	fmt.Printf("Start listen at:http://localhost:8800")
	http.Handle("/", Page{})
	if err := http.ListenAndServe(":8800", nil); err != nil {
		log.Print(err)
	}
}

func (m Page) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
	return
}
