package main

import (
	"fmt"
	"sort"
)

func main() {
	numbers := []int{10, 20, 4, 45, 99, 50}
	sort.Ints(numbers)
	fmt.Println("Second largest number:", numbers[len(numbers)-2])
}
