package src

import (
	"testing"
)

func TestValidateYaml(t *testing.T) {
	// Test valid YAML
	validYaml := "name: John Doe"
	if !validateYaml(validYaml) {
		t.Errorf("'%s' is a valid YAML, but it was not recognized as such", validYaml)
	}

	// Test invalid YAML
	invalidYaml := "name: John Doe:"
	if validateYaml(invalidYaml) {
		t.Errorf("'%s' is an invalid YAML, but it was recognized as such", invalidYaml)
	}
}

func TestValidateJson(t *testing.T) {
	// Test valid JSON
	validJson := `{"name": "John Doe"}`
	if !validateJson(validJson) {
		t.Errorf("'%s' is a valid JSON, but it was not recognized as such", validJson)
	}

	// Test invalid JSON
	invalidJson := `<html>Foo</html>"`
	if validateJson(invalidJson) {
		t.Errorf("'%s' is an invalid JSON, but it was recognized as such", invalidJson)
	}
}
