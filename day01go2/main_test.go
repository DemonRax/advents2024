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
			in: `3   4
4   3
2   5
1   3
3   9
3   3`,
			want: 31,
		},
	} {
		t.Run(test.in, func(t *testing.T) {
			if got := similarity(util.ReadString(test.in)); got != test.want {
				t.Errorf("got = %d, want %d", got, test.want)
			}
		})
	}
}
