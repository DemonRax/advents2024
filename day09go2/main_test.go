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
			in:   `2333133121414131402`,
			want: 2858,
		},
	} {
		t.Run(test.in, func(t *testing.T) {
			if got := checksum(util.ReadString(test.in)); got != test.want {
				t.Errorf("got = %d, want %d", got, test.want)
			}
		})
	}
}

func Test_convert(t *testing.T) {
	for _, test := range []struct {
		in   string
		want []file
	}{
		{
			in:   "12345",
			want: []file{{0, 1, false}, {0, 2, true}, {1, 3, false}, {0, 4, true}, {2, 5, false}},
		},
		{
			in:   "90909",
			want: []file{{0, 9, false}, {1, 9, false}, {2, 9, false}},
		},
		{
			in: "2333133121414131402",
			want: []file{
				{0, 2, false},
				{0, 3, true},
				{1, 3, false},
				{0, 3, true},
				{2, 1, false},
				{0, 3, true},
				{3, 3, false},
				{0, 1, true},
				{4, 2, false},
				{0, 1, true},
				{5, 4, false},
				{0, 1, true},
				{6, 4, false},
				{0, 1, true},
				{7, 3, false},
				{0, 1, true},
				{8, 4, false},
				{9, 2, false},
			},
		},
	} {
		t.Run(test.in, func(t *testing.T) {
			got := convert(test.in)
			if diff := cmp.Diff(test.want, got, cmp.AllowUnexported(file{})); diff != "" {
				t.Errorf("did not get expected result (-want/+got):\n%s", diff)
			}
		})
	}
}

func Test_compress(t *testing.T) {
	for _, test := range []struct {
		in, want []file
	}{
		{
			in:   []file{{0, 1, false}, {0, 2, true}, {1, 3, false}, {0, 4, true}, {2, 5, false}},
			want: []file{{0, 1, false}, {0, 2, true}, {1, 3, false}, {0, 4, true}, {2, 5, false}},
		},
		{
			in: []file{
				{0, 2, false},
				{0, 3, true},
				{1, 3, false},
				{0, 3, true},
				{2, 1, false},
				{0, 3, true},
				{3, 3, false},
				{0, 1, true},
				{4, 2, false},
				{0, 1, true},
				{5, 4, false},
				{0, 1, true},
				{6, 4, false},
				{0, 1, true},
				{7, 3, false},
				{0, 1, true},
				{8, 4, false},
				{9, 2, false},
			},
			want: []file{ // 00992111777.44.333....5555.6666.....8888..
				{0, 2, false},
				{9, 2, false},
				{2, 1, false},
				{1, 3, false},
				{7, 3, false},
				{0, 1, true},
				{4, 2, false},
				{0, 1, true},
				{3, 3, false},
				{0, 1, true},
				{0, 2, true},
				{0, 1, true},
				{5, 4, false},
				{0, 1, true},
				{6, 4, false},
				{0, 1, true},
				{0, 3, true},
				{0, 1, true},
				{8, 4, false},
				{0, 2, true},
			},
		},
	} {
		t.Run("", func(t *testing.T) {
			got := compress(test.in)
			if diff := cmp.Diff(test.want, got, cmp.AllowUnexported(file{})); diff != "" {
				t.Errorf("did not get expected result (-want/+got):\n%s", diff)
			}
		})
	}
}
