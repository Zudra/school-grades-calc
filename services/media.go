package services

import "school-grades-calc/models"

func CalcularMedia(entrada models.NotaEntrada) models.NotaMedia {
	var valorTotal float64

	for _, valor := range entrada.Valores {
		valorTotal += valor
	}

	media := valorTotal / float64(len(entrada.Valores))

	media = round(media)

	return ClassificarAprovacao(media)
}

func ClassificarAprovacao(media float64) models.NotaMedia {
	switch {
	case media < 5:
		return models.NotaMedia{Media: media, Status: "Reprovado"}
	case media >= 5 && media < 7:
		return models.NotaMedia{Media: media, Status: "Recuperação"}
	case media >= 7:
		return models.NotaMedia{Media: media, Status: "Aprovado"}
	default:
		return models.NotaMedia{}
	}
}

func round(x float64) float64 {
	return float64(int(x*100)) / 100
}
