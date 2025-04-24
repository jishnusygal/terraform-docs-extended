package cmd

import (
	"testing"
)

func TestVersionCommand(t *testing.T) {
	// Simple test to verify that the version command exists
	if versionCmd == nil {
		t.Fatal("versionCmd is nil")
	}

	if versionCmd.Use != "version" {
		t.Errorf("expected command Use to be 'version', got %s", versionCmd.Use)
	}
}

func TestRootCommand(t *testing.T) {
	// Simple test to verify that the root command exists
	if rootCmd == nil {
		t.Fatal("rootCmd is nil")
	}
}
