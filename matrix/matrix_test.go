package matrix

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	var newMatrix *Matrix = New(2, 2)

	assert.Equal(t, 2, newMatrix.nbCols)
	assert.Equal(t, 2, newMatrix.nbRows)
	assert.Equal(t, newMatrix.data, [][]float64{{0.0, 0.0}, {0.0, 0.0}})
}

func TestGetData(t *testing.T) {
	data := [][]float64{{1.0, 2.0}, {3.0, 4.0}}
	newMatrix, err := NewFromData(data)

	assert.NoError(t, err)
	assert.Equal(t, data, newMatrix.GetData())
}

func TestGetNbRows(t *testing.T) {
	var newMatrix *Matrix = New(2, 2)

	assert.Equal(t, 2, newMatrix.GetNbRows())
}

func TestGetNbCols(t *testing.T) {
	var newMatrix *Matrix = New(2, 2)

	assert.Equal(t, 2, newMatrix.GetNbCols())
}

func TestIsSquareMatrix(t *testing.T) {
	var newMatrix *Matrix = New(2, 3)
	var newSquareMatrix *Matrix = New(2, 2)

	assert.Equal(t, false, newMatrix.IsSquareMatrix())
	assert.Equal(t, true, newSquareMatrix.IsSquareMatrix())
}

func TestSetElementAt(t *testing.T) {
	col := 2
	row := 0
	elt := 8.0

	var newMatrix *Matrix = New(4, 4)
	newMatrix.SetElementAt(row, col, elt)

	assert.Equal(t, elt, newMatrix.GetElementAt(row, col))
}

func TestNewFromData_ValidInput(t *testing.T) {
	input := [][]float64{{1.0, 2.0}, {3.0, 4.0}}

	mat, err := NewFromData(input)
	assert.NoError(t, err)

	assert.Equal(t, 2, mat.nbRows)
	assert.Equal(t, 2, mat.nbCols)
	assert.Equal(t, [][]float64{{1.0, 2.0}, {3.0, 4.0}}, mat.data)
}

func TestNewFromData_EmptyInput(t *testing.T) {
	input := [][]float64{}

	mat, err := NewFromData(input)
	assert.NoError(t, err)

	assert.Equal(t, 0, mat.nbRows)
	assert.Equal(t, 0, mat.nbCols)
	assert.Empty(t, mat.data)
}

func TestNewFromData_DeepCopy(t *testing.T) {
	input := [][]float64{{1.0, 2.0}, {3.0, 4.0}}

	mat, err := NewFromData(input)
	assert.NoError(t, err)

	input[0][0] = 99.0
	// Should be the old value still
	assert.Equal(t, 1.0, mat.GetElementAt(0, 0))
}

func TestNewFromData_InconsistentCols_ReturnsError(t *testing.T) {
	input := [][]float64{{1.0, 2.0}, {3.0}}

	mat, err := NewFromData(input)
	assert.Nil(t, mat)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "inconsistent number of columns")
}

func TestIsZero_ShouldBeTrue(t *testing.T) {
	var newMatrix *Matrix = New(2, 3)

	assert.Equal(t, true, newMatrix.IsZero())
}

func TestIsZero_ShouldBeFalse(t *testing.T) {
	var newMatrix *Matrix = New(2, 3)
	newMatrix.data[0][0] = 1.0

	assert.Equal(t, false, newMatrix.IsZero())
}

func TestMatrix_Add_Success(t *testing.T) {
	m1, err1 := NewFromData([][]float64{
		{1, 2},
		{3, 4},
	})
	m2, err2 := NewFromData([][]float64{
		{5, 6},
		{7, 8},
	})

	assert.NoError(t, err1)
	assert.NoError(t, err2)

	result, err := m1.Add(m2)
	assert.NoError(t, err)

	expected, errExpected := NewFromData([][]float64{
		{6, 8},
		{10, 12},
	})
	assert.NoError(t, errExpected)

	assert.Equal(t, expected.data, result.data)
}

func TestMatrix_Add_DimensionMismatch_ShouldFail(t *testing.T) {
	m1, err1 := NewFromData([][]float64{
		{1, 2},
	})
	m2, err2 := NewFromData([][]float64{
		{3, 4},
		{5, 6},
	})

	assert.NoError(t, err1)
	assert.NoError(t, err2)

	_, err := m1.Add(m2)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "mismatch in the number of rows or columns")
}

