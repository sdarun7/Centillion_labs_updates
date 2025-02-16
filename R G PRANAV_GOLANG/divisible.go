package main

import "fmt"

func main() {
    var num int
    fmt.Print("Enter a number: ")
    fmt.Scan(&num)

    if num%3 == 0 && num%5 == 0 {
        fmt.Println(num, "is divisible by both 3 and 5")
    } else if num%3 == 0 {
        fmt.Println(num, "is divisible by 3")
    } else if num%5 == 0 {
        fmt.Println(num, "is divisible by 5")
    } else {
        fmt.Println(num, "is NOT divisible by 3 or 5")
    }
}
