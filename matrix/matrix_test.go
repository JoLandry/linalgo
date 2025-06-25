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

func TestMulScalar_ZeroMatrix(t *testing.T) {
	m, err := NewFromData([][]float64{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	})

	require.NoError(t, err)
	assert.Equal(t, m, m.MulScalar(2))
}

func TestMulScalar_ScalarValueIsOne(t *testing.T) {
	m, err := NewFromData([][]float64{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	})

	require.NoError(t, err)
	assert.Equal(t, m, m.MulScalar(1))
}

func TestMulScalar_ShouldSucceed(t *testing.T) {
	m, err := NewFromData([][]float64{
		{2, 2, 2},
		{2, 2, 2},
		{2, 2, 2},
	})

	expected, err2 := NewFromData([][]float64{
		{4, 4, 4},
		{4, 4, 4},
		{4, 4, 4},
	})

	require.NoError(t, err)
	require.NoError(t, err2)
	assert.Equal(t, expected.data, m.MulScalar(2).data)
}

func TestMul_ShouldSucceed(t *testing.T) {
	m1, _ := NewFromData([][]float64{
		{1, 2, 3},
		{4, 5, 6},
	})
	m2, _ := NewFromData([][]float64{
		{7, 8},
		{9, 10},
		{11, 12},
	})

	expected, _ := NewFromData([][]float64{
		{58, 64},
		{139, 154},
	})

	result, err := m1.Mul(m2)

	assert.NoError(t, err)
	assert.Equal(t, expected.data, result.data)
}

func TestMul_ShouldFail_DimensionMismatch(t *testing.T) {
	m1, _ := NewFromData([][]float64{
		{1, 2},
		{3, 4},
	})
	m2, _ := NewFromData([][]float64{
		{5, 6},
		{7, 8},
		{9, 10},
	})

	_, err := m1.Mul(m2)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "mismatch between number of columns of first matrix")
}

func TestMul_WithZeroMatrix_ShouldReturnZeroMatrix(t *testing.T) {
	m1, _ := NewFromData([][]float64{
		{1, 2},
		{3, 4},
	})
	zero, _ := NewFromData([][]float64{
		{0, 0},
		{0, 0},
	})

	result, err := m1.Mul(zero)

	assert.NoError(t, err)
	expected, _ := NewFromData([][]float64{
		{0, 0},
		{0, 0},
	})
	assert.Equal(t, expected.data, result.data)
}

func TestDeterminant_ShouldFail_NonSquareMatrix(t *testing.T) {
	m, err := NewFromData([][]float64{
		{1, 2, 3},
		{4, 5, 6},
	})
	assert.NoError(t, err)

	_, detErr := m.Determinant()
	assert.Error(t, detErr)
	assert.Contains(t, detErr.Error(), "cannot compute the determinant of a non-square matrix")
}

func TestDeterminant_EmptyMatrix_ShouldReturnZero(t *testing.T) {
	m, err := NewFromData([][]float64{})
	assert.NoError(t, err)

	det, err := m.Determinant()
	assert.NoError(t, err)
	assert.Equal(t, 0.0, det)
}

func TestDeterminant_1x1Matrix(t *testing.T) {
	m, err := NewFromData([][]float64{
		{5},
	})
	assert.NoError(t, err)

	det, err := m.Determinant()
	assert.NoError(t, err)
	assert.Equal(t, 5.0, det)
}

func TestDeterminant_2x2Matrix(t *testing.T) {
	m, err := NewFromData([][]float64{
		{4, 6},
		{3, 8},
	})
	assert.NoError(t, err)

	det, err := m.Determinant()
	assert.NoError(t, err)
	assert.Equal(t, 14.0, det)
}

func TestDeterminant_3x3Matrix(t *testing.T) {
	m, err := NewFromData([][]float64{
		{6, 1, 1},
		{4, -2, 5},
		{2, 8, 7},
	})
	assert.NoError(t, err)

	det, err := m.Determinant()
	assert.NoError(t, err)
	assert.Equal(t, -306.0, det)
}

func TestDeterminant_IdentityMatrix_ShouldReturnOne(t *testing.T) {
	m, err := NewFromData([][]float64{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	})
	assert.NoError(t, err)

	det, err := m.Determinant()
	assert.NoError(t, err)
	assert.Equal(t, 1.0, det)
}

func TestDeterminant_ZeroMatrix_ShouldReturnZero(t *testing.T) {
	m, err := NewFromData([][]float64{
		{0, 0},
		{0, 0},
	})
	assert.NoError(t, err)

	det, err := m.Determinant()
	assert.NoError(t, err)
	assert.Equal(t, 0.0, det)
}

