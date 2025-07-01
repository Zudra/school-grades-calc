package models

type NotaEntrada struct {
	Valores []float64 `json:"valores"`
}

type ResultadoMedia struct {
	Notas  []float64 `json:"notas"`
	Media  float64   `json:"media"`
	Status string    `json:"status"`
}
