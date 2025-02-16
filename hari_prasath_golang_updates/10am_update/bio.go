package main
import "fmt"

func main(){
	var name string
	var age int
	fmt.Print("Enter your name here: ")
	fmt.Scan(&name)
	fmt.Print("Enter your age here: ")
	fmt.Scan(&age)

	fmt.Printf("My name is %s and my age is %d",name, age)

	
}