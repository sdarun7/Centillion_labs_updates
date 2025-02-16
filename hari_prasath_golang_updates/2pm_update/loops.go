package main 

import "fmt"

func main(){

	//for style
	for i:=0; i<10;i++ {
		fmt.Print("Hello")
	}

	//while style
	n := 12
	for n < 21{
		fmt.Print("mama")
		n++
	}

	//Range based
	nums :=[]int{1,2,4,5,11,22}

	for i,j := range nums{
		fmt.Println(i,j)
		if i == 3{
			break
		}
	}




}