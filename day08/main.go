package main

import "fmt"

func main() {
	var n int
	m := make(map[string]int)

	fmt.Scan(&n)

	var name string
	var number int
	for i := 0; i < n; i++ {
		fmt.Scan(&name)
		fmt.Scan(&number)
		m[name] = number
	}

	var checkName string

	for i := 0; i < n; i++ {

		_, err := fmt.Scanf("%s", &checkName)
		if err != nil {
			break
		}

		if _, ok := m[checkName]; ok {
			fmt.Printf("%s=%d\n", checkName, m[checkName])
		} else {
			fmt.Println("Not found")
		}
	}
}
