package main

import (
	"fmt"
	"sort"
)

func main() {
	numbers := []int{7, 2, 1, 6, 8, 5, 3, 4}
	sort.Ints(numbers)
	n := len(numbers)

	var median float64
	if n%2 == 1 {
		median = float64(numbers[n/2])
	} else {
		median = float64(numbers[n/2-1]+numbers[n/2]) / 2.0
	}

	fmt.Println("Median:", median)
}
