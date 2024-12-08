package main

import (
	"fmt"

	"advents2024/util"
)

func main() {
	fmt.Println(antiNodes(util.ReadFile("./day08go1/input.txt")))
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

	for r, pts := range letters {
		for x := 0; x < len(pts)-1; x++ {
			for y := x + 1; y < len(pts); y++ {
				c, d := findAnti(pts[x], pts[y])
				if check(r, c, ss) {
					antis[c] = struct{}{}
				}
				if check(r, d, ss) {
					antis[d] = struct{}{}
				}
			}
		}
	}
	return len(antis)
}

func findAnti(a, b point) (point, point) {
	dx, dy := b.x-a.x, b.y-a.y
	return point{a.x - dx, a.y - dy}, point{b.x + dx, b.y + dy}
}

func check(c rune, p point, ss []string) bool {
	return p.x >= 0 && p.x < len(ss) && p.y >= 0 && p.y < len(ss[0]) && ss[p.x][p.y] != uint8(c)
}
