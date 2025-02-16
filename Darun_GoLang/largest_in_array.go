package main
import "fmt"

func main() {
    var n int
    fmt.Print("Enter the size of the array: ")
    fmt.Scan(&n)

    arr := make([]int, n)
    fmt.Println("Enter", n, "numbers:")
    for i := 0; i < n; i++ {
        fmt.Scan(&arr[i])
    }

    max := arr[0]
    for i := 1; i < n; i++ {
        if arr[i] > max {
            max = arr[i]
        }
    }

    fmt.Println("Largest number:",max)
}
