package main

import "fmt"

func main() {
    var a, b int
    fmt.Print("Enter two numbers: ")
    fmt.Scan(&a, &b)
    fmt.Printf("The sum of %d and %d is %d\n", a, b, a+b)
}
