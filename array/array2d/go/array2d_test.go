package arrays

import (
	"testing"
)

func TestMatrix_GetSet(t *testing.T) {
	matrix := NewMatrix(3, 3)

	matrix_SetAndGet(t, matrix)
	matrix_OutOfRange(t, matrix)
}

func TestMatrix_Exists(t *testing.T) {
	matrix := NewMatrix(3, 3)

	matrix_Exists(t, matrix)
}

func TestMatrix_GetSetEdgeCases(t *testing.T) {
	matrix := NewMatrix(3, 3)

	matrix_NegativeIndices(t, matrix)
	matrix_EmptyMatrix(t)
	matrix_OutOfRangeIndices(t, matrix)
}

func TestMatrix_ExistsEdgeCases(t *testing.T) {
	matrix := NewMatrix(3, 3)

	matrix_NegativeIndicesExists(t, matrix)
	matrix_EmptyMatrixExists(t)
	matrix_OutOfRangeIndicesExists(t, matrix)
}

func matrix_SetAndGet(t *testing.T, matrix *Matrix) {
	matrix.Set(1, 1, 5)
	value, err := matrix.Get(1, 1)
	if err != nil {
		t.Errorf("Error occurred while getting value: %v", err)
	}
	if value != 5 {
		t.Errorf("Expected value %d, got %d", 5, value)
	}
}

func matrix_OutOfRange(t *testing.T, matrix *Matrix) {
	_, err := matrix.Get(4, 4)
	if err == nil {
		t.Error("Expected an error for out-of-range index")
	}

	_, err = matrix.Set(4, 4, 10)
	if err == nil {
		t.Error("Expected an error for out-of-range index")
	}
}

func matrix_Exists(t *testing.T, matrix *Matrix) {
	if !matrix.exists(1, 1) {
		t.Error("Expected (1, 1) to exist")
	}

	if matrix.exists(4, 4) {
		t.Error("Expected (4, 4) to be out of range")
	}
}

func matrix_NegativeIndices(t *testing.T, matrix *Matrix) {
	_, err := matrix.Get(-1, 1)
	if err == nil {
		t.Error("Expected an error for negative row index")
	}

	_, err = matrix.Get(1, -1)
	if err == nil {
		t.Error("Expected an error for negative column index")
	}

	_, err = matrix.Set(-1, 1, 10)
	if err == nil {
		t.Error("Expected an error for negative row index")
	}

	_, err = matrix.Set(1, -1, 10)
	if err == nil {
		t.Error("Expected an error for negative column index")
	}
}

func matrix_EmptyMatrix(t *testing.T) {
	emptyMatrix := NewMatrix(0, 0)
	_, err := emptyMatrix.Get(0, 0)
	if err == nil {
		t.Error("Expected an error for accessing an element in an empty matrix")
	}

	_, err = emptyMatrix.Set(0, 0, 10)
	if err == nil {
		t.Error("Expected an error for setting an element in an empty matrix")
	}
}

func matrix_OutOfRangeIndices(t *testing.T, matrix *Matrix) {
	_, err := matrix.Get(10, 10)
	if err == nil {
		t.Error("Expected an error for out-of-range row and column indices")
	}

	_, err = matrix.Set(10, 10, 10)
	if err == nil {
		t.Error("Expected an error for out-of-range row and column indices")
	}
}

func matrix_NegativeIndicesExists(t *testing.T, matrix *Matrix) {
	if matrix.exists(-1, 1) {
		t.Error("Expected (row: -1, col: 1) to be out of range")
	}

	if matrix.exists(1, -1) {
		t.Error("Expected (row: 1, col: -1) to be out of range")
	}
}

func matrix_EmptyMatrixExists(t *testing.T) {
	emptyMatrix := NewMatrix(0, 0)
	if emptyMatrix.exists(0, 0) {
		t.Error("Expected (row: 0, col: 0) to be out of range in an empty matrix")
	}
}

func matrix_OutOfRangeIndicesExists(t *testing.T, matrix *Matrix) {
	if matrix.exists(10, 10) {
		t.Error("Expected (row: 10, col: 10) to be out of range")
	}
}
