package main

import (
	"slices"
	"testing"
)

func TestMatrix(t *testing.T) {
	A := [][]float64{
		{2, 1},
		{1, 4},
	}

	B := [][]float64{
		{1, 2, 0},
		{0, 1, 2},
	}

	want := [][]float64{
		{2, 5, 2},
		{1, 6, 8},
	}

	got, err := multiply(A, B)

	if err != nil {
		t.Fatalf("got %v", err)
	}

	if len(got) != len(want) {
		t.Fatalf("len(got) = %v, len(want) = %v", len(got), len(want))
	}

	for i := range len(got) {
		if !slices.Equal(got[i], want[i]) {
			t.Fatalf("got %v want %v", got, want)
		}
	}
}
