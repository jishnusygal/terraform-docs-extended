package terraform

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

// Variable represents a Terraform variable
type Variable struct {
	Name        string      `json:"name"`
	Type        string      `json:"type"`
	Description string      `json:"description"`
	Default     interface{} `json:"default"`
	Required    bool        `json:"required"`
}

// ExtractTerraformDocsInfo runs terraform-docs and extracts variable info
func ExtractTerraformDocsInfo(path string) (map[string]Variable, error) {
	cmd := exec.Command("terraform-docs", "json", path)
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to run terraform-docs: %v", err)
	}

	var tfDocsOutput map[string]interface{}
	if err := json.Unmarshal(output, &tfDocsOutput); err != nil {
		return nil, fmt.Errorf("failed to parse terraform-docs output: %v", err)
	}

	// Extract variables from terraform-docs output
	variables := make(map[string]Variable)

	// Navigate the JSON structure to find variables
	if inputs, ok := tfDocsOutput["inputs"].([]interface{}); ok {
		for _, input := range inputs {
			inputMap, ok := input.(map[string]interface{})
			if !ok {
				continue
			}

			name, ok := inputMap["name"].(string)
			if !ok {
				continue
			}
			
			typeStr, ok := inputMap["type"].(string)
			if !ok {
				typeStr = "any"
			} else {
				typeStr = FormatType(typeStr)
			}
			
			desc := ""
			if description, ok := inputMap["description"].(string); ok {
				desc = description
			}

			hasDefault := false
			if _, ok := inputMap["default"]; ok {
				hasDefault = true
			}

			variables[name] = Variable{
				Name:        name,
				Type:        typeStr,
				Description: desc,
				Default:     inputMap["default"],
				Required:    !hasDefault,
			}
		}
	}

	return variables, nil
}

// ParseModuleFiles parses Terraform module files directly for better variable type extraction
func ParseModuleFiles(modulePath string) (map[string]Variable, error) {
	variables := make(map[string]Variable)
	
	// Find all .tf files in the directory
	files, err := filepath.Glob(filepath.Join(modulePath, "*.tf"))
	if err != nil {
		return nil, fmt.Errorf("failed to list .tf files: %v", err)
	}
	
	// Process each file
	for _, file := range files {
		content, err := ioutil.ReadFile(file)
		if err != nil {
			return nil, fmt.Errorf("failed to read file %s: %v", file, err)
		}
		
		// Parse variables from the file
		fileVars, err := ParseVariablesFromContent(string(content))
		if err != nil {
			return nil, fmt.Errorf("failed to parse variables from %s: %v", file, err)
		}
		
		// Merge with existing variables
		for name, v := range fileVars {
			variables[name] = v
		}
	}
	
	return variables, nil
}

// ParseVariablesFromContent extracts variable definitions from HCL content
func ParseVariablesFromContent(content string) (map[string]Variable, error) {
	variables := make(map[string]Variable)
	
	// Regex to match variable blocks
	// This handles multi-line variable definitions
	varBlockRegex := regexp.MustCompile(`(?ms)variable\s+"([^"]+)"\s*{([^}]*)}`)
	
	// Find all variable blocks
	matches := varBlockRegex.FindAllStringSubmatch(content, -1)
	
	for _, match := range matches {
		if len(match) != 3 {
			continue
		}
		
		name := match[1]
		blockContent := match[2]
		
		variable := Variable{
			Name:     name,
			Required: true, // Default to required unless we find a default value
		}
		
		// Extract description
		descRegex := regexp.MustCompile(`(?m)description\s*=\s*"([^"]*)"`)
		descMatch := descRegex.FindStringSubmatch(blockContent)
		if len(descMatch) > 1 {
			variable.Description = descMatch[1]
		}
		
		// Extract type
		typeRegex := regexp.MustCompile(`(?ms)type\s*=\s*([^\n]+)`)
		typeMatch := typeRegex.FindStringSubmatch(blockContent)
		if len(typeMatch) > 1 {
			typeStr := typeMatch[1]
			// Handle multi-line type definitions
			variable.Type = CleanTypeString(typeStr)
		}
		
		// Check for default value
		defaultRegex := regexp.MustCompile(`(?m)default\s*=`)
		hasDefault := defaultRegex.MatchString(blockContent)
		variable.Required = !hasDefault
		
		variables[name] = variable
	}
	
	return variables, nil
}

// CleanTypeString formats a type string extracted from HCL
func CleanTypeString(typeStr string) string {
	// Remove trailing commas and whitespace
	typeStr = strings.TrimSpace(typeStr)
	typeStr = strings.TrimSuffix(typeStr, ",")
	
	// Handle multi-line type definitions by normalizing whitespace
	typeStr = NormalizeWhitespace(typeStr)
	
	// Format the type string for display
	return FormatType(typeStr)
}

// NormalizeWhitespace collapses multiple whitespace characters into a single space
func NormalizeWhitespace(s string) string {
	// Replace newlines with spaces
	s = strings.ReplaceAll(s, "\n", " ")
	
	// Collapse multiple spaces into one
	re := regexp.MustCompile(`\s+`)
	return re.ReplaceAllString(s, " ")
}

// FormatType cleans up Terraform type strings
func FormatType(typeStr string) string {
	// Remove any quotes around the type
	typeStr = strings.Trim(typeStr, "\"")
	
	// Remove newlines and extra whitespace
	re := regexp.MustCompile(`\s+`)
	typeStr = re.ReplaceAllString(typeStr, " ")
	
	// Handle complex types more concisely
	if strings.Contains(typeStr, "object({") && len(typeStr) > 50 {
		// Try to simplify complex object types
		objectStart := strings.Index(typeStr, "object({")
		if objectStart >= 0 {
			// Count braces to find the matching closing brace
			openBraces := 0
			closingPos := -1
			
			for i := objectStart + 8; i < len(typeStr); i++ {
				if typeStr[i] == '{' {
					openBraces++
				} else if typeStr[i] == '}' {
					if openBraces == 0 {
						closingPos = i
						break
					}
					openBraces--
				}
			}
			
			if closingPos > 0 {
				// Replace the content between braces with ...
				return typeStr[:objectStart+8] + "...}" + typeStr[closingPos+1:]
			}
		}
		
		// If we couldn't parse it properly, use a generic simplification
		return "object({...})"
	}
	
	// Simplify complex list/set/map types with nested objects
	for _, prefix := range []string{"list(", "set(", "map("} {
		if strings.HasPrefix(typeStr, prefix) && strings.Contains(typeStr, "object") && len(typeStr) > 50 {
			closingPos := len(typeStr) - 1
			if typeStr[closingPos] == ')' {
				return prefix + "..." + typeStr[closingPos:]
			}
		}
	}
	
	// Handle tuple types
	if strings.HasPrefix(typeStr, "tuple([") && len(typeStr) > 50 {
		return "tuple([...])"
	}
	
	return typeStr
}
