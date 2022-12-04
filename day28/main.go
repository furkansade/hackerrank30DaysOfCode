package main

import (
	"fmt"
	"regexp"
	"sort"
)

func main() {
	var n int
	var name, mail string
	var nameArr []string
	fmt.Scanln(&n)

	r, _ := regexp.Compile(`(@gmail\.com)$`)

	for i := 0; i < n; i++ {
		fmt.Scanln(&name, &mail)
		if r.MatchString(mail) {
			nameArr = append(nameArr, name)
		}
	}

	sort.Slice(nameArr, func(i, j int) bool {
		return nameArr[i] < nameArr[j]
	})

	for _, v := range nameArr {
		fmt.Println(v)
	}
}
