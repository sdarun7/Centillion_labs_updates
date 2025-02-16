package main
import "fmt"

func main() {
    var num int
    fmt.Print("Enter a number: ")
    fmt.Scan(&num)

    if num < 2 {
        fmt.Println(num,"is NOT a Prime Number")
        return
    }

    isPrime := true
    for i := 2; i*i <= num; i++ {
        if num%i == 0 {
            isPrime = false
            break
        }
    }

    if isPrime {
        fmt.Println(num,"is a Prime Number")
    }
    else {
        fmt.Println(num,"is NOT a Prime Number")
    }
}
