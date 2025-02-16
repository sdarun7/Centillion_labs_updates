package main
import "fmt"
func reverseNumber(n int) int {
	reversed := 0
	for n != 0 {
		reversed = reversed*10 + n%10
		n /= 10
	}
	return reversed
}
func main() {
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)
	fmt.Println("Reversed number:", reverseNumber(num))
}
