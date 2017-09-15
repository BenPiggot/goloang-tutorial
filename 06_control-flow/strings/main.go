package main 

import "fmt"

func main() {
  for i := 660; i < 760; i++ {
    fmt.Println(i, " - ", string(i), " - ", rune(i), " - ", []byte(string(i)))
  }
}