package matrix

import (
	"fmt"
	"math"
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
func (m *Matrix) IsSquare() bool {
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

// Creates a new identity matrix from a given size
func NewIdentity(size int) *Matrix {
	data := make([][]float64, size)
	for i := range data {
		data[i] = make([]float64, size)
	}
	idMatrix := &Matrix{
		data:   data,
		nbRows: size,
		nbCols: size,
	}

	for k := 0; k < size; k++ {
		data[k][k] = 1.0
	}

	return idMatrix
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
	if !m.IsSquare() {
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

// Tells if a matrix is a scalar matrix
//
// Returns true if so, false otherwise
// By convention, a matrix with no row or column is a scalar matrix
func (m *Matrix) IsScalar(scalar float64) bool {
	if !m.IsSquare() {
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
			if i == j && m.data[i][j] != scalar {
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
	if !m.IsSquare() {
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
	if !m.IsSquare() {
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
	if !m.IsSquare() {
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

// Returns a new matrix in Row Echelon Form (REF)
// Uses Gaussian elimination
//
// This function does not modify the original matrix
func (m *Matrix) ToRowEchelon() *Matrix {
	if m.nbRows == 0 || m.nbCols == 0 {
		return New(m.nbRows, m.nbCols)
	}

	// Copy matrix
	ref, err := NewFromData(m.data)
	if err != nil {
		panic(fmt.Sprintf("invalid matrix data: %v", err))
	}

	row := 0
	for col := 0; col < ref.nbCols && row < ref.nbRows; col++ {
		// Find pivot in the current column
		pivotRow := -1
		for r := row; r < ref.nbRows; r++ {
			if math.Abs(ref.data[r][col]) > 1e-10 {
				pivotRow = r
				// Pivot found so break here
				break
			}
		}

		// Skip column if all entries are zero
		if pivotRow == -1 {
			continue
		}
		// Swap current row with pivot row
		if pivotRow != row {
			ref.data[row], ref.data[pivotRow] = ref.data[pivotRow], ref.data[row]
		}
		// Eliminate below
		for r := row + 1; r < ref.nbRows; r++ {
			factor := ref.data[r][col] / ref.data[row][col]
			for c := col; c < ref.nbCols; c++ {
				ref.data[r][c] -= factor * ref.data[row][c]
			}
		}
		row++
	}

	return ref
}

// Returns the rank of the matrix
func (m *Matrix) Rank() int {
	ref := m.ToRowEchelon()
	rank := 0
	for _, row := range ref.data {
		notZero := false
		for _, val := range row {
			if math.Abs(val) > 1e-10 {
				notZero = true
				break
			}
		}
		if notZero {
			rank++
		}
	}

	return rank
}

// Tells if a matrix is said "full rank"
//
// Namely if its rank is the highest possible for a matrix of the same size
func (m *Matrix) IsFullRank() bool {
	rank := m.Rank()
	min := min(m.nbRows, m.nbCols)

	return rank == min
}

// Tells if a matrix is invertible
//
// Namely:
//   - it is full rank
//   - it is squared
//   - its determinant is not zero
func (m *Matrix) IsInvertible() bool {
	// Check if is full rank
	if !m.IsFullRank() {
		return false
	}
	// Must be squared
	if !m.IsSquare() {
		return false
	}
	// Determinant must be non-zero
	determinant, err := m.Determinant()
	if err != nil || math.Abs(determinant) < 1e-10 {
		return false
	}

	return true
}

// Returns the inverse of the matrix if it exists
//
// It returns an error if the matrix is singular (non-invertible) or not square
// For 2x2 matrices, it uses the direct "closed-form" formula
// For larger matrices, it uses Gaussian elimination with pivoting
//
// The original matrix is not modified
func (m *Matrix) Invert() (*Matrix, error) {
	if !m.IsInvertible() {
		return nil, fmt.Errorf("matrix is singular, cannot invert")

	}
	if !m.IsSquare() {
		return nil, fmt.Errorf("matrix is not squared, it cannot be inverted")
	}
	// 2x2 matrix
	if m.nbCols == 2 {
		a := m.data[0][0]
		b := m.data[0][1]
		c := m.data[1][0]
		d := m.data[1][1]
		det := a*d - b*c
		if math.Abs(det) < 1e-10 {
			return nil, fmt.Errorf("matrix determinant is zero, cannot invert")
		}
		invDet := 1.0 / det
		return NewFromData([][]float64{
			{d * invDet, -b * invDet},
			{-c * invDet, a * invDet},
		})
	}
	// General case
	matrixCopy, err := NewFromData(m.data)
	if err != nil {
		return nil, fmt.Errorf("invalid matrix: %w", err)
	}
	// Start with identity matrix that will be transformed into inverse at the end of the algorithm
	invMatrix := NewIdentity(m.nbRows)

	n := m.nbRows
	for i := 0; i < n; i++ {
		// Find pivot element (max absolute value in column i at or below row i)
		pivot := i
		maxVal := math.Abs(matrixCopy.data[i][i])
		for r := i + 1; r < n; r++ {
			if math.Abs(matrixCopy.data[r][i]) > maxVal {
				maxVal = math.Abs(matrixCopy.data[r][i])
				pivot = r
			}
		}

		// Matrix cannot be inverted
		if maxVal < 1e-10 {
			return nil, fmt.Errorf("matrix is singular, cannot invert")
		}

		// Swap rows in both matrixCopy and invMatrix if needed
		if pivot != i {
			matrixCopy.data[i], matrixCopy.data[pivot] = matrixCopy.data[pivot], matrixCopy.data[i]
			invMatrix.data[i], invMatrix.data[pivot] = invMatrix.data[pivot], invMatrix.data[i]
		}

		// Normalize pivot row
		pivotVal := matrixCopy.data[i][i]
		for col := 0; col < n; col++ {
			matrixCopy.data[i][col] /= pivotVal
			invMatrix.data[i][col] /= pivotVal
		}

		// Eliminate all other rows
		for row := 0; row < n; row++ {
			if row != i {
				factor := matrixCopy.data[row][i]
				for col := 0; col < n; col++ {
					matrixCopy.data[row][col] -= factor * matrixCopy.data[i][col]
					invMatrix.data[row][col] -= factor * invMatrix.data[i][col]
				}
			}
		}
	}

	return invMatrix, nil
}

// Performs division between two matrices and returns the resulting one
//
// In other words, performs multiplication between the matrix, and the inverse of the given one
func (m *Matrix) Div(other *Matrix) (*Matrix, error) {
	invert, err := other.Invert()
	if err != nil {
		return nil, fmt.Errorf("could not perform matrix division because given matrix cannot be inverted")
	}
	result, err := m.Mul(invert)
	if err != nil {
		return nil, fmt.Errorf("could not perform matrix multiplication")
	}

	return result, nil
}

// Helper method that compares two matrices for approximate equality (float64 type)
//
// It returns true if both matrices have the same dimensions and each pair of
// corresponding elements differ by no more than the specified epsilon value
func (m *Matrix) EqualsApprox(other *Matrix, epsilon float64) bool {
	if m.nbRows != other.nbRows || m.nbCols != other.nbCols {
		return false
	}
	for i := 0; i < m.nbRows; i++ {
		for j := 0; j < m.nbCols; j++ {
			if math.Abs(m.data[i][j]-other.data[i][j]) > epsilon {
				return false
			}
		}
	}
	return true
}

// Returns the matrix raised to the given integer power
//
// Power 0 returns the identity matrix (only for square matrices)
// Negative powers compute the inverse of the matrix first (only if invertible)
func (m *Matrix) Pow(power int) (*Matrix, error) {
	// Check if it's square first
	if !m.IsSquare() {
		return nil, fmt.Errorf("power %d cannot be computed for a non-square matrix", power)
	}
	// Identity matrix for power 0
	if power == 0 {
		return NewIdentity(m.nbRows), nil
	}
	// The matrix itself, make a copy
	if power == 1 {
		return NewFromData(m.data)
	}

	base := m
	var err error

	// Negative power: invert the matrix
	if power < 0 {
		base, err = m.Invert()
		if err != nil {
			return nil, fmt.Errorf("cannot raise to power %d: matrix is not invertible", power)
		}
		power = -power
	}

	// Exponentiation by repeated multiplication
	result := NewIdentity(m.nbRows)
	for i := 0; i < power; i++ {
		result, err = result.Mul(base)
		if err != nil {
			return nil, fmt.Errorf("matrix multiplication failed: %w", err)
		}
	}

	return result, nil
}

// Returns the transpose of the matrix (rows become columns and vice versa)
func (m *Matrix) Transpose() *Matrix {
	if m.nbRows == 0 || m.nbCols == 0 {
		return New(0, 0)
	}

	result := New(m.nbCols, m.nbRows)
	for i := 0; i < m.nbRows; i++ {
		for j := 0; j < m.nbCols; j++ {
			result.data[j][i] = m.data[i][j]
		}
	}

	return result
}
