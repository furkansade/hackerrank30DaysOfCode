package main

import (
	"fmt"
	"math"
)

func CheckPrime(n float64) string {
	var i int
	if n == 1 || (n != 2 && int(n)%2 == 0) {
		return "Not prime"
	}

	for i = 3; i <= int(math.Sqrt(n)); i += 2 {
		if int(n)%i == 0 {
			return "Not prime"
		}
	}
	return "Prime"
}

func main() {
	var t int
	fmt.Scan(&t)

	arr := make([]float64, t)

	for i := range arr {
		fmt.Scan(&arr[i])
	}

	for _, v := range arr {
		fmt.Println(CheckPrime(v))
	}
}
