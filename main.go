package main

import(
	"fmt"
	_ "github.com/lib/pq"
	"github.com/jinzhu/gorm"
	 _ "github.com/jinzhu/gorm/dialects/postgres"
)

type User struct {
  Id        uint32  `json:"id"`
  Name      string  `json:"name"`
  Password  string  `json:"name"`
}


func main() {
	db, err := gorm.Open("postgres", "host=localhost user=postgres dbname=sda sslmode=disable password=admin123")

	if err != nil {
		fmt.Printf("Works %v",err);
	}
	db.CreateTable(&User{})
	 	defer db.Close()
	// router test
}
