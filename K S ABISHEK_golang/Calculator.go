package main

import "fmt"

func main() {
    var a, b int
    fmt.Print("Enter two numbers: ")
    fmt.Scan(&a, &b)

    fmt.Println("Addition:", a+b)
    fmt.Println("Subtraction:", a-b)
    fmt.Println("Multiplication:", a*b)
    fmt.Println("Division:", a/b)
}
