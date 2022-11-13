package main

import "fmt"

func factorial(n int32) {

    result := int32(1)
    for i := int32(1); i <= n ; i++ {
        result *= i
    }
    fmt.Println(result)
}

func main() {
    var n int32
    fmt.Scanf("%d", &n)
    factorial(n)
}
