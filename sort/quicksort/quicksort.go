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
	quicksort(data, 0, len(data)-1)
	fmt.Println("Output:")
	fmt.Println(data)
}

func quicksort(data []int, min, max int) {
	if min >= max || min < 0 {
		return
	}
	pivotIdx := partition(data, min, max)
	quicksort(data, min, pivotIdx-1)
	quicksort(data, pivotIdx+1, max)
}

func partition(data []int, min, max int) int {
	pivot := data[max]
	idx := min
	for i := idx; i < max; i++ {
		if data[i] < pivot {
			data[idx], data[i] = data[i], data[idx]
			idx++
		}
	}
	data[idx], data[max] = data[max], data[idx]
	return idx
}
