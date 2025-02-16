package main

import "fmt"

func main() {
    var num int
    fmt.Print("Enter a number: ")
    fmt.Scan(&num)

    if num > 0 {
        fmt.Println(num, "is Positive")
    } else if num < 0 {
        fmt.Println(num, "is Negative")
    } else {
        fmt.Println("The number is Zero")
    }
}
