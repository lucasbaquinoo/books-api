package main

import (
	"github.com/lucasbaquinoo/webapi-with-go/database"
	"github.com/lucasbaquinoo/webapi-with-go/server"
)

func main() {
	database.StartDB()
	server := server.NewServer()

	server.Run()
}
