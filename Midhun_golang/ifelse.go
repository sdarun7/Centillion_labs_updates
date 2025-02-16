package main

import "fmt"

func main() {
    var number int
    fmt.Print("Enter a number: ")
    fmt.Scanln(&number)

    if number > 0 {
        fmt.Println("The number is positive.")
    } else if number < 0 {
        fmt.Println("The number is negative.")
    } else {
        fmt.Println("The number is zero.")
    }
}
