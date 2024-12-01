package main

import (
	"fmt"
	"sort"
	"strings"

	"advents2024/util"
)

func main() {
	fmt.Println(distance(util.ReadFile("./day01go1/input.txt")))
}

func distance(ss []string) int {
	var (
		left  = make([]int, len(ss))
		right = make([]int, len(ss))
		sum   int
	)

	for i, s := range ss {
		t := strings.Fields(s)
		left[i] = util.ToInt(t[0])
		right[i] = util.ToInt(t[1])
	}
	sort.Ints(left)
	sort.Ints(right)

	for i := range left {
		sum += util.Abs(left[i] - right[i])
	}
	return sum
}
