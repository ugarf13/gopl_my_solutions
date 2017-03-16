// Exercise 1.2 GOPL 
package main

import (
    "fmt"
    "os"
)

func main() {
    var i int
    var arg string
    for i, arg = range os.Args[1:] {
        fmt.Println(i + 1, arg)
    }
}
