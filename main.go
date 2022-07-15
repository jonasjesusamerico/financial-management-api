package main

import (
	"api-controle/src/config"
	"api-controle/src/contexto"
	"api-controle/src/database"
	"api-controle/src/routes"
)

func init() {
	config.Carregar()
	database.ConectarBanco()

}

func main() {
	contexto.CriaContextoGlobalAutenticacao()
	routes.Router{}.Route()
}
