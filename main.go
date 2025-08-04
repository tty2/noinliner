package main

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/singlechecker"
)

var Analyzer = &analysis.Analyzer{
	Name: "noinliner",
	Doc:  "disallows inline variable declarations in if statements",
	Run:  run,
}

func run(pass *analysis.Pass) (any, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			if stmt, ok := n.(*ast.IfStmt); ok && stmt.Init != nil {
				pass.Reportf(stmt.Pos(), "avoid inline variable declaration in if statement")
			}
			return true
		})
	}
	return nil, nil
}

func main() {
	singlechecker.Main(Analyzer)
}
