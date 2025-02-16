package main

import "fmt"
func calculateSum(arr []int) int {
    sum := 0
    for _, num := range arr {
        sum += num
    }
    return sum
}
func findLargest(arr []int) int {
    largest := arr[0]
    for _, num := range arr {
        if num > largest {
            largest = num
        }
    }
    return largest
}
func findSmallest(arr []int) int {
    smallest := arr[0]
    for _, num := range arr {
        if num < smallest {
            smallest = num
        }
    }
    return smallest
}

func main() {
    var size int
    fmt.Print("Enter the size of the array: ")
    fmt.Scanln(&size)
    arr := make([]int, size)
    fmt.Println("Enter the elements of the array:")
    for i := 0; i < size; i++ {
        fmt.Printf("Element %d: ", i+1)
        fmt.Scanln(&arr[i])
    }
    sum := calculateSum(arr)

    largest := findLargest(arr)
    smallest := findSmallest(arr)
    fmt.Printf("\nArray Elements: %v\n", arr)
    fmt.Printf("Sum of elements: %d\n", sum)
    fmt.Printf("Largest element: %d\n", largest)
    fmt.Printf("Smallest element: %d\n", smallest)
}
