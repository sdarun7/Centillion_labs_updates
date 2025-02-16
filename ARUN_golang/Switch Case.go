package main
import "fmt"
func main() {
	var num1, num2 float64
	var operator string
	fmt.Print("Enter first number: ")
	fmt.Scan(&num1)
	fmt.Print("Enter operator (+, -, *, /): ")
	fmt.Scan(&operator)
	fmt.Print("Enter second number: ")
	fmt.Scan(&num2)
	switch operator {
	case "+":
		fmt.Println("Result:", num1+num2)
	case "-":
		fmt.Println("Result:", num1-num2)
	case "*":
		fmt.Println("Result:", num1*num2)
	case "/":
		if num2 != 0 {
			fmt.Println("Result:", num1/num2)
		} else {
			fmt.Println("Error: Division by zero!")
		}
	default:
		fmt.Println("Invalid operator")
	}
}
