package main

import "fmt"

func main() {
    var a, b int
    fmt.Print("Enter two numbers: ")
    fmt.Scan(&a, &b)

    a, b = b, a // Swapping without using a temporary variable

    fmt.Println("After swapping:")
    fmt.Println("a =", a, "b =", b)
}
