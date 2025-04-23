package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/jishnusygal/terraform-docs-extended/pkg/processor"
	"github.com/spf13/cobra"
)

var (
	// Version is set during build
	Version = "0.1.0"

	// Command line flags
	modulePath   string
	outputFile   string
	recursive    bool
	outputFormat string
	moduleName   string
	moduleSource string
	quiet        bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "terraform-docs-extended",
	Short: "Generate extended Terraform module documentation",
	Long: `terraform-docs-extended is a CLI tool that extends terraform-docs functionality 
by adding a detailed "Usage" block with proper variable segregation.

It parses Terraform module files to extract variables and their types, and generates
documentation that shows users exactly how to use the module, with variables clearly
separated into "Required" and "Optional" sections.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Validate path
		if _, err := os.Stat(modulePath); os.IsNotExist(err) {
			fmt.Fprintf(os.Stderr, "Error: Module path does not exist: %s\n", modulePath)
			os.Exit(1)
		}

		// Validate output format
		if outputFormat != "markdown" && outputFormat != "json" {
			fmt.Fprintf(os.Stderr, "Error: Invalid output format: %s. Must be 'markdown' or 'json'\n", outputFormat)
			os.Exit(1)
		}

		// Check if terraform-docs is installed
		if !processor.IsTerraformDocsInstalled() {
			fmt.Fprintf(os.Stderr, "Error: terraform-docs is not installed or not found in PATH\n")
			os.Exit(1)
		}

		// Process directories based on recursive flag
		var err error
		if recursive {
			err = processor.ProcessRecursively(modulePath, outputFormat, outputFile, moduleName, moduleSource, quiet)
		} else {
			err = processor.ProcessDirectory(modulePath, outputFormat, outputFile, moduleName, moduleSource, quiet)
		}

		// Handle any errors
		if err != nil {
			errorExit(err)
		}
	},
}

// versionCmd displays version information
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Run: func(cmd *cobra.Command, args []string) {
		blue := color.New(color.FgBlue).SprintFunc()
		fmt.Printf("%s v%s\n", blue("terraform-docs-extended"), Version)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Add version command
	rootCmd.AddCommand(versionCmd)

	// Add command line flags
	rootCmd.Flags().StringVarP(&modulePath, "path", "p", ".", "Path to the Terraform module directory")
	rootCmd.Flags().StringVarP(&outputFile, "out", "o", "", "Output file path (defaults to stdout)")
	rootCmd.Flags().BoolVarP(&recursive, "recursive", "r", false, "Process directories recursively")
	rootCmd.Flags().StringVarP(&outputFormat, "format", "f", "markdown", "Output format (markdown or json)")
	rootCmd.Flags().StringVarP(&moduleName, "name", "n", "example", "Module name to use in the usage example")
	rootCmd.Flags().StringVarP(&moduleSource, "source", "s", "path/to/module", "Module source to use in the usage example")
	rootCmd.Flags().BoolVarP(&quiet, "quiet", "q", false, "Suppress informational output")
}
