package main 

import "fmt"

func main() {
  a := 43

  fmt.Println(a)
  fmt.Println(&a)

  var b *int = &a 

  fmt.Println(b)
  fmt.Println(*b) // deferencing

  *b = 42 // change the value at the memory address b represents
  fmt.Println(a)
}