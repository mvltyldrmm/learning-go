package main

import "fmt"

func main() {
	myArray := [5]int{}
	myArray[0] = 1
	myArray[1] = 2
	myArray[2] = 3

	fmt.Println(myArray)
	fmt.Println(len(myArray))

	//dizilerde eleman verilmelidir. ama oto size ozelligi var

	myArray2 := [...]int{1, 2, 3, 4, 5} //derleyici sayar.
	fmt.Println(cap(myArray2))          //cap: kapasite
	fmt.Println(len(myArray2))

	i := 0

	for i < len(myArray2) {
		fmt.Println(myArray2[i])
		i++
	}

	for i, value := range myArray2 {
		fmt.Println(i, value)
	}
}
