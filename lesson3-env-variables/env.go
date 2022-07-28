package main

import (
	"fmt"
	"os"
)

func main(){
	// for _, env := range os.Environ(){
	// 	fmt.Println(env)
	// }

	uName := os.Getenv("USERNAME")
	uDomain := os.Getenv("USERDOMAIN")
	processsArcitectur := os.Getenv("PROCESSOR_ARCHITECTURE")

	homePath := os.Getenv("HOMEPATH")

	fmt.Println("Username: ", uName)
	fmt.Println("User Domain: ", uDomain)
	fmt.Println("Processor Architecture: ", processsArcitectur)
	fmt.Println("Home Path: ", homePath)

	goPath := os.Getenv("GOPATH")

	fmt.Println("Go Path: ", goPath)

}