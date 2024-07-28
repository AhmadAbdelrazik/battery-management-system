package main

import (
	"errors"
	"slices"
)

var ErrDimensions = errors.New("columns of a and rows of b should be equal")

func MatMul(A, B [][]float64) [][]float64 {

	rowA, colA, rowB, colB := len(A), len(A[0]), len(B), len(B[0])

	if colA != rowB {
		panic("not equal")
	}

	result := make([][]float64, rowA)
	for i := range result {
		result[i] = make([]float64, colB)
	}

	for row := range result {
		for col := range result[row] {
			for i := range colA {
				result[row][col] += A[row][i] * B[i][col]
			}
		}
	}

	return result
}

func MatT(A [][]float64) [][]float64 {
	rowA, colA := len(A), len(A[0])

	result := make([][]float64, colA)
	for i := range colA {
		result[i] = make([]float64, rowA)
	}

	for row := range rowA {
		for col := range colA {
			result[col][row] = A[row][col]
		}
	}

	return result
}

func MatDiv(A, B [][]float64) [][]float64 {
	return MatMul(A, MatT(B))
}

func MatAdd(A, B [][]float64) [][]float64 {
	if len(A) != len(B) || len(A[0]) != len(B[0]) {
		panic("dimensions error")
	}

	result := slices.Clone(A)
	for i := range len(result) {
		for j := range len(result[0]) {
			result[i][j] += B[i][j]
		}
	}

	return result
}

func MatSub(A, B [][]float64) [][]float64 {
	if len(A) != len(B) || len(A[0]) != len(B[0]) {
		panic("dimensions error")
	}

	result := slices.Clone(A)
	for i := range len(result) {
		for j := range len(result[0]) {
			result[i][j] -= B[i][j]
		}
	}

	return result
}

func MatMulC(A [][]float64, C float64) [][]float64 {
	result := slices.Clone(A)

	for i := range len(result) {
		for j := range len(result[0]) {
			result[i][j] *= C
		}
	}
	return result
}

func MatI(Dimension int) [][]float64 {
	I := make([][]float64, Dimension)

	for i := range Dimension {
		I[i] = make([]float64, Dimension)
	}

	for i := range len(I) {
		for j := range len(I) {
			if i == j {
				I[i][j] = 1
			}
		}
	}

	return I
}
