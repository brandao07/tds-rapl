package main

import (
	"fmt"
	"os"
	"strconv"
)

func catalan(n int) int {
	if n == 0 {
		return 1
	}
	result := 0
	for i := 0; i < n; i++ {
		result += catalan(i) * catalan(n-1-i)
	}
	return result
}

func main() {
	n, err := strconv.Atoi(os.Args[1])

	if err != nil {
		panic(err)
	}

	fmt.Println(catalan(n))
}
