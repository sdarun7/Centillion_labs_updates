package main
import "fmt"
func fibonacci(n int) []int {
    fibSequence := make([]int, n)
    if n > 0 {
        fibSequence[0] = 0
    }
    if n > 1 {
        fibSequence[1] = 1
    }
    for i := 2; i < n; i++ {
        fibSequence[i] = fibSequence[i-1] + fibSequence[i-2]
    }
    return fibSequence
}
func main() {
    var num int
    fmt.Print("Enter the number of terms in the Fibonacci sequence: ")
    fmt.Scanln(&num)
    if num <= 0 {
        fmt.Println("Please enter a positive number greater than 0.")
        return
    }
    fibSeq := fibonacci(num)
    fmt.Printf("The first %d terms of the Fibonacci sequence are:\n", num)
    for _, value := range fibSeq {
        fmt.Print(value, " ")
    }
    fmt.Println()
}
