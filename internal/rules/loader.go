package rules

import (
	"os"

	"gopkg.in/yaml.v3"
)

func LoadRules(path string) ([]Rule, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var rules []Rule
	if err := yaml.Unmarshal(data, &rules); err != nil {
		return nil, err
	}

	return rules, nil
}