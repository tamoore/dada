package ui

import (
	"fmt"

	cl "github.com/fatih/color"
)

// Checkbox ...
func Checkbox(label string, checked bool) string {
	if checked {

		return cl.New(cl.Bold).Sprintf("[x] %s\n", label)
	}
	return fmt.Sprintf("[ ] %s\n", label)
}
