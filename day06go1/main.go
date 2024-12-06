package main

import (
	"fmt"

	"advents2024/util"
)

func main() {
	fmt.Println(positions(util.ReadFile("./day06go1/input.txt")))
}

type point struct {
	x, y int
}

var (
	du = point{-1, 0}
	dr = point{0, 1}
	dd = point{1, 0}
	dl = point{0, -1}
)

func positions(ss []string) int {
	var (
		guard, dir point
		obst       = map[point]struct{}{}
		visited    = map[point]struct{}{}
		mx         = len(ss) - 1
		my         = len(ss[0]) - 1
	)

	for i, s := range ss {
		for j, v := range s {
			if v == '^' {
				guard = point{i, j}
			} else if v == '#' {
				obst[point{i, j}] = struct{}{}
			}
		}
	}

	visited[guard] = struct{}{}
	dir = du

	for {
		x := guard.x + dir.x
		y := guard.y + dir.y
		if x < 0 || x > mx || y < 0 || y > my {
			break
		}
		if _, ok := obst[point{x, y}]; ok {
			dir = rotate(dir)
			continue
		}
		guard.x, guard.y = x, y
		visited[guard] = struct{}{}
	}

	return len(visited)
}

func rotate(d point) point {
	switch d {
	case du:
		return dr
	case dr:
		return dd
	case dd:
		return dl
	case dl:
		return du
	}
	return point{}
}
