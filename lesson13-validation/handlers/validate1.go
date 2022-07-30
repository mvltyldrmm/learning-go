package handlers

import (
	"fmt"

	"gopkg.in/go-playground/validator.v9"

	"net/http"
)

type User struct {
	Email    string `json:"email" validate:"required,email"`
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required,gte=8"`
}

func BasicDefaultValidate(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	v := validator.New()

	uName := r.FormValue("username")
	email := r.FormValue("email")
	pwd := r.FormValue("password")

	a := User{
		Email:    email,
		Name:     uName,
		Password: pwd,
	}
	err := v.Struct(a)

	fmt.Println(err)

	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			fmt.Println(e)
		}
		fmt.Fprintln(w, "Error:", err)
	} else {
		fmt.Fprintf(w, "Registered success. Username: %v\n", uName)

	}

}
