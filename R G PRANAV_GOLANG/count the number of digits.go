package main

import "fmt"

func countDigits(n int) int {
    count := 0
    for n != 0 {
        n /= 10
        count++
    }
    return count
}

func main() {
    var num int
    fmt.Print("Enter a number: ")
    fmt.Scan(&num)
    fmt.Println("Number of digits:", countDigits(num))
}
