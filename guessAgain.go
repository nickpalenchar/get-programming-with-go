package main

import (
  "fmt"
  "math/rand"
)

func main() {

  myNum := rand.Intn(100)

  for {
    theGuess := rand.Intn(100)
    if (theGuess > myNum) {
      fmt.Println("lower...")
    } else if (theGuess < myNum) {
      fmt.Println("higher...")
    } else {
      fmt.Printf("got it! (%v)", myNum)
      break
    }
  }
}
