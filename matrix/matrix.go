package matrix

import (
	"fmt"
)

type Matrix struct {
	data   [][]float64
	nbRows int
	nbCols int
}

// Get the data of the matrix
func (m *Matrix) GetData() [][]float64 {
	return m.data
}

// Get the numbers of Rows of the matrix
func (m *Matrix) GetNbRows() int {
	return m.nbRows
}

// Get the number of columns of the matrix
func (m *Matrix) GetNbCols() int {
	return m.nbCols
}

// Returns true if the matrix is a square matrix, false otherwise
func (m *Matrix) IsSquareMatrix() bool {
	return m.nbCols == m.nbRows
}

// Get the element at index idx
func (m *Matrix) GetElementAt(row int, col int) float64 {
	return m.data[row][col]
}

// Set the element called elt at indices (row,col)
func (m *Matrix) SetElementAt(row int, col int, elt float64) {
	m.data[row][col] = elt
}

// Create a new matrix from a given number of rows and a given number of columns
func New(nbRowsMat int, nbColsMat int) *Matrix {
	data := make([][]float64, nbRowsMat)
	for i := range data {
		data[i] = make([]float64, nbColsMat)
	}
	return &Matrix{
		data:   data,
		nbRows: nbRowsMat,
		nbCols: nbColsMat,
	}
}

// Create a new matrix from a given 2D slice (deep copy)
//
// If the input has no row, then a pointer to a new empty matrix
// Is returned
//
// Return an error if there is any inconsistency in the number of columns
func NewFromData(mData [][]float64) (*Matrix, error) {
	nbRows := len(mData)
	if nbRows == 0 {
		return &Matrix{
			data:   [][]float64{},
			nbRows: 0,
			nbCols: 0,
		}, nil
	}

	nbCols := len(mData[0])

	// Validate & deep copy
	copiedMatrix := make([][]float64, nbRows)
	for i := range mData {
		if len(mData[i]) != nbCols {
			return nil, fmt.Errorf("inconsistent number of columns in row %d: expected %d, got %d", i, nbCols, len(mData[i]))
		}
		copiedMatrix[i] = make([]float64, nbCols)
		copy(copiedMatrix[i], mData[i])
	}

	return &Matrix{
		data:   copiedMatrix,
		nbRows: nbRows,
		nbCols: nbCols,
	}, nil
}

// Tells if a matrix is a zero matrix
//
// Returns true if so, else otherwise
func (m *Matrix) IsZero() bool {
	for i := 0; i < m.nbRows; i++ {
		for j := 0; j < m.nbCols; j++ {
			if m.data[i][j] != 0 {
				return false
			}
		}
	}

	return true
}

// Performs the addition between two matrices and returns the resulting one
//
// Returns an error if both matrices do not share the same number of rows or columns
func (m *Matrix) Add(other *Matrix) (*Matrix, error) {
	if m.nbRows != other.nbRows || m.nbCols != other.nbCols {
		return nil, fmt.Errorf("mismatch in the number of rows or columns between matrices, cannot perform addition")
	}

	result := New(m.nbRows, m.nbCols)
	for i := 0; i < m.nbRows; i++ {
		for j := 0; j < m.nbCols; j++ {
			result.data[i][j] = m.data[i][j] + other.data[i][j]
		}
	}

	return result, nil
}

// Performs the substraction between two matrices and returns the resulting one
//
// Returns an error if both matrices do not share the same number of rows or columns
func (m *Matrix) Sub(other *Matrix) (*Matrix, error) {
	if m.nbRows != other.nbRows || m.nbCols != other.nbCols {
		return nil, fmt.Errorf("mismatch in the number of rows or columns between matrices, cannot perform substraction")
	}

	result := New(m.nbRows, m.nbCols)
	for i := 0; i < m.nbRows; i++ {
		for j := 0; j < m.nbCols; j++ {
			result.data[i][j] = m.data[i][j] - other.data[i][j]
		}
	}

	return result, nil
}

// Tells if a matrix is a diagonal matrix
//
// Returns true if so, false otherwise
// By convention, a matrix with no row or column is a diagonal matrix
func (m *Matrix) IsDiagonal() bool {
	if !m.IsSquareMatrix() {
		return false
	}
	// By convention
	if m.nbRows == 0 || m.nbCols == 0 {
		return true
	}

	for i := 0; i < m.nbRows; i++ {
		for j := 0; j < m.nbCols; j++ {
			if i != j && m.data[i][j] != 0.0 {
				return false
			}
		}
	}

	return true
}

// Tells if a matrix is a hollow matrix
//
// Returns true if so, false otherwise
// By convention, a matrix with no row or column is a hollow matrix
func (m *Matrix) IsHollow() bool {
	if !m.IsSquareMatrix() {
		return false
	}
	// By convention
	if m.nbRows == 0 || m.nbCols == 0 {
		return true
	}

	for i := 0; i < m.nbRows; i++ {
		if m.data[i][i] != 0.0 {
			return false
		}
	}

	return true
}

// Tells if a matrix is the identity matrix
//
// Returns true if so, false otherwise
// By convention, a matrix with no row or column is an identity matrix
func (m *Matrix) IsIdentity() bool {
	if !m.IsSquareMatrix() {
		return false
	}
	// By convention
	if m.nbRows == 0 || m.nbCols == 0 {
		return true
	}

	for i := 0; i < m.nbRows; i++ {
		for j := 0; j < m.nbCols; j++ {
			if i != j && m.data[i][j] != 0.0 {
				return false
			}
			if i == j && m.data[i][j] != 1.0 {
				return false
			}
		}
	}

	return true
}

/*
func (m *Matrix) Mul(other *Matrix) (*Matrix, error) {

}

func (m *Matrix) Div(other *Matrix) (*Matrix, error) {

}
*/
