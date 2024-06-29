package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"go-api/model"
)

type FrequenciaRepository struct {
	connection *sql.DB
}

func NewFrequenciaRepository(connection *sql.DB) FrequenciaRepository {
	return FrequenciaRepository{
		connection: connection,
	}
}

func (fr *FrequenciaRepository) GetFrequenciaByFilters(id_user int, startDate, endDate string) ([]model.Frequencia, error) {

	// Consulta SQL base
	query := "SELECT * FROM frequencias WHERE ID_user = $1"
	args := []interface{}{id_user} // Slice de interfaces para modar a consulta com base nos parametros passados
	argIndex := 2                  // Definindo o o numero do proximo placeholder da consulta sql

	// Condições adicionais baseadas nos parâmetros fornecidos
	if startDate != "" && endDate != "" {
		query += fmt.Sprintf(" AND data BETWEEN $%d AND $%d", argIndex, argIndex+1)
		args = append(args, startDate, endDate)
	} else if startDate != "" {
		query += fmt.Sprintf(" AND data >= $%d", argIndex)
		args = append(args, startDate)
	} else if endDate != "" {
		query += fmt.Sprintf(" AND data <= $%d", argIndex)
		args = append(args, endDate)
	}

	// Preparar a instrução SQL
	stmt, err := fr.connection.Prepare(query)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer stmt.Close()

	// Executar a consulta SQL
	rows, err := stmt.Query(args...)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	var listaFrequencia []model.Frequencia

	// Iterar sobre os resultados da consulta ao BD
	for rows.Next() {
		var frequenciaObj model.Frequencia
		err := rows.Scan( // Copiando os valores das colunas para os vampos da variavel frequenciasObj
			&frequenciaObj.T1,
			&frequenciaObj.T2,
			&frequenciaObj.T3,
			&frequenciaObj.Id_user,
			&frequenciaObj.Data)

		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		frequenciaObj.DiaSemana = frequenciaObj.Data.Weekday().String()

		listaFrequencia = append(listaFrequencia, frequenciaObj)
	}

	// Verificar erros durante a iteração
	if err = rows.Err(); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return listaFrequencia, nil
}

func (fr *FrequenciaRepository) CreateFrequencia(frequencia model.Frequencia) (int, error) {
	var id int

	query, err := fr.connection.Prepare("INSERT INTO frequencias (t1, t2, t3,id_user,data) VALUES ($1, $2, $3, $4, $5) RETURNING id_user")
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	// Execucao da query com os valoes vindos do parametros passado em  frequencia, escaneando e retorna o valor id_user
	err = query.QueryRow(frequencia.T1, frequencia.T2, frequencia.T3, frequencia.Id_user, frequencia.Data).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	defer query.Close()

	return id, nil
}

func (fr *FrequenciaRepository) UpdateFrequencia(frequencia model.Frequencia) (int, error) {

	query, err := fr.connection.Prepare("UPDATE frequencias SET t1 = $1, t2 = $2,  t3= $3 WHERE id_user = $4 AND data = $5")
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	defer query.Close()

	// Execute a consulta com os valores fornecidos
	result, err := query.Exec(frequencia.T1, frequencia.T2, frequencia.T3, frequencia.Id_user, frequencia.Data)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	// Obtenha o número de linhas afetadas pela operação de atualização
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	// Verifique se alguma linha foi afetada pela atualização
	if rowsAffected == 0 {
		return 0, errors.New("Nenhuma linha foi atualizada")
	}

	// Se uma linha foi atualizada com sucesso, retorne o ID do usuário
	return frequencia.Id_user, nil
}
