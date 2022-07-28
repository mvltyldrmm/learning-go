package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("dosyam.txt") //go da veriyi ve hatayi gonderir. (genelde)
	if err != nil {
		checkError(err)
		fmt.Println("Dosya acilamadi.")
	} else {
		fmt.Println(file)
		fmt.Println("Dosya acildi")
	}

	//eger err yi gormek istemiyorsan _ koyabilirsin. ornegin = file,_ := os.Open("dosyam.txt")
	//cunku bos tanımlayıcıdır.

	myError := errors.New("Hata oluştu.")
	fmt.Println(".Hata")
	fmt.Println(myError.Error)

	// i := -2
	// if i<0 {
	// 	return 0, fmt.Errorf("mat : hata %g", i)
	// }

}

func checkError(err error) {
	if err != nil {
		//buraya log ekleyebilirsin.
		log.Fatal("Error :", err)
		fmt.Println(err)
		os.Exit(1)
	}
}
