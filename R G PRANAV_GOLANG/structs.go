package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

func main() {
    p1 := Person{Name: "John Doe", Age: 25}
    fmt.Println("Person Details:")
    fmt.Println("Name:", p1.Name)
    fmt.Println("Age:", p1.Age)
}
