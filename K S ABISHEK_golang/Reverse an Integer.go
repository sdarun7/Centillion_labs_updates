package main

import "fmt"

func reverseNumber(n int) int {
    rev := 0
    for n > 0 {
        remainder := n % 10
        rev = rev*10 + remainder
        n /= 10
    }
    return rev
}

func main() {
    var num int
    fmt.Print("Enter a number: ")
    fmt.Scan(&num)
    fmt.Println("Reversed number:", reverseNumber(num))
}
