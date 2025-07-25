package matrix

import (
	"strings"
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

	assert.Equal(t, false, newMatrix.IsSquare())
	assert.Equal(t, true, newSquareMatrix.IsSquare())
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

func TestIsScalar_ShouldReturnTrue(t *testing.T) {
	matrix, err := NewFromData([][]float64{
		{5, 0, 0},
		{0, 5, 0},
		{0, 0, 5},
	})
	if err != nil {
		t.Fatalf("Failed to create matrix: %v", err)
	}

	if !matrix.IsScalar(5) {
		t.Errorf("Expected IsScalar(5) to return true")
	}
}

func TestIsScalar_ShouldReturnFalse_DiagonalNotSame(t *testing.T) {
	matrix, err := NewFromData([][]float64{
		{5, 0, 0},
		{0, 4, 0},
		{0, 0, 5},
	})
	if err != nil {
		t.Fatalf("Failed to create matrix: %v", err)
	}

	if matrix.IsScalar(5) {
		t.Errorf("Expected IsScalar(5) to return false (non-uniform diagonal)")
	}
}

func TestIsScalar_ShouldReturnFalse_OffDiagonalNonZero(t *testing.T) {
	matrix, err := NewFromData([][]float64{
		{5, 1, 0},
		{0, 5, 0},
		{0, 0, 5},
	})
	if err != nil {
		t.Fatalf("Failed to create matrix: %v", err)
	}

	if matrix.IsScalar(5) {
		t.Errorf("Expected IsScalar(5) to return false (non-zero off-diagonal)")
	}
}

func TestIsScalar_EmptyMatrix(t *testing.T) {
	matrix := New(0, 0)

	if !matrix.IsScalar(42) {
		t.Errorf("Expected IsScalar on empty matrix to return true (by convention)")
	}
}

func TestIsScalar_NonSquareMatrix(t *testing.T) {
	matrix, err := NewFromData([][]float64{
		{5, 0},
		{0, 5},
		{0, 0},
	})
	if err != nil {
		t.Fatalf("Failed to create matrix: %v", err)
	}

	if matrix.IsScalar(5) {
		t.Errorf("Expected IsScalar on non-square matrix to return false")
	}
}

func TestIsScalar_ZeroMatrixWithScalarZero(t *testing.T) {
	matrix, err := NewFromData([][]float64{
		{0, 0},
		{0, 0},
	})
	if err != nil {
		t.Fatalf("Failed to create matrix: %v", err)
	}

	if !matrix.IsScalar(0) {
		t.Errorf("Expected IsScalar(0) to return true on zero matrix")
	}
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

func TestPow_ZeroPower_ShouldReturnIdentity(t *testing.T) {
	m, _ := NewFromData([][]float64{
		{2, 3},
		{1, 4},
	})

	result, err := m.Pow(0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := NewIdentity(2)

	if !result.EqualsApprox(expected, 1e-6) {
		t.Errorf("expected identity matrix, got %v", result)
	}
}

func TestPow_OnePower_ShouldReturnSelf(t *testing.T) {
	m, _ := NewFromData([][]float64{
		{2, 3},
		{1, 4},
	})

	result, err := m.Pow(1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !result.EqualsApprox(m, 1e-6) {
		t.Errorf("expected original matrix, got %v", result)
	}
}

func TestPow_PositivePower_ShouldSucceed(t *testing.T) {
	m, _ := NewFromData([][]float64{
		{1, 1},
		{1, 0},
	})

	result, err := m.Pow(3)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// A^3
	expected, _ := NewFromData([][]float64{
		{3, 2},
		{2, 1},
	})

	if !result.EqualsApprox(expected, 1e-6) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestPow_NegativePower_ShouldSucceed(t *testing.T) {
	m, _ := NewFromData([][]float64{
		{4, 7},
		{2, 6},
	})

	result, err := m.Pow(-1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	inverse, _ := m.Invert()
	if !result.EqualsApprox(inverse, 1e-6) {
		t.Errorf("expected inverse matrix, got %v", result)
	}
}

func TestPow_NegativePowerGreaterThanOne(t *testing.T) {
	m, _ := NewFromData([][]float64{
		{2, 0},
		{0, 0.5},
	})

	result, err := m.Pow(-2)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected, _ := NewFromData([][]float64{
		{0.25, 0},
		{0, 4},
	})

	if !result.EqualsApprox(expected, 1e-6) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestPow_ShouldFail_NonSquareMatrix(t *testing.T) {
	m, _ := NewFromData([][]float64{
		{1, 2, 3},
		{4, 5, 6},
	})

	_, err := m.Pow(2)
	if err == nil {
		t.Errorf("expected error for non-square matrix, got none")
	}
}

func TestPow_ShouldFail_NonInvertibleMatrix(t *testing.T) {
	// Singular matrix
	m, _ := NewFromData([][]float64{
		{2, 4},
		{1, 2},
	})

	_, err := m.Pow(-3)
	if err == nil {
		t.Errorf("expected error for non-invertible matrix, got none")
	}
}

func TestMatrix_Transpose_Rectangular(t *testing.T) {
	m, err := NewFromData([][]float64{
		{1, 2, 3},
		{4, 5, 6},
	})
	assert.NoError(t, err)

	expected, errExpected := NewFromData([][]float64{
		{1, 4},
		{2, 5},
		{3, 6},
	})
	assert.NoError(t, errExpected)

	result := m.Transpose()

	assert.Equal(t, expected.data, result.data)
	assert.Equal(t, expected.nbRows, result.nbRows)
	assert.Equal(t, expected.nbCols, result.nbCols)
}

func TestMatrix_Transpose_Square(t *testing.T) {
	m, err := NewFromData([][]float64{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	})
	assert.NoError(t, err)

	expected, errExpected := NewFromData([][]float64{
		{1, 4, 7},
		{2, 5, 8},
		{3, 6, 9},
	})
	assert.NoError(t, errExpected)

	result := m.Transpose()

	assert.Equal(t, expected.data, result.data)
}

func TestMatrix_Transpose_RowVector(t *testing.T) {
	m, err := NewFromData([][]float64{
		{1, 2, 3, 4},
	})
	assert.NoError(t, err)

	expected, errExpected := NewFromData([][]float64{
		{1},
		{2},
		{3},
		{4},
	})
	assert.NoError(t, errExpected)

	result := m.Transpose()

	assert.Equal(t, expected.data, result.data)
}

func TestMatrix_Transpose_EmptyMatrix(t *testing.T) {
	m, err := NewFromData([][]float64{})
	assert.NoError(t, err)

	result := m.Transpose()

	assert.Equal(t, 0, result.nbRows)
	assert.Equal(t, 0, result.nbCols)
	assert.Empty(t, result.data)
}

func containsAll(str string, substrs []string) bool {
	for _, s := range substrs {
		if !strings.Contains(str, s) {
			return false
		}
	}
	return true
}

func TestMatrixString_Empty(t *testing.T) {
	m := New(0, 0)
	expected := "[]"

	if m.String() != expected {
		t.Errorf("expected %q, got %q", expected, m.String())
	}
}

func TestMatrixString_FormattedOutput(t *testing.T) {
	data := [][]float64{
		{1.0, 2.5, 3.1234},
		{4.56789, 5.0, 6.01},
	}
	m, _ := NewFromData(data)

	output := m.String()

	if !containsAll(output, []string{"1.0000", "2.5000", "3.1234", "4.5679", "6.0100"}) {
		t.Errorf("formatted output missing expected values: %q", output)
	}
	if output[0] != '[' {
		t.Errorf("expected output to start with '[', got %q", output[0])
	}
}

func TestNewFromFlat_Correct(t *testing.T) {
	values := []float64{1, 2, 3, 4, 5, 6}
	m := NewFromFlat(2, 3, values)

	expected := [][]float64{
		{1, 2, 3},
		{4, 5, 6},
	}

	for i := range expected {
		for j := range expected[i] {
			if m.data[i][j] != expected[i][j] {
				t.Errorf("expected m[%d][%d] = %f, got %f", i, j, expected[i][j], m.data[i][j])
			}
		}
	}
}

func TestNewFromFlat_PanicOnMismatch(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected panic due to mismatched dimensions, but did not panic")
		}
	}()

	// Should panic
	NewFromFlat(2, 3, []float64{1, 2, 3})
}

func TestNewDiagonal_Correct(t *testing.T) {
	diag := []float64{1.1, 2.2, 3.3}
	m := NewDiagonal(diag)

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == j && m.data[i][j] != diag[i] {
				t.Errorf("expected diagonal element m[%d][%d] = %f, got %f", i, j, diag[i], m.data[i][j])
			} else if i != j && m.data[i][j] != 0.0 {
				t.Errorf("expected off-diagonal element m[%d][%d] = 0.0, got %f", i, j, m.data[i][j])
			}
		}
	}
}
