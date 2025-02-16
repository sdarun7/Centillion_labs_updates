package main

import "fmt"

func main() {
    var a, b int
    fmt.Print("Enter two numbers: ")
    fmt.Scan(&a, &b)

    if a > b {
        fmt.Println(a, "is larger")
    } else {
        fmt.Println(b, "is larger")
    }
}
