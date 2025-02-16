package main

import (
	"fmt"
	"math"
)

func isArmstrong(n int) bool {
	temp, sum, digits := n, 0, 0
	for temp > 0 {
		digits++
		temp /= 10
	}

	temp = n
	for temp > 0 {
		sum += int(math.Pow(float64(temp%10), float64(digits)))
		temp /= 10
	}

	return sum == n
}

func main() {
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)

	if isArmstrong(num) {
		fmt.Println("Armstrong number")
	} else {
		fmt.Println("Not an Armstrong number")
	}
}
