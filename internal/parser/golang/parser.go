package golang

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"path/filepath"
	"strings"

	graph "github.com/andev0x/ctxd/internal/graph/contracts"
)

type Parser struct {
	fset *token.FileSet
}

func NewParser() *Parser {
	return &Parser{
		fset: token.NewFileSet(),
	}
}

func (p *Parser) ParseFile(path string) ([]*graph.Node, []*graph.Edge, error) {
	f, err := parser.ParseFile(p.fset, path, nil, parser.ParseComments)
	if err != nil {
		return nil, nil, err
	}

	var nodes []*graph.Node
	var edges []*graph.Edge

	// Use directory as package path for now
	pkgPath := filepath.Dir(path)
	pkgID := fmt.Sprintf("pkg:%s", pkgPath)

	nodes = append(nodes, &graph.Node{
		ID:   pkgID,
		Type: graph.NodePackage,
		Name: f.Name.Name,
		File: pkgPath,
	})

	ast.Inspect(f, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.FuncDecl:
			funcName := x.Name.Name
			nodeType := graph.NodeFunction
			id := fmt.Sprintf("func:%s:%s", pkgPath, funcName)

			if x.Recv != nil {
				nodeType = graph.NodeMethod
				recvType := p.getReceiverType(x.Recv)
				if recvType != "" {
					id = fmt.Sprintf("method:%s:%s.%s", pkgPath, recvType, funcName)
					// Add USES edge from method to its receiver struct/interface
					edges = append(edges, &graph.Edge{
						FromID: id,
						ToID:   fmt.Sprintf("type:%s:%s", pkgPath, recvType),
						Type:   graph.EdgeBelongsTo,
					})
				}
			}

			nodes = append(nodes, &graph.Node{
				ID:   id,
				Type: nodeType,
				Name: funcName,
				File: path,
				Line: p.fset.Position(x.Pos()).Line,
			})

			if x.Recv == nil {
				edges = append(edges, &graph.Edge{
					FromID: id,
					ToID:   pkgID,
					Type:   graph.EdgeBelongsTo,
				})
			}

		case *ast.TypeSpec:
			typeName := x.Name.Name
			var nodeType graph.NodeType
			switch x.Type.(type) {
			case *ast.StructType:
				nodeType = graph.NodeStruct
			case *ast.InterfaceType:
				nodeType = graph.NodeInterface
			default:
				return true
			}

			id := fmt.Sprintf("type:%s:%s", pkgPath, typeName)
			nodes = append(nodes, &graph.Node{
				ID:   id,
				Type: nodeType,
				Name: typeName,
				File: path,
				Line: p.fset.Position(x.Pos()).Line,
			})

			edges = append(edges, &graph.Edge{
				FromID: id,
				ToID:   pkgID,
				Type:   graph.EdgeBelongsTo,
			})
		}
		return true
	})

	return nodes, edges, nil
}

func (p *Parser) getReceiverType(recv *ast.FieldList) string {
	if recv == nil || len(recv.List) == 0 {
		return ""
	}
	typ := recv.List[0].Type
	for {
		if star, ok := typ.(*ast.StarExpr); ok {
			typ = star.X
			continue
		}
		break
	}
	if ident, ok := typ.(*ast.Ident); ok {
		return ident.Name
	}
	return ""
}

func (p *Parser) ExtractCalls(path string) ([]*graph.Edge, error) {
	f, err := parser.ParseFile(p.fset, path, nil, 0)
	if err != nil {
		return nil, err
	}

	var edges []*graph.Edge
	pkgPath := filepath.Dir(path)

	var currentFunc string

	ast.Inspect(f, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.FuncDecl:
			if x.Recv != nil {
				recvType := p.getReceiverType(x.Recv)
				currentFunc = fmt.Sprintf("method:%s:%s.%s", pkgPath, recvType, x.Name.Name)
			} else {
				currentFunc = fmt.Sprintf("func:%s:%s", pkgPath, x.Name.Name)
			}
		case *ast.CallExpr:
			if currentFunc == "" {
				return true
			}
			callTarget := p.getCallTarget(x)
			if callTarget != "" {
				// This is tricky because we don't know the package of the call target without type info
				// For now, assume it's in the same package or it's a qualified call
				targetID := ""
				if strings.Contains(callTarget, ".") {
					// Likely pkg.Func or receiver.Method
					// Real implementation needs go/types
					targetID = fmt.Sprintf("unknown:%s", callTarget)
				} else {
					targetID = fmt.Sprintf("func:%s:%s", pkgPath, callTarget)
				}

				edges = append(edges, &graph.Edge{
					FromID: currentFunc,
					ToID:   targetID,
					Type:   graph.EdgeCalls,
				})
			}
		}
		return true
	})

	return edges, nil
}

func (p *Parser) getCallTarget(ce *ast.CallExpr) string {
	switch x := ce.Fun.(type) {
	case *ast.Ident:
		return x.Name
	case *ast.SelectorExpr:
		if ident, ok := x.X.(*ast.Ident); ok {
			return fmt.Sprintf("%s.%s", ident.Name, x.Sel.Name)
		}
	}
	return ""
}
