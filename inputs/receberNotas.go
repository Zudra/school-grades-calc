package inputs

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"school-grades-calc/models"
)

func ColetarNotas() (models.NotaEntrada, error) {
	scanner := bufio.NewReader(os.Stdin)

	fmt.Println("Bem-vindo ao sistema de coleta de notas!")

	var notas []float64

	for {
		fmt.Print("Digite a nota do aluno (ou 'fim' para encerrar): ")
		valorStr, _ := scanner.ReadString('\n')
		valorStr = strings.TrimSpace(valorStr)

		if strings.ToLower(valorStr) == "fim" {
			fmt.Println("Coleta de notas encerrada.")
			break
		}

		valor, err := strconv.ParseFloat(valorStr, 64)
		if err != nil {
			fmt.Println("Entrada inválida. Por favor, insira um número válido ou 'fim' para encerrar.")
			continue
		}

		if valor < 0 || valor > 10 {
			fmt.Println("Nota inválida. Por favor, insira um valor entre 0 e 10.")
			continue
		}

		notas = append(notas, valor)
	}

	if len(notas) == 0 {
		return models.NotaEntrada{}, errors.New("nenhuma nota foi inserida")
	}

	return models.NotaEntrada{Valores: notas}, nil
}
