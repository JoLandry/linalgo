// Package matrix provides basic matrix operations including creation,
// access, arithmetic, and linear algebra utilities such as determinant,
// rank, transpose, inversion and others.
package matrix

import (
	"fmt"
	"math"
	"strings"
)

// Matrix represents a two-dimensional matrix of float64 values.
//
// It includes metadata for the number of rows and columns to simplify
// operations and validations.
type Matrix struct {
	data   [][]float64
	nbRows int
	nbCols int
}

// String returns a human-readable string representation of the matrix.
func (m *Matrix) String() string {
	if m.nbRows == 0 || m.nbCols == 0 {
		return "[]"
	}

	var builder strings.Builder
	builder.WriteString("[\n")
	for i := 0; i < m.nbRows; i++ {
		builder.WriteString("  [")
		for j := 0; j < m.nbCols; j++ {
			builder.WriteString(fmt.Sprintf("%8.4f", m.data[i][j]))
			if j < m.nbCols-1 {
				builder.WriteString(", ")
			}
		}
		builder.WriteString("]")
		if i < m.nbRows-1 {
			builder.WriteString(",\n")
		}
	}
	builder.WriteString("\n]")

	return builder.String()
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

// Creates and returns a new identity matrix of a given size.
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

// Creates and returns a matrix from a flat slice of values.
//
// The matrix will have the specified number of rows and columns.
// If the number of values is not exactly rows * cols, it panics.
func NewFromFlat(rows, cols int, values []float64) *Matrix {
	if len(values) != rows*cols {
		panic("number of values does not match matrix dimensions")
	}

	data := make([][]float64, rows)
	for i := 0; i < rows; i++ {
		data[i] = make([]float64, cols)
		for j := 0; j < cols; j++ {
			data[i][j] = values[i*cols+j]
		}
	}

	return &Matrix{
		data:   data,
		nbRows: rows,
		nbCols: cols,
	}
}

// Creates and returns a square diagonal matrix with the provided diagonal values.
//
// All non-diagonal elements are set to zero.
func NewDiagonal(values []float64) *Matrix {
	size := len(values)
	data := make([][]float64, size)
	for i := 0; i < size; i++ {
		data[i] = make([]float64, size)
		data[i][i] = values[i]
	}

	return &Matrix{
		data:   data,
		nbRows: size,
		nbCols: size,
	}
}

// Tells whether all elements of the matrix are zero.
//
// Returns true if the matrix is a zero matrix, false otherwise.
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

// Returns a new matrix that is the element-wise sum of m and other.
//
// Returns an error if the matrices do not have the same dimensions.
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

// Returns a new matrix that is the element-wise difference of m and other.
//
// Returns an error if the matrices do not have the same dimensions.
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

// Tells whether the matrix is a diagonal matrix.
//
// By convention, an empty (0x0) matrix is considered diagonal.
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

// Tells whether the matrix is a scalar matrix with the given scalar value.
//
// By convention, an empty (0x0) matrix is considered scalar.
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

// Tells whether the matrix is a hollow matrix.
//
// By convention, an empty (0x0) matrix is considered hollow.
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

// Tells whether the matrix is the identity matrix.
//
// By convention, an empty (0x0) matrix is considered an identity matrix.
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
// of the current matrix with the given scalar value.
//
// Each element of the matrix is multiplied by the scalar. The original matrix
// is not modified. If the scalar is 1.0 or the matrix is a zero matrix,
// the original matrix is returned as-is (no copy is made).
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

// Performs standard matrix multiplication between the receiver matrix and another given matrix.
//
// It returns a new matrix representing the product m * other. If the number of
// columns in the calling matrix does not match the number of rows in the other matrix,
// an error is returned. If either matrix is a zero matrix, a new zero matrix of
// the correct dimensions is returned.
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

// Returns the determinant of the calling matrix.
//
// It returns an error if the matrix is not square.
// Internally, it uses a recursive Laplace expansion approach.
func (m *Matrix) Determinant() (float64, error) {
	if !m.IsSquare() {
		return 0.0, fmt.Errorf("cannot compute the determinant of a non-square matrix")
	}
	return m.determinantRecursive(), nil
}

// Computes the determinant of a square matrix recursively
// using Laplace expansion along the first row.
//
// This helper method assumes the matrix is square and valid.
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

// Returns the minor matrix obtained by removing the specified
// row and column from the current matrix.
//
// It is used as a helper in determinant computations.
// The resulting minor is a deep copy of the reduced matrix.
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

// Returns a new matrix that is the Row Echelon Form (REF)
// of the calling matrix.
//
// The method applies Gaussian elimination without row scaling,
// and does not modify the original matrix.
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

// Returns the rank of the matrix.
//
// The rank is computed by transforming the matrix into its Row Echelon Form (REF)
// using Gaussian elimination and counting the number of non-zero rows.
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

// Tells whether the matrix is full rank.
//
// A matrix is full rank if its rank is equal to the smaller of its number
// of rows and columns.
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

// Invert returns the inverse of the matrix.
//
// For 2x2 matrices, it uses a "closed-form" formula.
// For larger square matrices, it applies Gaussian elimination with pivoting.
//
// Returns an error if the matrix is not square or is singular (non-invertible).
// The original matrix is not modified.
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

// Performs matrix division by multiplying the current matrix by the inverse of the given matrix.
//
// Returns an error if the given matrix is not invertible or if multiplication fails.
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

// Helper method that compares two matrices for approximate equality (float64 type).
//
// It returns true if both matrices have the same dimensions and each pair of
// corresponding elements differ by no more than the specified epsilon value.
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

// Returns the matrix raised to the given integer power.
//
// Power 0 returns the identity matrix (only for square matrices).
// Negative powers compute the inverse of the matrix first (only if invertible).
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

// Transpose returns the transpose of the matrix.
//
// The transpose flips the matrix over its diagonal, converting rows to columns and vice versa.
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
