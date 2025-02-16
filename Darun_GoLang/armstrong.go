package main
import "fmt"

func main() {
    var num, temp, remainder, result int
    fmt.Print("Enter a number: ")
    fmt.Scan(&num)

    temp = num
    for temp > 0 {
        remainder = temp % 10
        result += remainder * remainder * remainder
        temp /= 10
    }

    if result == num {
        fmt.Println("Armstrong Number")
    }
    else {
        fmt.Println("Not an Armstrong Number")
    }
}
