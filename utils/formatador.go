package utils

import (
	"fmt"
	"school-grades-calc/models"
)

func ExibirRetorno(saida models.ResultadoMedia) {
	fmt.Printf("\nMédia das suas notas:\n")
	fmt.Printf("%.2f\n\n", saida.Media)
	fmt.Printf("Status de aprovação:\n")
	fmt.Printf("%s\n\n", saida.Status)
}
