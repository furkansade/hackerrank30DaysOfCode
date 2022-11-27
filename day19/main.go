package main

import "fmt"

func Solve(n int) {
	if n < 1 {
		return
	}

	sum := 0
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			sum += i
		}
	}
	fmt.Println("I implemented: AdvancedArithmetic")
	fmt.Println(sum)
}

func main() {
	var n int
	fmt.Scanln(&n)
	Solve(n)
}
