package util

import (
	"strconv"
	"strings"
)

func ToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func IsInt(c rune) bool {
	return c >= 48 && c <= 57
}

func Pow(n, m int) int {
	if m == 0 {
		return 1
	}
	result := n
	for i := 2; i <= m; i++ {
		result *= n
	}
	return result
}

func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func IntArray(s string) []int {
	ss := strings.Fields(s)
	if len(ss) > 0 {
		return intArray(ss)
	}
	ss = strings.Split(s, ", ")
	if len(ss) > 0 {
		return intArray(ss)
	}
	ss = strings.Split(s, "\n")
	if len(ss) > 0 {
		return intArray(ss)
	}
	return nil
}

func intArray(ss []string) []int {
	res := make([]int, len(ss))
	for i, s := range ss {
		res[i] = ToInt(s)
	}
	return res
}
