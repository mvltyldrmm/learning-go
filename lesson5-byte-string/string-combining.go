package main

import(
	"fmt"
	"strings"
)

func main(){
	var x strings.Builder
	for i:=0;i<100;i++{
		x.WriteString(rastgeleString()) //buffer a string ekle
	}

	fmt.Println(x.String())
}

func rastgeleString() string {
	return ".::&*#@"
}