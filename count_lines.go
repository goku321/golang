package main

import (
  "os"
  "fmt"
  "bufio"
)

func main() {
  display_lines()
}

func display_lines() {
  counts := make(map[string]int)
  input := bufio.NewScanner(os.Stdin)

  for input.Scan() {
    counts[input.Text()]++
  }

  for lines, n := range counts {
    if n > 1 {
      fmt.Printf("%d\t%s\n", n, lines)
    }
  }
}
