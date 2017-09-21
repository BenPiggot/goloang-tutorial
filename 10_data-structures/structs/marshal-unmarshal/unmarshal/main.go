package main 

import (
  "encoding/json"
  "fmt"
)

type Person struct {
  First string
  Last string
  Age int
}

func main() {
  var p1 Person

  byteSlice := []byte(`{"First":"Ben","Last":"Piggot","Age":20}`)
  json.Unmarshal(byteSlice, &p1)

  fmt.Println(p1.Last)
  fmt.Println(p1.Age)
  fmt.Printf("%T \n", p1)
}