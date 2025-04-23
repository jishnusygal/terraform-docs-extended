package cmd

import (
	"fmt"
	"os"
)

// errorExit prints an error message and exits with a non-zero status
func errorExit(err error) {
	fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	os.Exit(1)
}