package main

import (
	"fmt"
)

type person struct {
	name    string
	surname string
}

//struct (yapi) type bildirimine yarar.

//constructor func varsayılan ve boş yapıcı metod

func NewHuman() *person {
	h := new(person)
	return h
}

//parametreli constructor

func NewHumanParams(name, surname string) *person {
	h := new(person)
	h.name = name
	h.surname = surname
	return h
}

//method overloading i desteklemez. Parametresiz-parametreli aynı isimden fonksiyon oluşturulmaz. Yeni isim farklı method gerekli.

func main() {
	new_person := person{"ali", "veli"}
	name_pointer := &new_person.name
	fmt.Println(name_pointer)
	fmt.Println(new_person.name)
	fmt.Println("Hello World")

	// veya

	y := new(person)
	fmt.Println(y)

	//boş yapıcı metod kullanımı

	z := NewHuman()

	z.name = "ali"

	fmt.Println(z.name)

	new_human_params := NewHumanParams("ali", "veli")

	fmt.Println(new_human_params.name)

}
