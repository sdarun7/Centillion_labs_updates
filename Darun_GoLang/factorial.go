package main
import "fmt"

func main() {
    var num int
    fmt.Print("Enter a number: ")
    fmt.Scan(&num)

    if num < 0 {
        fmt.Println("Factorial is not defined for negative numbers.")
        return
    }

    fact := 1
    for i := 1; i <= num; i++ {
        fact *= i
    }

    fmt.Println("Factorial of",num,"is",fact)
}
