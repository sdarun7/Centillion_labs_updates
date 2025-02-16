package main

import "fmt"

func decimalToBinary(n int) string {
	binary := ""

	for n > 0 {
		binary = fmt.Sprintf("%d", n%2) + binary
		n /= 2
	}

	if binary == "" {
		binary = "0"
	}

	return binary
}

func main() {
	var num int
	fmt.Print("Enter a decimal number: ")
	fmt.Scan(&num)

	fmt.Println("Binary:", decimalToBinary(num))
}
