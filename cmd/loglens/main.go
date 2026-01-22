package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/filipe/loglens/internal/engine"
)

func main() {
	logFile := flag.String("f", "", "Arquivo de log para análise")
	rulesFile := flag.String("r", "configs/rules.yaml", "Arquivo de regras")
	jsonOutput := flag.Bool("json", false, "Saída em JSON")
	reportFile := flag.String("report", "", "Gerar relatório Markdown")
	flag.Parse()

	if *logFile == "" {
		fmt.Println("Uso: loglens -f <arquivo.log> [-r regras.yaml] [--json]")
		os.Exit(1)
	}

	engine.Run(*logFile, *rulesFile, *jsonOutput, *reportFile)
}