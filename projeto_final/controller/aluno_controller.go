package controller

import (
	"go-api/model"
	"go-api/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//nesse arquivo vou receber a resquisição, tratar a entrada da requisição
//e retornar o que quero

//definir estrutura para o controller

// estrutura de controle vai ter um usecase
type alunoController struct {
	alunosUsecase usecase.AlunosUsecase
}

// função para inicializar essa struct recebe como parametro
func NewAlunoController(usecase usecase.AlunosUsecase) alunoController {
	return alunoController{
		alunosUsecase: usecase,
	} //para injetar o alunoController no usacase
}

// função que vai tratar a requisição
func (p *alunoController) GetAluno(ctx *gin.Context) {
	//mapear a rota de get aluno
	alunos, err := p.alunosUsecase.GetAluno()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, alunos)
}

// funcao para cirar o aluno no controller
// método post da api
func (a *alunoController) CreateAluno(ctx *gin.Context) {
	//receber no body da requisição e fazer o json virar a estrutura aluno

	var aluno model.Aluno

	err := ctx.BindJSON(&aluno) //essa função que vai pegar o json e transformar no aluno

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	//se não der erro, chama o useCase, que por sua vez, chama o repository
	// que inseri os dados no banco de dados
	insertedAluno, err := a.alunosUsecase.CreateAluno(aluno)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	// se não deu erro, significa que o produto foi criado
	ctx.JSON(http.StatusCreated, insertedAluno)
}

func (a *alunoController) GetAlunoById(ctx *gin.Context) {

	id := ctx.Param("alunoId")

	if id == "" {
		response := model.Response{
			Message: "Id não pode ser nulo",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	alunoId, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response{
			Message: "Id precisa ser um número",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	aluno, err := a.alunosUsecase.GetAlunoById(alunoId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if aluno == nil {
		response := model.Response{
			Message: "Aluno não encontrado na base de dados",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}
	ctx.JSON(http.StatusOK, aluno)
}

func (a *alunoController) DeleteAluno(ctx *gin.Context) {
	id := ctx.Param("alunoId")

	if id == "" {
		response := model.Response{
			Message: "Id não pode ser nulo",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	alunoId, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response{
			Message: "Id precisa ser um número",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	err = a.alunosUsecase.DeleteAluno(alunoId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	response := model.Response{
		Message: "Aluno deletado com sucesso",
	}
	ctx.JSON(http.StatusOK, response)
}

func (a *alunoController) UpdateAluno(ctx *gin.Context) {
	// Capturar o ID do aluno da URL
	id := ctx.Param("alunoId")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID do aluno não especificado"})
		return
	}

	alunoID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID do aluno inválido"})
		return
	}

	// Bind dos dados JSON para o objeto Aluno
	var aluno model.Aluno
	if err := ctx.BindJSON(&aluno); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Verificar se o ID no JSON corresponde ao ID na URL
	if aluno.ID != alunoID {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "IDs de aluno na URL e no corpo da requisição não correspondem"})
		return
	}

	// Chamar o caso de uso para atualizar o aluno
	if err := a.alunosUsecase.UpdateAluno(aluno); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Responder com o aluno atualizado
	ctx.JSON(http.StatusOK, aluno)
}
