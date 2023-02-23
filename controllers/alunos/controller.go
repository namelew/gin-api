package alunos

import (
	"github.com/namelew/gin-api/database"
	"github.com/namelew/gin-api/models"
)

type AlunosController struct {
	alunos []models.Aluno
}

func (a *AlunosController) GetALL() []models.Aluno {
	return append(a.alunos, models.Aluno{Nome: "Diogo", CPF: "00000000000", RG: "111111"}, models.Aluno{Nome: "Ana", CPF: "00000000000", RG: "111111"})
}

func (a *AlunosController) Get(key models.Aluno) []models.Aluno {
	return a.alunos
}

func (a *AlunosController) Add(aluno *models.Aluno) error {
	database.DB.Create(aluno)
	return nil
}

func (a *AlunosController) Update(old models.Aluno, new models.Aluno) {

}

func (a *AlunosController) Delete(key models.Aluno) {

}
