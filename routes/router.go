package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/namelew/gin-api/routes/alunos"
)

func HandleRequests() {
	r := gin.Default()
	alunos.HandleRequests(r)
	r.Run()
}
