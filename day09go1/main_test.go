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
			want: 1928,
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
		want []int
	}{
		{
			in:   "12345",
			want: []int{0, -1, -1, 1, 1, 1, -1, -1, -1, -1, 2, 2, 2, 2, 2},
		},
		{
			in:   "90909",
			want: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2},
		},
		{
			in:   "2333133121414131402",
			want: []int{0, 0, -1, -1, -1, 1, 1, 1, -1, -1, -1, 2, -1, -1, -1, 3, 3, 3, -1, 4, 4, -1, 5, 5, 5, 5, -1, 6, 6, 6, 6, -1, 7, 7, 7, -1, 8, 8, 8, 8, 9, 9},
		},
	} {
		t.Run(test.in, func(t *testing.T) {
			got := convert(test.in)
			if diff := cmp.Diff(test.want, got); diff != "" {
				t.Errorf("did not get expected result (-want/+got):\n%s", diff)
			}
		})
	}
}

func Test_compress(t *testing.T) {
	for _, test := range []struct {
		in, want []int
	}{
		{
			in:   []int{0, -1, -1, 1, 1, 1, -1, -1, -1, -1, 2, 2, 2, 2, 2},
			want: []int{0, 2, 2, 1, 1, 1, 2, 2, 2, -1, -1, -1, -1, -1, -1},
		},
		{
			in:   []int{0, 0, -1, -1, -1, 1, 1, 1, -1, -1, -1, 2, -1, -1, -1, 3, 3, 3, -1, 4, 4, -1, 5, 5, 5, 5, -1, 6, 6, 6, 6, -1, 7, 7, 7, -1, 8, 8, 8, 8, 9, 9},
			want: []int{0, 0, 9, 9, 8, 1, 1, 1, 8, 8, 8, 2, 7, 7, 7, 3, 3, 3, 6, 4, 4, 6, 5, 5, 5, 5, 6, 6, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
		},
	} {
		t.Run("", func(t *testing.T) {
			got := compress(test.in)
			if diff := cmp.Diff(test.want, got); diff != "" {
				t.Errorf("did not get expected result (-want/+got):\n%s", diff)
			}
		})
	}

}
