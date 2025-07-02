package storage

import (
	"encoding/json"
	"os"
	"path/filepath"

	"school-grades-calc/models"
)

const arquivoHistorico = "historico.json"

func SalvarHistorico(resultado models.ResultadoMedia) error {

	pastaDados := os.Getenv("DATA_PATH")

	if err := os.MkdirAll(pastaDados, os.ModePerm); err != nil {
		return err
	}

	path := filepath.Join(pastaDados, arquivoHistorico)

	var historico []models.ResultadoMedia

	if dadosAntigos, err := os.ReadFile(path); err == nil {
		_ = json.Unmarshal(dadosAntigos, &historico)
	}

	historico = append(historico, resultado)

	novoConteudo, err := json.Marshal(historico)
	if err != nil {
		return err
	}

	return os.WriteFile(path, novoConteudo, os.ModePerm)
}
