package main 

import "fmt"

func forEach(numbers []int, callback func(int)) {
  for _, n := range numbers {
    callback(n)
  }
}

func filter(numbers []int, callback func(int) bool) []int {
  filtered := []int{}
  for _, n := range numbers {
    if callback(n) {
      filtered = append(filtered, n)
    }
  }

  return filtered
}

func mapIt(numbers []int, callback func(int) int) []int {
  mapped := []int{}
  for _, n := range numbers {
    mapped = append(mapped, callback(n))
  }

  return mapped
}

func main() {
  forEach([]int{1,2,3,4}, func(x int) {
    fmt.Println(x)
  })
}