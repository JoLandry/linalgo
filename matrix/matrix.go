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

// Returns a new matrix resulting from the scalar multiplication
// of the current matrix with the given scalar value
//
// Each element of the original matrix is multiplied by the scalar value
// The original matrix is not modified
//
// If the scalar is 1.0, the original matrix is returned as-is
// If the matrix is a zero matrix, it is returned as-is
func (m *Matrix) MulScalar(scalar float64) *Matrix {
	if m.IsZero() || scalar == 1.0 {
		return m
	}

	result := New(m.nbRows, m.nbCols)
	for i := 0; i < m.nbRows; i++ {
		for j := 0; j < m.nbCols; j++ {
			result.data[i][j] = m.data[i][j] * scalar
		}
	}
	return result
}

// Performs matrix multiplication between the current matrix and the given matrix
//
// Returns a new matrix result = m * other, or an error if dimensions are incompatible
func (m *Matrix) Mul(other *Matrix) (*Matrix, error) {
	if m.nbCols != other.nbRows {
		return nil, fmt.Errorf("mismatch between number of columns of first matrix (%d) and number of rows of second matrix (%d), cannot perform multiplication", m.nbCols, other.nbRows)
	}
	if m.IsZero() || other.IsZero() {
		return New(m.nbRows, other.nbCols), nil
	}

	result := New(m.nbRows, other.nbCols)
	for i := 0; i < m.nbRows; i++ {
		for j := 0; j < other.nbCols; j++ {
			idx := 0
			value := 0.0
			for idx < m.nbCols {
				value += m.data[i][idx] * other.data[idx][j]
				idx++
			}

			result.data[i][j] = value
		}
	}

	return result, nil
}

// Returns the determinant of a matrix
//
// Returns an error is the matrix is not squared
func (m *Matrix) Determinant() (float64, error) {
	if !m.IsSquareMatrix() {
		return 0.0, fmt.Errorf("cannot compute the determinant of a non-square matrix")
	}
	return m.determinantRecursive(), nil
}

// Helper function
// Returns the determinant of a matrix by breaking it down
// Into smaller matrices recursively
func (m *Matrix) determinantRecursive() float64 {
	if m.nbRows == 1 {
		return m.data[0][0]
	}
	if m.nbRows == 2 {
		return m.data[0][0]*m.data[1][1] - m.data[0][1]*m.data[1][0]
	}

	determinant := 0.0
	for col := 0; col < m.nbCols; col++ {
		cofactor := m.getMinor(0, col)
		sign := 1.0
		if col%2 != 0 {
			sign = -1.0
		}
		determinant += sign * m.data[0][col] * cofactor.determinantRecursive()
	}

	return determinant
}

// Helper function
// Returns a new Matrix with row i and column j removed
func (m *Matrix) getMinor(rowToRemove, colToRemove int) *Matrix {
	minorData := [][]float64{}
	for i := 0; i < m.nbRows; i++ {
		if i == rowToRemove {
			continue
		}
		row := []float64{}
		for j := 0; j < m.nbCols; j++ {
			if j == colToRemove {
				continue
			}
			row = append(row, m.data[i][j])
		}
		minorData = append(minorData, row)
	}
	minor, _ := NewFromData(minorData)
	return minor
}

/*
func (m *Matrix) Div(other *Matrix) (*Matrix, error) {

}
*/
