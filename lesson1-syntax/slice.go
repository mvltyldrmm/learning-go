package main

import "fmt"


func main(){
	myArray := []int{45,23,43}

	mySlice := myArray[:] // tum elemanlari al
	//slice ı degisirsek, array de degisir yani gercek veri degisir. slice veri tutmaz.
	fmt.Println(mySlice)

	new_list := append(myArray,12) //kullanılarak yeni liste olusturulabilir.

	fmt.Println(new_list)
}