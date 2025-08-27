package main

import (
	"bytes"
	"strings"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestYAMLToJSON(t *testing.T) {
	conv := DefaultConverter{}

	yamlData := `
app:
  name: my-service
  version: 1.0
  database:
    host: localhost
    port: 5432
`
	jsonData, err := conv.YAMLToJSON([]byte(yamlData))
	if err != nil {
		t.Fatalf("YAMLToJSON failed: %v", err)
	}

	// Check if JSON contains expected keys
	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"name": "my-service"`) {
		t.Errorf("JSON output missing name field: %s", jsonStr)
	}
	if !strings.Contains(jsonStr, `"port": 5432`) {
		t.Errorf("JSON output missing port field: %s", jsonStr)
	}
}

func TestJSONToYAML(t *testing.T) {
	conv := DefaultConverter{}

	jsonData := []byte(`{
  "app": {
    "name": "my-service",
    "version": 1.0,
    "database": {
      "host": "localhost",
      "port": 5432
    }
  }
}`)

	yamlData, err := conv.JSONToYAML(jsonData)
	if err != nil {
		t.Fatalf("JSONToYAML failed: %v", err)
	}

	// Instead of string match, unmarshal YAML back
	var result map[string]interface{}
	if err := yaml.Unmarshal(yamlData, &result); err != nil {
		t.Fatalf("Failed to unmarshal YAML: %v", err)
	}

	app, ok := result["app"].(map[string]interface{})
	if !ok {
		t.Fatalf("app key missing or wrong type")
	}

	db, ok := app["database"].(map[string]interface{})
	if !ok {
		t.Fatalf("database key missing or wrong type")
	}

	if app["name"] != "my-service" {
		t.Errorf("Expected name=my-service, got %v", app["name"])
	}

	if db["port"] != "5432" && db["port"] != 5432 {
		t.Errorf("Expected port=5432, got %v", db["port"])
	}

	if app["version"] != "1.0" && app["version"] != 1.0 {
		t.Errorf("Expected version=1.0, got %v", app["version"])
	}
}

func TestYAMLToJSONReader(t *testing.T) {
	conv := DefaultConverter{}

	yamlData := `
app:
  name: my-service
  database:
    host: localhost
`
	r := bytes.NewReader([]byte(yamlData))
	var buf bytes.Buffer

	err := conv.YAMLToJSONReader(r, &buf)
	if err != nil {
		t.Fatalf("YAMLToJSONReader failed: %v", err)
	}

	if !strings.Contains(buf.String(), `"name": "my-service"`) {
		t.Errorf("JSON output missing name field: %s", buf.String())
	}
}

func TestJSONToYAMLReader(t *testing.T) {
	conv := DefaultConverter{}

	jsonData := `{"app": {"name": "my-service","database":{"host":"localhost"}}}`
	r := strings.NewReader(jsonData)
	var buf bytes.Buffer

	err := conv.JSONToYAMLReader(r, &buf)
	if err != nil {
		t.Fatalf("JSONToYAMLReader failed: %v", err)
	}

	if !strings.Contains(buf.String(), "name: my-service") {
		t.Errorf("YAML output missing name field: %s", buf.String())
	}
}
