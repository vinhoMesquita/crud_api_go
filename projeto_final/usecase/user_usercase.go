package usecase

import (
	"go-api/model"
	"go-api/repository"
)

// Estrutura UserUsecase que contém o repositório de usuário
type UserUsecase struct {
	repository repository.UserRepository
}

// Função de instância de UserUsecase
func NewUserUseCase(repo repository.UserRepository) UserUsecase {
	return UserUsecase{
		repository: repo,
	}
}

// Função de use case que chama GetUsers no repositório nao utilizada na pagina de freqeuncias
func (us *UserUsecase) GetUsers() ([]model.User, error) {
	return us.repository.GetUsers()
}

// Função de use case que chama CreateUser no repositório nao utilizada na pagina de frequencias
func (us *UserUsecase) CreateUser(user model.User) (model.User, error) {
	userid, err := us.repository.CreateUser(user)
	if err != nil {
		return model.User{}, err
	}

	user.Id = userid

	return user, nil

}

// Função de use case que chama GetUserById no repositório utilizada para pegar os dados do user e atualizar a pagina de graficos
func (us *UserUsecase) GetUserById(id_user int) (*model.User, error) {

	user, err := us.repository.GetUserById(id_user)
	if err != nil {
		return nil, nil
	}

	return user, nil
}
