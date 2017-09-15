package main 

import "fmt"

func main() {
  n := average(323, 23423, 322, 64, 94340)
  fmt.Println(n)

  // now with variadic arguments
  data := []float64{42, 41, 90, 34}
  // data... spreads out slice into individual args
  m := average(data...)
  fmt.Println(m)
}

func average(nums ...float64) float64 {
  var total float64
  for _, n := range nums {
    total += n
  } 
  return total / float64(len(nums))
}