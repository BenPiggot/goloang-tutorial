package main 

import "fmt"

func main() {
  doubleReturn := func(x int) (int, bool) {
    return x / 2, x % 2 == 0
  }

  fmt.Println(doubleReturn(12))
}