package main

import (
  "os" 
  "github.com/Gabriel-Ivarsson/code2vec-demo/data-preprocessing-GO/src/ASTparser"
)

func main() {
  dir := os.Args[1]
  astParser.ParseDir(dir) 
  return
}
