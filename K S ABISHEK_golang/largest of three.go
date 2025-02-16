package main

import "fmt"

func main() {
    var a, b, c int
    fmt.Print("Enter three numbers: ")
    fmt.Scan(&a, &b, &c)

    if a >= b && a >= c {
        fmt.Println(a, "is the largest number")
    } else if b >= a && b >= c {
        fmt.Println(b, "is the largest number")
    } else {
        fmt.Println(c, "is the largest number")
    }
}
