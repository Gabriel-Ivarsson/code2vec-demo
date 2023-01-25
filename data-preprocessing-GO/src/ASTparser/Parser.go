package astParser

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
  "path/filepath"
)

func generateContext(node *ast.File) ([]string) {
  var context []string
  var parentPackage *ast.File

  ast.Inspect(node, func(n ast.Node) bool {

    switch x := n.(type) {
      case *ast.File:
        context = append(context, fmt.Sprintf("%s is package", x.Name.Name))
        parentPackage = x     
      case *ast.FuncDecl:
        context = append(context, fmt.Sprintf("%s is function in %s", x.Name.Name, parentPackage.Name.Name))
    }
    return true
  })
  return context 
}

func context2json(contextList []string) []string {
  var json []string
  for i := 0;i < len(contextList);i++ {
    jsonLine := "{\"context\":" + "\"" + contextList[i] + "\"}"
    json = append(json, jsonLine)
  }
  return json
}

func ParseDir(dir string) error {
  filepath.Walk(dir, func (path string, info os.FileInfo, err error) error {
    if filepath.Ext(path) == ".go" {
      fset := token.NewFileSet()

      node, err := parser.ParseFile(fset, path, nil, 0)
      if err != nil {
       log.Println(err)
        return err
      }

      contextList := generateContext(node)
      json := context2json(contextList) 
      f, err := os.OpenFile("data.json",
        os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
      if err != nil {
        log.Println(err)
      }

      for i := 0;i < len(json);i++ {
        if _, err := f.WriteString(string(json[i]) + "\n"); err != nil {
          log.Println(err)
          return err
        }
      }

      f.Close()
    }
    return nil
  })
  return nil
}
