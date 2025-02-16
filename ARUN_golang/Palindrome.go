package main

import "fmt"
func isPalindrome(str string) bool {
    runes := []rune(str)
    for i := 0; i < len(runes)/2; i++ {
        if runes[i] != runes[len(runes)-1-i] {
            return false
        }
    }
    return true
}

func main() {
    var input string

    fmt.Print("Enter a string to check if it's a palindrome: ")
    fmt.Scanln(&input)

    if isPalindrome(input) {
        fmt.Printf("\"%s\" is a palindrome.\n", input)
    } else {
        fmt.Printf("\"%s\" is not a palindrome.\n", input)
    }
}
