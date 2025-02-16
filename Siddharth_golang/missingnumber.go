package main

import "fmt"

func findMissingNumber(arr []int, n int) int {
	total := (n * (n + 1)) / 2
	sum := 0
	for _, num := range arr {
		sum += num
	}
	return total - sum
}

func main() {
	numbers := []int{1, 2, 3, 5}
	fmt.Println("Missing number:", findMissingNumber(numbers, 5))
}
