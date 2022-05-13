package main

import (
	"github.com/jonasjesusamerico/poupancudo-api/database"
	"github.com/jonasjesusamerico/poupancudo-api/routes"
)

func main() {
	database.ConectaComBancoDados()
	routes.Route()
}
