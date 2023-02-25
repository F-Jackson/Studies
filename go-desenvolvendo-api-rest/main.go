package main

import (
	"fmt"

	"go-desenvolvendo-api-rest/database"
	"go-desenvolvendo-api-rest/models"
	"go-desenvolvendo-api-rest/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Nome 1", Historia: "Historia 1"},
		{Id: 2, Nome: "Nome 2", Historia: "Historia 2"},
	}

	database.ConectaDB()

	fmt.Println("iniciando")
	routes.HandlleRequest()
}
