package main

import (
	"testing"
)

var testCases = []struct {
	in        [][]int
	out       [][]int
	direction int
}{
	{
		[][]int{
			{2, 0, 0, 2},
			{4, 16, 8, 2},
			{2, 64, 32, 4},
			{1024, 1024, 64, 0},
		},
		[][]int{
			{4, 0, 0, 0},
			{4, 16, 8, 2},
			{2, 64, 32, 4},
			{2048, 64, 0, 0},
		},
		0,
	},
	{
		[][]int{
			{2, 0, 0, 2},
			{4, 16, 8, 2},
			{2, 64, 32, 4},
			{1024, 1024, 64, 0},
		},
		[][]int{
			{2, 16, 8, 4},
			{4, 64, 32, 4},
			{2, 1024, 64, 0},
			{1024, 0, 0, 0},
		},
		1,
	},
	{
		[][]int{
			{2, 0, 0, 2},
			{4, 16, 8, 2},
			{2, 64, 32, 4},
			{1024, 1024, 64, 0},
		},
		[][]int{
			{0, 0, 0, 4},
			{4, 16, 8, 2},
			{2, 64, 32, 4},
			{0, 0, 2048, 64},
		},
		2,
	},
	{
		[][]int{
			{2, 0, 0, 2},
			{4, 16, 8, 2},
			{2, 64, 32, 4},
			{1024, 1024, 64, 0},
		},
		[][]int{
			{2, 0, 0, 0},
			{4, 16, 8, 0},
			{2, 64, 32, 4},
			{1024, 1024, 64, 4},
		},
		3,
	},
	{
		[][]int{
			{2, 2, 4, 8},
			{4, 0, 4, 4},
			{16, 16, 16, 16},
			{32, 16, 16, 32},
		},
		[][]int{
			{4, 4, 8, 0},
			{8, 4, 0, 0},
			{32, 32, 0, 0},
			{32, 32, 32, 0},
		},
		0,
	},
	{
		[][]int{
			{2, 2, 4, 8},
			{4, 0, 4, 4},
			{16, 16, 16, 16},
			{32, 16, 16, 32},
		},
		[][]int{
			{0, 4, 4, 8},
			{0, 0, 4, 8},
			{0, 0, 32, 32},
			{0, 32, 32, 32},
		},
		2,
	},
}

func Equal(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, vA := range a {
		for j, vE := range vA {
			if vE != b[i][j] {
				return false
			}
		}
	}
	return true
}

func TestStarter(t *testing.T) {
	for _, tt := range testCases {
		t.Run("test", func(t *testing.T) {
			res := start(tt.in, tt.direction)
			if !Equal(res, tt.out) {
				t.Errorf("got %v, want %v", res, tt.out)
			}
		})
	}
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		start(testCases[0].in, testCases[0].direction)
	}
}
