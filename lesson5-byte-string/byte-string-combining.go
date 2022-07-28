package main

import (
	"bytes"
	"fmt"
)

func main() {

	//iki string + ile toplama buyuk boyutlarda performanssızdır o yuzden asagıdakı gibi birlestilir.

	var x bytes.Buffer
	for i := 0; i < 100; i++ {
		x.WriteString(rastgeleString()) //buffer a string ekle
	}

	fmt.Println(x.String())
}

func rastgeleString() string {
	return ".::&*#@"
}
