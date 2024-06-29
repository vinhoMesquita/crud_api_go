package model

type User struct {
	Id    int    `json:"id"`
	Nome  string `json:"nome"`
	Senha string `json:"senha"`
}

type Aluno struct {
	ID          int     `json:"id_aluno"`
	Name        string  `json:"name_aluno"`
	Age         int     `json:"age_aluno"`
	Body_fat    float64 `json:"bf_aluno"`
	Muscle_mass float64 `json:"mm_aluno"`
	Altura      float64 `json:"altura_aluno"`
	Peso        float64 `json:"peso_aluno"`
	// Imc         float64 `json:"imc_aluno"`
}
