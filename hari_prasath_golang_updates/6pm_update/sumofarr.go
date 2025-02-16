package main
import "fmt"

func main(){
	nums := []int{2,1,3,5,6}
	tot := 0

	for _,val := range nums{
		tot += val
	}

	fmt.Println(tot)
}