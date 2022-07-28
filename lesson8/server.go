package main

import (
	"net/http"
)

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("About page"))

	x := r.URL.Path[1:]
	data := "Merhaba " + x
	w.Write([]byte(data))

}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Index page"))
}

func main() {
	//temel golang server, her handle da response ve request ister. Daha sonra http dinlenir. İstemci sunucu iletişimi.
	//byte isteme sebebi, çevirmesine gerek kalmıyor. Zaten normalde byte ile iletişim kurulur.

	http.HandleFunc("/index", indexHandler)

	http.HandleFunc("/about", aboutHandler)

	http.ListenAndServe(":8080", nil) //localhost:8080  ve tls için https dir.

}
