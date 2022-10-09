package main

import (
	"log"

	"csuProject/server"
	"csuProject/server/storage"
	_ "github.com/lib/pq"
)

func main() {
	db, err := storage.NewPostgresDb()
	if err != nil {
		log.Fatalln(err)
	}
	repo := storage.NewRepository(db)
	handler := server.Router{Repo: repo}
	serv := new(server.HttpService)
	if err := serv.Run("8080", handler.InitRoutes()); err != nil {
		log.Fatalln(err)
	}
}
