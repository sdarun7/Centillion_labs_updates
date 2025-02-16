package main

import "fmt"

func main() {
    var str string
    fmt.Print("Enter a string: ")
    fmt.Scanln(&str)
    byteLength := len(str)
    runeLength := len([]rune(str))
    fmt.Printf("Length of the string in bytes: %d\n", byteLength)
    fmt.Printf("Length of the string in characters: %d\n", runeLength)
}
