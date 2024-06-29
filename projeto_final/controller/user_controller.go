package controller

import (
	"github.com/gin-gonic/gin"
	"go-api/model"
	"go-api/usecase"
	"net/http"
	"strconv"
)

// Estrutura do UserControler
type UserController struct {
	userUseCase usecase.UserUsecase
}

// Função de instanciação
func NewUserController(usecase usecase.UserUsecase) UserController {
	return UserController{
		userUseCase: usecase,
	}
}

// Função que lida com a requisição GET para obter dados de todos os usuario, nao utilizada na pagina de graficos
func (u *UserController) GetUsers(ctx *gin.Context) {

	users, err := u.userUseCase.GetUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, users)

}

// // Função que lida com a requisicao POST para criar novos usuario, nao utilizada na pagina degraficos
func (u *UserController) CreateUser(ctx *gin.Context) {

	var user model.User
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	insertedUser, err := u.userUseCase.CreateUser(user)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, insertedUser)

}

// Função que lida com a requisição GET para obter dados de um usuario com base no seu id
func (u *UserController) GetUsersById(ctx *gin.Context) {

	id := ctx.Query("userid")
	if id == "" {
		response := model.Response{
			Message: "ID do usuario nao pode ser nulo",
		}

		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	userid, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response{
			Message: "ID do usuario precisa ser  um numero",
		}
		ctx.JSON(http.StatusInternalServerError, response)
	}

	user, err := u.userUseCase.GetUserById(userid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if user == nil {
		response := model.Response{
			Message: "Nao foi encontrado usuario na base de dade",
		}

		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	ctx.JSON(http.StatusOK, user)

}
