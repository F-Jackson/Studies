package main

import (
	"go-validacoes-testes-paginas-html/database"
	"go-validacoes-testes-paginas-html/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
