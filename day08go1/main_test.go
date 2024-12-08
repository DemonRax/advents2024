package main

import (
	"testing"

	"advents2024/util"
)

func Test_code(t *testing.T) {
	for _, test := range []struct {
		in   string
		want int
	}{
		{
			in: `..........
..........
..........
....a.....
..........
.....a....
..........
..........
..........
..........`,
			want: 2,
		},
		{
			in: `..........
..........
..........
....a.....
........a.
.....a....
..........
..........
..........
..........`,
			want: 4,
		},
		{
			in: `..........
..........
..........
....a.....
........a.
.....a....
..........
......A...
..........
..........`,
			want: 4,
		},
		{
			in: `............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............`,
			want: 14,
		},
	} {
		t.Run(test.in, func(t *testing.T) {
			if got := antiNodes(util.ReadString(test.in)); got != test.want {
				t.Errorf("got = %d, want %d", got, test.want)
			}
		})
	}
}

func Test_findAnti(t *testing.T) {
	for _, test := range []struct {
		desc string
		a, b point
		c, d point
	}{
		{
			desc: "empty",
		},
		{
			desc: "LR",
			a:    point{3, 3},
			b:    point{4, 4},
			c:    point{2, 2},
			d:    point{5, 5},
		},
		{
			desc: "UP",
			a:    point{3, 3},
			b:    point{4, 3},
			c:    point{2, 3},
			d:    point{5, 3},
		},
		{
			desc: "RL",
			a:    point{4, 6},
			b:    point{3, 3},
			c:    point{5, 9},
			d:    point{2, 0},
		},
	} {
		t.Run(test.desc, func(t *testing.T) {
			c, d := findAnti(test.a, test.b)
			if c.x != test.c.x || c.y != test.c.y {
				t.Errorf("c wrong: %v, want %v", c, test.c)
			}
			if d.x != test.d.x || d.y != test.d.y {
				t.Errorf("d wrong: %v, want %v", d, test.d)
			}
		})
	}
}
