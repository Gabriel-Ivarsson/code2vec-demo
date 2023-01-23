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
  var typeSpecifier = make(map[string]string)
  var name []string
  var context = make(map[string]string)
  var parentPackage *ast.File
  var parentFunc *ast.FuncDecl

  ast.Inspect(node, func(n ast.Node) bool {

    switch x := n.(type) {
      case *ast.File:
        typeSpecifier[x.Name.Name] = "PackageDecl"
        name = append(name, x.Name.Name)
        context[x.Name.Name] += fmt.Sprintf("%s is the parent package\n", x.Name.Name) 
        parentPackage = x     
      case *ast.FuncDecl:
        typeSpecifier[x.Name.Name] = "FuncDecl"
        name = append(name, x.Name.Name)
        context[x.Name.Name] +=  fmt.Sprintf("%s is a child of package %s\n", x.Name.Name, parentPackage.Name.Name)
        parentFunc = x
      case *ast.CallExpr:
        methodName :=  fmt.Sprintf("%s", x.Fun)
        typeSpecifier[methodName] = "MethodCall"
        name = append(name, methodName)
        context[methodName] += fmt.Sprintf("%s is a child of function %s\n", methodName, parentFunc.Name.Name)
    }
        return true
    })
  for i := 0;i < len(name);i++ {
    fmt.Printf("Name: %s\n", name[i])
    fmt.Printf("Type: %s\n", typeSpecifier[name[i]])
    fmt.Printf("Context: %s\n", context[name[i]])
  }
  /*
  file, err := os.OpenFile("code-data.json", os.O_APPEND | os.O_CREATE, 0644)  
  if err != nil {
    return err
  }
  */
  return nil
}

func main() {
  fmt.Printf("Starting parsing...\n")
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, os.Args[1], nil, 0)
	if err != nil {
		log.Println(err)
		return
	}
  generateJSON(node)
}
