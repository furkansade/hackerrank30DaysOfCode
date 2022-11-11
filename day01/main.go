package main

import "fmt"

func main() {
    var i int
    var d float64
    var s string = "Hackerrank "
    var s1 string

    fmt.Scanln(&i)
    fmt.Scanln(&d)
    fmt.Scanln(&s1)

    fmt.Println(i + int(d))
    fmt.Println(2*d)
    fmt.Print(s, s1)
}