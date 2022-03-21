package main

import (
	"github.com/pedr0diniz/alura-go-5/database"
	"github.com/pedr0diniz/alura-go-5/routes"
)

func main() {
	database.ConnectToDatabase()
	routes.HandleRequests()
}
