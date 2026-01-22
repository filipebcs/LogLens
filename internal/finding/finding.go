package finding

type Finding struct {
	RuleID      string `json:"rule_id"`
	Description string `json:"description"`
	Severity    string `json:"severity"`

	Level     string `json:"level,omitempty"`
	Count     int    `json:"count,omitempty"`
	Threshold int    `json:"threshold,omitempty"`

	// 🔹 Novo
	LineNumber int    `json:"line_number,omitempty"`
	Message    string `json:"message,omitempty"`
}