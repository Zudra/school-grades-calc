package utils

import (
	"fmt"
	"school-grades-calc/models"
	"strings"
)

func ExibirMenu() {
	fmt.Print("\033[H\033[2J")

	fmt.Println(strings.Repeat("=", 60))
	fmt.Println("🎓 CALCULADORA DE MÉDIAS ESCOLARES")
	fmt.Println(strings.Repeat("=", 60))

	fmt.Println("📚 Escolha uma opção:")
	fmt.Println()
	fmt.Println("  1️⃣  Calcular nova média")
	fmt.Println("  2️⃣  Ver histórico de médias")
	fmt.Println("  3️⃣  Sair")
	fmt.Println()
	fmt.Println(strings.Repeat("-", 60))
	fmt.Print("👉 Digite sua escolha: ")
}

func ExibirRetorno(saida models.ResultadoMedia) {
	fmt.Printf("\nMédia das suas notas:\n")
	fmt.Printf("%.2f\n\n", saida.Media)
	fmt.Printf("Status de aprovação:\n")
	fmt.Printf("%s\n\n", saida.Status)
}
