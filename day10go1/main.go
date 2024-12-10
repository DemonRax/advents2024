package main

import (
	"fmt"

	"advents2024/util"
)

func main() {
	fmt.Println(trails(util.ReadFile("./day10go1/input.txt")))
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
				nines := map[point]int{}
				score(ss, point{i, j}, 0, nines)
				res += len(nines)
			}
		}
	}
	return res
}

func score(ss []string, p point, h int, res map[point]int) {
	if p.x < 0 || p.x >= len(ss) || p.y < 0 || p.y >= len(ss[0]) {
		return
	}

	if ss[p.x][p.y] != ns[h] {
		return
	}
	if h == 9 {
		res[p]++
		return
	}

	for _, d := range ds {
		score(ss, point{p.x + d.x, p.y + d.y}, h+1, res)
	}
}
