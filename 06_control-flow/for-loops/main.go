package main 

import "fmt"

func main() {
  for i := 0; i <= 100; i++ {
    fmt.Println(i)
  }

  for i := 0; i <= 10; i++ {
    for j := 0; j <=10; j++ {
      fmt.Println(i, " - ", j)
    }
  }

  k := 0
  for k < 10 {
    fmt.Println(k)
    k++
  }
}