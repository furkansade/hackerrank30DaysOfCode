package main

import (
	"fmt"
	"sort"
)

type Difference struct {
	elements []int
}

func (d *Difference) computeDifference() {
	sort.Ints(d.elements)
	n := len(d.elements)
	fmt.Println(d.elements[n-1] - d.elements[0])
}

func main() {
	var n int
	fmt.Scanln(&n)
	var arr = make([]int, n)

	for i := range arr {
		_, err := fmt.Scan(&arr[i])
		if err != nil {
			break
		}
	}
	difference := &Difference{elements: arr}
	difference.computeDifference()
}