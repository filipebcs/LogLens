package rules

type Rule struct {
	ID          string `yaml:"id"`
	Description string `yaml:"description"`
	Severity    string `yaml:"severity"`

	Level          string `yaml:"level,omitempty"`
	Threshold      int    `yaml:"threshold,omitempty"`
	Pattern        string `yaml:"pattern,omitempty"`
	WindowSeconds  int    `yaml:"window_seconds,omitempty"`
}