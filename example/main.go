package main

import (
	"fmt"
	"log"

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
}
