package routes

import (
	"gin-api-rest/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")
	r.GET("/index", controllers.ExibePaginaIndex)
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	r.GET("/:nome", controllers.Saudacao)
	r.GET("/alunos/:id", controllers.ExibeAlunoPorId)
	r.POST("/alunos", controllers.CriaNovoAluno)
	r.DELETE("alunos/:id", controllers.DeletaAluno)
	r.PATCH("alunos/:id", controllers.EditarAluno)
	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoPorCPF)

	r.NoRoute(controllers.RotaNaoEncontrada)
	r.Run()
}
