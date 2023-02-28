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
		var aluno models.Aluno

		if err := c.ShouldBindJSON(&aluno); err != nil && err.Error() != "EOF" {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		alunos, err := alunosCRL.GetALL(aluno)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, alunos)
	}
}

func Get(c *gin.Context) {
	url := c.Request.URL.String()
	switch {
	case url[:7] == "/alunos":
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

		c.JSON(http.StatusOK, gin.H{"data": "aluno criado com sucesso"})
	}
}

func Delete(c *gin.Context) {
	url := c.Request.URL.String()
	switch {
	case url[:7] == "/alunos":
		id := c.Params.ByName("id")
		key, err := strconv.Atoi(id)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := alunosCRL.Delete(key); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": "aluno deletado com sucesso"})
	}
}

func Update(c *gin.Context) {
	url := c.Request.URL.String()

	switch {
	case url[:7] == "/alunos":
		var new models.Aluno
		id := c.Params.ByName("id")

		key, err := strconv.Atoi(id)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := c.ShouldBindJSON(&new); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := alunosCRL.Update(key, new); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": "aluno atualizado com sucesso"})
	}
}
