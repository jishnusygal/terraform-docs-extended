# Release v0.1.0

Initial release of terraform-docs-extended.

## New Features

- Enhanced "Usage" section that segregates variables into "Required" and "Optional" categories
- Support for all Terraform data types including complex nested types
- Multi-line variable definition handling
- Proper formatting of complex types (objects, lists, maps, etc.)
- Command-line interface with flexible options
- Recursive processing capability
- Multiple output formats (Markdown and JSON)
- Docker support for containerized execution
- Homebrew formula integration

## Technical Improvements

- Built on top of terraform-docs to leverage existing parsing capabilities
- Extended with custom HCL parsing for improved type information extraction
- Clean architecture with separation of concerns:
  - Parser package for Terraform file handling
  - Formatter package for output generation
  - Processor package for core logic
- CI/CD pipeline for automated testing and releases
- Comprehensive test suite

## Installation

### Using Go

```bash
go install github.com/jishnusygal/terraform-docs-extended@v0.1.0
```

### Using Homebrew

```bash
brew tap jishnusygal/tap
brew install jishnusygal/tap/terraform-docs-extended
```

### Using Docker

```bash
docker pull ghcr.io/jishnusygal/terraform-docs-extended:v0.1.0
docker run -v $(pwd):/workspace ghcr.io/jishnusygal/terraform-docs-extended:v0.1.0 --path=/workspace
```

## Notes

This is the initial release. Feedback and contributions are welcome!