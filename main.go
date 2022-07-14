package main

import (
	"api-controle/src/contexto"
	"api-controle/src/database"
	"api-controle/src/routes"
)

func init() {
	database.ConnectWithDatabase()

}

func main() {
	contexto.CriaContextoGlobalAutenticacao()
	routes.Router{}.Route()
}
