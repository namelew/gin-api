package alunos

import (
	"errors"

	"github.com/namelew/gin-api/database"
	"github.com/namelew/gin-api/models"
)

type AlunosController struct {
}

func (a *AlunosController) GetALL(filter models.Aluno) ([]models.Aluno, error) {
	var alunos []models.Aluno
	if err := (database.DB.Where(&filter).Find(&alunos)).Error; err != nil {
		return nil, err
	}
	return alunos, nil
}

func (a *AlunosController) Get(id int) (*models.Aluno, error) {
	var aluno models.Aluno

	if err := (database.DB.Find(&aluno, id)); err != nil {
		return nil, errors.New("aluno não encontrado")
	}

	return &aluno, nil
}

func (a *AlunosController) Add(aluno models.Aluno) error {
	if err := (database.DB.Create(&aluno)).Error; err != nil {
		return err
	}
	return nil
}

func (a *AlunosController) Update(id int, new models.Aluno) error {
	var aluno models.Aluno

	if err := (database.DB.Find(&aluno, id)).Error; err != nil {
		return err
	}

	return (database.DB.Model(&aluno).UpdateColumns(new)).Error
}

func (a *AlunosController) Delete(id int) error {
	var aluno models.Aluno
	return (database.DB.Delete(&aluno, id)).Error
}
