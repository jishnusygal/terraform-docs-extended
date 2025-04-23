package cmd

import (
	"os/exec"

	"github.com/jishnusygal/terraform-docs-extended/pkg/processor"
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
	// Delegate to the processor package
	err := processor.ProcessRecursively(path, format, outputFile, moduleName, moduleSource, quiet)
	if err != nil {
		// Print error and exit
		errorExit(err)
	}
}

// processDirectory handles a single directory
func processDirectory(path string, format string, outputFile string, moduleName string, moduleSource string, quiet bool) {
	// Delegate to the processor package
	err := processor.ProcessDirectory(path, format, outputFile, moduleName, moduleSource, quiet)
	if err != nil {
		// Print error and exit
		errorExit(err)
	}
}
