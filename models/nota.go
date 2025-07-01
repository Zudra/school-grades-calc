package models

type NotaEntrada struct {
	Valores []float64 `json:"valores"`
}

type NotaMedia struct {
	Media  float64 `json:"media"`
	Status string  `json:"status"`
}
