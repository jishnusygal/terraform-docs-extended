package formatter

import (
	"strings"
	"testing"
)

func TestFormatMarkdown(t *testing.T) {
	// Create test variables
	variables := map[string]Variable{
		"required_string": {
			Name:        "required_string",
			Type:        "string",
			Description: "A required string variable",
			Required:    true,
		},
		"optional_number": {
			Name:        "optional_number",
			Type:        "number",
			Description: "An optional number variable",
			Default:     42,
			Required:    false,
		},
		"complex_object": {
			Name:        "complex_object",
			Type:        "object({...})",
			Description: "A complex object variable",
			Default:     map[string]interface{}{},
			Required:    false,
		},
	}

	// Create a formatter
	formatter := NewUsageFormatter(variables, "test-module", "terraform-registry/module")

	// Generate markdown
	output := formatter.FormatMarkdown()

	// Check expected lines
	expectedLines := []string{
		"## Usage",
		"```hcl",
		"module \"test-module\" {",
		"  source = \"terraform-registry/module\"",
		"  # Required variables",
		"  required_string = string",
		"  # Optional variables",
		"  complex_object = object({...})",
		"  optional_number = number",
	}

	for _, line := range expectedLines {
		if !strings.Contains(output, line) {
			t.Errorf("Expected output to contain '%s', but it didn't", line)
		}
	}
}

func TestFormatJSON(t *testing.T) {
	// Create test variables
	variables := map[string]Variable{
		"required_string": {
			Name:        "required_string",
			Type:        "string",
			Description: "A required string variable",
			Required:    true,
		},
		"optional_number": {
			Name:        "optional_number",
			Type:        "number",
			Description: "An optional number variable",
			Default:     42,
			Required:    false,
		},
	}

	// Create a formatter
	formatter := NewUsageFormatter(variables, "test-module", "terraform-registry/module")

	// Generate JSON structure
	output := formatter.FormatJSON()

	// Check for expected values
	if output["module_name"] != "test-module" {
		t.Errorf("Expected module_name to be 'test-module', got '%s'", output["module_name"])
	}

	if output["source"] != "terraform-registry/module" {
		t.Errorf("Expected source to be 'terraform-registry/module', got '%s'", output["source"])
	}

	// Check required variables
	requiredVars, ok := output["required"].([]map[string]string)
	if !ok {
		t.Fatalf("Expected required to be a slice of maps, got %T", output["required"])
	}

	if len(requiredVars) != 1 {
		t.Errorf("Expected 1 required variable, got %d", len(requiredVars))
	}

	// Check optional variables
	optionalVars, ok := output["optional"].([]map[string]string)
	if !ok {
		t.Fatalf("Expected optional to be a slice of maps, got %T", output["optional"])
	}

	if len(optionalVars) != 1 {
		t.Errorf("Expected 1 optional variable, got %d", len(optionalVars))
	}
}

func TestFormatTypeForUsage(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Simple string", "string", "string"},
		{"Simple number", "number", "number"},
		{"Complex object", "object({name = string, age = number, address = string})", "object({...})"},
		{"List of objects", "list(object({id = string, value = number}))", "list(...)"},
		{"Map of strings", "map(string)", "map(string)"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := formatTypeForUsage(test.input)
			if test.expected != result {
				t.Errorf("Expected '%s', got '%s'", test.expected, result)
			}
		})
	}
}
