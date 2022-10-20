package main

import (
	"log"

	"csuProject/storage"
	server "csuProject/sys"
	_ "github.com/lib/pq"
)

func main() {
	db, err := storage.NewPostgresDb()
	if err != nil {
		log.Fatalln(err)
	}
	repo := storage.NewRepository(db)
	provider := server.NewProvider()
	handler := server.Router{Repo: repo, Provider: provider}
	serv := new(server.HttpService)
	if err := serv.Run("8081", handler.InitRoutes()); err != nil {
		log.Fatalln(err)
	}
}
