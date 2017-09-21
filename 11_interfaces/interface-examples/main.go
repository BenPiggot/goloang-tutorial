package main 

import (
  "fmt"
  "math"
)

type Square struct {
  side float64
}

type Circle struct {
  radius float64
}

type Shape interface {
  area() float64
}

// Square implements Shape interface
func (z Square) area() float64 {
  return z.side * z.side
}

// Circle implements Shape interface
func (c Circle) area() float64 {
  return math.Pi * c.radius * c.radius
}

func info(z Shape) {
  fmt.Println(z)
  fmt.Println(z.area())
}

func main() {
  s := Square{10}
  info(s)

  c := Circle{5}
  info(c)
}