package main

import "fmt"  

func main() {
  x := 42

  fmt.Println("x - ", x)
  fmt.Println("Address of x - ", &x)
  fmt.Printf("%d\n", &x)
}