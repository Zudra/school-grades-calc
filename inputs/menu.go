package inputs

import (
	"bufio"
	"fmt"
	"os"
	"school-grades-calc/services"
	"school-grades-calc/utils"
	"strings"
)

func Menu() {
	scanner := bufio.NewReader(os.Stdin)

	for {
		utils.ExibirMenu()
		escolha, _ := scanner.ReadString('\n')
		escolha = strings.TrimSpace(escolha)
		escolha = strings.ToLower(escolha)

		switch escolha {
		case "1":
			notas, err := ColetarNotas()
			if err != nil {
				fmt.Println("Erro ao receber suas notas:", err)
				continue
			}

			resultado, err := services.CalcularMedia(notas)
			if err != nil {
				fmt.Println("Erro ao calcular e salvar sua média:", err)
				continue
			}

			utils.ExibirRetorno(resultado)

		case "2":
			err := services.ExibirHistorico()
			if err != nil {
				fmt.Println("Erro ao exibir histórico:", err)
			}

		case "3", "sair", "quit", "exit":
			fmt.Println("Saindo da calculadora de médias...")
			return

		default:
			fmt.Println("Opção inválida! Por favor, escolha uma opção válida.")
		}

		fmt.Println("\nPressione Enter para continuar...")
		scanner.ReadString('\n')
	}
}
