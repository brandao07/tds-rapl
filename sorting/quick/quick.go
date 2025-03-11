package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func quickSort(array []string) []string {
	if len(array) <= 1 {
		return array
	}

	pivot := array[len(array)-1]
	var left, right []string

	for i := 0; i < len(array)-1; i++ {
		if array[i] <= pivot {
			left = append(left, array[i])
		} else {
			right = append(right, array[i])
		}
	}

	left = quickSort(left)
	right = quickSort(right)

	return append(append(left, pivot), right...)
}

func randomString(n int) string {
	letterRunes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func main() {
	n, err := strconv.Atoi(os.Args[1])

	if err != nil {
		panic(err)
	}

	arr := []string{}
	for i := 0; i < n; i++ {
		arr = append(arr, randomString(8))
	}

	fmt.Println(quickSort(arr))
}
