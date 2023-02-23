package main

import "github.com/gin-gonic/gin"

func getAlunos(c *gin.Context) {
	c.JSON(200, gin.H{"id": 1, "nome": "Diogo"})
}

func main() {
	r := gin.Default()
	r.GET("/alunos", getAlunos)
	r.Run()
}
