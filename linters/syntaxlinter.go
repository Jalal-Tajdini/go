package linters

import (
	"go/parser"
	"go/token"
)

type SyntaxLinter struct {
	BaseLinter
}

func (l *SyntaxLinter) Apply(input string) (string, error) {
	// Parse the input source code
	fset := token.NewFileSet()
	_, err := parser.ParseFile(fset, "source.go", input, 0)
	if err != nil {
		return "", err // Return the syntax error
	}

	// No syntax errors, pass the input to the next linter
	if l.next != nil {
		return l.next.Apply(input)
	}

	return input, nil
}
