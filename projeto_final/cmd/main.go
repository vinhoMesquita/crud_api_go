package main

import (
	"go-api/controller"
	"go-api/db"
	"go-api/repository"
	"go-api/usecase"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Inicializa o roteador do Gin
	server := gin.Default()

	// Configuração do CORS para permitir apenas a origem do seu frontend
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://127.0.0.1:5501"}
	server.Use(cors.New(config))

	// Conexão com o banco de dados
	dbConnection, err := db.ConectDB() // Ajuste o nome da função conforme a sua implementação
	if err != nil {
		panic(err)
	}

	// Inicialização das camadas da aplicação para Alunos
	AlunoRepository := repository.NewAlunoRepository(dbConnection)
	AlunosUsecase := usecase.NewAlunoUseCase(AlunoRepository)
	AlunoController := controller.NewAlunoController(AlunosUsecase)

	// Inicializar os repositórios
	userRepository := repository.NewUserRepository(dbConnection)
	frequenciaRepository := repository.NewFrequenciaRepository(dbConnection)
	// Inicializar os casos de uso
	userUseCase := usecase.NewUserUseCase(userRepository)
	frequenciaUseCase := usecase.NewFrequenciaCase(frequenciaRepository)
	// Inicializar os controladores
	userController := controller.NewUserController(userUseCase)
	frequenciaController := controller.NewFrequenciaController(frequenciaUseCase)
	server.Static("/static", "./web/static/")
	// // Carregar os arquivos estáticos
	// server.LoadHTMLGlob("web/dashboard.html")

	// // Rota para a página dashboard
	// server.GET("web/dashboard.html", func(ctx *gin.Context) {
	// 	ctx.HTML(200, "dashboard.html", nil)
	// })

	// Rotas da API para Alunos
	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "pong"})
	})
	server.GET("/alunos", AlunoController.GetAluno)
	server.POST("/alunos", AlunoController.CreateAluno)
	server.GET("/alunos/:alunoId", AlunoController.GetAlunoById)
	server.DELETE("/alunos/:alunoId", AlunoController.DeleteAluno)
	server.PUT("/alunos/:alunoId", AlunoController.UpdateAluno)

	// Rotas da API para Frequência
	server.GET("/user/", userController.GetUsersById)
	server.GET("/user/frequencia/", frequenciaController.GetFrequenciaByFilters)
	server.POST("/user/frequencia/", frequenciaController.CreateFrequencia)
	server.PATCH("/user/frequencia/update/", frequenciaController.UpdateFrequencia)

	// Inicia o servidor na porta 5501
	if err := server.Run(":5501"); err != nil {
		panic(err)
	}
}
