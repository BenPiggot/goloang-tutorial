package main 

import "fmt"

func main() {
  b := true

  // can initialize variable right before running
  // the variable 'food' only in scope in if block
  // which prevents scope pollution
  if food := "Pizza"; b {
    fmt.Println(food)
  }
}