package main

import (
	"fmt"

	"advents2024/util"
)

func main() {
	fmt.Println(checksum(util.ReadFile("./day09go1/input.txt")))
}

type file struct {
	value,
	len int
	isEmpty bool
}

func checksum(ss []string) int {
	return calculate(compress(convert(ss[0])))
}

func convert(s string) []file {
	var (
		res     []file
		isEmpty = false
		index   = 0
		x       int
	)
	for _, c := range s {
		i := util.ToInt(string(c))
		x = index
		if i != 0 {
			res = append(res, file{
				value: func() int {
					if isEmpty {
						return 0
					}
					return x
				}(),
				len:     i,
				isEmpty: isEmpty,
			})
		}
		if isEmpty {
			index++
		}
		isEmpty = !isEmpty
	}
	return res
}

func compress(s []file) []file {
	var (
		ie = 0
	)
	for i := len(s) - 1; i > ie; i-- {
		if s[i].isEmpty {
			continue
		}
		for j := ie; j < i; j++ {
			if !s[j].isEmpty {
				continue
			}
			if s[j].len < s[i].len {
				continue
			}
			if s[j].len > s[i].len {
				s = append(s[:j+1], append([]file{{0, s[j].len - s[i].len, true}}, s[j+1:]...)...)
				i++
				s[j].len = s[i].len
			}
			s[j], s[i] = s[i], s[j]
			break
		}
	}
	return s
}

func calculate(f []file) int {
	var (
		res int
		s   []int
	)
	for _, v := range f {
		x := v.value
		if v.isEmpty {
			x = -1
		}
		for j := 0; j < v.len; j++ {
			s = append(s, x)
		}
	}
	for i, v := range s {
		if v < 0 {
			continue
		}
		res += i * v
	}
	return res
}
