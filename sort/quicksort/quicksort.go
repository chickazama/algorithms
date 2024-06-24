package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

const (
	minArgc = 2
)

func main() {
	if len(os.Args) < minArgc {
		log.Fatalf("invalid argument count. expected: %d. actual: %d.\n", minArgc, len(os.Args))
	}
	var data []int
	argv := os.Args[1:]
	for _, arg := range argv {
		v, err := strconv.Atoi(arg)
		if err != nil {
			log.Fatalf("invalid argument: %s. expects integers.\n", arg)
		}
		data = append(data, v)
	}
	fmt.Println(data)
}
