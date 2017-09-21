package main 

import "fmt"

type Person struct {
  First string
  Last string
  Age int
}

type DoubleZero struct {
  // inherits from person
  Person
  First string
  LicenseToKill bool
}

func main() {
  // note embedded types are specified
  p1 := DoubleZero{
    Person: Person{
      First: "James",
      Last: "Bond",
      Age: 20,
    },
    First: "Double Zero",
    LicenseToKill: true,
  }

  fmt.Println(p1.First, p1.Person.First)
}