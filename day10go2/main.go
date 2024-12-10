package main

import (
	"fmt"

	"advents2024/util"
)

func main() {
	fmt.Println(trails(util.ReadFile("./day10go2/input.txt")))
}

type point struct {
	x, y int
}

var ds = []point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
var ns = []uint8{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

func trails(ss []string) int {
	var res int
	for i, s := range ss {
		for j, c := range s {
			if c == '0' {
				res += score(ss, point{i, j}, 0)
			}
		}
	}
	return res
}

func score(ss []string, p point, h int) int {
	if p.x < 0 || p.x >= len(ss) || p.y < 0 || p.y >= len(ss[0]) {
		return 0
	}

	if ss[p.x][p.y] != ns[h] {
		return 0
	}
	if h == 9 {
		return 1
	}

	var sum int
	for _, d := range ds {
		sum += score(ss, point{p.x + d.x, p.y + d.y}, h+1)
	}
	return sum
}
