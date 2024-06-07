package linters

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

type FunctionLengthLinter struct {
	BaseLinter
	MaxFunctionLength int
}

func (l *FunctionLengthLinter) Apply(input string) (string, error) {
	// Parse the input source code
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "source.go", input, 0)
	if err != nil {
		return "", err
	}

	// Assume MaxFunctionLength as 50 if not provided
	if l.MaxFunctionLength == 0 {
		l.MaxFunctionLength = 50
	}

	// Check the length of each function
	for _, decl := range file.Decls {
		if fn, ok := decl.(*ast.FuncDecl); ok {
			length := len(fn.Body.List)
			if length > l.MaxFunctionLength {
				// Function length exceeds the limit
				funcName := fn.Name.Name
				fmt.Printf("Function '%s' exceeds the maximum length of %d lines (current length: %d)\n", funcName, l.MaxFunctionLength, length)
			}
		}
	}

	// Pass the input to the next linter
	if l.next != nil {
		return l.next.Apply(input)
	}

	return input, nil
}
