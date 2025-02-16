package main

import "fmt"

func main() {
    var n int
    fmt.Print("Enter the number of elements: ")
    fmt.Scan(&n)

    arr := make([]int, n)
    fmt.Println("Enter the elements:")
    for i := 0; i < n; i++ {
        fmt.Scan(&arr[i])
    }

    max := arr[0]
    for _, value := range arr {
        if value > max {
            max = value
        }
    }

    fmt.Println("Largest element:", max)
}
