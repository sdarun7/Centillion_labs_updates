package main
import "fmt"

func add(a,b int)int{ 
	return a+b
}

func sub(a,b int) int{
	return a-b
}

func div(a,b int)int{
	return a/b
}

func mul(a,b int)int{
	return a*b
}

func main(){

	var a,b,n int

	n = 5

	for n < 6{
	
		fmt.Println("Enter any of these options below\n1.Add two numbers\n2.Subtract two numbers\n3.Multiply two number\n4.Divide two numbers\n5.Exit\n")
		fmt.Scan(&n)
	
		if n == 5{
			break
		}

		fmt.Print("Enter the first number: ") 
		fmt.Scan(&a)
		fmt.Print("Enter the Second number: ") 
		fmt.Scan(&b)

    if n == 1{
		fmt.Println("The addition of the above two numbers is: ",add(a,b))
	} else if n == 2{
		fmt.Println("The subtraction of the above two numbers is: ",sub(a,b))
	} else if n == 3{
		fmt.Println("The multiplication of the above two numbers is: ",mul(a,b))
	} else if n == 4{
		fmt.Println("The Division of the above two numbers is: ",div(a,b))
	}
}
}