package core

import (
	"log"
	"strconv"
)

func Convert(argv []string) []int {
	var data []int
	for _, arg := range argv {
		v, err := strconv.Atoi(arg)
		if err != nil {
			log.Fatalf("invalid argument: %s. expects integers only.\n", arg)
		}
		data = append(data, v)
	}
	return data
}
