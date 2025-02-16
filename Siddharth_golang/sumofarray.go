package main
import "fmt"

func sumArray(arr []int) int {
	sum := 0
	for _, num := range arr {
		sum += num
	}
	return sum
}

func main() {
	arr := []int{2, 4, 6, 8, 10}
	fmt.Println("Sum:", sumArray(arr))
}
