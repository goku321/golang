package main

import (
  "os"
  "bufio"
  "fmt"
)

func main() {
  counts := make(map[string]int)
  files := os.Args[1:]
  if len(files) == 0 {
    countLines(os.Stdin, counts)
  } else
}
