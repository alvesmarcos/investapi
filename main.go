package main

import "github.com/alvesmarcos/investapi/app/server"


func main() {
	s := server.NewServer()
	s.Start()
}
