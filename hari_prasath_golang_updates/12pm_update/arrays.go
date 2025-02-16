package main
import "fmt"

func main(){
	var scores [3]int
	scores[0] = 50
	scores[1] = 75
	scores[2] = 90
	fmt.Println(scores)

	scores2 := [...]int{50,75,90}

	names := [4]string{"hari","divya","kavin","kavi"}
	fmt.Println(names,scores2)

	fmt.Println(names[1],scores2[2])

}