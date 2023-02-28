package alunos

import (
	"errors"

	"github.com/namelew/gin-api/database"
	"github.com/namelew/gin-api/models"
)

type AlunosController struct {
}

var emptyAluno models.Aluno

func (a *AlunosController) GetALL(filter models.Aluno) []models.Aluno {
	var alunos []models.Aluno
	database.DB.Where(&filter).Find(&alunos)
	return alunos
}

func (a *AlunosController) Get(id int) (*models.Aluno, error) {
	var aluno models.Aluno

	database.DB.Find(&aluno, id)

	if aluno == emptyAluno {
		return nil, errors.New("aluno não encontrado")
	}

	return &aluno, nil
}

func (a *AlunosController) Add(aluno models.Aluno) error {
	go database.DB.Create(&aluno)
	return nil
}

func (a *AlunosController) Update(id int, new models.Aluno) {
	var aluno models.Aluno
	database.DB.Find(&aluno, id)
	go database.DB.Model(&aluno).UpdateColumns(new)
}

func (a *AlunosController) Delete(id int) {
	var aluno models.Aluno
	go database.DB.Delete(&aluno, id)
}
