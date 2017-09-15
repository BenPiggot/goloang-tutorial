package main 

import "fmt"

func main() {
  fmt.Println(multipleGreet("Ben ", "Piggot"))
}

func greet(fname string, lname string) string {
  return fmt.Sprint(fname, lname)
}

// named returned
func namedGreet(fname string, lname string) (s string) {
  s = fmt.Sprint(fname, lname)
  return
}

// return multiple values
func multipleGreet(fname string, lname string) (string, string) {
  return fmt.Sprint(fname, lname), fmt.Sprint(lname, fname)
}