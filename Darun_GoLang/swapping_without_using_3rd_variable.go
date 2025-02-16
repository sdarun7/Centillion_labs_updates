package main
import "fmt"

func main() {
    var a, b int
    fmt.Print("Enter two numbers: ")
    fmt.Scan(&a, &b)

    fmt.Println("Before Swap: a =", a, "b =", b)

    a = a + b
    b = a - b
    a = a - b

    fmt.Println("After Swap: a =", a, "b =", b)
}
