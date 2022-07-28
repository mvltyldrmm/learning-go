package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Option struct {
	length    int
	upperCase bool
	lowerCase bool
	specials  bool
	numbers   bool
}

var source = rand.NewSource(time.Now().UnixNano()) //rastgele veri üretme
// const _charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789" //kullanılacak karakterler temel kurallar.

const _charsetLowerCase = "abcdefghijklmnopqrstuvwxyz"    //kullanılacak karakterler temel kurallar.
const _charsetUpperCase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"    //kullanılacak karakterler temel kurallar.
const _charsetNumbers = "0123456789"                      //kullanılacak karakterler temel kurallar.
const _charsetSpecials = "!@#$%^&*()_+-=[]{};':|,.<>/?~`" //kullanılacak karakterler temel kurallar.

func GeneratePassword(opt Option) (string, error) {
	charset := "."

	if opt.upperCase {
		charset += _charsetUpperCase
	}
	if opt.lowerCase {
		charset += _charsetLowerCase
	}
	if opt.specials {
		charset += _charsetSpecials
	}
	if opt.numbers {
		charset += _charsetNumbers
	}

	if charset == "." {
		return "", fmt.Errorf("no charset selected")
	}

	x := make([]byte, opt.length) //byte oluşturdu
	for i := range x {
		x[i] = charset[source.Int63()%int64(len(charset))] //rastgele karakterleri üret
	}
	return string(x), nil
}

func main() {
	password, _ := GeneratePassword(Option{length: 10, upperCase: false, lowerCase: false, specials: false, numbers: false})
	fmt.Println(password)
}
