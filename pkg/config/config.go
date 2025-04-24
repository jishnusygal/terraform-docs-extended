package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// Config represents the configuration for terraform-docs-extended
type Config struct {
	TypeFormatting TypeFormattingConfig `yaml:"type_formatting"`
}

// TypeFormattingConfig contains settings for how types are formatted
type TypeFormattingConfig struct {
	// DetailLevel determines how much detail to include in type descriptions
	// Possible values: "minimal", "moderate", "detailed"
	DetailLevel string `yaml:"detail_level"`

	// ShowFieldNames determines whether to show field names in complex objects
	ShowFieldNames bool `yaml:"show_field_names"`

	// MaxFieldsToShow is the maximum number of fields to show in complex objects
	MaxFieldsToShow int `yaml:"max_fields_to_show"`

	// CustomFormats allows for custom formatting of specific types
	CustomFormats map[string]string `yaml:"custom_formats"`
}

// DefaultConfig returns the default configuration
func DefaultConfig() Config {
	return Config{
		TypeFormatting: TypeFormattingConfig{
			DetailLevel:     "moderate",
			ShowFieldNames:  true,
			MaxFieldsToShow: 3,
			CustomFormats:   make(map[string]string),
		},
	}
}

// LoadConfig loads configuration from file
func LoadConfig(modulePath string) (Config, error) {
	config := DefaultConfig()

	// Check for configuration file in module path
	configPaths := []string{
		filepath.Join(modulePath, ".terraform-docs-extended.yml"),
		filepath.Join(modulePath, ".terraform-docs-extended.yaml"),
		filepath.Join(modulePath, "terraform-docs-extended.yml"),
		filepath.Join(modulePath, "terraform-docs-extended.yaml"),
	}

	for _, path := range configPaths {
		if fileExists(path) {
			data, err := ioutil.ReadFile(path)
			if err != nil {
				return config, fmt.Errorf("error reading config file: %w", err)
			}

			err = yaml.Unmarshal(data, &config)
			if err != nil {
				return config, fmt.Errorf("error parsing config file: %w", err)
			}

			// Config file found and loaded successfully
			return config, nil
		}
	}

	// No config file found, return default config
	return config, nil
}

// fileExists checks if a file exists
func fileExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// Validate checks if the configuration is valid
func (c *Config) Validate() error {
	// Validate DetailLevel
	validDetailLevels := map[string]bool{
		"minimal":  true,
		"moderate": true,
		"detailed": true,
	}
	
	if !validDetailLevels[c.TypeFormatting.DetailLevel] {
		return fmt.Errorf("invalid detail_level: %s. Valid values are 'minimal', 'moderate', 'detailed'", 
			c.TypeFormatting.DetailLevel)
	}

	// Validate MaxFieldsToShow
	if c.TypeFormatting.MaxFieldsToShow < 0 {
		return fmt.Errorf("max_fields_to_show must be non-negative")
	}

	return nil
}
