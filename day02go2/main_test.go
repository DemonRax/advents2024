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
			in: `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`,
			want: 4,
		},
	} {
		t.Run(test.in, func(t *testing.T) {
			if got := safe(util.ReadString(test.in)); got != test.want {
				t.Errorf("got = %d, want %d", got, test.want)
			}
		})
	}
}

func Test_isSafe(t *testing.T) {
	for _, test := range []struct {
		desc   string
		levels []int
		want   bool
	}{
		{
			desc:   "happy asc",
			levels: []int{1, 3, 6, 7},
			want:   true,
		},
		{
			desc:   "happy desc",
			levels: []int{7, 6, 4, 2, 1},
			want:   true,
		},
		{
			desc:   "happy asc skip",
			levels: []int{1, 2, 18, 4, 7},
			want:   true,
		},
		{
			desc:   "happy desc skip",
			levels: []int{7, 6, 1, 5, 2},
			want:   true,
		},
		{
			desc:   "happy asc skip first correct",
			levels: []int{3, 1, 2, 4, 7},
			want:   true,
		},
		{
			desc:   "happy asc skip last correct",
			levels: []int{1, 2, 4, 7, 6},
			want:   true,
		},
		{
			desc:   "happy desc skip first",
			levels: []int{1, 7, 6, 5, 2},
			want:   true,
		},
		{
			desc:   "happy desc skip last",
			levels: []int{7, 6, 5, 12},
			want:   true,
		},
		{
			desc:   "happy desc skip second",
			levels: []int{7, 1, 6, 5, 2},
			want:   true,
		},
		{
			desc:   "happy desc skip penultimate",
			levels: []int{7, 6, 5, 12, 2},
			want:   true,
		},
		{
			desc:   "fail two",
			levels: []int{17, 6, 5, 12, 2},
			want:   false,
		},
	} {
		t.Run(test.desc, func(t *testing.T) {
			if got := isSafe(test.levels, false); got != test.want {
				t.Errorf("got = %t, want %t", got, test.want)
			}
		})
	}
}
