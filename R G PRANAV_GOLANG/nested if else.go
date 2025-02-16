package main

import "fmt"

func main() {
    var age int
    fmt.Print("Enter your age: ")
    fmt.Scan(&age)

    if age >= 18 {
        fmt.Println("You are eligible to vote.")
        if age >= 60 {
            fmt.Println("You are a Senior Citizen.")
        } else {
            fmt.Println("You are an Adult.")
        }
    } else {
        fmt.Println("You are NOT eligible to vote.")
    }
}
