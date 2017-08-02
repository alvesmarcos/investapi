package main

import(
	"fmt"
)

type User struct {
  Id        uint32  `json:"id"`
  Name      string  `json:"name"`
  Password  string  `json:"name"`
}


func main() {
	var user User
	fmt.Printf("%v", user)
}
