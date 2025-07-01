package tests

import (
	"school-grades-calc/models"
	"school-grades-calc/services"
	"testing"
)

func testMediaService(t *testing.T) {
	casos := []struct {
		notas          models.NotaEntrada
		mediaEsperada  float64
		statusEsperado string
	}{
		{models.NotaEntrada{Valores: []float64{10, 9, 8}}, 9.0, "Aprovado"},
		{models.NotaEntrada{Valores: []float64{7, 6, 6}}, 6.34, "Recuperação"},
		{models.NotaEntrada{Valores: []float64{3, 4, 2}}, 3.0, "Reprovado"},
	}

	for _, c := range casos {
		resultado := services.CalcularMedia(c.notas)

		if c.mediaEsperada != resultado.Media {
			t.Errorf("Média esperada: %.2f, obtido: %.2f", c.mediaEsperada, resultado.Media)
		}

		if c.statusEsperado != resultado.Status {
			t.Errorf("Status esperado: %s, obtido: %s", c.statusEsperado, resultado.Status)
		}
	}
}
