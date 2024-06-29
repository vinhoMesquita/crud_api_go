package repository

import (
	"database/sql"
	"fmt"
	"go-api/model"
)

// Estrutura UserRepository que contém o connection que é um ponteiro de um sql.DB
// fornecendo uma maneira de lidar com as conexoes com o banco de dados SQL
type UserRepository struct {
	connection *sql.DB
}

// Funcao de instanciacao de UserRepository
func NewUserRepository(connection *sql.DB) UserRepository {
	return UserRepository{
		connection: connection,
	}
}

// funcao nao utilizada na pagina de graficos
func (us *UserRepository) GetUsers() ([]model.User, error) {

	query := "SELECT id, nome, senha FROM users"
	rows, err := us.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.User{}, err
	}

	var listaUser []model.User
	var userObj model.User

	for rows.Next() {
		err = rows.Scan(
			&userObj.Id,
			&userObj.Nome,
			&userObj.Senha)

		if err != nil {
			fmt.Println(err)
			return []model.User{}, err
		}

		listaUser = append(listaUser, userObj)

	}

	rows.Close()

	return listaUser, nil

}

// funcao de criacao de ussuario nao utilizada
func (us *UserRepository) CreateUser(user model.User) (int, error) {

	var id int

	query, err := us.connection.Prepare("INSERT INTO users" +
		"(nome, senha)" +
		"Values ($1, $2) RETURNING id")
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	err = query.QueryRow(user.Nome, user.Senha).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	query.Close()
	return id, nil
}

// Funcao para obter dados do usuario a partir do ID
func (us *UserRepository) GetUserById(id int) (*model.User, error) {
	query, err := us.connection.Prepare("SELECT * FROM users WHERE ID = $1")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var user model.User

	err = query.QueryRow(id).Scan(
		&user.Id,
		&user.Nome,
		&user.Senha)

	if err != nil {

		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	query.Close()
	return &user, nil

}
