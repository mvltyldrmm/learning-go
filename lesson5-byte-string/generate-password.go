package main

import(
	"fmt"
	"math/rand"
	"time"
)

var source = rand.NewSource(time.Now().UnixNano()) //rastgele veri üretme
const _charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789" //kullanılacak karakterler temel kurallar.

func GeneratePassword(length int) string {
	x:= make([]byte, length) //byte oluşturdu
	for i:= range x{
		x[i] = _charset[source.Int63() % int64(len(_charset))] //rastgele karakterleri üret
	}

	return string(x)
}

func main(){
	password := GeneratePassword(10)
	fmt.Println(password)
}