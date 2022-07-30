package main

import (
	"fmt"
	"net/http"

	. "validateExample/handlers"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/login", SpecialValidate).Methods(("POST"))
	r.HandleFunc("/signup", BasicDefaultValidate).Methods(("POST"))

	server := &http.Server{Addr: ":8080", Handler: r}
	fmt.Println("Server started at port 8080")
	server.ListenAndServe()

}
