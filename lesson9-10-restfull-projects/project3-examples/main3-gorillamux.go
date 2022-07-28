package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"github.com/gorilla/mux"
)

//gorilla web toolkit

type API struct {
	Message string "json:message"
}

func Hello(w http.ResponseWriter, r *http.Request){
	urlParams := mux.Vars(r) // all params get the

	id:= urlParams["id"]

	messageConcat := "Kullanıcı ID : " + id

	message := API{messageConcat}

	output, err := json.Marshal(message)

	checkError(err)

	fmt.Fprintf(w, string(output))


}

func main(){
	gorillaRoute := mux.NewRouter()

	gorillaRoute.HandleFunc("/api/user/{id:[0-9]+}", Hello)  //sadece numeric deger alir.
	http.Handle("/", gorillaRoute)
	http.ListenAndServe(":8080", nil)
}

func checkError(err error){
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
}

