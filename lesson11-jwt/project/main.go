package main

import (
	"log"
	"net/http"
	. "jwtProject/handlers"
)


func main(){
	http.HandleFunc("/login", Login)
	http.HandleFunc("/home", Home)
	http.HandleFunc("/refresh", Refresh)
	log.Printf("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}