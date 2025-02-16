package main

import "fmt"

func main() {
    arr1 := []int{1, 2, 3, 4, 5}
    arr2 := make([]int, len(arr1))
    copy(arr2, arr1) 
    fmt.Println("Original array:", arr1)
    fmt.Println("Copied array:", arr2)
}
