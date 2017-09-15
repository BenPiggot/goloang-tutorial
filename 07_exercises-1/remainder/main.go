package main 

import "fmt"

func main() {
  var x int
  var y int

  fmt.Println("Enter a large number: ")
  fmt.Scan(&x)
  fmt.Println("Enter a small number: ")
  fmt.Scan(&y)
  remainder := x % y
  fmt.Println("The remainder of large number divided is", remainder)
}