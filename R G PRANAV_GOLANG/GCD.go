package main

import "fmt"

func gcd(a, b int) int {
    for b != 0 {
        a, b = b, a%b
    }
    return a
}

func main() {
    fmt.Println("GCD of 36 and 60:", gcd(36, 60))
}
