package alunos

import (
	"github.com/gin-gonic/gin"
	"github.com/namelew/gin-api/controllers"
)

func HandleRequests(r *gin.Engine) {
	r.GET("/alunos", controllers.GetALL)
	r.GET("/alunos/:id", controllers.Get)
	r.DELETE("/alunos/:id", controllers.Delete)
	r.PATCH("/alunos/:id", controllers.Update)
	r.POST("/alunos", controllers.Create)
}
