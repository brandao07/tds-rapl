package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func bubble(array []int, size int) []int {
	for i := 0; i < size-1; i++ {
		for j := i + 1; j < size; j++ {
			if array[i] > array[j] {
				array[i], array[j] = array[j], array[i]
			}
		}
	}
	return array
}

func main() {
	n, err := strconv.Atoi(os.Args[1])

	if err != nil {
		panic(err)
	}

	arr := []int{}

	for i := 0; i < n; i++ {
		arr = append(arr, rand.Intn(100000)+1)
	}

	fmt.Println(bubble(arr, len(arr)))
}
