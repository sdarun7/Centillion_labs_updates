package main
import "fmt"
func countVowels(str string) int {
    vowels := "aeiouAEIOU"
    count := 0
    for _, char := range str {
        if string(char) == string(vowels) {
            count++
        }
    }
    return count
}
func main() {
    var input string
    fmt.Print("Enter a string: ")
    fmt.Scanln(&input)
    vowelCount := countVowels(input)
    fmt.Printf("The number of vowels in the string is: %d\n", vowelCount)
}
