package main

import (
	"advents2024/util"
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println(incorrectMiddle(util.ReadFile("./day05go2/input.txt")))
}

type rule struct {
	a, b int
}

type sortable struct {
	ints  []int
	rules []rule
}

func (s sortable) Len() int {
	return len(s.ints)
}

func (s sortable) Less(i, j int) bool {
	for _, r := range s.rules {
		if s.ints[i] == r.a && s.ints[j] == r.b {
			return true
		}
		if s.ints[i] == r.b && s.ints[j] == r.a {
			return false
		}
	}
	return false
}

func (s sortable) Swap(i, j int) {
	s.ints[i], s.ints[j] = s.ints[j], s.ints[i]
}

func incorrectMiddle(ss []string) int {
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
				s := sortable{
					ints:  ints,
					rules: rules,
				}
				sort.Sort(s)
				res += s.ints[len(ints)/2]
				continue LINE
			}
		}
	}
	return res
}
