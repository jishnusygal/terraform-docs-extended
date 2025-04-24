package processor

import (
	"os/exec"
	"testing"
)

func TestIsTerraformDocsInstalled(t *testing.T) {
	// Skip this test if we're in CI environment since it depends on external program
	if _, err := exec.LookPath("terraform-docs"); err != nil {
		t.Skip("terraform-docs not installed, skipping test")
	}
	
	// The actual function just checks if terraform-docs is installed
	result := IsTerraformDocsInstalled()
	if !result {
		t.Error("IsTerraformDocsInstalled returned false but terraform-docs is installed")
	}
}

// Simple unit test that doesn't depend on external commands
func TestGetTerraformDocsVersion(t *testing.T) {
	// Mock executing command by replacing function
	origExecCommand := execCommand
	defer func() { execCommand = origExecCommand }()
	
	// Mock version output
	execCommand = func(command string, args ...string) *exec.Cmd {
		return exec.Command("echo", "terraform-docs v0.16.0")
	}
	
	version := GetTerraformDocsVersion()
	expected := "v0.16.0"
	
	if version != expected {
		t.Errorf("Expected version %s, got %s", expected, version)
	}
}
