package main

import (
	"go-gin-api-rest-simplicidade/database"
	"go-gin-api-rest-simplicidade/routes"
)

func main() {
	database.ConectaDB()
	routes.HandleRequests()
}
