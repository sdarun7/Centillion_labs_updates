package main
import "fmt"

func main(){
	num := 281
	rev := 0
	for num > 0{
		temp := num % 10
		rev = rev*10 + temp
		num = num/10
	}
	// rev += num%10
	fmt.Println(rev)
}