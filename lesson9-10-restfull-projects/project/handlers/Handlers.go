package handlers

import (
	"encoding/json"
	"log"
	. "project/dataloaders"
	. "project/models"
	"net/http"
)

func Run() {
	http.HandleFunc("/", Handler)
	log.Print("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func Handler(w http.ResponseWriter, r *http.Request) {

	page := Page{ID: 7, Name: "Kullanıcılar", Description: "Kullanıcı Listesi", URI: "/users"}

	users := LoadUsers()
	interests := LoadInterests()
	interestMappings := LoadInterestsMappigns()

	var newUsers []User

	for _, user := range users {
		for _, interestMapping := range interestMappings {
			if user.ID == interestMapping.UserID {
				for _, interest := range interests {
					if interest.ID == interestMapping.InterestID {
						user.Interests = append(user.Interests, interest)
					}
				}
			}
		}
		newUsers = append(newUsers, user)

	}

	viewModel := UserViewModel{Page: page, Users: newUsers}

	data, _ := json.Marshal(viewModel)

	w.Write([]byte(data))

}
