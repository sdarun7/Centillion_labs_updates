package main
import "fmt"

func main() {
    var a, b, temp int
    fmt.Print("Enter two numbers: ")
    fmt.Scan(&a, &b)

    fmt.Println("Before Swap: a =", a, "b =", b)

    temp = a
    a = b
    b = temp

    fmt.Println("After Swap: a =", a, "b =", b)
}
