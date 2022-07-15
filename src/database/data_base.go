package database

import (
	"api-controle/src/config"
	"api-controle/src/model"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	Connect Db
	err     error
)

type Db struct {
	db_postgres *gorm.DB
}

func (d Db) GetInstance() (bancoDados *gorm.DB) {
	bancoDados = d.db_postgres
	return
}

func ConectarBanco() {
	Connect.db_postgres, err = gorm.Open(postgres.Open(config.StringConexaoBanco), &gorm.Config{})
	if err != nil {
		log.Panic("Erro ao conectar ao banco de dados")
	}

	err = Connect.db_postgres.AutoMigrate(&model.Usuario{})

	if err != nil {
		log.Panic(err.Error())
	}

}
