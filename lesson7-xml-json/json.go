package main


import (
	"encoding/json"
	"fmt"
)
func main(){
	jsonStr := `{"name":"John", "age":30, "city":"New York"}`
	fmt.Println(jsonStr)

	var m map[string]interface{}  // ic ice olsa daha fazla map eklenmeli. 
	//json data map e donusturuluyor. Ve artık map olarak kullanılabiliyor.

	if err := json.Unmarshal([]byte(jsonStr), &m); err != nil {
		panic(err)
	}

	fmt.Println(m)


	b, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
}