package converter

import "io"

// Converter defines the contract for YAML ↔ JSON conversions.
type Converter interface {
	YAMLToJSON(yamlData []byte) ([]byte, error)
	JSONToYAML(jsonData []byte) ([]byte, error)
	YAMLToJSONReader(r io.Reader, w io.Writer) error
	JSONToYAMLReader(r io.Reader, w io.Writer) error
}
