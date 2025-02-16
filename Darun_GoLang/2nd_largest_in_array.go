package main
import "fmt"

func main() {
    var n int
    fmt.Print("Enter the size of the array: ")
    fmt.Scan(&n)

    if n < 2 {
        fmt.Println("Array should have at least 2 elements.")
        return
    }

    arr := make([]int, n)
    fmt.Println("Enter", n, "numbers:")
    for i := 0; i < n; i++ {
        fmt.Scan(&arr[i])
    }

    largest, secondLargest := arr[0], -1

    for i := 1; i < n; i++ {
        if arr[i] > largest {
            secondLargest = largest
            largest = arr[i]
        }
        else if arr[i] > secondLargest && arr[i] != largest {
            secondLargest = arr[i]
        }
    }

    if secondLargest == -1 {
        fmt.Println("No second largest number found.")
    } else {
        fmt.Println("Second largest number:",secondLargest)
    }
}
