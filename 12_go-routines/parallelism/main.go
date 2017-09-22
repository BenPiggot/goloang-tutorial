package main 

import (
  "fmt"
  "sync"
  "time"
  "runtime"
)

var wg sync.WaitGroup

// init is a special function that runs before any other code executes
// no longer necessary to specify more than one core after Go 1.5
func init() {
  runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
  wg.Add(2)
  go foo()
  go bar() 
  wg.Wait()
}

func foo() {
  for i := 0; i < 50; i++ {
    fmt.Println("Foo: ", i)
    time.Sleep(time.Duration(3 * time.Millisecond))
  }
  wg.Done()
}

func bar() {
  for i := 0; i < 50; i++ {
    fmt.Println("Bar: ", i)
    time.Sleep(time.Duration(6 * time.Millisecond))
  }
  wg.Done()
}