package src

import (
	"encoding/json"

	"gopkg.in/yaml.v3"
)

// Checks whether the file matches the file format that is expected
func ValidateContent(text, filetype string) bool {
	if text == "OK" || text == "" {
		return false
	}
	if filetype == "yaml" {
		return validateYaml(text)
	} else if filetype == "json" {
		return validateJson(text)
	}
	return true
}

// Validates whether the given text is a valid YAML
func validateYaml(text string) bool {
	yamlMap := make(map[interface{}]interface{})
	err := yaml.Unmarshal([]byte(text), &yamlMap)
	return err == nil
}

// Checks whether the given text is a valid JSON
func validateJson(text string) bool {
	jsonMap := make(map[string]interface{})
	err := json.Unmarshal([]byte(text), &jsonMap)
	return err == nil
}