func TestInvert_ShouldSucceed_2x2(t *testing.T) {
	m, _ := NewFromData([][]float64{
		{4, 7},
		{2, 6},
	})

	inv, err := m.Invert()

	expected, _ := NewFromData([][]float64{
		{0.6, -0.7},
		{-0.2, 0.4},
	})

	assert.NoError(t, err)
	assert.True(t, inv.EqualsApprox(expected, 1e-9), "expected %v, got %v", expected, inv)
}

func TestInvert_ShouldFail_SingularMatrix(t *testing.T) {
	m, _ := NewFromData([][]float64{
		{1, 2},
		{2, 4},
	})

	_, err := m.Invert()

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "singular")
}

func TestInvert_ShouldSucceed_3x3(t *testing.T) {
	m, _ := NewFromData([][]float64{
		{2, -1, 0},
		{-1, 2, -1},
		{0, -1, 2},
	})

	inv, err := m.Invert()

	expected, _ := NewFromData([][]float64{
		{0.75, 0.5, 0.25},
		{0.5, 1.0, 0.5},
		{0.25, 0.5, 0.75},
	})

	assert.NoError(t, err)
	assert.True(t, inv.EqualsApprox(expected, 1e-9), "expected %v, got %v", expected, inv)
}

func TestDiv_ShouldSucceed(t *testing.T) {
	a, _ := NewFromData([][]float64{
		{3, 3},
		{2, 1},
	})
	b, _ := NewFromData([][]float64{
		{1, 2},
		{1, 1},
	})

	result, err := a.Div(b)

	expected, _ := NewFromData([][]float64{
		{0, 3},
		{-1, 3},
	})

	assert.NoError(t, err)
	assert.True(t, result.EqualsApprox(expected, 1e-9), "expected %v, got %v", expected, result)
}

func TestDiv_ShouldFail_NonInvertible(t *testing.T) {
	a, _ := NewFromData([][]float64{
		{1, 2},
		{3, 4},
	})
	b, _ := NewFromData([][]float64{
		{2, 4},
		{1, 2},
	})
	// singular

	_, err := a.Div(b)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "cannot be inverted")
}

func TestToRowEchelon_ShouldSucceed(t *testing.T) {
	matrix, _ := NewFromData([][]float64{
		{1, 2, -1, -4},
		{2, 3, -1, -11},
		{-2, 0, -3, 22},
	})

	expectedREF, _ := NewFromData([][]float64{
		{1, 2, -1, -4},
		{0, -1, 1, -3},
		{0, 0, -1, 2},
	})

	result := matrix.ToRowEchelon()

	if !result.EqualsApprox(expectedREF, 1e-9) {
		t.Errorf("expected REF %v, got %v", expectedREF, result)
	}
}

func TestToRowEchelon_ZeroMatrix(t *testing.T) {
	matrix, _ := NewFromData([][]float64{
		{0, 0},
		{0, 0},
	})
	expected, _ := NewFromData([][]float64{
		{0, 0},
		{0, 0},
	})

	ref := matrix.ToRowEchelon()
	if !ref.EqualsApprox(expected, 1e-9) {
		t.Errorf("expected zero matrix, got %v", ref)
	}
}

func TestRank_ShouldBeCorrect(t *testing.T) {
	matrix, _ := NewFromData([][]float64{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	})

	// This matrix has rank 2
	rank := matrix.Rank()
	if rank != 2 {
		t.Errorf("expected rank 2, got %d", rank)
	}
}

func TestRank_ZeroMatrix(t *testing.T) {
	matrix, _ := NewFromData([][]float64{
		{0, 0},
		{0, 0},
	})

	rank := matrix.Rank()
	if rank != 0 {
		t.Errorf("expected rank 0, got %d", rank)
	}
}

func TestIsFullRank_ShouldReturnTrue(t *testing.T) {
	matrix, _ := NewFromData([][]float64{
		{1, 2},
		{3, 4},
	})

	if !matrix.IsFullRank() {
		t.Errorf("expected matrix to be full rank")
	}
}

func TestIsFullRank_ShouldReturnFalse(t *testing.T) {
	matrix, _ := NewFromData([][]float64{
		{1, 2},
		{2, 4},
	})

	if matrix.IsFullRank() {
		t.Errorf("expected matrix NOT to be full rank")
	}
}

func TestRank_NonSquareMatrix(t *testing.T) {
	matrix, _ := NewFromData([][]float64{
		{1, 2, 3},
		{4, 5, 6},
	})

	rank := matrix.Rank()
	if rank != 2 {
		t.Errorf("expected rank 2 for full row-rank 2x3 matrix, got %d", rank)
	}
}
