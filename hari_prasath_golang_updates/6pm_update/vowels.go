package main

import (
	"fmt"
	"unicode"
)

func main() {
	var s string
	fmt.Print("Enter the string val: ")
	fmt.Scan(&s)
	var vow, cons int

	for _, ch := range s {
		if unicode.IsLetter(ch) {
			low := unicode.ToLower(ch)
			if low == 'a' || low == 'e' || low == 'i' || low == 'o' || low == 'u' {
				vow++
			} else {
				cons++
			}
		}
	}

	fmt.Printf("Vowels: %d\nConsonants: %d\n", vow, cons)
}
