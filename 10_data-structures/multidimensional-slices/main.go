package main 

import "fmt"

func main() {
  records := make([][]string, 0)

  fmt.Println(records)

  student1 := make([]string, 4)
  student1[0] = "Piggot"
  student1[1] = "Ben"
  student1[2] = "100.00"
  student1[3] = "74.00"

  records = append(records, student1)

  fmt.Println(records)

  student2 := make([]string, 4)
  student2[0] = "Plett"
  student2[1] = "Bethany"
  student2[2] = "97.00"
  student2[3] = "87.00"

  records = append(records, student2)

  fmt.Println(records)
}