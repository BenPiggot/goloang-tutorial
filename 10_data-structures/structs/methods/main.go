package main 

import "fmt"

type person struct {
  first string
  last string
  age int
}

// note this function has a receiver, which attaches a function to a type
func (p person) fullName() string {
  return p.first + " " + p.last
}

func main() {
  p1 := person{"Ben", "Piggot", 41}

  fmt.Println(p1)
  fmt.Println(p1.fullName())
}