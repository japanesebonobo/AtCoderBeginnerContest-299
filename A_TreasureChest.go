package main

import (
	"fmt"
	"strings"
)

func main() {
	var N int
	var S string
	fmt.Scan(&N, &S)

	s := strings.Split(S, "")

	var count int = 0
	for i := 0; i < N; i++ {
		if s[i] == "|" {
			count++
		}

		if s[i] == "*" && count == 1 {
			fmt.Print("in")
			break
		}

		if count == 2 {
			fmt.Print("out")
			break
		}
	}
}
