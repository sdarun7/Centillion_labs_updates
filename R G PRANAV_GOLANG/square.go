package main

import "fmt"

func square(n int) int {
    return n * n
}

func main() {
    num := 5
    fmt.Println("Square of", num, "is", square(num))
}
