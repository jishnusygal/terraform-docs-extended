package cmd

import (
	"os/exec"
)

// isTerraformDocsInstalled checks if terraform-docs is available
func isTerraformDocsInstalled() bool {
	cmd := exec.Command("terraform-docs", "--version")
	if err := cmd.Run(); err != nil {
		return false
	}
	return true
}

// processRecursively handles recursive directory traversal
func processRecursively(path string, format string, outputFile string, moduleName string, moduleSource string, quiet bool) {
	// In a real implementation, this would call into the processor package
	// For now, just print a message so we can test the command-line interface
	if !quiet {
		println("Processing recursively:", path)
	}
}

// processDirectory handles a single directory
func processDirectory(path string, format string, outputFile string, moduleName string, moduleSource string, quiet bool) {
	// In a real implementation, this would call into the processor package
	// For now, just print a message so we can test the command-line interface
	if !quiet {
		println("Processing directory:", path)
	}
}
