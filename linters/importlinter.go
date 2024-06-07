package linters

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
)

type ImportLinter struct {
	BaseLinter
}

func (l *ImportLinter) Apply(input string) (string, error) {
	// Parse the input source code
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "source.go", input, 0)
	if err != nil {
		return "", err
	}

	// Get a list of imported packages
	imports := make([]string, 0, len(file.Imports))
	for _, imp := range file.Imports {
		imports = append(imports, strings.Trim(imp.Path.Value, `"`))
	}

	// Check if all imported packages are used
	usedImports := make(map[string]bool)
	ast.Inspect(file, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.ImportSpec:
			// Mark imported package as used
			usedImports[strings.Trim(x.Path.Value, `"`)] = true
		case *ast.SelectorExpr:
			// If the selector expression refers to an imported package, mark it as used
			if x, ok := x.X.(*ast.Ident); ok {
				if isImportedPackage(imports, x.Name) {
					usedImports[x.Name] = true
				}
			}
		}
		return true
	})

	// Check for unused imports
	var unusedImports []string
	for _, imp := range imports {
		if !usedImports[imp] {
			unusedImports = append(unusedImports, imp)
		}
	}

	if len(unusedImports) > 0 {
		return input, fmt.Errorf("unused imports: %v", unusedImports)
	}

	// Pass the input to the next linter
	if l.next != nil {
		return l.next.Apply(input)
	}

	return input, nil
}

func isImportedPackage(imports []string, name string) bool {
	for _, imp := range imports {
		if imp == name {
			return true
		}
	}
	return false
}
