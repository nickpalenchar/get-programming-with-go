package main

import (
  "fmt"
  "math/rand"
)

func main() {

  var myNum = rand.Intn(100)

  for {
    var guess = rand.Intn(100)
    if guess == myNum {
      fmt.Println("You got it!")
      break
    } else if guess > myNum {
      fmt.Println("too high...")
    } else {
      fmt.Println("too low...")
    }
  }
}
