package main

import (
	"net/http"
	"fmt"
	helper "project2/helpers"
)

func main(){

	uName, email, pwd, pwdConfirm := "", "", "", ""

	mux := http.NewServeMux()

	//sign up
	mux.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		uName = r.FormValue("username")
		email = r.FormValue("email")
		pwd = r.FormValue("password")
		pwdConfirm = r.FormValue("confirm")
		
		uNameCheck := helper.IsEmpty(uName)
		emailCheck := helper.IsEmpty(email)
		pwdCheck := helper.IsEmpty(pwd)
		pwdConfirmCheck := helper.IsEmpty(pwdConfirm)

		if uNameCheck || emailCheck || pwdCheck || pwdConfirmCheck {
			fmt.Fprintf(w, "Lütfen tüm alanları doldurunuz.")
			return
		}

		if pwd == pwdConfirm {
			fmt.Fprintf(w,"Kayit basarili.")
		}else{
			fmt.Fprintf(w,"Parolalar eşleşmiyor.")
		}

		// for key, value := range r.Form {
		// 	fmt.Printf("key:", key)
		// 	fmt.Printf("value:", value)
		// }
		
	})

	//login
	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {

		r.ParseForm()

		uName = r.FormValue("username")
		pwd = r.FormValue("password")

		uNameCheck := helper.IsEmpty(uName)
		pwdCheck := helper.IsEmpty(pwd)

		if uNameCheck || pwdCheck {
			fmt.Fprintf(w, "Lütfen tüm alanları doldurunuz.")
			return
		}

		if uName == "admin" && pwd == "123456" {
			fmt.Fprintf(w, "Giris basarili.")
		}else{
			fmt.Fprintf(w, "Giris basarisiz.")
		}
	})

	http.ListenAndServe(":8080", mux)
}