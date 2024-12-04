package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"advents2024/util"
)

func Test_code(t *testing.T) {
	for _, test := range []struct {
		in   string
		want int
	}{
		{
			in: `..X...
.SAMX.
.A..A.
XMAS.S
.X....`,
			want: 4,
		},
		{
			in: `....XXMAS.
.SAMXMS...
...S..A...
..A.A.MS.X
XMASAMX.MM
X.....XA.A
S.S.S.S.SS
.A.A.A.A.A
..M.M.M.MM
.X.X.XMASX`,
			want: 18,
		},
		{
			in: `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`,
			want: 18,
		},
	} {
		t.Run(test.in, func(t *testing.T) {
			if got := xmas(util.ReadString(test.in)); got != test.want {
				t.Errorf("got = %d, want %d", got, test.want)
			}
		})
	}
}

func Test_diag1(t *testing.T) {
	for _, test := range []struct {
		desc string
		in,
		want [][]byte
	}{
		{
			in: [][]byte{
				{'0', '1', '2'},
				{'3', '4', '5'},
				{'6', '7', '8'},
				{'9', 'A', 'B'},
			},
			want: [][]byte{
				{'0', '4', '8'},
				{'0'},
				{'1', '5'},
				{'1', '3'},
				{'2'},
				{'2', '4', '6'},
				{'9', '7', '5'},
				{'9'},
				{'A', '8'},
				{'A', '6'},
				{'B'},
				{'B', '7', '3'},
			},
		},
	} {
		t.Run(test.desc, func(t *testing.T) {
			if diff := cmp.Diff(test.want, diag(test.in)); diff != "" {
				t.Errorf("did not get expected result (-want/+got):\n%s", diff)
			}
		})
	}
}
