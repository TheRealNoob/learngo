package routes

import (
	"fmt"
	"io"
	"net/http"
)

func NewMux() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", root)
	mux.HandleFunc("/hello", hello)

	return mux
}

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Println("got request: /")
	io.WriteString(w, "responding on: /")
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("got request: /hello")
	io.WriteString(w, "responding on: /hello")
}
