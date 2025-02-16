package main
import "fmt"
func findLargest(arr []int) int {
	largest := arr[0]
	for _, num := range arr {
		if num > largest {
			largest = num
		}
	}
	return largest
}
func main() {
	arr := []int{12, 45, 23, 78, 56, 89, 90}
	fmt.Println("Largest number:", findLargest(arr))
}
