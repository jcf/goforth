package main

import (
  "os"
  "fmt"
  "io/ioutil"
  "strings"
)

func noNulls(s string) string {
  return strings.ToUpper(s)
}

func main() {
  bytes, err := ioutil.ReadAll(os.Stdin)
  if err != nil {
    fmt.Println("ERR!?!?!?!: %v", err)
  }

  fmt.Println(noNulls(string(bytes)))
}
