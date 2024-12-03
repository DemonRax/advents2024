package main

import (
	"fmt"
	"regexp"
	"strings"

	"advents2024/util"
)

const (
	dont = "don't()"
	do   = "do()"
)

var r = regexp.MustCompile("mul\\(\\d{1,3},\\d{1,3}\\)")

func main() {
	fmt.Println(mul(util.ReadFile("./day03go2/input.txt")))
}

func mul(ss []string) int {
	var (
		res int
		s   = strings.Join(ss, "")
	)

	for {
		j := strings.Index(s, dont)
		if j < 0 {
			res += parseN(s)
			break
		}
		res += parseN(s[:j+1])
		s = s[j+len(dont):]
		j = strings.Index(s, do)
		if j < 0 {
			break
		}
		s = s[j+len(do):]
	}
	return res
}

func parseN(s string) int {
	var res int
	for _, cmd := range r.FindAll([]byte(s), -1) {
		res += parse(cmd)
	}
	return res
}

func parse(c []byte) int {
	s := strings.TrimLeft(strings.TrimRight(string(c), ")"), "mul(")
	ss := strings.Split(s, ",")
	return util.ToInt(ss[0]) * util.ToInt(ss[1])
}
