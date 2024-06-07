package linters

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

type CommentLinter struct {
	BaseLinter
}

func (l *CommentLinter) Apply(input string) (string, error) {
	// Parse the input source code
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "source.go", input, parser.ParseComments)
	if err != nil {
		return "", err
	}

	// Check if each function has a comment
	for _, decl := range file.Decls {
		if fn, ok := decl.(*ast.FuncDecl); ok {
			if fn.Doc == nil {
				// Function has no comment
				funcName := fn.Name.Name
				fmt.Printf("Warning: Function '%s' does not have a comment.\n", funcName)
			}
		}
	}

	// Pass the input to the next linter (if any)
	if l.next != nil {
		return l.next.Apply(input)
	}

	return input, nil
}
