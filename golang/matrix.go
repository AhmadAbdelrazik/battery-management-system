package main

// Matrix Operations to ease kalman filter.

import (
	"errors"
)

var ErrDimensions = errors.New("columns of a and rows of b should be equal")

// MatMul: Multiplies to Matrices
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

// MatT: Transpose Matrix
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

// MatDiv: Divides two matrices
func MatDiv(A, B [][]float64) [][]float64 {
	return MatMul(A, MatT(B))
}

// MatAdd: Adds two matrices
func MatAdd(A, B [][]float64) [][]float64 {
	if len(A) != len(B) || len(A[0]) != len(B[0]) {
		panic("dimensions error")
	}

	result := clone(A)
	for i := range len(result) {
		for j := range len(result[0]) {
			result[i][j] += B[i][j]
		}
	}

	return result
}

// MatSub: Subtract two matrices
func MatSub(A, B [][]float64) [][]float64 {
	if len(A) != len(B) || len(A[0]) != len(B[0]) {
		panic("dimensions error")
	}

	result := clone(A)
	for i := range len(result) {
		for j := range len(result[0]) {
			result[i][j] -= B[i][j]
		}
	}

	return result
}

// MatMulC: Multiples a matrix with a constant
func MatMulC(A [][]float64, C float64) [][]float64 {
	result := clone(A)

	for i := range len(result) {
		for j := range len(result[0]) {
			result[i][j] *= C
		}
	}
	return result
}

// MatI: Produce Identity Matrix
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

// MatDiffuse: Converts a 1 * 1 matrix to a float
// Same working as A[0][0], used for abstraction purposes.
func MatDiffuse(A [][]float64) float64 {
	return A[0][0]
}

// clone: makes a new copy of the matrix.
// this prevents editing the copied matrix in the operations.
func clone(A [][]float64) [][]float64 {
	result := make([][]float64, len(A))
	for i := range result {
		result[i] = make([]float64, len(A[0]))
	}

	for i := range A {
		for j := range A[i] {
			result[i][j] = A[i][j]
		}
	}

	return result
}
