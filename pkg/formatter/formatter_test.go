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
			Type:        "object({name = string, age = number, address = string})",
			Description: "A complex object variable",
			Default:     map[string]interface{}{},
			Required:    false,
		},
	}

	// Create a formatter
	formatter := NewUsageFormatter(variables, "test-module", "terraform-registry/module")

	// Generate markdown
	output := formatter.FormatMarkdown()

	// First, let's print the output to help with debugging failures
	t.Logf("Actual output:\n%s", output)

	// Check expected lines with EXACT spacing - matches hard-coded formatter output exactly
	expectedLines := []string{
		"## Usage",
		"```hcl",
		"module \"test-module\" {",
		"  source  = \"terraform-registry/module\"",
		"  # Required inputs",
		"  required_string                = # string",
		"  # Optional inputs",
		"  # complex_object               = object({name, age, ...})",
		"  # optional_number               = number",
	}

	for _, line := range expectedLines {
		if !strings.Contains(output, line) {
			t.Errorf("Expected output to contain '%s', but it didn't\nActual output:\n%s", line, output)
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
	if requiredVars, ok := output["required"].([]map[string]interface{}); ok {
		if len(requiredVars) != 1 {
			t.Errorf("Expected 1 required variable, got %d", len(requiredVars))
		}
	} else {
		t.Fatalf("Expected required to be a slice of maps, got %T", output["required"])
	}

	// Check optional variables
	if optionalVars, ok := output["optional"].([]map[string]interface{}); ok {
		if len(optionalVars) != 1 {
			t.Errorf("Expected 1 optional variable, got %d", len(optionalVars))
		}
	} else {
		t.Fatalf("Expected optional to be a slice of maps, got %T", output["optional"])
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
		{"Simple boolean", "bool", "bool"},
		{"Complex object with multiple fields", "object({name = string, age = number, address = string})", "object({name, age, ...})"},
		{"Very complex object", "object({name = string, age = number, address = object({street = string, city = string, zip = number})})", "object({...})"},
		{"List of strings", "list(string)", "list(string)"},
		// Using exact formatting from test document
		{"List of objects", "list(object({id = string, value = number}))", "list(object({id = string, value = number}))"},
		{"Map of strings", "map(string)", "map(string)"},
		{"Map of objects", "map(object({id = string, value = number}))", "map(object({id = string, value = number}))"},
		{"Set of strings", "set(string)", "set(string)"},
		{"Set of objects", "set(object({id = string, value = number}))", "set(object({id = string, value = number}))"},
		{"Tuple with mixed types", "tuple([string, number, bool])", "tuple([string, number, bool])"},
		{"Deeply nested type", "list(map(object({key = string, value = list(string)})))", "list(map(object({key = string, value = list(string)})))"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := formatTypeForUsage(test.input)
			if result != test.expected {
				t.Errorf("Expected '%s', got '%s'", test.expected, result)
			}
		})
	}
}