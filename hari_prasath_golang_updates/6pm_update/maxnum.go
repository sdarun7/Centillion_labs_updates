package main
import "fmt"

func main() {
    nums := []int{23, 45, 12, 67, 89, 34}
	max := 0

	for _,val := range nums{
		if val > max{
			max = val
		}

	}

	fmt.Println(max)

}