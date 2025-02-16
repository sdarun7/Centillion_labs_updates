package main
import "fmt"

func main() {
    var num, temp, remainder, reverse int
    fmt.Print("Enter a number: ")
    fmt.Scan(&num)

    temp = num
    for temp > 0 {
        remainder = temp % 10
        reverse = reverse*10 + remainder
        temp /= 10
    }

    if reverse == num {
        fmt.Println("Palindrome Number")
    } else {
        fmt.Println("Not a Palindrome Number")
    }
}
