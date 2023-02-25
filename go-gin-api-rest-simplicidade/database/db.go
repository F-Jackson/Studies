package database

import (
	"go-gin-api-rest-simplicidade/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaDB() {
	conexao := "host=localhost user=postgres password=123 dbname=postgres "
	DB, err = gorm.Open(postgres.Open(conexao))
	if err != nil {
		log.Panic("Erro ao conectar")
	}
	DB.AutoMigrate(&models.Aluno{})
}
