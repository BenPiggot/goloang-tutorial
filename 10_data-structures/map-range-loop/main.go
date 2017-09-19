package main 

import "fmt"

func main() {
  greetings := map[int]string{
    0: "Hello",
    1: "Hola",
    2: "Bonjour",
    3: "Buongiorno",
  }

  for key, val := range greetings {
    fmt.Println(key, " - ", val)
  }
}