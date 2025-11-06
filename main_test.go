package main

import "testing"

func TestGenerateRandomElements(t *testing.T) {
	result := generateRandomElements(0)
	if len(result) != 0 {
		t.Errorf("generateRandomElements(0) = %d; want 0", len(result))
	}
}

func TestMaximum(t *testing.T) {
	if got := maximum([]int{}); got != 0 {
		t.Errorf("maximum([]int{}) = %d; want 0", got)
	}

	if got := maximum([]int{5}); got != 5 {
		t.Errorf("maximum([]int{5}) = %d; want 5", got)
	}

	if got := maximum([]int{1, 3, 2}); got != 3 {
		t.Errorf("maximum([]int{1, 3, 2}) = %d; want 3", got)
	}
}
