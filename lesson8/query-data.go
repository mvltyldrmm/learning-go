package main

import (
	"net/http"
)

func main(){
	
	//veri geçme - veri gönderme

	http.HandleFunc("/search", search)

	http.ListenAndServe(":8080", nil)
}

func search(w http.ResponseWriter, r *http.Request){
	h1 := r.FormValue("h1")
	source := r.FormValue("source")

	data := "Merhaba " + h1 + " " + source

	w.Write([]byte(data))

}

//google.com7search?h1=hello&source=google vs gibi. form value ile çekilebilir.