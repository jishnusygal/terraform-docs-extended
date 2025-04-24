package processor

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/jishnusygal/terraform-docs-extended/pkg/formatter"
	"github.com/jishnusygal/terraform-docs-extended/pkg/terraform"
)

// ProcessRecursively handles recursive directory traversal
func ProcessRecursively(root string, format string, outputFile string, moduleName string, moduleSource string, quiet bool) error {
	return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip .git, .terraform directories
		if info.IsDir() && (info.Name() == ".git" || info.Name() == ".terraform") {
			return filepath.SkipDir
		}

		// Skip non-directories
		if !info.IsDir() {
			return nil
		}

		// Check if this directory contains .tf files (potential module)
		files, err := filepath.Glob(filepath.Join(path, "*.tf"))
		if err != nil {
			return err
		}

		if len(files) > 0 {
			// Generate output filename based on directory if not specified
			outputPath := outputFile
			if outputFile == "" {
				// Use stdout if not recursive
				if path == root && !filepath.HasPrefix(path, root+string(os.PathSeparator)) {
					outputPath = ""
				} else {
					// Create output filename based on directory name
					outputPath = filepath.Join(path, fmt.Sprintf("README.%s", format))
				}
			}
			
			// Use directory name as module name if processing recursively
			modName := moduleName
			if path != root {
				modName = filepath.Base(path)
			}
			
			// Process this directory as a module
			if err := ProcessDirectory(path, format, outputPath, modName, moduleSource, quiet); err != nil {
				return err
			}
		}

		return nil
	})
}

// ProcessDirectory handles a single directory
func ProcessDirectory(path string, format string, outputPath string, moduleName string, moduleSource string, quiet bool) error {
	if !quiet {
		fmt.Printf("Processing module: %s\n", path)
	}
	
	// Extract module information
	module, err := ExtractModuleInfo(path, moduleName)
	if err != nil {
		return fmt.Errorf("failed to extract module info: %v", err)
	}

	// Generate the documentation with our extended usage section
	docContent := formatter.GenerateDoc(module, format, moduleSource)

	// Output the documentation
	if outputPath != "" {
		if err := os.WriteFile(outputPath, []byte(docContent), 0644); err != nil {
			return fmt.Errorf("failed to write output file: %v", err)
		}
		if !quiet {
			fmt.Printf("Documentation written to: %s\n", outputPath)
		}
	} else {
		fmt.Println(docContent)
	}
	
	return nil
}

// ExtractModuleInfo collects information about a Terraform module
func ExtractModuleInfo(path string, moduleName string) (formatter.Module, error) {
	// Run terraform-docs to get base information
	tfDocsVars, err := terraform.ExtractTerraformDocsInfo(path)
	if err != nil {
		log.Printf("Warning: Failed to extract info from terraform-docs: %v", err)
		// Continue with empty variables map
		tfDocsVars = make(map[string]terraform.Variable)
	}

	// Parse Terraform files directly for better type extraction
	parsedVars, err := terraform.ParseModuleFiles(path)
	if err != nil {
		log.Printf("Warning: Failed to parse module files directly: %v", err)
		// Continue with terraform-docs variables only
	}

	// Convert terraform.Variable to formatter.Variable
	formatterVars := make(map[string]formatter.Variable)
	for name, v := range MergeVariables(tfDocsVars, parsedVars) {
		formatterVars[name] = formatter.Variable{
			Name:        v.Name,
			Type:        v.Type,
			Description: v.Description,
			Default:     v.Default,
			Required:    v.Required,
		}
	}

	// Create the module with merged variable information
	module := formatter.Module{
		Path:      path,
		Name:      moduleName,
		Variables: formatterVars,
	}

	return module, nil
}

// IsTerraformDocsInstalled checks if terraform-docs is available
func IsTerraformDocsInstalled() bool {
	cmd := exec.Command("terraform-docs", "--version")
	if err := cmd.Run(); err != nil {
		return false
	}
	return true
}

// MergeVariables combines variable information from terraform-docs with direct parsing
func MergeVariables(tfDocsVars map[string]terraform.Variable, parsedVars map[string]terraform.Variable) map[string]terraform.Variable {
	result := make(map[string]terraform.Variable)
	
	// Start with terraform-docs variables
	for name, v := range tfDocsVars {
		result[name] = v
	}
	
	// Enhance with our parsed variables
	for name, v := range parsedVars {
		if existing, ok := result[name]; ok {
			// Keep most fields from terraform-docs but use our type parsing
			if v.Type != "" {
				existing.Type = v.Type
			}
			result[name] = existing
		} else {
			// Add any variables we found that terraform-docs didn't
			result[name] = v
		}
	}
	
	return result
}
