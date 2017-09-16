package main 

import "fmt"

func main() {
  fmt.Println(greatest([]int{4324, 232, 55, 23255, 3223}))
}

func greatest(numbers []int) int {

  var largest int

  for _, n := range numbers {
    if n >= largest {
      largest = n
    }
  }

  return largest
}
