package main

import (
	"fmt"

	"advents2024/util"
)

func main() {
	fmt.Println(checksum(util.ReadFile("./day09go1/input.txt")))
}

func checksum(ss []string) int {
	return calculate(compress(convert(ss[0])))
}

func convert(s string) []int {
	var (
		res    []int
		isFile = true
		index  = 0
		x      int
	)
	for _, c := range s {
		i := util.ToInt(string(c))
		x = index
		if !isFile {
			x = -1
		}
		for j := 0; j < i; j++ {
			res = append(res, x)
		}
		if isFile {
			index++
		}
		isFile = !isFile
	}
	return res
}

func compress(s []int) []int {
	var (
		ie = 0
	)
	for i := len(s) - 1; i > ie; i-- {
		if s[i] == -1 {
			continue
		}
		for j := ie; j < i; j++ {
			if s[j] == -1 {
				s[j], s[i] = s[i], s[j]
				ie = j
				break
			}
		}
	}
	return s
}

func calculate(s []int) int {
	var (
		res int
	)
	for i, v := range s {
		if v < 0 {
			break
		}
		res += i * v
	}
	return res
}
