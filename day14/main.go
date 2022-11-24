package main

import (
	"fmt"
	"sort"
)

func Solve(arr []int) {
	sort.Ints(arr)
	// fmt.Println(arr)
	n := len(arr)
	fmt.Println(arr[n-1] - arr[0])
}

func main() {
	var n int
	fmt.Scanln(&n)
	var arr = make([]int, n)

	for i := range arr {
		_, err := fmt.Scan(&arr[i])
		if err != nil {
			fmt.Println(err)
		}
	}

	Solve(arr)

}
