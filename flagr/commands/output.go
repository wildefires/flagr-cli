package commands

import (
	"fmt"
	"os"
)

// Helper functions for printing
func logAndDie(msg string) {
	fmt.Fprintln(os.Stderr, msg)
	os.Exit(1)
}
