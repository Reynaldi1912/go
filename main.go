package main

import (
	"net/http"

	"github.com/reynaldi1912/go/database"
	"github.com/reynaldi1912/go/routes"
)

func main() {
	db := database.InitDatabase()

	server := http.NewServeMux()
	routes.MapRoute(server, db)

	http.ListenAndServe(":8080", server)
}
