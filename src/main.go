package main

import(
	"fmt"
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	api := router.PathPrefix("/api/v1").Subrouter()
	
	// handlers test
	api.HandleFunc("/", Index)
	api.HandleFunc("/users", GetNameUser)

	log.Fatal(http.ListenAndServe(":8080", router))
}

// example function handle
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

// get users mockup
func GetNameUser(w http.ResponseWriter, r *http.Request) {
	var users [2]User
	users[0] = User{ Name: "Marcos", Password: "08quinho" }
	users[1] = User{ Name: "Rubenita", Password: "mainha" }

	if err := json.NewEncoder(w).Encode(users); err != nil {
		panic(err)
	}
}