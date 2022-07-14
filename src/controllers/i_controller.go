package controllers

import (
	"api-controle/src/repository"

	"github.com/gin-gonic/gin"
)

type IController interface {
	New(repository repository.IRepository) IController

	NameGroupRoute() string

	FindAll(c *gin.Context)

	FindById(c *gin.Context)

	Create(c *gin.Context)

	Update(c *gin.Context)

	Delete(c *gin.Context)
}
