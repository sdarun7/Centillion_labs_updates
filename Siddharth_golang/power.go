package main

import "fmt"

func power(base, exp int) int {
	result := 1
	for i := 0; i < exp; i++ {
		result *= base
	}
	return result
}

func main() {
	var base, exp int
	fmt.Print("Enter base and exponent: ")
	fmt.Scan(&base, &exp)
	fmt.Println("Result:", power(base, exp))
}
