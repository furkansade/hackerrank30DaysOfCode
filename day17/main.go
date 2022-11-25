package main

import (
	"fmt"
	"math"
)

func Solve(a, b int) {

	if a >= 0 && b >= 0 {
		fmt.Println(int(math.Pow(float64(a), float64(b))))
	} else {
		fmt.Println("n and p should be non-negative.")
	}
}

func main() {
	var a, b int
	fmt.Scan(&a)
	fmt.Scanln(&b)
	Solve(a, b)

}
