package main

import (
  "fmt"
  "go/ast"
  "go/parser"
  "go/token"
  "log"
  "os"
)

func generateJSON(node *ast.File) (error) {
  var data []string
  ast.Inspect(node, func(n ast.Node) bool {
    switch x := n.(type) {
      case *ast.FuncDecl:
        fmt.Println("Function Name:", x.Name.Name)
        append(data, "\"FuncDecl\":")
        append(data, x.Name.Name)
        fmt.Println("Function Parameters:")
        for _, param := range x.Type.Params.List {
            for _, name := range param.Names {
                fmt.Println("\t", name.Name)
            }
        }
       
      case *ast.CallExpr:
        fmt.Println("Call expression:", x.Fun)
        for _, arg := range x.Args {
          fmt.Println("\tArg:", arg)
        }
      case *ast.AssignStmt:
        fmt.Println("Assignment statement:")
        for _, lhs := range x.Lhs {
          fmt.Println("\tLHS:", lhs)
        }
        for _, rhs := range x.Rhs {
          fmt.Println("\tRHS:", rhs)
        }
    }
        return true
    })
  return nil
}

func main() {
  fmt.Printf("Starting parsing...")
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, os.Args[1], nil, 0)
	if err != nil {
		log.Println(err)
		return
	}

	ast.Print(fset, node)
  generateJSON(node)
}
