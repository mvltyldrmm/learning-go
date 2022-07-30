package handlers

import (
	"fmt"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	en_translations "gopkg.in/go-playground/validator.v9/translations/en"
	"log"
	"net/http"
)

type UserTest struct {
	Email    string `json:"email" validate:"required,email"`
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"passwd"`
}

func SpecialValidate(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	translator := en.New()
	uni := ut.New(translator, translator)

	trans, found := uni.GetTranslator("en")
	if !found {
		log.Fatal("translator not found")
	}

	v := validator.New()

	if err := en_translations.RegisterDefaultTranslations(v, trans); err != nil {
		log.Fatal(err)
	}

	_ = v.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "{0} is a required field", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())
		return t
	})

	_ = v.RegisterTranslation("email", trans, func(ut ut.Translator) error {
		return ut.Add("email", "{0} must be a valid email", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("email", fe.Field())
		return t
	})

	_ = v.RegisterTranslation("passwd", trans, func(ut ut.Translator) error {
		return ut.Add("passwd", "{0} is not strong enough", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("passwd", fe.Field())
		return t
	})

	_ = v.RegisterValidation("passwd", func(fl validator.FieldLevel) bool {
		return len(fl.Field().String()) > 6
	})

	uName := r.FormValue("username")
	email := r.FormValue("email")
	pwd := r.FormValue("password")

	a := UserTest{
		Email:    email,
		Name:     uName,
		Password: pwd,
	}
	err := v.Struct(a)

	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			fmt.Fprintf(w,"Error:",e.Translate(trans))
		}

	} else {
		fmt.Fprintf(w, "Registered success. Username: %v\n", uName)

	}

	log.Printf("Server is running on port 8080")

}
