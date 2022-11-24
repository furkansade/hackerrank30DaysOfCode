package main

import (
	"fmt"
	"strconv"
)

func Solve(s string) {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		fmt.Println("Bad String")
		return
	}

	fmt.Println(i)
}

func main() {
	var s string
	fmt.Scanln(&s)
	Solve(s)
}
