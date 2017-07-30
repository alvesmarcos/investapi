package main

type User struct {
	Id				uint32	`bson:"id" json:"id"`
	Username	string	`bson:"username" json:"username"`
	Password 	string	`bson:"password" json:"password"`
}

type Report struct {
	Id			uint32		`bson:"id" json:"id"`
	Title		string		`bson:"title" json:"title"`
	Body		string		`bson:"body" json:"body"`
	Images	[]string	`bson:"images" json:"images"`
}

type Sample struct {
	Id		uint32	`bson:"id" json:"id"`
	Date	string	`bson:"date" json:"date"`
	Value	string	`bson:"value" json:"value"`
}

type Indicator struct {
	Id					uint32		`bson:"id" json:"id"`
	Name				string		`bson:"name" json:"name"`
	Description	string		`bson:"description" json:"description"`
	Metric			string 		`bson:"metric" json:"metric"`
	Status			string 		`bson:"status" json:"status"`
	Date				string 		`bson:"date" json:"date"`
	Samples			[]Sample	`bson:"samples" json:"samples"`
}
