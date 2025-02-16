package main
import "fmt"

func toBinary(n int) string {
	if n == 0 {
		return "0"
	}
	bin := ""
	for n > 0 {
		bin = fmt.Sprintf("%d%s", n%2, bin)
		n /= 2
	}
	return bin
}

func main() {
	var num int
	fmt.Print("Enter number: ")
	fmt.Scan(&num)
	fmt.Println("Binary:", toBinary(num))
}
