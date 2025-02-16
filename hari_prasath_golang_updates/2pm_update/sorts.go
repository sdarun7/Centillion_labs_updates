package main
import (
	"fmt"
	"sort"
)

func main(){
	// ages := []int{3,65,9}
	// sort.Ints(ages)
	// fmt.Println(ages)

	// index := sort.SearchInts(ages,30)
	// fmt.Println(index)

	names := []string{"hari","divya","dharshini","prasath"}
	sort.Strings(names)
	fmt.Println(names)

}