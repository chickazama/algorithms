package main

import (
	"fmt"
	"log"
	"os"

	"github.com/chickazama/algorithms/sort/core"
)

const (
	minArgc = 2
)

func main() {
	if len(os.Args) < minArgc {
		log.Fatalf("invalid argument count. expected: >= %d. actual: %d.\n", minArgc, len(os.Args))
	}
	argv := os.Args[1:]
	data := core.Convert(argv)
	fmt.Println("Input:")
	fmt.Println(data)
	data = mergeSort(data)
	fmt.Println("Output:")
	fmt.Println(data)
}

func mergeSort(data []int) []int {
	n := len(data)
	if n < 2 {
		return data
	}
	var left []int
	var right []int
	for i, v := range data {
		if i < n/2 {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}
	left = mergeSort(left)
	right = mergeSort(right)
	return merge(left, right)
}

func merge(left, right []int) []int {
	var ret []int
	for len(left) > 0 && len(right) > 0 {
		if left[0] <= right[0] {
			ret = append(ret, left[0])
			if len(left) > 1 {
				left = left[1:]
			} else {
				left = []int{}
			}
		} else {
			ret = append(ret, right[0])
			if len(right) > 1 {
				right = right[1:]
			} else {
				right = []int{}
			}
		}
	}
	if len(left) > 0 {
		ret = append(ret, left...)
	}
	if len(right) > 0 {
		ret = append(ret, right...)
	}
	return ret
}
