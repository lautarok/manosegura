package docs

import (
	"fmt"
	"go/ast"
)

func ResolveAstType(expr ast.Expr) string {
	switch t := expr.(type) {

	case *ast.Ident:
		return t.Name

	case *ast.StarExpr:
		return ResolveAstType(t.X) // puntero, ignoramos el *

	case *ast.SelectorExpr:
		if pkgIdent, ok := t.X.(*ast.Ident); ok {
			return pkgIdent.Name + "." + t.Sel.Name
		}
		return t.Sel.Name

	case *ast.ArrayType:
		return "[]" + ResolveAstType(t.Elt)

	case *ast.MapType:
		keyType := ResolveAstType(t.Key)
		valType := ResolveAstType(t.Value)
		return fmt.Sprintf("map[%s]%s", keyType, valType)

	default:
		return "unknown"
	}
}
