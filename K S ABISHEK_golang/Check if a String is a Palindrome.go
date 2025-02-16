package main

import (
    "fmt"
    "strings"
)

func isPalindrome(s string) bool {
    s = strings.ToLower(s)
    length := len(s)
    for i := 0; i < length/2; i++ {
        if s[i] != s[length-1-i] {
            return false
        }
    }
    return true
}

func main() {
    var str string
    fmt.Print("Enter a string: ")
    fmt.Scan(&str)
    if isPalindrome(str) {
        fmt.Println(str, "is a palindrome")
    } else {
        fmt.Println(str, "is not a palindrome")
    }
}
