package main

import (
	"fmt"
	"log"
	"os"

	converter "github.com/arshad404/data_format_converter"
)

func main() {
	conv := converter.DefaultConverter{}

	yamlData := []byte(`
app:
  name: my-service
  version: 1.0
  database:
    host: localhost
    port: 5432
`)

	// YAML → JSON
	jsonData, err := conv.YAMLToJSON(yamlData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("YAML → JSON:\n", string(jsonData))

	// JSON → YAML
	backToYAML, err := conv.JSONToYAML(jsonData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\nJSON → YAML:\n", string(backToYAML))

	// Example 1: Convert YAML file → JSON file using streaming
	yamlFile, err := os.Open("config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	defer yamlFile.Close()

	jsonFile, err := os.Create("config.json")
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	if err := conv.YAMLToJSONReader(yamlFile, jsonFile); err != nil {
		log.Fatal("YAML → JSON failed:", err)
	}
	log.Println("Converted config.yaml → config.json")
}
