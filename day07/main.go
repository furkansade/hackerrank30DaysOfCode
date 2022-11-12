package main

import "fmt"

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

	for i := (n - 1); i >= 0; i-- {
		fmt.Printf("%v ", arr[i])
	}
}
