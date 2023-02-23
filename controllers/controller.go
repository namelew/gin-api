package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/namelew/gin-api/controllers/alunos"
	"github.com/namelew/gin-api/models"
)

var alunosCRL alunos.AlunosController

func Get(c *gin.Context) {
	switch c.Request.URL.String() {
	case "/alunos":
		c.JSON(200, alunosCRL.GetALL())
	}
}

func Create(c *gin.Context) {
	switch c.Request.URL.String() {
	case "/alunos":
		var aluno models.Aluno

		if err := c.ShouldBindJSON(&aluno); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := alunosCRL.Add(&aluno); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, aluno)
	}
}
