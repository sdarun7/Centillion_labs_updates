package main

import (
	"fmt"
	"math/rand"
)

func main() {
	secret_num := rand.Intn(100)
	// fmt.Println(secret_num)

	var usr int
	for i := 0; i < 10; i++ {
		fmt.Printf("Enter your guess, you have only %d attempts! ", i)
		fmt.Scan(&usr)

		if usr == secret_num {
			fmt.Print("You are a legend! the answer is right!")
			break
		}

		if usr > secret_num {
			fmt.Print("Greater")
		} else {
			fmt.Print("Lesser")
		}
	}
}
