package main

import(
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	api := router.PathPrefix("/api/v1").Subrouter()

	// router test
	api.HandleFunc("/", Index)

	// routers users
	api.HandleFunc("/users", GetNameAllUsers)
	api.HandleFunc("/user", GetUserAndPassword)
	api.HandleFunc("/user/{id}", DelUser)

	// routers report
	api.HandleFunc("/reports", GetAllReports)
	api.HandleFunc("/report", AddReport)

	log.Fatal(http.ListenAndServe(":8080", router))
}

// example function handle
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}
