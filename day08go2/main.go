package main

import (
	"fmt"

	"advents2024/util"
)

func main() {
	fmt.Println(antiNodes(util.ReadFile("./day08go2/input.txt")))
}

type point struct {
	x, y int
}

func antiNodes(ss []string) int {
	var (
		letters = map[rune][]point{}
		antis   = map[point]struct{}{}
	)

	for x, s := range ss {
		for y, c := range s {
			if c == '.' {
				continue
			}
			letters[c] = append(letters[c], point{x, y})
		}
	}

	for _, pts := range letters {
		for x := 0; x < len(pts)-1; x++ {
			for y := x + 1; y < len(pts); y++ {
				for _, anti := range findAnti(pts[x], pts[y], ss) {
					antis[anti] = struct{}{}
				}
			}
		}
	}
	return len(antis)
}

func findAnti(a, b point, ss []string) []point {
	var res []point
	res = append(res, a)
	res = append(res, b)
	dx, dy := b.x-a.x, b.y-a.y
	for {
		a.x -= dx
		a.y -= dy
		if a.x < 0 || a.x >= len(ss) || a.y < 0 || a.y >= len(ss[0]) {
			break
		}
		res = append(res, a)
	}
	for {
		b.x += dx
		b.y += dy
		if b.x < 0 || b.x >= len(ss) || b.y < 0 || b.y >= len(ss[0]) {
			break
		}
		res = append(res, b)
	}
	return res
}
