package routes

import (
	"github.com/gin-gonic/gin"
	carteirahandler "github.com/jonasjesusamerico/poupancudo-api/routes/handlers/carteira"
	estabelecimentohandler "github.com/jonasjesusamerico/poupancudo-api/routes/handlers/estabelecimento"
	lancamentohandler "github.com/jonasjesusamerico/poupancudo-api/routes/handlers/lancamento"
)

func Route() {
	// gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.SetTrustedProxies(nil)

	carteirahandler.MakeCarteiraHandlers(r)
	estabelecimentohandler.MakeEstabelecimentoHandlers(r)
	lancamentohandler.MakeLancamentoHandlers(r)

	r.Run(":8000")
}
