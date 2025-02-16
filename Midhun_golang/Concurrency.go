package main
import (
	"fmt"
	"time"
)
func task(id int, ch chan string) {
	time.Sleep(2 * time.Second)
	ch <- fmt.Sprintf("Task %d complete", id)
}
func main() {
	ch := make(chan string)
	for i := 1; i <= 5; i++ {
		go task(i, ch)
	}
	for i := 1; i <= 5; i++ {
		fmt.Println(<-ch)
	}
}
