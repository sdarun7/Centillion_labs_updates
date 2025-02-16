package main
import "fmt"

func main() {
    names := []string{"Hari", "Div", "Go", "Mo"}
	big := ""
	large := 0
	for _,val := range names{
		curlen := len(val)
		if curlen > large{
			large = curlen
			big = val
		}
	}
	fmt.Println(big,large)
}