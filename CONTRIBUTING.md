# Contributing to terraform-docs-extended

Thank you for your interest in contributing to terraform-docs-extended! This document provides guidelines and instructions for contributing.

## Code of Conduct

Please be respectful and considerate of others when contributing to this project.

## Getting Started

1. Fork the repository
2. Clone your fork to your local machine
3. Set up the development environment
4. Make your changes
5. Submit a pull request

## Development Environment Setup

### Prerequisites

- Go 1.16 or later
- terraform-docs installed and available in PATH

### Setup

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/terraform-docs-extended.git
   cd terraform-docs-extended
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Build the project:
   ```bash
   make build
   ```

## Development Workflow

1. Create a new branch for your feature or bugfix:
   ```bash
   git checkout -b feature/your-feature-name
   ```

2. Make your changes, following the coding conventions below
3. Add tests for your changes
4. Run the tests to ensure they pass:
   ```bash
   make test
   ```

5. Commit your changes with a clear and descriptive commit message
6. Push your branch to your fork
7. Create a pull request to the main repository

## Coding Conventions

- Follow Go best practices and idiomatic Go
- Use meaningful variable and function names
- Add comments to explain complex logic
- Update documentation for any changed functionality
- Write tests for new features or bugfixes

## Pull Request Process

1. Ensure your code passes all tests
2. Update the README.md with details of changes if appropriate
3. The PR should clearly describe what changes were made and why
4. Wait for a maintainer to review your PR
5. Address any feedback or requested changes
6. Once approved, a maintainer will merge your PR

## Reporting Bugs

Please report bugs by opening a new issue. Include:

- A clear descriptive title
- A detailed description of the bug
- Steps to reproduce the bug
- Expected and actual behavior
- Version information (Go version, OS, etc.)
- Any relevant logs or error messages

## Feature Requests

Feature requests are welcome! Please submit them as issues with:

- A clear descriptive title
- A detailed description of the feature
- Any relevant background or context
- Potential implementation ideas (optional)

## Questions or Need Help?

If you have questions or need help, please open an issue with the "question" label.

Thank you for contributing to terraform-docs-extended!
