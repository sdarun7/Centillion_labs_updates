package main

import "fmt"

func main() {
    var numCount int
    fmt.Print("Enter the number of elements: ")
    fmt.Scanln(&numCount)
    numbers := make([]float64, numCount)
    for i := 0; i < numCount; i++ {
        fmt.Printf("Enter number %d: ", i+1)
        fmt.Scanln(&numbers[i])
    }
    sum := 0.0
    for _, number := range numbers {
        sum += number
    }
    average := sum / float64(numCount)
    fmt.Printf("The average of the numbers is: %.2f\n", average)
}
