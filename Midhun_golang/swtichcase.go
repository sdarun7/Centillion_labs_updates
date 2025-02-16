package main

import "fmt"

func main() {
    var day int
    fmt.Print("Enter the day number (1-7): ")
    fmt.Scanln(&day)

    switch day {
    case 1:
        fmt.Println("Sunday")
    case 2:
        fmt.Println("Monday")
    case 3:
        fmt.Println("Tuesday")
    case 4:
        fmt.Println("Wednesday")
    case 5:
        fmt.Println("Thursday")
    case 6:
        fmt.Println("Friday")
    case 7:
        fmt.Println("Saturday")
    default:
        fmt.Println("Invalid input!")
    }
}
