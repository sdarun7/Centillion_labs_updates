package main
import "fmt"



func main() {
    var n int
    fmt.Print("enter n: ")
    fmt.Scan(&n)

    sieve := make([]bool, n+1)
    for i := 2; i <= n; i++ {
        sieve[i] = true
    }

    for i := 2; i*i <= n; i++ {
        if sieve[i] {
            for j := i * i; j <= n; j += i {
                sieve[j] = false
            }
        }
    }

    var primes []int
    for i := 2; i <= n; i++ {
        if sieve[i] {
            primes = append(primes, i)
        }
    }

    fmt.Println("Prime numbers:", primes)
}
