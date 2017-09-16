package main 

import "fmt"

func main() {
  greeting := func(name string) {
    fmt.Println("Hello " + name)
  }

  greeting("Ben")
}