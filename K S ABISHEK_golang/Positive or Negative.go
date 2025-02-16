package main

import "fmt"

func main() {
    var num int
    fmt.Print("Enter a number: ")
    fmt.Scan(&num)

    if num > 0 {
        fmt.Println(num, "is positive")
    } else if num < 0 {
        fmt.Println(num, "is negative")
    } else {
        fmt.Println("The number is zero")
    }
}
