package main

import(
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
)

// HANDLERS USER

func GetNameAllUsers(w http.ResponseWriter, r *http.Request) {
	var users [2]User
	users[0] = User{ Id: 1, Username: "Marcos", Password: "08quinho" }
	users[1] = User{ Id: 2, Username: "Rubenita", Password: "mainha" }
	// TODO
	// users = database.GetUsers()

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

// HANDLERS REPORT

func GetAllReports(w http.ResponseWriter, r *http.Request) {
	report := Report{ Id: 1, Title: "Semana Economica", Body: "Eh isso aí", Images: []string{"path1", "path2"}}

	// TODO
	// reports := database.GetReports()

	if err := json.NewEncoder(w).Encode(report); err != nil {
		panic(report)
	}
}

func AddReport(w http.ResponseWriter, r *http.Request) {
	// https://golang.org/pkg/net/http/#Request.ParseForm
	r.ParseForm() // For all requests, ParseForm parses the raw query from the URL and updates r.Form.
	query := r.URL.Query()

	title, body, images := query.Get("title"), query.Get("body"), r.Form["images"]
	report := Report{ Id: 0, Title: title, Body: body, Images: images }
	// TODO
	// response := database.SaveReport(report)

	if err := json.NewEncoder(w).Encode(report); err != nil {
		panic(report)
	}
}

// HANDLERS INDICATOR

func GetAllIndicators(w http.ResponseWriter, r *http.Request) {
	indicator := Indicator{ Id: 1, Name: "IPCA", Description: "Nao eh aquilo", Metric: "%",
													Status: "Subiu", Date: "02/03/2015", Samples: []Sample{ Sample{ Date: "02/32", Value: "32.32"},
													Sample { Date: "01/01", Value: "100"} }}

	// TODO
	// indicators := database.GetIndicators()

	if err := json.NewEncoder(w).Encode(indicator); err != nil {
		panic(indicator)
	}
}

func AddIndicator(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	query := r.URL.Query()

	name, description, metric := query.Get("name"), query.Get("description"), query.Get("metric")
	status, date, sample_date, sample_value := query.Get("status"), query.Get("date"), r.Form["sample-date"], r.Form["sample-value"]

	indicator := Indicator{ Id: 0, Name: name, Description: description, Metric: metric,
													Status: status, Date: date, Samples: make([]Sample, 2) }
	indicator.Samples[0].Date = sample_date[0]
	indicator.Samples[0].Value = sample_value[0]
	indicator.Samples[1].Date = sample_date[1]
	indicator.Samples[1].Value = sample_value[1]
	// TODO
	// response := database.SaveReport(indicator)

	if err := json.NewEncoder(w).Encode(indicator); err != nil {
		panic(indicator)
	}
}
