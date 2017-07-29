package main

import(
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
)

// handlers user

func GetNameAllUsers(w http.ResponseWriter, r *http.Request) {
	var users [2]User
	users[0] = User{ Id: 1, Username: "Marcos", Password: "08quinho" }
	users[1] = User{ Id: 2, Username: "Rubenita", Password: "mainha" }

	if err := json.NewEncoder(w).Encode(users); err != nil {
		panic(err)
	}
}

func GetUserAndPassword(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	username, password := query.Get("username"), query.Get("password")
	// TODO
	// response := database.FindUserByUsernameAndPassword(username, password)

	if err := json.NewEncoder(w).Encode(username+password); err != nil {
		panic(err)
	}
}

func DelUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"] 
	// TODO
	// response := database.DeleteUserById(user)

	if err := json.NewEncoder(w).Encode(id); err != nil {
		panic(err)
	}
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	username, password := query.Get("username"), query.Get("password")
	user := User{ Id: 0, Username: username, Password: password }
	// TODO
	// reponse := database.SaveUser(user)

	if err := json.NewEncoder(w).Encode(user); err != nil {
		panic(err)
	}
}