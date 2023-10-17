package main

import (
	"belajar-clean/app/routes"
	"belajar-clean/helpers"
)

func main() {
	db := helpers.DatabaseConnect()

	server := routes.RouteService(db)

	server.Logger.Fatal(server.Start(":8001"))
}