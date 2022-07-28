package dataloaders

//go mod dosyası, workdir dir. ordan itibaren alınması gerekli (. ile herşeyi aktarır.)

import (
	"encoding/json"
	. "project/models"
	util "project/utils"
)

func LoadUsers() []User {
	bytes, _ := util.ReadFile("json/users.json")
	var users []User
	json.Unmarshal([]byte(bytes), &users)
	return users
}

func LoadInterests() []Interest {
	bytes, _ := util.ReadFile("json/interest.json")
	var interest []Interest
	json.Unmarshal([]byte(bytes), &interest)
	return interest
}

func LoadInterestsMappigns() []InterestMapping {
	bytes, _ := util.ReadFile("json/userInterestMappings.json")
	var interestMapping []InterestMapping
	json.Unmarshal([]byte(bytes), &interestMapping)
	return interestMapping
}