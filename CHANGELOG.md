# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.1.0] - 2025-04-23

### Added
- Initial release of terraform-docs-extended
- Core functionality to parse Terraform module files and extract variable information
- Enhanced "Usage" section that segregates variables into "Required" and "Optional" categories
- Support for all Terraform data types including complex nested types
- Multi-line variable definition handling
- Proper formatting of complex types (objects, lists, maps, etc.)
- Command-line interface with the following features:
  - Path specification (`--path`)
  - Output file selection (`--out`)
  - Recursive processing (`--recursive`)
  - Output format selection (`--format`)
  - Module name customization (`--name`)
  - Module source customization (`--source`)
  - Quiet mode (`--quiet`)
- Multiple output formats (Markdown and JSON)
- Docker support for containerized execution
- Installation script for easy deployment
- Support for building on multiple platforms (Linux, macOS, Windows)
- Comprehensive test suite
- Homebrew formula for easy installation

### Technical Details
- Built on top of terraform-docs to leverage existing parsing capabilities
- Extended with custom HCL parsing for improved type information extraction
- Implemented with clean architecture and separation of concerns
- CI/CD pipeline for automated testing and releases
