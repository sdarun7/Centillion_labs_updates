package main

import "fmt"

func main() {
    var terms int
    fmt.Print("Enter the number of terms: ")
    fmt.Scanln(&terms)

    a, b := 0, 1
    fmt.Println("Fibonacci Sequence:")
    for i := 0; i < terms; i++ {
        fmt.Print(a, " ")
        a, b = b, a+b
    }
    fmt.Println()
}
