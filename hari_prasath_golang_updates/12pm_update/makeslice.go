//slice with make syntax

package main
import "fmt"

func main(){
	nums := make([]int, 3, 5)
	nums = append(nums,20)
	nums = append(nums,21)
	nums = append(nums,22)
	nums = append(nums,23)
	nums = append(nums,22)
	nums = append(nums,212)
	nums = append(nums,24)

	fmt.Println(nums)

}
