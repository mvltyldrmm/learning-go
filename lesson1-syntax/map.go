package main

import "fmt"

func main() {
	//map nesnesi olusturma yontem 1
	myMap := make(map[int]string) //anahtar int, karsısı string

	fmt.Println(myMap)

	myMap[43] = "Foo"

	fmt.Println(myMap[43])

	//yontem2

	states := make(map[string]string)

	states["TX"] = "Texas"
	states["FL"] = "Florida"
	states["NY"] = "New York"
	states["OR"] = "Oregon"
	states["CA"] = "California"

	fmt.Println(states)

	delete(states, "OR")

	new_york := states["NY"]

	fmt.Println(new_york)

	keys:= make([]string, len(states))

	i := 0

	for k:= range states {
		keys[i] = k
		i++
	}
}