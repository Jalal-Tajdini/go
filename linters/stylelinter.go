package linters

import (
	"strings"
)

type StyleLinter struct {
	BaseLinter
}

func (l *StyleLinter) Apply(input string) (string, error) {
	// Check for tabs and replace them with spaces
	output := strings.ReplaceAll(input, "\t", "    ")

	// Pass the processed input to the next linter
	if l.next != nil {
		return l.next.Apply(output)
	}

	return output, nil
}
