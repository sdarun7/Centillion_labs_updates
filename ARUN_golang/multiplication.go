package main

import "fmt"
func generateTable(number, rangeLimit int) {
    sum := 0
    fmt.Printf("\nMultiplication table for %d (up to %d):\n", number, rangeLimit)
    for i := 1; i <= rangeLimit; i++ {
        result := number * i
        sum += result
        fmt.Printf("%d x %d = %d\n", number, i, result)
    }

    fmt.Printf("\nThe sum of the results from the table is: %d\n", sum)
}

func main() {
    var number, rangeLimit int

    fmt.Print("Enter a number to generate its multiplication table: ")
    fmt.Scanln(&number)

    fmt.Print("Enter the range limit (up to which number you want the table): ")
    fmt.Scanln(&rangeLimit)

    generateTable(number, rangeLimit)
}
