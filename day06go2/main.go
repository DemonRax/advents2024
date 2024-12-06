package main

import (
	"fmt"

	"advents2024/util"
)

func main() {
	fmt.Println(newLoops(util.ReadFile("./day06go2/input.txt")))
}

type point struct {
	x, y int
}

type guard struct {
	p, d point
}

var (
	du = point{-1, 0}
	dr = point{0, 1}
	dd = point{1, 0}
	dl = point{0, -1}
)

func newLoops(ss []string) int {
	var res int
	g, obst, visited, mx, my := scan(ss)
	for p := range visited {
		res += checkLoop(g, p, obst, mx, my)
	}
	return res
}

func checkLoop(g guard, p point, obst map[point]struct{}, mx, my int) int {
	defer func() {
		delete(obst, p)
	}()
	obst[p] = struct{}{}
	visited := map[guard]struct{}{}

	for {
		x := g.p.x + g.d.x
		y := g.p.y + g.d.y
		if x < 0 || x > mx || y < 0 || y > my {
			return 0
		}
		if _, ok := obst[point{x, y}]; ok {
			g.d = rotate(g.d)
			continue
		}
		g.p.x, g.p.y = x, y
		if _, ok := visited[g]; ok {
			break
		}
		visited[g] = struct{}{}
	}

	return 1
}

func scan(ss []string) (guard, map[point]struct{}, map[point]struct{}, int, int) {
	var (
		g, dir  point
		obst    = map[point]struct{}{}
		visited = map[point]struct{}{}
		mx      = len(ss) - 1
		my      = len(ss[0]) - 1
	)

	for i, s := range ss {
		for j, v := range s {
			if v == '^' {
				g = point{i, j}
			} else if v == '#' {
				obst[point{i, j}] = struct{}{}
			}
		}
	}

	visited[g] = struct{}{}
	dir = du
	res := guard{g, dir}

	for {
		x := g.x + dir.x
		y := g.y + dir.y
		if x < 0 || x > mx || y < 0 || y > my {
			break
		}
		if _, ok := obst[point{x, y}]; ok {
			dir = rotate(dir)
			continue
		}
		g.x, g.y = x, y
		visited[g] = struct{}{}
	}

	return res, obst, visited, mx, my
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
