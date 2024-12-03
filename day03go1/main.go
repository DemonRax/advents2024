package main

import (
	"fmt"
	"regexp"
	"strings"

	"advents2024/util"
)

func main() {
	fmt.Println(mul(util.ReadFile("./day03go1/input.txt")))
}

func mul(ss []string) int {
	var res int
	r := regexp.MustCompile("mul\\(\\d{1,3},\\d{1,3}\\)")
	cmds := r.FindAll([]byte(strings.Join(ss, "")), -1)
	for _, cmd := range cmds {
		res += parse(cmd)
	}
	return res
}

func parse(c []byte) int {
	s := strings.TrimLeft(strings.TrimRight(string(c), ")"), "mul(")
	ss := strings.Split(s, ",")
	return util.ToInt(ss[0]) * util.ToInt(ss[1])
}
