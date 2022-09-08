package main

import (
	"csuProject/server"
	"log"
)

func main() {
	handler := server.Router{}
	serv := new(server.HttpService)
	if err := serv.Run("8080", handler.InitRoutes()); err != nil {
		log.Fatalln(err)
	}
}
