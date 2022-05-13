package database

import (
	"fmt"
	"log"

	"github.com/jonasjesusamerico/poupancudo-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDados() {
	dsn := "host=localhost user=postgres password=postgres dbname=poupancudo port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic("Erro ao conectar ao banco de dados")
	}

	// DB.AutoMigrate(&models.User{})
	// DB.AutoMigrate(&models.CreditCard{})

	err := DB.AutoMigrate(&models.Lancamento{})
	if err != nil {
		fmt.Println(err.Error())
	}

	// DB.AutoMigrate(&models.Carteira{})
	DB.AutoMigrate(&models.Conta{})
	// DB.AutoMigrate(&models.Estabelecimento{})
	// DB.AutoMigrate(&models.Setor{})
}
