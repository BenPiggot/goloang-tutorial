package main 

import (
  "encoding/json"
  "os"
)

type Person struct {
  First string
  Last string
  Age int
  notExported int
}

func main() {
  p1 := Person{"Ben", "Piggot", 41, 20}
  json.NewEncoder(os.Stdout).Encode(p1)
}