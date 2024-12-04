package main

import (
	"advents2024/util"
	"fmt"
)

func main() {
	fmt.Println(xmas(util.ReadFile("./day04go2/input.txt")))
}

func xmas(ss []string) int {
	var res int
	for i := 1; i < len(ss)-1; i++ {
		for j := 1; j < len(ss[i])-1; j++ {
			if ss[i][j] != 'A' {
				continue
			}
			if ss[i-1][j-1] == 'M' && ss[i+1][j-1] == 'M' && ss[i-1][j+1] == 'S' && ss[i+1][j+1] == 'S' {
				res++
			}
			if ss[i-1][j-1] == 'S' && ss[i+1][j-1] == 'M' && ss[i-1][j+1] == 'S' && ss[i+1][j+1] == 'M' {
				res++
			}
			if ss[i-1][j-1] == 'S' && ss[i+1][j-1] == 'S' && ss[i-1][j+1] == 'M' && ss[i+1][j+1] == 'M' {
				res++
			}
			if ss[i-1][j-1] == 'M' && ss[i+1][j-1] == 'S' && ss[i-1][j+1] == 'M' && ss[i+1][j+1] == 'S' {
				res++
			}
		}
	}
	return res
}
