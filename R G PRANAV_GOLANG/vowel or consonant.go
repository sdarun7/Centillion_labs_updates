package main

import "fmt"

func main() {
    var char rune
    fmt.Print("Enter a character: ")
    fmt.Scanf("%c", &char)

    if char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' ||
        char == 'A' || char == 'E' || char == 'I' || char == 'O' || char == 'U' {
        fmt.Println("Vowel")
    } else {
        fmt.Println("Consonant")
    }
}
