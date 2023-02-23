package alunos

import "github.com/namelew/gin-api/models"

type AlunosController struct {
	students []models.Aluno
}

func (a *AlunosController) GetALL() []models.Aluno {
	return append(a.students, models.Aluno{Id: 1, Name: "Diogo"})
}

func (a *AlunosController) Get(key models.Aluno) []models.Aluno {
	return a.students
}

func (a *AlunosController) Add(student models.Aluno) {

}

func (a *AlunosController) Update(old models.Aluno, new models.Aluno) {

}

func (a *AlunosController) Delete(key models.Aluno) {

}
