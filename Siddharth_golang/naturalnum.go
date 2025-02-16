package main
import "fmt"

func sumNatural(n int) int {
	return n * (n + 1) / 2
}

func main() {
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)
	fmt.Println("Sum:", sumNatural(num))
}
