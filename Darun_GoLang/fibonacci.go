package main
import "fmt"

func main() {
    var n int
    fmt.Print("Enter the number of terms: ")
    fmt.Scan(&n)

    if n <= 0 {
        fmt.Println("Please enter a positive number.")
        return
    }

    a, b := 0, 1
    fmt.Print(a," ",b," ")

    for i := 2; i < n; i++ {
        next := a + b
        fmt.Print(next, " ")
        a, b = b, next
    }
    fmt.Println()
}
