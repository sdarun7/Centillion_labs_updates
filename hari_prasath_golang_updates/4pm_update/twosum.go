package main
import "fmt"

func main(){
	nums := []int{2,7,11,15}
	target := 9
	hashmap := map[int]int{	}

	for i, num := range nums{
		comp := target-num

		if j, ok := hashmap[comp]; ok{
			fmt.Println([]int{j,i})
			return 
		}
		hashmap[num] = 1
	}
	// return nil



	

	// fmt.Print(hashmap)
}