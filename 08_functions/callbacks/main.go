package main 

import "fmt"

func forEach(numbers []int, callback func(int)) {
  for _, n := range numbers {
    callback(n)
  }
}

func main() {
  forEach([]int{1,2,3,4}, func(x int) {
    fmt.Println(x)
  })
}