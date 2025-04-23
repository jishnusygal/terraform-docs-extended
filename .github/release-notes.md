# terraform-docs-extended v0.1.0

This is the initial release of terraform-docs-extended, a CLI tool that extends terraform-docs with a detailed Usage section for Terraform modules.

## Features

- Generate extended documentation for Terraform modules
- Automatically separate variables into "Required" and "Optional" sections
- Format output in Markdown or JSON
- Support for recursive processing of directories
- Customizable module name and source path in usage examples

## Installation

Download the appropriate binary for your platform:
- Linux: `terraform-docs-extended_linux_amd64`
- macOS: `terraform-docs-extended_darwin_amd64`
- Windows: `terraform-docs-extended_windows_amd64.exe`

Make it executable (Linux/macOS):
```bash
chmod +x terraform-docs-extended_*
```

## Usage

```bash
# Basic usage
terraform-docs-extended --path /path/to/module

# Output to file
terraform-docs-extended --path /path/to/module --out README.md

# Process directories recursively
terraform-docs-extended --path /path/to/modules --recursive

# Custom module name and source path
terraform-docs-extended --path /path/to/module --name my-module --source github.com/myorg/my-module
```

## Requirements

- Requires terraform-docs to be installed and available in PATH
