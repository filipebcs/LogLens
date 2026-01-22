package engine

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"time"

	"github.com/filipe/loglens/internal/event"
	"github.com/filipe/loglens/internal/finding"
	"github.com/filipe/loglens/internal/parser"
	"github.com/filipe/loglens/internal/report"
	"github.com/filipe/loglens/internal/rules"
)

func Run(logFile, rulesFile string, jsonOutput bool, reportFile string) {
	file, err := os.Open(logFile)
	if err != nil {
		fmt.Println("Erro ao abrir arquivo:", err)
		return
	}
	defer file.Close()

	// 🔹 Carregar regras
	ruleSet, err := rules.LoadRules(rulesFile)
	if err != nil {
		fmt.Println("Erro ao carregar regras:", err)
		return
	}

	scanner := bufio.NewScanner(file)

	stats := make(map[string]int)
	var events []event.Event
	lineNumber := 0

	baseTime := time.Now()

	for scanner.Scan() {
		lineNumber++

		e := event.Event{
			LineNumber: lineNumber,
			Raw:        scanner.Text(),
			Timestamp:  baseTime.Add(time.Duration(lineNumber-1) * time.Second),
		}

		parser.ParseLine(&e)
		events = append(events, e)

		if e.Level != "" {
			stats[e.Level]++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Erro durante leitura:", err)
		return
	}

	var findings []finding.Finding

	// 🔹 Regras por threshold
	for _, rule := range ruleSet {
		if rule.Level == "" || rule.Threshold == 0 {
			continue
		}

		count := stats[rule.Level]
		if count >= rule.Threshold {
			findings = append(findings, finding.Finding{
				RuleID:      rule.ID,
				Description: rule.Description,
				Severity:    rule.Severity,
				Level:       rule.Level,
				Count:       count,
				Threshold:   rule.Threshold,
			})
		}
	}

	// 🔹 Regras por regex
	for _, rule := range ruleSet {
		if rule.Pattern == "" {
			continue
		}

		re, err := regexp.Compile(rule.Pattern)
		if err != nil {
			fmt.Println("Regex inválida na regra:", rule.ID)
			continue
		}

		for _, e := range events {
			if re.MatchString(e.Message) {
				findings = append(findings, finding.Finding{
					RuleID:      rule.ID,
					Description: rule.Description,
					Severity:    rule.Severity,
					Level:       e.Level,
					LineNumber:  e.LineNumber,
					Message:     e.Message,
				})
			}
		}
	}

	// 🔹 Correlação temporal
	for _, rule := range ruleSet {
		if rule.Level == "" || rule.WindowSeconds == 0 || rule.Threshold == 0 {
			continue
		}

		var filtered []event.Event
		for _, e := range events {
			if e.Level == rule.Level {
				filtered = append(filtered, e)
			}
		}

		for i := 0; i < len(filtered); i++ {
			start := filtered[i].Timestamp
			count := 1

			for j := i + 1; j < len(filtered); j++ {
				if filtered[j].Timestamp.Sub(start) <= time.Duration(rule.WindowSeconds)*time.Second {
					count++
				} else {
					break
				}
			}

			if count >= rule.Threshold {
				findings = append(findings, finding.Finding{
					RuleID:      rule.ID,
					Description: rule.Description,
					Severity:    rule.Severity,
					Level:       rule.Level,
					Count:       count,
					Threshold:   rule.Threshold,
				})
				break
			}
		}
	}

	// 🔹 Deduplicação FINAL
	findings = deduplicateFindings(findings)

	// 🔹 Geração de relatório Markdown
	if reportFile != "" {
		err := report.GenerateMarkdown(
			reportFile,
			logFile,
			findings,
			len(events),
		)
		if err != nil {
			fmt.Println("Erro ao gerar relatório:", err)
			return
		}

		fmt.Println("Relatório gerado em:", reportFile)
	}

	// 🔹 Output
	if jsonOutput {
		out, err := json.MarshalIndent(findings, "", "  ")
		if err != nil {
			fmt.Println("Erro ao gerar JSON:", err)
			return
		}
		fmt.Println(string(out))
		return
	}

	fmt.Println("\nAchados:")
	if len(findings) == 0 {
		fmt.Println("Nenhum achado identificado")
	} else {
		for _, f := range findings {
			if f.Threshold > 0 {
				fmt.Printf(
					"- [%s] %s (regra=%s, nível=%s, contagem=%d, limite=%d)\n",
					f.Severity, f.Description, f.RuleID, f.Level, f.Count, f.Threshold,
				)
				continue
			}

			fmt.Printf(
				"- [%s] %s (regra=%s, nível=%s, linha=%d, mensagem=\"%s\")\n",
				f.Severity, f.Description, f.RuleID, f.Level, f.LineNumber, f.Message,
			)
		}
	}

	fmt.Printf("\nTotal de eventos analisados: %d\n", len(events))
}

func deduplicateFindings(findings []finding.Finding) []finding.Finding {
	dedup := make(map[string]bool)
	var unique []finding.Finding

	for _, f := range findings {
		if f.Threshold > 0 && f.LineNumber == 0 {
			if dedup[f.RuleID] {
				continue
			}
			dedup[f.RuleID] = true
			unique = append(unique, f)
			continue
		}
		unique = append(unique, f)
	}

	return unique
}