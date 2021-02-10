package main

import (
  "fmt"

)

func main() {

  fmt.Printf("%-17v %+4v %-12v %+5v\n", "Spaceline", "Days", "Trip type", "Price")
  fmt.Printf("=========================================\n")
  fmt.Printf("%-17v %+4v %-12v $%+4v\n", "Virgin Galactic", "1", "rount-trip", "5")

}
