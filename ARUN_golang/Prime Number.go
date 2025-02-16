package main
import "fmt"
func isPrime(n int) bool {
    if n <= 1 {
        return false 
    }
    for i := 2; i*i <= n; i++ {
        if n%i == 0 {
            return false 
        }
    }
    return true 
}
func main() {
    var number int
    fmt.Print("Enter a number to check if it is prime: ")
    fmt.Scanln(&number)
    if isPrime(number) {
        fmt.Printf("%d is a prime number.\n", number)
    } else {
        fmt.Printf("%d is not a prime number.\n", number)
    }
}