func TestMatrix_Sub_Success(t *testing.T) {
	m1, err1 := NewFromData([][]float64{
		{9, 8},
		{7, 6},
	})
	m2, err2 := NewFromData([][]float64{
		{1, 2},
		{3, 4},
	})

	assert.NoError(t, err1)
	assert.NoError(t, err2)

	result, err := m1.Sub(m2)
	assert.NoError(t, err)

	expected, errExpected := NewFromData([][]float64{
		{8, 6},
		{4, 2},
	})
	assert.NoError(t, errExpected)

	assert.Equal(t, expected.data, result.data)
}

func TestMatrix_Sub_DimensionMismatch_ShouldFail(t *testing.T) {
	m1, err1 := NewFromData([][]float64{
		{1, 2, 3},
	})
	m2, err2 := NewFromData([][]float64{
		{4, 5},
	})

	assert.NoError(t, err1)
	assert.NoError(t, err2)

	_, err := m1.Sub(m2)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "mismatch in the number of rows or columns")
}

func TestIsDiagonal_ShouldReturnTrue_ForEmptyMatrix(t *testing.T) {
	m, err := NewFromData([][]float64{})

	assert.NoError(t, err)
	assert.True(t, m.IsDiagonal())
}

func TestIsDiagonal_ShouldReturnTrue_ForValidDiagonal(t *testing.T) {
	m, err := NewFromData([][]float64{
		{1, 0, 0},
		{0, 2, 0},
		{0, 0, 3},
	})

	assert.NoError(t, err)
	assert.True(t, m.IsDiagonal())
}

func TestIsDiagonal_ShouldReturnFalse_ForNonSquareMatrix(t *testing.T) {
	m, err := NewFromData([][]float64{
		{1, 0, 0},
		{0, 2, 0},
	})

	assert.NoError(t, err)
	assert.False(t, m.IsDiagonal())
}

func TestIsDiagonal_ShouldReturnFalse_IfNonDiagonalElementIsNonZero(t *testing.T) {
	m, err := NewFromData([][]float64{
		{1, 2, 0},
		{0, 2, 0},
		{0, 0, 3},
	})

	assert.NoError(t, err)
	assert.False(t, m.IsDiagonal())
}

func TestIsHollow_ShouldReturnTrue_ForEmptyMatrix(t *testing.T) {
	m, err := NewFromData([][]float64{})

	assert.NoError(t, err)
	assert.True(t, m.IsHollow())
}

func TestIsHollow_ShouldReturnTrue_WhenAllDiagonalElementsAreZero(t *testing.T) {
	m, err := NewFromData([][]float64{
		{0, 1, 2},
		{3, 0, 4},
		{5, 6, 0},
	})

	assert.NoError(t, err)
	assert.True(t, m.IsHollow())
}

func TestIsHollow_ShouldReturnFalse_IfOneDiagonalElementIsNonZero(t *testing.T) {
	m, err := NewFromData([][]float64{
		{0, 1, 2},
		{3, 9, 4},
		{5, 6, 0},
	})

	assert.NoError(t, err)
	assert.False(t, m.IsHollow())
}

func TestIsHollow_ShouldReturnFalse_ForNonSquareMatrix(t *testing.T) {
	m, err := NewFromData([][]float64{
		{0, 1},
		{1, 0},
		{2, 2},
	})

	assert.NoError(t, err)
	assert.False(t, m.IsHollow())
}

func TestIsIdentity_ShouldReturnTrue_On3x3IdentityMatrix(t *testing.T) {
	m, err := NewFromData([][]float64{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	})

	require.NoError(t, err)
	assert.True(t, m.IsIdentity())
}

func TestIsIdentity_ShouldReturnFalse_OnNonSquareMatrix(t *testing.T) {
	m, err := NewFromData([][]float64{
		{1, 0, 0},
		{0, 1, 0},
	})

	require.NoError(t, err)
	assert.False(t, m.IsIdentity())
}

func TestIsIdentity_ShouldReturnFalse_WhenDiagonalNot1(t *testing.T) {
	m, err := NewFromData([][]float64{
		{2, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	})

	require.NoError(t, err)
	assert.False(t, m.IsIdentity())
}

func TestIsIdentity_ShouldReturnFalse_WhenOffDiagonalNot0(t *testing.T) {
	m, err := NewFromData([][]float64{
		{1, 1, 0},
		{0, 1, 0},
		{0, 0, 1},
	})

	require.NoError(t, err)
	assert.False(t, m.IsIdentity())
}

func TestIsIdentity_ShouldReturnTrue_OnEmptyMatrix(t *testing.T) {
	m, err := NewFromData([][]float64{})

	require.NoError(t, err)
	assert.True(t, m.IsIdentity())
}
