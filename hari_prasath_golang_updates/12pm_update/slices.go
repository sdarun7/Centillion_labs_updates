package main 
import "fmt"

func main(){
	// scores := []int{50,60,70}
	// fmt.Println(scores)
	// scores = append(scores,30)
	// fmt.Println(scores)
	// scores[3] = 23
	// fmt.Println(scores)
	// fmt.Println(len(scores))

	var fun []int
	fun = append(fun,23)
	fun = append(fun,24)
	fun = append(fun,25)
	fun = append(fun,26)
	fun = append(fun,278)
	fmt.Println(fun)

	a := fun[0:2]
	b := fun[0:3]

	fmt.Println(a,b)

}