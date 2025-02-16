package main
import (
	"fmt"
	"sort"
	"strings"
)
func sortString(s string) string {
	r := []rune(strings.ToLower(s))
	sort.Slice(r, func(i, j int) bool { return r[i] < r[j] })
	return string(r)
}
func isAnagram(s1, s2 string) bool {
	return sortString(s1) == sortString(s2)
}
func main() {
	var str1, str2 string
	fmt.Print("Enter first string: ")
	fmt.Scan(&str1)
	fmt.Print("Enter second string: ")
	fmt.Scan(&str2)
	if isAnagram(str1, str2) {
		fmt.Println("The strings are anagrams")
	} else {
		fmt.Println("The strings are not anagrams")
	}
}
