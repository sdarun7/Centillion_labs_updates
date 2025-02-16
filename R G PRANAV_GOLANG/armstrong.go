package main

import (
	"fmt"
	"math"
)

func isArmstrong(num int) bool {
	originalNum, sum, temp := num, 0, num
	count := len(fmt.Sprint(num)) // Count digits

	for temp > 0 {
		digit := temp % 10
		sum += int(math.Pow(float64(digit), float64(count)))
		temp /= 10
	}

	return sum == originalNum
}

func main() {
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)

	if isArmstrong(num) {
		fmt.Println(num, "is an Armstrong number.")
	} else {
		fmt.Println(num, "is NOT an Armstrong number.")
	}
}
