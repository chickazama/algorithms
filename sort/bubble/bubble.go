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
	bubbleSort(data)
	fmt.Println("Output:")
	fmt.Println(data)
}

func bubbleSort(data []int) {
	if len(data) < 2 {
		return
	}
	pass := 0
	for {
		pass++
		fmt.Printf("Pass: %d\n", pass)
		swapped := false
		for i := 0; i < len(data)-1; i++ {
			if data[i] > data[i+1] {
				data[i], data[i+1] = data[i+1], data[i]
				swapped = true
				fmt.Println(data)
			}
		}
		if !swapped {
			fmt.Println("No swaps")
			return
		}
	}
}
