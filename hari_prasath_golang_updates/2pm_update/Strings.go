package main
import (
	"fmt"
	"strings"
)

func main(){
	greet := "Hello Guys! I hate girls but not Guys"
	fmt.Println(strings.Contains(greet,"Hello"))
	fmt.Println(strings.ReplaceAll(greet, "Guys","Girls"))
	fmt.Println(strings.Index(greet,"but"))
	fmt.Println(strings.Split(greet,"G"))
} 