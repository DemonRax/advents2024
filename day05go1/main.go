package main

import (
	"advents2024/util"
	"fmt"
	"strings"
)

func main() {
	fmt.Println(correctMiddle(util.ReadFile("./day05go1/input.txt")))
}

type rule struct {
	a, b int
}

func correctMiddle(ss []string) int {
	var (
		res   int
		rules []rule
		cut   int
	)

	for i, r := range ss {
		if strings.TrimSpace(r) == "" {
			cut = i + 1
			break
		}
		s := strings.Split(r, "|")
		rules = append(rules, rule{util.ToInt(s[0]), util.ToInt(s[1])})
	}

LINE:
	for _, line := range ss[cut:] {
		var (
			order = map[int]int{}
			ints  = util.IntArray(line)
		)
		for i, v := range ints {
			order[v] = i + 1
		}
		for _, r := range rules {
			if order[r.a] == 0 || order[r.b] == 0 {
				continue
			}
			if order[r.a] > order[r.b] {
				continue LINE
			}
		}
		res += ints[len(ints)/2]
	}

	return res
}
