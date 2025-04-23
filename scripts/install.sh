#!/bin/bash
set -e

# terraform-docs-extended installer script
#
# This script installs terraform-docs-extended and its dependencies
# on Linux and macOS systems.

# Check if terraform-docs is installed
check_terraform_docs() {
  if ! command -v terraform-docs >/dev/null 2>&1; then
    echo "terraform-docs is not installed or not in PATH."
    echo "terraform-docs is required for terraform-docs-extended to function."
    echo "Please install terraform-docs first: https://github.com/terraform-docs/terraform-docs#installation"
    exit 1
  fi
}

# Check if Go is installed
check_go() {
  if ! command -v go >/dev/null 2>&1; then
    echo "Go is not installed or not in PATH."
    echo "Go is required to build terraform-docs-extended."
    echo "Please install Go first: https://golang.org/doc/install"
    exit 1
  fi
}

# Install from source
install_from_source() {
  check_go
  check_terraform_docs

  # Clone the repository
  TEMP_DIR=$(mktemp -d)
  echo "Cloning terraform-docs-extended repository..."
  git clone https://github.com/jishnusygal/terraform-docs-extended.git "$TEMP_DIR"
  cd "$TEMP_DIR"

  # Build and install
  echo "Building terraform-docs-extended..."
  go build -o terraform-docs-extended

  # Move to destination directory
  INSTALL_DIR="/usr/local/bin"
  if [[ "$OSTYPE" == "darwin"* ]]; then
    # macOS might need sudo for /usr/local/bin
    echo "Installing to $INSTALL_DIR (may require sudo)..."
    sudo mv terraform-docs-extended "$INSTALL_DIR/"
  else
    # Linux
    echo "Installing to $INSTALL_DIR (may require sudo)..."
    sudo mv terraform-docs-extended "$INSTALL_DIR/"
  fi

  # Clean up
  cd - > /dev/null
  rm -rf "$TEMP_DIR"

  echo "terraform-docs-extended has been installed successfully!"
  echo "Run 'terraform-docs-extended --help' to get started."
}

# Main installation logic
echo "terraform-docs-extended installer"
echo "--------------------------------"

install_from_source

echo "Installation complete!"