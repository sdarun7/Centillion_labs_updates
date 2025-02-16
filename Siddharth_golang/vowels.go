package main
import (
	"fmt"
	"strings"
)

func countVowels(s string) int {
	count := 0
	vowels := "aeiouAEIOU"
	for _, char := range s {
		if strings.ContainsRune(vowels, char) {
			count++
		}
	}
	return count
}

func main() {
	var input string
	fmt.Print("Enter string: ")
	fmt.Scanln(&input)
	fmt.Println("Vowel count:", countVowels(input))
}
