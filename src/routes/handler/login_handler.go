package handler

import (
	"api-controle/src/controllers"
	"api-controle/src/repository"

	"github.com/gin-gonic/gin"
)

type LoginHandler struct {
	Repo  repository.IRepository
	Route *gin.RouterGroup
}

func (l *LoginHandler) New(repo repository.IRepository, rota *gin.RouterGroup) IHandler {
	l.Repo = repo
	l.Route = rota
	return l
}

func (l LoginHandler) RotasAutenticadas() IHandler {
	return &l
}

func (l LoginHandler) RotasNaoAutenticadas() IHandler {

	controller := controllers.LoginController{Repo: l.Repo}

	route := l.Route.Group("/" + controller.NameGroupRoute())
	{
		route.POST("", controller.Login)
	}
	return &l
}
