package main
import "fmt"
func main() {
    var a, b int
    fmt.Print("Enter two numbers: ")
    fmt.Scan(&a, &b)
    a, b = b, a
    fmt.Println("After swapping: a =", a, "b =", b)
}
