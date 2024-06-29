package usecase

import (
	"go-api/model"
	"go-api/repository"
)

// Estrutura FrequenciaUsecase que contém o repositório de frequência
type FrequenciaUsecase struct {
	repository repository.FrequenciaRepository
}

// Função de instância de FrequenciaUsecase
func NewFrequenciaCase(repo repository.FrequenciaRepository) FrequenciaUsecase {
	return FrequenciaUsecase{
		repository: repo,
	}
}

// Função de use case que chama GetFrequenciaByFilters no repositório
func (fr *FrequenciaUsecase) GetFrequenciaByFilters(id_user int, startDate, endDate string) ([]model.Frequencia, error) {
	return fr.repository.GetFrequenciaByFilters(id_user, startDate, endDate)
}

// Função de use case que chama CreateFrequencia no repositório
func (fr *FrequenciaUsecase) CreateFrequencia(frequencia model.Frequencia) (model.Frequencia, error) {
	id, err := fr.repository.CreateFrequencia(frequencia)
	if err != nil {
		return model.Frequencia{}, err
	}

	frequencia.Id_user = id

	return frequencia, nil
}

// Função de use case que chama UpdateFrequencia no repositório
func (fr *FrequenciaUsecase) UpdateFrequencia(frequencia model.Frequencia) (model.Frequencia, error) {
	id, err := fr.repository.UpdateFrequencia(frequencia)
	if err != nil {
		return model.Frequencia{}, err
	}

	frequencia.Id_user = id

	return frequencia, nil
}
