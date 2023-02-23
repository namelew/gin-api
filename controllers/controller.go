package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/namelew/gin-api/controllers/alunos"
)

var alunosCRL alunos.AlunosController

func Get(c *gin.Context) {
	switch c.Request.URL.String() {
	case "/alunos":
		c.JSON(200, alunosCRL.GetALL())
	}
}
