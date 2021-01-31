package main

import (
  "fmt"
  "rand"
)

func main() {
  var degrees = 0

  for {
    fmt.Println(degrees)
    if degrees >= 360 {
      degrees = 0
      if rand.Intn(2) == 0 {
        break
      }
    }
  }
}
