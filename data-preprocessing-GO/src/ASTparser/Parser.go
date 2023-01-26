package astParser

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"strings"
  "regexp"
)

func sanitizeName(name string) string {
  name = strings.ToLower(name)
  name = regexp.MustCompile(`[^a-zA-Z0-9 ]+`).ReplaceAllString(name, "")
  return name
}

func generateContext(node *ast.File) ([]string) {
  var context []string
  var parentPackage *ast.File

  ast.Inspect(node, func(n ast.Node) bool {

    switch x := n.(type) {
      case *ast.File:
        parentPackage = x     
      case *ast.FuncDecl:
    context = append(context, fmt.Sprintf("FuncDecl%s package%s", sanitizeName(x.Name.Name), sanitizeName(parentPackage.Name.Name)))
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

func getGoFileNode(path string) (*ast.File, error) {
  if filepath.Ext(path) == ".go" {
    fset := token.NewFileSet()

    node, err := parser.ParseFile(fset, path, nil, 0)
    if err != nil {
      return nil, err
    }
    return node, nil
  }
  return nil, nil
}

func write2DataFile(json []string, fileName string) error {
  f, err := os.OpenFile(fileName,
    os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
  if err != nil {
    return err
  }
  
  for i := 0;i < len(json);i++ {
    if _, err := f.WriteString(string(json[i]) + "\n"); err != nil {
      return err
    }
  }
  
  f.Close()
  return nil
}

func ParseDir(dir string) error {
  err := filepath.Walk(dir, func (path string, info os.FileInfo, err error) error {
    node, err := getGoFileNode(path)
    if node != nil {
      contextList := generateContext(node)
      json := context2json(contextList) 
      write2DataFile(json, "data.json")
    } else {
      if err != nil {
        log.Println(err)
      }
    }
    return nil
  })
  if err != nil {
  }
  return nil
}
