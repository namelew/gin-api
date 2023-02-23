package alunos

import (
	"github.com/namelew/gin-api/database"
	"github.com/namelew/gin-api/models"
)

type AlunosController struct {
}

func (a *AlunosController) GetALL() []models.Aluno {
	var alunos []models.Aluno
	database.DB.Find(&alunos)
	return alunos
}

func (a *AlunosController) Get(id int) *models.Aluno {
	var aluno models.Aluno

	database.DB.Find(&aluno, id)

	return &aluno
}

func (a *AlunosController) Add(aluno *models.Aluno) error {
	database.DB.Create(aluno)
	return nil
}

func (a *AlunosController) Update(old models.Aluno, new models.Aluno) {

}

func (a *AlunosController) Delete(key models.Aluno) {

}
