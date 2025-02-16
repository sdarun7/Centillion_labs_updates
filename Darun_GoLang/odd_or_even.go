package main
import "fmt"

func main() {
    var num int
    fmt.Print("Enter a number: ")
    fmt.Scan(&num) // Take input from user

    if num%2 == 0 {
        fmt.Println(num, "is Even")
    } else {
        fmt.Println(num, "is Odd")
    }
}
