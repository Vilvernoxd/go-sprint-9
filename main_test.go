package main

import "testing"

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		size     int
		wantSize int
	}{
		{0, 0},
		{1, 1},
		{10, 10},
	}

	for _, tt := range tests {
		result := generateRandomElements(tt.size)
		if len(result) != tt.wantSize {
			t.Errorf("generateRandomElements(%d) = %d; want %d", tt.size, len(result), tt.wantSize)
		}
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		data []int
		want int
	}{
		{[]int{}, 0},
		{[]int{5}, 5},
		{[]int{1, 3, 2}, 3},
		{[]int{7, 7, 7}, 7},
		{[]int{10, 2, 5}, 10},
		{[]int{1, 4, 9}, 9},
		{[]int{2, 8, 3}, 8},
	}

	for _, tt := range tests {
		got := maximum(tt.data)
		if got != tt.want {
			t.Errorf("maximum(%v) = %d; want %d", tt.data, got, tt.want)
		}
	}
}
