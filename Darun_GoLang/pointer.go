package main
import "fmt"

func main() {
    var num int
    fmt.Print("Enter a number: ")
    fmt.Scan(&num)

    var ptr *int = &num

    fmt.Println("You entered:",*ptr)
}
