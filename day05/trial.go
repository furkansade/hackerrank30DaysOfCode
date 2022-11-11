package main

import "fmt"

func Solve(n int) {
	var result int = 1

	for i := 1; i <= 10; i++ {
		result = n * i
		fmt.Printf("%d x %d = %d\n", n, i, result)
	}
}

func main() {
	var n int
	fmt.Scanln(&n)
	Solve(n)
}
