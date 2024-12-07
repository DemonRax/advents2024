package main

import (
	"fmt"
	"strings"

	"advents2024/util"
)

func main() {
	fmt.Println(calibration(util.ReadFile("./day07go1/input.txt")))
}

func calibration(ss []string) uint64 {
	var res uint64
	for _, s := range ss {
		res += calibrate(s)
	}
	return res
}

func calibrate(s string) uint64 {
	var (
		ss   = strings.Split(s, ": ")
		res  = util.ToInt(ss[0])
		ints = util.IntArray(ss[1])
	)

	return match(ints, 0, uint64(res))
}

func match(ints []int, sum, res uint64) uint64 {
	if sum > res {
		return 0
	}
	if len(ints) == 0 {
		if sum == res {
			return res
		}
		return 0
	}

	if match(ints[1:], sum+uint64(ints[0]), res) == res {
		return res
	}
	if sum == 0 {
		sum = 1
	}
	return match(ints[1:], sum*uint64(ints[0]), res)
}
