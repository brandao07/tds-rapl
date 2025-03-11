package main

import (
	"fmt"
	"os"
	"strconv"
)

func fib(n int) []int {
	if n <= 0 {
		return []int{}
	}
	if n == 1 {
		return []int{0}
	}

	fib := make([]int, n)
	fib[0], fib[1] = 0, 1

	for i := 2; i < n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}

	return fib
}

func main() {
	n, err := strconv.Atoi(os.Args[1])

	if err != nil {
		panic(err)
	}
	result := fib(n)
	fmt.Println(result[len(result)-1])
}
