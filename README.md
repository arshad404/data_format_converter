# YAML ↔ JSON Converter (Go)

A lightweight Go library for converting **nested YAML ↔ JSON** with support for
in-memory, file, and streaming (`io.Reader`/`io.Writer`) conversions.

## Features
- Convert YAML → JSON and JSON → YAML
- Works with deeply nested structures
- Safe: ensures JSON keys are always strings
- Provides both `[]byte` and `io.Reader/io.Writer` APIs
- Ready-to-use `DefaultConverter` implementation

## Install
```bash
go get github.com/yourname/yamljson-converter
