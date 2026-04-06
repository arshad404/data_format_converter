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
go get github.com/arshad404/data_format_converter
# reviewme test
trigger review Mon 06 Apr 2026 06:32:52 PM IST
summary test Mon 06 Apr 2026 06:37:22 PM IST
verbose log test Mon 06 Apr 2026 06:39:38 PM IST
