package main 

import "fmt"

func doubleReturn(x int) (int, bool) {
  return x / 2, x % 2 == 0
}

func main() {
  fmt.Println(doubleReturn(10))
}