package carteirahandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jonasjesusamerico/poupancudo-api/database"
	"github.com/jonasjesusamerico/poupancudo-api/models"
)

func MakeCarteiraHandlers(route *gin.Engine) {
	r := route.Group("/carteira")
	r.GET("/", buscarTodos)
	r.GET("/:id", buscaPorId)
	r.POST("/", criar)
	r.PATCH("/:id", atualizar)
}

func buscarTodos(c *gin.Context) {
	var carteiras []models.Carteira
	database.DB.Find(&carteiras)
	c.JSON(http.StatusOK, carteiras)
}

func buscaPorId(c *gin.Context) {
	var carteira models.Carteira
	id := c.Params.ByName("id") // Pega os contexto da requisição

	database.DB.First(&carteira, id)

	if carteira.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Aluno não encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, carteira)
}

func criar(c *gin.Context) {
	var carteira models.Carteira

	if err := c.ShouldBindJSON(&carteira); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	database.DB.Create(&carteira)
	c.JSON(http.StatusOK, carteira)
}

func atualizar(c *gin.Context) {
	var carteira models.Carteira
	id := c.Params.ByName("id")

	database.DB.First(&carteira, id)

	if err := c.ShouldBindJSON(&carteira); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	database.DB.Model(&carteira).UpdateColumns(carteira)
	c.JSON(http.StatusOK, carteira)
}

func deleta(c *gin.Context) {
	var carteira models.Carteira
	id := c.Params.ByName("id")
	database.DB.Delete(&carteira, id)
	c.JSON(http.StatusOK, gin.H{"data": "Aluno deletado com sucesso"})
}
