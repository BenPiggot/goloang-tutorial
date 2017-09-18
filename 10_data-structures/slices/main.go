package main 

import "fmt"

func main() {
  mySlice := make([]int, 0, 5)

  fmt.Println("----------")
  fmt.Println(mySlice)
  fmt.Println(len(mySlice))
  fmt.Println(cap(mySlice))
  fmt.Println("----------")

  for i := 0; i < 80; i++ {
    mySlice = append(mySlice, i)
    fmt.Println("Length: ", len(mySlice), "Capacity: ", cap(mySlice), "Value: ", mySlice[i])
  }

  // append a slice to a slice
  myOtherSlice := []int{500, 600, 700}
  fmt.Println(myOtherSlice)
  mySlice = append(mySlice, myOtherSlice...)
  fmt.Println(mySlice)

  // delete from a slice
  myOtherSlice = append(myOtherSlice[:1], myOtherSlice[2:]...)
  fmt.Println(myOtherSlice)
}