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
			in: `...0...
...1...
...2...
6543456
7.....7
8.....8
9.....9`,
			want: 2,
		},
		{
			in: `..90..9
...1.98
...2..7
6543456
765.987
876....
987....`,
			want: 4,
		},
		{
			in: `10..9..
2...8..
3...7..
4567654
...8..3
...9..2
.....01`,
			want: 3,
		},
		{
			in: `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`,
			want: 36,
		},
	} {
		t.Run(test.in, func(t *testing.T) {
			if got := trails(util.ReadString(test.in)); got != test.want {
				t.Errorf("got = %d, want %d", got, test.want)
			}
		})
	}
}
