package main

import (
	"log"
	"net/http"

	"github.com/apiGO/internal/config"
	"github.com/apiGO/internal/routes"
)

func main() {
	db := config.ConnectDB()
	defer db.Close()

	r:= routes.RegisterRoutes(db)

	log.Println("Server running on: 8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}
