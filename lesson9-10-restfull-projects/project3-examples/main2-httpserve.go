// package main
package main
// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"os"
// )

// type API struct {
// 	Message string "json:message"
// }

// type User struct {
// 	ID        int    "json:id"
// 	FirstName string "json:first_name"
// 	LastName  string "json:last_name"
// 	Age       int    "json:age"
// }

// func main() {
// 	apiRoot := "/api"

// 	// .../api/user/1

// 	http.HandleFunc(apiRoot, func(w http.ResponseWriter, r *http.Request) {
// 		message := API{"API home"}

// 		output, err  := json.Marshal(message) //convert to JSON
// 		checkError(err)

// 		w.Header().Set("Content-Type", "application/json") //set header
// 		fmt.Fprintf(w, string(output))

// 	})

// 	http.HandleFunc(apiRoot + "/users", func(w http.ResponseWriter, r *http.Request) {
// 		users := []User{
// 			User{ID: 1, FirstName: "John", LastName: "Doe", Age: 30},
// 			User{ID: 2, FirstName: "Jane", LastName: "Doe", Age: 25},
// 			User{ID: 3, FirstName: "Jack", LastName: "Doe", Age: 20},
// 		}

// 		message := users

// 		output, err := json.Marshal(message)

// 		checkError(err)

// 		fmt.Fprintf(w, string(output))


// 	})

// 	http.HandleFunc(apiRoot + "/me", func(w http.ResponseWriter, r *http.Request) {
// 		user := User{ID: 1, FirstName: "John", LastName: "Doe", Age: 30}

// 		message := user

// 		output, err := json.Marshal(message)

// 		checkError(err)

// 		fmt.Fprintf(w, string(output))

// 	})

// 	http.ListenAndServe(":8080", nil)
// }

// func checkError(err error) {
// 	if err != nil {
// 		panic(err)
// 		os.Exit(1) //close app
// 	}
// }