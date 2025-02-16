package main
import "fmt"

func mergeSortedArrays(arr1, arr2 []int) []int {
	i, j := 0, 0
	var merged []int
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			merged = append(merged, arr1[i])
			i++
		} else {
			merged = append(merged, arr2[j])
			j++
		}
	}
	return append(merged, arr1[i:]...)
}

func main() {
	arr1 := []int{1, 3, 5}
	arr2 := []int{2, 4, 6}
	fmt.Println("Merged:", mergeSortedArrays(arr1, arr2))
}
