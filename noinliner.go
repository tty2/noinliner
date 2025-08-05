package noinliner

import (
	"go/ast"

	"github.com/golangci/plugin-module-register/register"
	"golang.org/x/tools/go/analysis"
)

type NoInliner struct{}

func init() {
	register.Plugin("noinliner", NewPlugin)
}

func NewPlugin(settings any) (register.LinterPlugin, error) {
	return &NoInliner{}, nil
}

func (ni *NoInliner) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{
		{
			Name: "noinliner",
			Doc:  "disallows inline variable declarations in if statements",
			Run:  ni.run,
		},
	}, nil
}

func (ni *NoInliner) run(pass *analysis.Pass) (any, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			stmt, ok := n.(*ast.IfStmt)
			if ok && stmt.Init != nil {
				pass.Reportf(stmt.Pos(), "avoid inline variable declaration in if statement")
			}
			return true
		})
	}

	return nil, nil
}

func (ni *NoInliner) GetLoadMode() string {
	return register.LoadModeSyntax
}
