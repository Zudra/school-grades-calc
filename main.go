package main

import (
	"fmt"

	"school-grades-calc/inputs"
	"school-grades-calc/services"
	"school-grades-calc/utils"
)

func main() {
	notas, err := inputs.ColetarNotas()
	if err != nil {
		fmt.Println("Ocorreu um erro ao receber suas notas: ", err)
		return
	}

	resultado := services.CalcularMedia(notas)

	utils.ExibirRetorno(resultado)
}
