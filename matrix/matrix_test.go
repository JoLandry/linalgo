package matrix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	var newMatrix *Matrix = New(2, 2)

	assert.Equal(t, newMatrix.nbCols, 2)
	assert.Equal(t, newMatrix.nbRows, 2)
	assert.Equal(t, [][]float64{{0.0, 0.0}, {0.0, 0.0}}, newMatrix.data)
}

func TestGetData(t *testing.T) {
	data := [][]float64{{1.0, 2.0}, {3.0, 4.0}}
	var newMatrix *Matrix = NewFromData(data)

	assert.Equal(t, newMatrix.GetData(), data)
}

func TestGetNbRows(t *testing.T) {
	var newMatrix *Matrix = New(2, 2)

	assert.Equal(t, newMatrix.GetNbRows(), 2)
}

func TestGetNbCols(t *testing.T) {
	var newMatrix *Matrix = New(2, 2)

	assert.Equal(t, newMatrix.GetNbCols(), 2)
}

func TestIsSquareMatrix(t *testing.T) {
	var newMatrix *Matrix = New(2, 3)
	var newSquareMatrix *Matrix = New(2, 2)

	assert.Equal(t, newMatrix.IsSquareMatrix(), false)
	assert.Equal(t, newSquareMatrix.IsSquareMatrix(), true)
}

func TestSetElementAt(t *testing.T) {
	col := 2
	row := 0
	elt := 8.0

	var newMatrix *Matrix = New(4, 4)
	newMatrix.SetElementAt(row, col, elt)

	assert.Equal(t, newMatrix.GetElementAt(row, col), elt)
}

func TestNewFromData_ValidInput(t *testing.T) {
	input := [][]float64{{1.0, 2.0}, {3.0, 4.0}}

	mat := NewFromData(input)

	assert.Equal(t, 2, mat.nbRows)
	assert.Equal(t, 2, mat.nbCols)
	assert.Equal(t, [][]float64{{1.0, 2.0}, {3.0, 4.0}}, mat.data)
}

func TestNewFromData_EmptyInput(t *testing.T) {
	input := [][]float64{}

	mat := NewFromData(input)

	assert.Equal(t, 0, mat.nbRows)
	assert.Equal(t, 0, mat.nbCols)
	assert.Empty(t, mat.data)
}

func TestNewFromData_DeepCopy(t *testing.T) {
	input := [][]float64{{1.0, 2.0}, {3.0, 4.0}}

	mat := NewFromData(input)
	input[0][0] = 99.0

	// Value was not changed (still the old one)
	assert.Equal(t, 1.0, mat.GetElementAt(0, 0))
}

func TestNewFromData_InconsistentCols_Panics(t *testing.T) {
	input := [][]float64{{1.0, 2.0}, {3.0}}

	// Defer and then assess what's returned (or not)
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic due to inconsistent column lengths, but none occurred")
		}
	}()

	_ = NewFromData(input)
}
