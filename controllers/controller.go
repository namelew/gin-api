package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/namelew/gin-api/controllers/alunos"
	"github.com/namelew/gin-api/models"
)

var alunosCRL alunos.AlunosController

func GetALL(c *gin.Context) {
	switch c.Request.URL.String() {
	case "/alunos":
		c.JSON(200, alunosCRL.GetALL())
	}
}

func Get(c *gin.Context) {
	switch c.Request.URL.String()[:7] {
	case "/alunos":
		id := c.Params.ByName("id")
		key, err := strconv.Atoi(id)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ret, err := alunosCRL.Get(key)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, ret)
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

		if err := alunosCRL.Add(aluno); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, aluno)
	}
}

func Delete(c *gin.Context) {
	switch c.Request.URL.String()[:7] {
	case "/alunos":
		id := c.Params.ByName("id")
		key, err := strconv.Atoi(id)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		alunosCRL.Delete(key)

		c.JSON(http.StatusOK, gin.H{"data": "aluno deletado com sucesso"})
	}
}
