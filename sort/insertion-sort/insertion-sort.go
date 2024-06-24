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
	insertionSort(data)
	fmt.Println("Output:")
	fmt.Println(data)
}

func insertionSort(data []int) {
	for i := 0; i < len(data); i++ {
		j := i
		for j > 0 && data[j] < data[j-1] {
			data[j], data[j-1] = data[j-1], data[j]
			j--
		}
	}
}
