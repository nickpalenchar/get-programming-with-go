package main

import (
  "fmt"
  "math/rand"
)

var era = "AD"

func main() {
  switch num := rand.Intn(10); num {
  case 0:
    fmt.Println("Space Adventures")
  case 1:
    fmt.Println("Asteroids")
  case 2:
    fmt.Println("Mario")
  case 4:
    fmt.Println("sonig")
  case 5:
    fmt.Println("Saniccc")
  default:
    fmt.Println("chores...")
  }

  fmt.Println(num)
  fmt.Println(era)
}
