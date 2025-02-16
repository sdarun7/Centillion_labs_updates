package main
import "fmt"

func main() {
	var c float64
	fmt.Print("Enter Celsius: ")
	fmt.Scan(&c)
	f := (c * 9 / 5) + 32
	fmt.Println("Fahrenheit:", f)
}
