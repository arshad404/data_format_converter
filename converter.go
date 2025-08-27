package converter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	"gopkg.in/yaml.v3"
)

// DefaultConverter is a ready-to-use implementation of Converter
type DefaultConverter struct{}

// YAMLToJSON converts YAML bytes into JSON bytes
func (DefaultConverter) YAMLToJSON(yamlData []byte) ([]byte, error) {
	var obj interface{}
	if err := yaml.Unmarshal(yamlData, &obj); err != nil {
		return nil, fmt.Errorf("yaml unmarshal failed: %w", err)
	}
	obj = cleanMap(obj)
	jsonData, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("json marshal failed: %w", err)
	}
	return jsonData, nil
}

// JSONToYAML converts JSON bytes into YAML bytes
func (DefaultConverter) JSONToYAML(jsonData []byte) ([]byte, error) {
	var obj interface{}
	dec := json.NewDecoder(bytes.NewReader(jsonData))
	dec.UseNumber()
	if err := dec.Decode(&obj); err != nil {
		return nil, fmt.Errorf("json unmarshal failed: %w", err)
	}
	yamlData, err := yaml.Marshal(obj)
	if err != nil {
		return nil, fmt.Errorf("yaml marshal failed: %w", err)
	}
	return yamlData, nil
}

// YAMLToJSONReader reads YAML from r and writes JSON to w
func (dc DefaultConverter) YAMLToJSONReader(r io.Reader, w io.Writer) error {
	yamlData, err := io.ReadAll(r)
	if err != nil {
		return fmt.Errorf("read yaml input: %w", err)
	}
	jsonData, err := dc.YAMLToJSON(yamlData)
	if err != nil {
		return err
	}
	_, err = w.Write(jsonData)
	return err
}

// JSONToYAMLReader reads JSON from r and writes YAML to w
func (dc DefaultConverter) JSONToYAMLReader(r io.Reader, w io.Writer) error {
	jsonData, err := io.ReadAll(r)
	if err != nil {
		return fmt.Errorf("read json input: %w", err)
	}
	yamlData, err := dc.JSONToYAML(jsonData)
	if err != nil {
		return err
	}
	_, err = w.Write(yamlData)
	return err
}

// cleanMap ensures all keys are strings (needed for JSON)
func cleanMap(i interface{}) interface{} {
	switch x := i.(type) {
	case map[interface{}]interface{}:
		m2 := make(map[string]interface{})
		for k, v := range x {
			m2[fmt.Sprintf("%v", k)] = cleanMap(v)
		}
		return m2
	case []interface{}:
		for i, v := range x {
			x[i] = cleanMap(v)
		}
	}
	return i
}
