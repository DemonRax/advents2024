package main

import (
	"testing"

	"advents2024/util"
)

func Test_code(t *testing.T) {
	for _, test := range []struct {
		in   string
		want uint64
	}{
		{
			in:   `190: 10 19`,
			want: 190,
		},
		{
			in:   `192: 17 8 14`,
			want: 192,
		},
		{
			in:   `156: 15 6`,
			want: 156,
		},
		{
			in:   `7290: 6 8 6 15`,
			want: 7290,
		},
		{
			in:   `292: 11 6 16 20`,
			want: 292,
		},
		{
			in:   `21037: 9 7 18 13`,
			want: 0,
		},
		{
			in: `190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`,
			want: 11387,
		},
	} {
		t.Run(test.in, func(t *testing.T) {
			if got := calibration(util.ReadString(test.in)); got != test.want {
				t.Errorf("got = %d, want %d", got, test.want)
			}
		})
	}
}
