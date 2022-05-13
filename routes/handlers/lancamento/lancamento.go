package lancamentohandler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jonasjesusamerico/poupancudo-api/database"
	"github.com/jonasjesusamerico/poupancudo-api/models"
)

func MakeLancamentoHandlers(route *gin.Engine) {
	r := route.Group("/lancamento")
	r.GET("/", buscarTodos)
	r.GET("/:id", buscaPorId)
	r.POST("/", criar)
	r.PATCH("/:id", atualizar)
}

func buscarTodos(c *gin.Context) {
	var lancamentos []models.Lancamento

	database.DB.Preload("Conta").Find(&lancamentos)
	c.JSON(http.StatusOK, lancamentos)
}

func buscaPorId(c *gin.Context) {
	var lancamento models.Lancamento
	id := c.Params.ByName("id") // Pega os contexto da requisição

	database.DB.First(&lancamento, id)

	if lancamento.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Aluno não encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, lancamento)
}

func criar(c *gin.Context) {
	var lancamento models.Lancamento

	err := c.ShouldBindJSON(&lancamento)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		fmt.Println("Passei aqui")
		return
	}

	teste := database.DB.Create(&lancamento)

	if teste.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": teste.Statement.Error,
		})
		return
	}
	c.JSON(http.StatusOK, lancamento)
}

func atualizar(c *gin.Context) {
	var lancamento models.Lancamento
	id := c.Params.ByName("id")

	database.DB.First(&lancamento, id)

	if err := c.ShouldBindJSON(&lancamento); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	database.DB.Model(&lancamento).UpdateColumns(lancamento)
	c.JSON(http.StatusOK, lancamento)
}

func deleta(c *gin.Context) {
	var lancamento models.Lancamento
	id := c.Params.ByName("id")
	database.DB.Delete(&lancamento, id)
	c.JSON(http.StatusOK, gin.H{"data": "Aluno deletado com sucesso"})
}
