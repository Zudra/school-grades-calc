package tests

import (
	"school-grades-calc/models"
	"school-grades-calc/services"
	"testing"
)

func TestMediaService(t *testing.T) {
	casos := []struct {
		notas          models.NotaEntrada
		mediaEsperada  float64
		statusEsperado string
		esperaErro     bool
	}{
		{models.NotaEntrada{Valores: []float64{10, 9, 8}}, 9.0, "Aprovado", false},
		{models.NotaEntrada{Valores: []float64{7, 6, 6}}, 6.33, "Recuperação", false},
		{models.NotaEntrada{Valores: []float64{3, 4, 2}}, 3.0, "Reprovado", false},
	}

	for _, c := range casos {
		resultado, err := services.CalcularMedia(c.notas)

		if c.esperaErro && err == nil {
			t.Errorf("esperava erro, mas não houve para a entrada: %+v", c.notas)
		}

		if !c.esperaErro && err != nil {
			t.Errorf("erro inesperado para a entrada: %+v: %v", c.notas, err)
		}

		if !c.esperaErro {
			if c.mediaEsperada != resultado.Media {
				t.Errorf("Média esperada: %.2f, obtido: %.2f", c.mediaEsperada, resultado.Media)
			}

			if c.statusEsperado != resultado.Status {
				t.Errorf("Status esperado: %s, obtido: %s", c.statusEsperado, resultado.Status)
			}
		}
	}
}
