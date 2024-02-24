package arrays

import "errors"

type Matrix struct {
	data [][]int
}

func NewMatrix(row_count, col_count int) *Matrix {
	matrix := make([][]int, row_count)
	for i := range matrix {
		matrix[i] = make([]int, col_count)
	}
	return &Matrix{
		data: matrix,
	}
}

func (m *Matrix) Get(row, col int) (int, error) {
	if !m.exists(row, col) {
		return 0, errors.New("index out of range")
	}
	return m.data[row][col], nil
}

func (m *Matrix) Set(row, col, value int) (int, error) {
	if !m.exists(row, col) {
		return 0, errors.New("index out of range")
	}
	m.data[row][col] = value
	return value, nil
}

func (m *Matrix) exists(row, col int) bool {
	return row >= 0 && row < len(m.data) && col >= 0 && col < len(m.data[row])
}
