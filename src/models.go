type User struct {
	Name			string	`bson:"name", json:"name"`
	Password 	string	`bson:"password", json:"password`
}

type Report struct {
	Title		string		`bson:"title", json:"title"`
	Body 		string		`bson:"body", json:"body"`
	Images	[]string	`bson:"images", json:"images"`
}

type Indicator struct {
	Name				string	`bson:"name", json:"name"`
	Description	string	`bson:"description", json:"description"`
	Metric			string 	`bson:"", json:""`
	Status			string 	`bson:"status", json:"status"`
	Date 				string 	`bson:"date", json:"date"`
	Samples			[] struct {
		Date	string	`bson:"date", json:"date"`
		Value	string	`bson:"value", json:"value"`
	} `bson:"samples", json:"samples"`
}