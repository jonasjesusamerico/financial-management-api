package routes

import (
	"api-controle/src/routes/handler"

	"github.com/gin-gonic/gin"
)

type Router struct {
}

func (route Router) Route() {

	r := gin.Default()
	r.SetTrustedProxies(nil)
	// gin.SetMode(gin.ReleaseMode)

	handler.Handler{Route: r}.MakeHandlers()

	r.Run(":8000")
}
