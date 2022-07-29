package starter

import (
	"fmt"
	"math"
	"net/http"
	"reflect"
	"strings"
)

type Register struct {
	username string
	password string
}

func SayHello(name string) string {
	return fmt.Sprintf("Hello %v. Welcome!", name)
}

func OddOrEven(num int) string {
	criteria := math.Mod(float64(num), 2)
	if criteria == 1 || criteria == -1 {
		return fmt.Sprintf("%v is an odd number", num)
	}
	return fmt.Sprintf("%v is an even number", num)
}

func Checkhealth(writer http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(writer, "health check passed")
}

func RegisterExample(name string, surname string, password string) string {

	if name == "" || surname == "" || password == "" {
		return fmt.Sprintf("Name, Surname or password is empty")
	}

	name_type, surname_type, password_type := reflect.TypeOf(name), reflect.TypeOf(surname), reflect.TypeOf(password)
	if name_type != reflect.TypeOf("") || surname_type != reflect.TypeOf("") || password_type != reflect.TypeOf("") {
		return fmt.Sprintf("Name, Surname or password is not string")
	}

	username := strings.ToLower(name[0:1] + surname)

	var register Register = Register{username, password}

	return fmt.Sprintf("%v is registered", register.username)
}
