package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

var sc *bufio.Scanner

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	N, T := readInt(), readInt()

	C := make([]int, N)
	for i := range C {
		C[i] = readInt()
	}
	R := make([]int, N)
	for i := range R {
		R[i] = readInt()
	}
	var ans int = checkT(C, R, T)
	if ans != -1 {
		fmt.Print(ans)
	} else {
		ans = checkOne(C[0], R[0], C, R)
		fmt.Print(ans)
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}

func checkT(C []int, R []int, T int) int {
	var ans int = -1
	var value int = -1
	for i := 0; i < len(C); i++ {
		if C[i] == T && R[i] > value {
			value = R[i]
			ans = i + 1
		}
	}
	return ans
}

func checkOne(C1 int, R1 int, C []int, R []int) int {
	var ans int = -1
	var value int = R1
	for i := 0; i < len(C); i++ {
		if C[i] == C1 {
			if R[i] > value {
				value = R[i]
				ans = i + 1
			}
		}
	}
	if ans != -1 {
		return ans
	} else {
		return 1
	}
}
