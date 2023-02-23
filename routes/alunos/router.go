package alunos

import (
	"github.com/gin-gonic/gin"
	"github.com/namelew/gin-api/controllers"
)

func HandleRequests(r *gin.Engine) {
	r.GET("/alunos", controllers.Get)
	r.POST("/alunos", controllers.Create)
}
