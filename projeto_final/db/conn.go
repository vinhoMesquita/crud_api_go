package db

//arquivo responsavel por criar a conexão com o banco de dados

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 1234
	user     = "postgres"
	password = "post123"
	dbname   = "postgres"
)

// criar a conexão de fato usando as constantes que definimos
func ConectDB() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host= %s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo) //abrindo conexão com o postgres
	if err != nil {
		panic(err)
	}

	err = db.Ping() //fazendo um ping para ver se a conexão está abrindo com sucesso
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to " + dbname)

	return db, nil //nossa conexão com o banco
}
