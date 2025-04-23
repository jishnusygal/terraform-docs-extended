# terraform-docs-extended

A command-line tool that extends `terraform-docs` functionality by adding a detailed "Usage" block that clearly shows how to use a Terraform module with proper variable segregation between required and optional variables.

## Features

- Generates standard documentation similar to `terraform-docs`
- Adds a specialized "Usage" section that shows how to instantiate the module
- Segregates module variables into "Required" and "Optional" sections
- Properly handles all Terraform variable types including complex types
- Supports multi-line variable definitions and type constraints
- Can process modules recursively
- Outputs in Markdown or JSON format

## Installation

### Prerequisites

- Go 1.16 or later
- `terraform-docs` installed and available in your PATH

### Install from source

```bash
git clone https://github.com/jishnusygal/terraform-docs-extended.git
cd terraform-docs-extended
go install
```

## Usage

```bash
# Generate documentation for a module in the current directory
terraform-docs-extended

# Generate documentation for a specific module directory
terraform-docs-extended --path=/path/to/module

# Output to a specific file
terraform-docs-extended --path=/path/to/module --out=README.md

# Process directories recursively
terraform-docs-extended --recursive --path=/path/to/modules

# Generate JSON output instead of Markdown
terraform-docs-extended --format=json --path=/path/to/module
```

### Command-line options

| Option | Description | Default |
|--------|-------------|---------|
| `--path` | Path to the Terraform module directory | Current directory |
| `--out` | Output file path (defaults to stdout) | stdout |
| `--recursive` | Process directories recursively | false |
| `--format` | Output format (markdown or json) | markdown |
| `--name` | Module name to use in the usage example | example |
| `--source` | Module source to use in the usage example | path/to/module |

## Example Output

Given a Terraform module with variables like:

```hcl
variable "instance_count" {
  description = "Number of instances to create"
  type        = number
}

variable "instance_type" {
  description = "EC2 instance type"
  type        = string
  default     = "t2.micro"
}
```

The generated documentation will include a usage section like:

```hcl
module "example" {
  source = "path/to/module"

  # Required variables
  instance_count = number

  # Optional variables
  instance_type = string
}
```

## How It Works

The tool works by:

1. Using `terraform-docs` to extract basic module information
2. Parsing the module's HCL files directly to handle complex types and multi-line definitions
3. Merging and processing the variable information
4. Generating a rich "Usage" section that properly segregates variables
5. Combining this with the standard documentation from `terraform-docs`

## Development

### Project Structure

- `main.go` - Main CLI application, flags, and program flow
- `cmd/` - Command handling using Cobra
- `pkg/terraform/` - Logic for parsing Terraform HCL files
- `pkg/formatter/` - Specialized formatting for the "Usage" section
- `pkg/processor/` - Core processing logic

### Building from source

```bash
go build -o terraform-docs-extended
```

### Running tests

```bash
go test ./...
```

## License

MIT License