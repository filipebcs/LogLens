package report

import (
	"fmt"
	"os"
	"time"

	"github.com/filipe/loglens/internal/finding"
)

func GenerateMarkdown(
	outputPath string,
	logFile string,
	findings []finding.Finding,
	totalEvents int,
) error {

	f, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer f.Close()

	now := time.Now().Format("2006-01-02 15:04:05")

	// 🔹 Cabeçalho
	fmt.Fprintf(f, "# Relatório de Análise de Logs\n\n")
	fmt.Fprintf(f, "**Arquivo analisado:** `%s`\n\n", logFile)
	fmt.Fprintf(f, "**Data da análise:** %s\n\n", now)
	fmt.Fprintf(f, "**Total de eventos analisados:** %d\n\n", totalEvents)

	// 🔹 Resumo executivo
	fmt.Fprintf(f, "## Resumo Executivo\n\n")
	if len(findings) == 0 {
		fmt.Fprintf(f, "Nenhum achado relevante foi identificado durante a análise.\n\n")
	} else {
		fmt.Fprintf(
			f,
			"Foram identificados **%d achados relevantes**, conforme detalhado a seguir.\n\n",
			len(findings),
		)
	}

	// 🔹 Achados
	fmt.Fprintf(f, "## Achados\n\n")

	if len(findings) == 0 {
		fmt.Fprintf(f, "_Nenhum achado identificado._\n")
		return nil
	}

	for i, fd := range findings {
		fmt.Fprintf(f, "### Achado %d\n\n", i+1)
		fmt.Fprintf(f, "- **Regra:** `%s`\n", fd.RuleID)
		fmt.Fprintf(f, "- **Descrição:** %s\n", fd.Description)
		fmt.Fprintf(f, "- **Severidade:** %s\n", fd.Severity)

		if fd.Threshold > 0 {
			fmt.Fprintf(f, "- **Tipo:** Agregado\n")
			fmt.Fprintf(f, "- **Nível:** %s\n", fd.Level)
			fmt.Fprintf(f, "- **Contagem:** %d\n", fd.Count)
			fmt.Fprintf(f, "- **Limite:** %d\n", fd.Threshold)
		} else {
			fmt.Fprintf(f, "- **Tipo:** Evidência pontual\n")
			fmt.Fprintf(f, "- **Nível:** %s\n", fd.Level)
			fmt.Fprintf(f, "- **Linha:** %d\n", fd.LineNumber)
			fmt.Fprintf(f, "- **Mensagem:** `%s`\n", fd.Message)
		}

		fmt.Fprintf(f, "\n")
	}

	return nil
}