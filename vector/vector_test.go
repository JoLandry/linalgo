package vector

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	var newVector *Vector = New(2)

	assert.Equal(t, 2, newVector.dim)
	assert.Equal(t, []float64{0.0, 0.0}, newVector.data)
}

func TestGetSize(t *testing.T) {
	size := 2
	var newVector *Vector = New(size)

	assert.Equal(t, size, newVector.GetSize())
}

func TestGetData(t *testing.T) {
	data := []float64{2.0, 3.0}
	var newVector *Vector = NewFromData(data)

	assert.Equal(t, data, newVector.GetData())
}

func TestSetElementAt(t *testing.T) {
	var newVector *Vector = New(3)
	idx := 1
	elt := 4.0
	newVector.SetElementAt(idx, elt)

	assert.Equal(t, elt, newVector.GetElementAt(idx))
}

func TestEquals_ShouldReturnTrue(t *testing.T) {
	var newVector1 *Vector = New(3)
	var newVector2 *Vector = New(3)
	sameElt := 3.0

	for i := 0; i < newVector1.dim; i++ {
		newVector1.SetElementAt(i, sameElt)
		newVector2.SetElementAt(i, sameElt)
	}

	assert.Equal(t, true, newVector1.Equals(newVector2))
}

func TestEquals_ShouldReturnFalse_DifferentDimensions(t *testing.T) {
	var newVector1 *Vector = New(3)
	var newVector2 *Vector = New(2)

	assert.Equal(t, false, newVector1.Equals(newVector2))
}

func TestEquals_ShouldReturnFalse_NotSameData(t *testing.T) {
	var newVector1 *Vector = New(3)
	var newVector2 *Vector = New(3)

	sameElt := 3.0

	for i := 0; i < newVector1.dim; i++ {
		newVector1.SetElementAt(i, sameElt)
		newVector2.SetElementAt(i, sameElt)
	}

	newVector2.SetElementAt(2, 1.0)

	assert.Equal(t, false, newVector1.Equals(newVector2))
}

func TestAdd_ShouldBeSuccessful(t *testing.T) {
	v1 := New(3)
	v2 := New(3)
	sameElt := 3.0

	for i := 0; i < v1.dim; i++ {
		v1.SetElementAt(i, sameElt)
		v2.SetElementAt(i, sameElt)
	}

	res, err := v1.Add(v2)

	assert.NoError(t, err)
	assert.Equal(t, []float64{6.0, 6.0, 6.0}, res.data)
}

func TestAdd_ShoulFail_IfNotSameDimensions(t *testing.T) {
	v1 := New(3)
	v2 := New(2)

	_, err := v1.Add(v2)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "cannot add vectors of different dimensions")
}

func TestSub_ShouldBeSuccessful(t *testing.T) {
	v1 := New(3)
	v2 := New(3)
	sameElt := 3.0

	for i := 0; i < v1.dim; i++ {
		v1.SetElementAt(i, sameElt)
		v2.SetElementAt(i, sameElt)
	}

	res, err := v1.Sub(v2)

	assert.NoError(t, err)
	assert.Equal(t, []float64{0.0, 0.0, 0.0}, res.data)
}

func TestSub_ShoulFail_IfNotSameDimensions(t *testing.T) {
	v1 := New(3)
	v2 := New(2)

	_, err := v1.Sub(v2)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "cannot sub vectors of different dimensions")
}

func TestMul_ShouldBeSuccessful(t *testing.T) {
	v1 := New(3)
	v2 := New(3)
	sameElt := 3.0

	for i := 0; i < v1.dim; i++ {
		v1.SetElementAt(i, sameElt)
		v2.SetElementAt(i, sameElt)
	}

	res, err := v1.Mul(v2)

	assert.NoError(t, err)
	assert.Equal(t, []float64{9.0, 9.0, 9.0}, res.data)
}

func TestMul_ShoulFail_IfNotSameDimensions(t *testing.T) {
	v1 := New(3)
	v2 := New(2)

	_, err := v1.Mul(v2)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "cannot multiply vectors of different dimensions")
}

func TestMDiv_ShouldBeSuccessful(t *testing.T) {
	v1 := New(3)
	v2 := New(3)
	sameElt := 3.0

	for i := 0; i < v1.dim; i++ {
		v1.SetElementAt(i, sameElt)
		v2.SetElementAt(i, sameElt)
	}

	res, err := v1.Div(v2)

	assert.NoError(t, err)
	assert.Equal(t, []float64{1.0, 1.0, 1.0}, res.data)
}

func TestDiv_ShoulFail_IfNotSameDimensions(t *testing.T) {
	v1 := New(3)
	v2 := New(2)

	_, err := v1.Div(v2)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "cannot divide vectors of different dimensions")
}

func TestAddScalar(t *testing.T) {
	vector := New(3)
	newVector := vector.AddScalar(4.0)

	assert.Equal(t, []float64{4.0, 4.0, 4.0}, newVector.data)
}

func TestSubScalar(t *testing.T) {
	vector := New(3)
	newVector := vector.SubScalar(4.0)

	assert.Equal(t, []float64{-4.0, -4.0, -4.0}, newVector.data)
}

func TestMulScalar(t *testing.T) {
	vector := New(3)
	for i := 0; i < vector.dim; i++ {
		vector.SetElementAt(i, 2.0)
	}
	newVector := vector.MulScalar(4.0)

	assert.Equal(t, []float64{8.0, 8.0, 8.0}, newVector.data)
}

func TestDivScalar(t *testing.T) {
	vector := New(3)
	for i := 0; i < vector.dim; i++ {
		vector.SetElementAt(i, 2.0)
	}
	newVector := vector.DivScalar(2.0)

	assert.Equal(t, []float64{1.0, 1.0, 1.0}, newVector.data)
}

func TestNorm_ZeroVector(t *testing.T) {
	v := New(0)
	assert.Equal(t, 0.0, v.Norm())
}

func TestNorm_UnitVector(t *testing.T) {
	v := NewFromData([]float64{3.0, 4.0})
	// 5 = sqrt(3^2 + 4^2)
	expected := 5.0
	assert.InDelta(t, expected, v.Norm(), 1e-9)
}

func TestNormalize_ZeroVector(t *testing.T) {
	v := New(0)
	normalized := v.Normalize()

	assert.Equal(t, 0, normalized.GetSize())
	assert.Empty(t, normalized.GetData())
}

func TestNormalize_AlreadyUnitVector(t *testing.T) {
	v := NewFromData([]float64{1.0 / math.Sqrt(2), 1.0 / math.Sqrt(2)})
	normalized := v.Normalize()

	assert.InDelta(t, 1.0, normalized.Norm(), 1e-9)
}

func TestNormalize_NonUnitVector(t *testing.T) {
	v := NewFromData([]float64{3.0, 4.0})
	normalized := v.Normalize()

	assert.InDelta(t, 1.0, normalized.Norm(), 1e-9)

	// Check values are correctly scaled (cuz of float type)
	assert.InDelta(t, 0.6, normalized.GetElementAt(0), 1e-9)
	assert.InDelta(t, 0.8, normalized.GetElementAt(1), 1e-9)
}
