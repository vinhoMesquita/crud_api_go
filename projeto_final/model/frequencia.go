package model

import (
	"encoding/json"
	"time"
)

type Frequencia struct {
	T1        int       `json:"turno manha"`
	T2        int       `json:"turno tarde"`
	T3        int       `json:"turno noite"`
	Id_user   int       `json:"id_User"`
	Data      time.Time `json:"data"`
	DiaSemana string    `json:"dia_semana"`
}

// Implementa a desserialização personalizada para a estrutura Frequencia
func (f *Frequencia) UnmarshalJSON(data []byte) error {
	// Define um alias para evitar recursão infinita durante o unmarshal
	type Alias Frequencia

	// Struct auxiliar para capturar o campo "data" do JSON e manter os outros campos de Frequencia
	aux := &struct {
		Data string `json:"data"`
		*Alias
	}{
		Alias: (*Alias)(f), // Converte a estrutura Frequencia para o Alias
	}

	// Faz o unmarshal do JSON recebido para a estrutura auxiliar
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	// Faz o parsing da string da data para o formato time.Time
	parsedData, err := time.Parse("2006-01-02", aux.Data)
	if err != nil {
		return err
	}
	// Atribui a data parseada à estrutura Frequencia
	f.Data = parsedData
	return nil // Retorna nil indicando que a operação foi bem-sucedida
}
