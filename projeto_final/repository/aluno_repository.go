package repository

import (
	"database/sql"
	"go-api/model"
)

// essa struct vai ser a conexão com o banco de dados
type AlunoRepository struct {
	connection *sql.DB
}

// funcoa de inicialização
func NewAlunoRepository(connection *sql.DB) AlunoRepository {
	return AlunoRepository{
		connection: connection,
	}
}

// função que inicializa no banco
func (ar *AlunoRepository) GetAluno() ([]model.Aluno, error) {
	query := "SELECT id_aluno, nome, age, body_fat, muscle_mass, altura, peso FROM aluno"
	rows, err := ar.connection.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var alunos []model.Aluno
	for rows.Next() {
		var aluno model.Aluno
		if err := rows.Scan(&aluno.ID, &aluno.Name, &aluno.Age, &aluno.Body_fat, &aluno.Muscle_mass, &aluno.Altura, &aluno.Peso); err != nil {
			return nil, err
		}
		// &aluno.Imc = aluno.Peso / (aluno.Altura * aluno.Altura)

		alunos = append(alunos, aluno)
	}

	return alunos, nil
}

func (ar *AlunoRepository) CreateAluno(aluno model.Aluno) (int, error) {
	query := `
		INSERT INTO aluno (nome, age, body_fat, muscle_mass, altura, peso)
		VALUES ($1, $2, $3, $4, $5, $6) RETURNING id_aluno
	`
	stmt, err := ar.connection.Prepare(query)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	var id int
	err = stmt.QueryRow(aluno.Name, aluno.Age, aluno.Body_fat, aluno.Muscle_mass, aluno.Altura, aluno.Peso).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (ar *AlunoRepository) GetAlunoById(id int) (*model.Aluno, error) {
	query := "SELECT id_aluno, nome, age, body_fat, muscle_mass, altura, peso  FROM aluno WHERE id_aluno = $1"
	row := ar.connection.QueryRow(query, id)

	var aluno model.Aluno
	err := row.Scan(&aluno.ID, &aluno.Name, &aluno.Age, &aluno.Body_fat, &aluno.Muscle_mass, &aluno.Altura, &aluno.Peso)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &aluno, nil
}

func (ar *AlunoRepository) DeleteAluno(id int) error {
	query := "DELETE FROM aluno WHERE id_aluno = $1"
	_, err := ar.connection.Exec(query, id)
	return err
}

func (ar *AlunoRepository) UpdateAluno(aluno model.Aluno) error {
	query := `
		UPDATE aluno SET nome = $2, age = $3, body_fat = $4, muscle_mass = $5, altura = $6, peso = $7
		WHERE id_aluno = $1
	`
	_, err := ar.connection.Exec(query, aluno.ID, aluno.Name, aluno.Age, aluno.Body_fat, aluno.Muscle_mass, aluno.Altura, aluno.Peso)
	return err
}
