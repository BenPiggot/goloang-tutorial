package main

import "fmt"

const y string = "Hello World"

func main() {
  const x = 42

  fmt.Printf("x -", x)
  fmt.Printf("y -", y)
}