package main

import (
  "fmt"
  "path/filepath"
)

func main() {
  files, err := filepath.Glob("src/*.go")
  if err != nil {
    fmt.Printf("Stop the world! Things gone bad! %v\n", err)
  }

  for _, file := range files  {
    fmt.Printf("Found %+v\n", file)
  }
}
