package renderer

import (
	"testing"

	"gopkg.in/yaml.v3"
)

func TestNewRenderer(t *testing.T) {
	tests := []struct {
		name       string
		output     string
		expectType string
	}{
		{
			name:       "json renderer",
			output:     "json",
			expectType: "*renderer.jsonRenderer",
		},
		{
			name:       "yaml renderer",
			output:     "yaml",
			expectType: "*renderer.yamlRenderer",
		},
		{
			name:       "yml renderer",
			output:     "yml",
			expectType: "*renderer.yamlRenderer",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			renderer := NewRenderer(tt.output)
			if renderer == nil {
				t.Errorf("NewRenderer(%s) returned nil", tt.output)
			}
		})
	}
}

func TestIsRenderer(t *testing.T) {
	tests := []struct {
		name     string
		output   string
		expected bool
	}{
		{"json", "json", true},
		{"yaml", "yaml", true},
		{"yml", "yml", true},
		{"xml", "xml", false},
		{"csv", "csv", false},
		{"pdf", "pdf", false},
		{"empty", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsRenderer(tt.output)
			if result != tt.expected {
				t.Errorf("IsRenderer(%s) = %v, expected %v", tt.output, result, tt.expected)
			}
		})
	}
}

func TestJsonRenderer(t *testing.T) {
	renderer := &jsonRenderer{}

	tests := []struct {
		name     string
		input    interface{}
		expected string
	}{
		{
			name:     "simple object",
			input:    map[string]interface{}{"key": "value"},
			expected: "{\n\t\"key\": \"value\"\n}",
		},
		{
			name:     "array",
			input:    []string{"a", "b", "c"},
			expected: "[\n\t\"a\",\n\t\"b\",\n\t\"c\"\n]",
		},
		{
			name:     "number",
			input:    42,
			expected: "42",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := renderer.Render(tt.input)
			if err != nil {
				t.Errorf("jsonRenderer.Render() error = %v", err)
			}
			if result != tt.expected {
				t.Errorf("jsonRenderer.Render() = %s, expected %s", result, tt.expected)
			}
		})
	}
}

func TestJsonRenderer_RenderBytes(t *testing.T) {
	renderer := &jsonRenderer{}

	testData := []byte(`{"test": "data"}`)
	result, err := renderer.RenderBytes(testData)
	if err != nil {
		t.Errorf("jsonRenderer.RenderBytes() error = %v", err)
	}

	expected := "{\n\t\"test\": \"data\"\n}"
	if result != expected {
		t.Errorf("jsonRenderer.RenderBytes() = %s, expected %s", result, expected)
	}
}

func TestYamlRenderer(t *testing.T) {
	renderer := &yamlRenderer{}

	tests := []struct {
		name  string
		input interface{}
	}{
		{
			name:  "simple object",
			input: map[string]interface{}{"key": "value", "number": 42},
		},
		{
			name:  "array",
			input: []string{"a", "b", "c"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := renderer.Render(tt.input)
			if err != nil {
				t.Errorf("yamlRenderer.Render() error = %v", err)
			}

			// Verify that the result is valid YAML by unmarshaling it
			var unmarshaled interface{}
			err = yaml.Unmarshal([]byte(result), &unmarshaled)
			if err != nil {
				t.Errorf("yamlRenderer.Render() produced invalid YAML: %v", err)
			}
		})
	}
}

func TestYamlRenderer_RenderBytes(t *testing.T) {
	renderer := &yamlRenderer{}

	testData := []byte(`{"test": "data"}`)
	result, err := renderer.RenderBytes(testData)
	if err != nil {
		t.Errorf("yamlRenderer.RenderBytes() error = %v", err)
	}

	expected := "test: data\n"
	if result != expected {
		t.Errorf("yamlRenderer.RenderBytes() = %s, expected %s", result, expected)
	}
}

func TestNewCsvRenderer(t *testing.T) {
	settings := &CsvSettings{
		Price:         0.35,
		TimeFormat:    "15:04:05",
		TimeZone:      "Europe/Berlin",
		Comma:         ";",
		HeaderEnabled: true,
		FilePath:      "/tmp/test.csv",
	}

	renderer := NewCsvRenderer(settings)
	if renderer == nil {
		t.Error("NewCsvRenderer() returned nil")
	}
	if renderer.settings != settings {
		t.Error("NewCsvRenderer() settings not set correctly")
	}
}

func TestNewPdfRenderer(t *testing.T) {
	settings := &PdfSettings{
		Price:       0.35,
		TimeFormat:  "15:04:05",
		TimeZone:    "Europe/Berlin",
		Comma:       ";",
		PrintHeader: true,
		LogoHeader:  "/tmp/logo.png",
		Settings: GlobalSettings{
			Firstname: "John",
			Lastname:  "Doe",
			Street:    "Main St",
			Postcode:  "12345",
			City:      "Berlin",
		},
	}

	renderer := NewPdfRenderer(settings)
	if renderer == nil {
		t.Error("NewPdfRenderer() returned nil")
	}
	if renderer.settings != settings {
		t.Error("NewPdfRenderer() settings not set correctly")
	}
}
