package main

import (
    "fmt"
    "time"
)

func main() {
    var count = 10

    for count > 0 {
        fmt.Println(count)
        time.Sleep(time.Second)
        count--
        if count % 5 == 0 {
            fmt.Println("btw, weird how a for loop can completely omit other conditions.. but I like it")
        }
    }
}
