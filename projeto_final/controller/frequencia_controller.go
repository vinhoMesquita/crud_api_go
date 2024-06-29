package controller

import (
	"github.com/gin-gonic/gin"
	"go-api/model"
	"go-api/usecase"
	"net/http"
	"strconv"
)

// FrequenciaController estrutura que contém o caso de uso de frequência
type FrequenciaController struct {
	frequenciaUseCase usecase.FrequenciaUsecase
}

// Função de instância de FrequenciaController
func NewFrequenciaController(usecase usecase.FrequenciaUsecase) FrequenciaController {
	return FrequenciaController{
		frequenciaUseCase: usecase,
	}
}

// Função que lida com a requisição GET para obter frequências filtradas por ID de usuário, data inicial e final
func (f *FrequenciaController) GetFrequenciaByFilters(ctx *gin.Context) {

	// Extraindo parametros de consulta
	id_userStr := ctx.Query("id_user")
	startDate := ctx.Query("start_date")
	endDate := ctx.Query("end_date")

	// convertendo o paramatro id_user para inteiro
	id_user, err := strconv.Atoi(id_userStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "parametro id_user Invalido"}) // Retorna um erro de Bad Request com o detalhe do erro JSON
		return
	}

	// chamando metodo do usecase para obter os dados de frequencia
	frequencias, err := f.frequenciaUseCase.GetFrequenciaByFilters(id_user, startDate, endDate)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}) // Retorna um erro de InternalServeError com o detalhe do erro JSON
		return
	}

	ctx.JSON(http.StatusOK, frequencias) // Retorna um Ok e as frequencias. Rodou File!

}

// Função que lida com a requisicao POST para criar novas entradas no BD referentes a um dia,
func (f *FrequenciaController) CreateFrequencia(ctx *gin.Context) {
	var frequencia model.Frequencia
	err := ctx.BindJSON(&frequencia)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	insertedFrequencia, err := f.frequenciaUseCase.CreateFrequencia(frequencia)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, insertedFrequencia)

}

// Função que lida com a requisicao PATCH para atualizar uma entrada ja existente no banco de dados
func (f *FrequenciaController) UpdateFrequencia(ctx *gin.Context) {
	var frequencia model.Frequencia
	err := ctx.BindJSON(&frequencia)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	updatedFreqeuncia, err := f.frequenciaUseCase.UpdateFrequencia(frequencia)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, updatedFreqeuncia)

}
