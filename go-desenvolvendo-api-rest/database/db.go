package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaDB() {
	stringConexao := "host=localhost user=postgres password=123 dbname=go-loja port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringConexao))

	if err != nil {
		log.Panic("Error ao conectar")
	}
}
