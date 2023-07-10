// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package ioreadall

import (
	"fmt"
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const (
	// Doc for the timeafter check
	Doc = `This checks for "io.ReadAll" instances.`

	readAllFunc = "ReadAll"
)

var ioReadAllPkgs = []string{"io", "ioutil"}

// Analyzer is the global for the multichecker
var Analyzer = &analysis.Analyzer{
	Name:     "readall",
	Doc:      Doc,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
	Run:      run,
}

var ignoreArg string

func init() {
	Analyzer.Flags.StringVar(&ignoreArg, "ignore", "", `list of packages to ignore (e.g. "readall,config")`)
}

type visitor func(ast.Node) bool

func (v visitor) Visit(node ast.Node) ast.Visitor {
	if v(node) {
		return v
	}
	return nil
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspct, ok := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	if !ok {
		return nil, fmt.Errorf("analyzer is not type *inspector.Inspector")
	}

	ignoreMap := make(map[string]struct{})
	for _, ign := range strings.Split(ignoreArg, ",") {
		ignoreMap[strings.TrimSpace(ign)] = struct{}{}
	}

	var (
		pkgAliases []string
		ignore     = false
		nodeFilter = []ast.Node{
			(*ast.CallExpr)(nil),
			(*ast.File)(nil),
			(*ast.ImportSpec)(nil),
		}
	)
	inspct.Preorder(nodeFilter, func(n ast.Node) {
		switch stmt := n.(type) {
		case *ast.File:
			_, ignore = ignoreMap[stmt.Name.Name]
			pkgAliases = ioReadAllPkgs
		case *ast.ImportSpec:
			if ignore {
				return
			}
			// Collect aliases.
			pkg := stmt.Path.Value
			for _, originPkg := range ioReadAllPkgs {
				if pkg == fmt.Sprintf("%q", originPkg) {
					if stmt.Name != nil {
						pkgAliases = append(pkgAliases, stmt.Name.Name)
					}
				}
			}
		case *ast.CallExpr:
			if ignore {
				return
			}
			for _, pkg := range pkgAliases {
				if isPkgDot(stmt.Fun, pkg, readAllFunc) {
					pass.Reportf(n.Pos(), "use of %s.ReadAll is prohibited, use safeio.ReadAllLimit instead", pkg)
				}
			}
		}
	})
	return nil, nil
}

func isPkgDot(expr ast.Expr, pkg, name string) bool {
	sel, ok := expr.(*ast.SelectorExpr)
	res := ok && isIdent(sel.X, pkg) && isIdent(sel.Sel, name)
	return res
}

func isIdent(expr ast.Expr, ident string) bool {
	id, ok := expr.(*ast.Ident)
	return ok && id.Name == ident
}
