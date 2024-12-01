package main

import (
	"fmt"
	"strings"

	"advents2024/util"
)

func main() {
	fmt.Println(similarity(util.ReadFile("./day01go1/input.txt")))
}

func similarity(ss []string) int {
	var (
		left  = make([]int, len(ss))
		right = make(map[int]int, len(ss))
		sum   int
	)

	for i, s := range ss {
		t := strings.Fields(s)
		left[i] = util.ToInt(t[0])
		right[util.ToInt(t[1])]++
	}

	for _, k := range left {
		sum += k * right[k]
	}

	return sum
}
