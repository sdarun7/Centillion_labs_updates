package main

import "fmt"

func main() {
    var num int
    fmt.Print("Enter a number: ")
    fmt.Scan(&num)

    square := num * num
    fmt.Println("Square of", num, "is", square)
}
