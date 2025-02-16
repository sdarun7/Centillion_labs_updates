package main
import "fmt"

func main() {
    var arr [3]int
    fmt.Print("Enter 3 numbers: ")
    for i := 0; i < 3; i++ {
        fmt.Scan(&arr[i])
    }

    ptr := &arr // Pointer to array
    fmt.Println("First element using pointer:",(*ptr)[0])
}
