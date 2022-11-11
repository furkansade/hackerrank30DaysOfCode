package main

import "fmt"

func Solve(n int) {
    switch{
    case n % 2 == 1:
        fmt.Println("Weird")
    case n >= 2 && n <= 5 && n % 2 == 0:
        fmt.Println("Not Weird")
    case n >= 6 && n <= 20 && n % 2 == 0:
        fmt.Println("Weird")
    case n > 20 && n % 2 == 0:
        fmt.Println("Not Weird")
    default:
        break
    }
}

func main() {
    var n int
    fmt.Scanln(&n)
    Solve(n)
}
