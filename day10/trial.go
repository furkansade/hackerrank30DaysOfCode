package main

import (
	"fmt"
	"math"
)

func Solve(n int) {
	var kalan, sayac, max int

	for n > 0 {
		kalan = n % 2 // ikilik sisteme ceviri kurali
		n /= 2

		if kalan == 1 {
			sayac++
			max = int(math.Max(float64(max), float64(sayac)))
			// max - sayac ikilisinden hangisi buyukse max degiskenine onu ata.
		} else if kalan == 0 {
			sayac = 0
		}
	}
	fmt.Println(max)
}

func main() {
	var n int
	fmt.Scanln(&n)
	Solve(n)
}
