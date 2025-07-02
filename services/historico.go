package services

import (
	"encoding/json"
	"fmt"
	"os"
	"school-grades-calc/models"
	"strings"
)

func ExibirHistorico() error {
	arquivo, err := os.ReadFile("./data/historico.json")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("📋 Nenhum histórico encontrado ainda.")
			return nil
		}
		return fmt.Errorf("erro ao ler arquivo de histórico: %v", err)
	}

	var historico []models.ResultadoMedia
	if err := json.Unmarshal(arquivo, &historico); err != nil {
		return fmt.Errorf("erro ao decodificar histórico: %v", err)
	}

	if len(historico) == 0 {
		fmt.Println("📋 Histórico vazio.")
		return nil
	}

	fmt.Println("\n" + strings.Repeat("=", 80))
	fmt.Println("📊 HISTÓRICO DE MÉDIAS")
	fmt.Println(strings.Repeat("=", 80))

	for i := len(historico) - 1; i >= 0; i-- {
		registro := historico[i]

		fmt.Printf("📝 Notas: ")

		notasStr := make([]string, len(registro.Notas))
		for j, nota := range registro.Notas {
			notasStr[j] = fmt.Sprintf("%.1f", nota)
		}
		fmt.Printf("[%s]\n", strings.Join(notasStr, ", "))

		fmt.Printf("🎯 Média: %.2f\n", registro.Media)

		var statusEmoji string
		switch strings.ToLower(registro.Status) {
		case "aprovado":
			statusEmoji = "✅ APROVADO"
		case "reprovado":
			statusEmoji = "❌ REPROVADO"
		case "recuperação":
			statusEmoji = "⚠️  RECUPERAÇÃO"
		default:
			statusEmoji = "❓ " + registro.Status
		}
		fmt.Printf("📊 Status: %s\n", statusEmoji)

		fmt.Println(strings.Repeat("-", 50))
	}

	exibirEstatisticas(historico)

	return nil
}

func exibirEstatisticas(historico []models.ResultadoMedia) {
	if len(historico) == 0 {
		return
	}

	var somaMedias float64
	aprovados := 0
	reprovados := 0
	recuperacao := 0

	for _, registro := range historico {
		somaMedias += registro.Media

		switch strings.ToLower(registro.Status) {
		case "aprovado":
			aprovados++
		case "reprovado":
			reprovados++
		case "recuperação":
			recuperacao++
		}
	}

	mediaGeral := somaMedias / float64(len(historico))

	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("📈 ESTATÍSTICAS GERAIS")
	fmt.Println(strings.Repeat("=", 50))
	fmt.Printf("🔢 Total de registros: %d\n", len(historico))
	fmt.Printf("📊 Média geral: %.2f\n", mediaGeral)
	fmt.Printf("✅ Aprovações: %d\n", aprovados)
	fmt.Printf("❌ Reprovações: %d\n", reprovados)
	fmt.Printf("⚠️  Recuperações: %d\n", recuperacao)

	if len(historico) > 0 {
		porcentagemAprovacao := float64(aprovados) / float64(len(historico)) * 100
		fmt.Printf("📈 Taxa de aprovação: %.1f%%\n", porcentagemAprovacao)
	}

	fmt.Println(strings.Repeat("=", 50))
}
