package main

import (
	"bytes"
	"fmt"
	"strings"

	"advents2024/util"
)

func main() {
	fmt.Println(xmas(util.ReadFile("./day04go1/input.txt")))
}

func xmas(ss []string) int {
	cc := make([][]byte, len(ss))
	for i, s := range ss {
		cc[i] = []byte(strings.TrimSpace(s))
	}
	var res int
	res += scan(cc)
	res += scan(rotate(cc))
	res += scan(diag(cc))
	return res
}

func scan(ss [][]byte) int {
	var res int
	for _, s := range ss {
		res += bytes.Count(s, []byte("XMAS"))
		res += bytes.Count(s, []byte("SAMX"))
	}
	return res
}

func rotate(cc [][]byte) [][]byte {
	res := make([][]byte, len(cc))
	for x := range cc {
		res[x] = make([]byte, len(cc))
	}
	for x := range cc {
		for y := range cc {
			res[y][x] = cc[x][y]
		}
	}
	return res
}

func diag(cc [][]byte) [][]byte {
	var res [][]byte
	for y := range cc[0][:len(cc)-1] {
		var row1, row2 []byte
		for i, j, k := 0, y, y; i < len(cc); i, j, k = i+1, j+1, k-1 {
			if j < len(cc[0]) {
				row1 = append(row1, cc[i][j])
			}
			if k >= 0 {
				row2 = append(row2, cc[i][k])
			}
		}
		res = append(res, row1)
		res = append(res, row2)
	}
	for y := range cc[0][:len(cc)-1] {
		var row1, row2 []byte
		for i, j, k := len(cc)-1, y, y; i >= 0; i, j, k = i-1, j+1, k-1 {
			if j < len(cc[0]) {
				row1 = append(row1, cc[i][j])
			}
			if k >= 0 {
				row2 = append(row2, cc[i][k])
			}
		}
		res = append(res, row1)
		res = append(res, row2)
	}
	return res
}
