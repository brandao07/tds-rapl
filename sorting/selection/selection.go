package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func selection(array []float32, size int) []float32 {
	for i := 0; i < size; i++ {
		min := i
		for j := i + 1; j < size; j++ {
			if array[min] > array[j] {
				min = j
			}
		}
		temp := array[min]
		array[min] = array[i]
		array[i] = temp
	}
	return array
}

func main() {
	n, err := strconv.Atoi(os.Args[1])

	if err != nil {
		panic(err)
	}

	arr := []float32{}

	for i := 0; i < n; i++ {
		arr = append(arr, rand.Float32())
	}

	fmt.Println(selection(arr, len(arr)))
}
