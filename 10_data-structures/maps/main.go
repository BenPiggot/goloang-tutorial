package main 

import "fmt"

func main() {
  myGreeting := make(map[string]string)

  myGreeting["Ben"] = "Hello"
  myGreeting["Bethany"] = "Hola"
  myGreeting["Soren"] = "Bonjour"

  fmt.Println(myGreeting)

  // delete entry
  delete(myGreeting, "Soren")

  fmt.Println(myGreeting)
}