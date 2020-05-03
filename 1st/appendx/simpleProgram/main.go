package main

import (
	"fmt"
	"log"
	"net/http"
)

//Sample Sample struct
type Sample struct{}

func main() {
	http.Handle("/", &Sample{})
	if err := http.ListenAndServe("0.0.0.0:8080", nil); err != nil {
		log.Fatalln("Error")
	}
}

func (s *Sample) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
	return
}
