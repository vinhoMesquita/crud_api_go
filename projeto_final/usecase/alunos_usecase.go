package usecase

import (
	"go-api/model"
	"go-api/repository"
)

type AlunosUsecase struct {
	//repository
	repository repository.AlunoRepository
}

// função que inicializa a struct
func NewAlunoUseCase(repo repository.AlunoRepository) AlunosUsecase {
	return AlunosUsecase{
		repository: repo,
	}
}

// função responsavel por tratar as regras de negócios
func (au *AlunosUsecase) GetAluno() ([]model.Aluno, error) {
	return au.repository.GetAluno()
}

// criar uma função de criar alunos
// a função retorna um model.aluno para retorna ao controller qual objeto criado no banco de dados
// o objeto esta parcialmente aqui porque recebe ele por parametro do controller, junta as duas informções
// o objeto que o controller passou e o id gerado pelo banco e retorna pro controlle gerar com o JSON
func (au *AlunosUsecase) CreateAluno(aluno model.Aluno) (model.Aluno, error) {
	// aluno.Imc = aluno.Peso / (aluno.Altura * aluno.Altura)
	alunoID, err := au.repository.CreateAluno(aluno)
	if err != nil {
		return model.Aluno{}, err
	}
	aluno.ID = alunoID
	return aluno, nil
}

func (au *AlunosUsecase) GetAlunoById(id int) (*model.Aluno, error) {
	return au.repository.GetAlunoById(id)
}

func (au *AlunosUsecase) DeleteAluno(id int) error {
	return au.repository.DeleteAluno(id)
}

func (au *AlunosUsecase) UpdateAluno(aluno model.Aluno) error {

	return au.repository.UpdateAluno(aluno)
}
