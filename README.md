# terraform-docs-extended

An extension to [terraform-docs](https://github.com/terraform-docs/terraform-docs) that provides enhanced variable type formatting and improved documentation structure.

## Features

### Enhanced Type Formatting

The standard terraform-docs output simplifies complex variable types without providing much context. This extension offers:

- **Intelligent type simplification** - Preserves important type information while making it readable
- **Context-aware formatting** - Different container types (lists, maps, sets) are formatted appropriately
- **Nested type handling** - Complex nested types are simplified in an intelligent way
- **Example values** - Generated example values based on variable type and name

### Improved Documentation Structure

- **Usage section at the end** - Places the Usage section at the end of the documentation for better flow
- **Header and footer support** - Inherits header and footer from terraform-docs configuration
- **Better variable organization** - Clear separation of required and optional variables

## Installation

```bash
# Clone this repository
git clone https://github.com/jishnusygal/terraform-docs-extended.git

# Build the binary
cd terraform-docs-extended
go build -o terraform-docs-extended

# Move to a directory in your PATH
sudo mv terraform-docs-extended /usr/local/bin/
```

## Usage

```bash
# Generate documentation for a module
terraform-docs-extended -p /path/to/module -o README.md

# Recursively generate documentation for all modules
terraform-docs-extended -p /path/to/modules -r
```

## Configuration

`terraform-docs-extended` respects the same configuration files as terraform-docs (.terraform-docs.yml), and additionally inherits header and footer content.

Example configuration:

```yaml
formatter: markdown

sections:
  show:
    - requirements
    - providers
    - inputs
    - outputs

content: |-
  {{ .Header }}

  {{ .Body }}

  {{ .Footer }}

header: |-
  # My Terraform Module
  
  This header will be included at the top of the documentation.

footer: |-
  ## License
  
  MIT
```

## Requirements

- Go 1.16 or later
- terraform-docs installed and available in your PATH

## License

MIT
