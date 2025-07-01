package services

import "school-grades-calc/models"

func CalcularMedia(entrada models.NotaEntrada) models.ResultadoMedia {
	var valorTotal float64

	for _, valor := range entrada.Valores {
		valorTotal += valor
	}

	media := valorTotal / float64(len(entrada.Valores))

	media = round(media)

	resultado := ClassificarAprovacao(entrada.Valores, media)

	return resultado
}

func ClassificarAprovacao(notas []float64, media float64) models.ResultadoMedia {
	switch {
	case media < 5:
		return models.ResultadoMedia{Notas: notas, Media: media, Status: "Reprovado"}
	case media >= 5 && media < 7:
		return models.ResultadoMedia{Notas: notas, Media: media, Status: "Recuperação"}
	case media >= 7:
		return models.ResultadoMedia{Notas: notas, Media: media, Status: "Aprovado"}
	default:
		return models.ResultadoMedia{}
	}
}

func round(x float64) float64 {
	return float64(int(x*100)) / 100
}
