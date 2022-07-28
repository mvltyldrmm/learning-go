package main

import (
	"io"
	"net/http"
)

type ironman int

func (x ironman) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello from Ironman")
}
type wolverine int

func (x wolverine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello from Wolverine")
}

func main(){
	var i ironman
	var w wolverine

	mux := http.NewServeMux()

	mux.Handle("/ironman", i)
	mux.Handle("/wolverine", w)

	http.ListenAndServe(":8080", mux)
}