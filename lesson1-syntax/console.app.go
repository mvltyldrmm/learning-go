package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	str1 := "bir"
	str2 := "iki"
	text, _ := fmt.Println(str1, str2)
	fmt.Println(text)

	fmt.Println(text)

	//PLACE HOLDER

	fmt.Printf("Value of anumber: %v\n", str1) //float icin %.2f, tip için %T

	//veri girisi

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	str, _ := reader.ReadString('\n')
	fmt.Println("Your name is: ", str)

	fmt.Print("Enter your age: ")
	//veri string gelir
	str, _ = reader.ReadString('\n')

	f, err := strconv.ParseFloat(strings.TrimSpace(str), 64)

	if err != nil { //error handling
		fmt.Println(err)
	} else {
		fmt.Println("Your age is: ", f)
	}

}
