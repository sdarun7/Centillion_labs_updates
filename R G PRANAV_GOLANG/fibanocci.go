package main

import "fmt"

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	var n int
	fmt.Print("Enter the number of terms: ")
	fmt.Scan(&n)

	fmt.Println("Fibonacci Series:")
	for i := 0; i < n; i++ {
		fmt.Print(fibonacci(i), " ")
	}
	fmt.Println()
}
