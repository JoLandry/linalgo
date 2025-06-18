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
