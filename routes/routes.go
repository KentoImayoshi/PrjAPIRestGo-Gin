package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kentoimayoshi/PrjAPIRestGo-Gin/controllers"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	r.Run()
}
