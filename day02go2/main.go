package main

import (
	"fmt"

	"advents2024/util"
)

func main() {
	fmt.Println(safe(util.ReadFile("./day02go2/input.txt")))
}

func safe(ss []string) int {
	var sum int

	for _, report := range ss {
		if isSafe(util.IntArray(report), false) {
			sum++
		}
	}
	return sum
}

func isSafe(levels []int, failed bool) bool {
	var asc *bool
	for i := 1; i < len(levels); i++ {
		d := levels[i] - levels[i-1]
		if asc == nil {
			b := d > 0
			asc = &b
		}
		if d == 0 || util.Abs(d) > 3 || *asc != (d > 0) {
			if failed {
				return false
			}
			add := false
			if i == 2 {
				add = isSafe(remove(levels, 0), true)
			}
			return add || isSafe(remove(levels, i-1), true) || isSafe(remove(levels, i), true)
		}
	}
	return true
}

func remove(src []int, k int) []int {
	res := make([]int, 0, len(src)-1)
	for i := range src {
		if i != k {
			res = append(res, src[i])
		}
	}
	return res
}
