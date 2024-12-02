package main

import (
	"fmt"

	"advents2024/util"
)

func main() {
	fmt.Println(safe(util.ReadFile("./day02go1/input.txt")))
}

func safe(ss []string) int {
	var sum int

	for _, report := range ss {
		var (
			levels = util.IntArray(report)
			asc    *bool
			ok     = true
		)

		for i := 1; i < len(levels); i++ {
			d := levels[i] - levels[i-1]
			if asc == nil {
				b := d > 0
				asc = &b
			}
			if d == 0 || util.Abs(d) > 3 || *asc != (d > 0) {
				ok = false
				break
			}
		}
		if ok {
			sum++
		}
	}
	return sum
}
