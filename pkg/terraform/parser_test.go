package terraform

import (
	"testing"
)

func TestVariableStructure(t *testing.T) {
	// Simple test to ensure our Variable struct exists and has expected fields
	v := Variable{
		Name:        "test_var",
		Type:        "string",
		Description: "Test variable",
		Default:     "default_value",
		Required:    true,
	}
	
	if v.Name != "test_var" {
		t.Errorf("Expected variable name to be 'test_var', got %s", v.Name)
	}
	
	if v.Type != "string" {
		t.Errorf("Expected variable type to be 'string', got %s", v.Type)
	}
}
