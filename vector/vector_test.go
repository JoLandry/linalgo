package vector

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	var newVector *Vector = New(2)

	assert.Equal(t, newVector.dim, 2)
	assert.Equal(t, []float64{0.0, 0.0}, newVector.data)
}

func TestGetSize(t *testing.T) {
	size := 2
	var newVector *Vector = New(size)

	assert.Equal(t, newVector.GetSize(), size)
}

func TestGetData(t *testing.T) {
	data := []float64{2.0, 3.0}
	var newVector *Vector = NewFromData(data)

	assert.Equal(t, newVector.GetData(), data)
}

func TestSetElementAt(t *testing.T) {
	var newVector *Vector = New(3)
	idx := 1
	elt := 4.0
	newVector.SetElementAt(idx, elt)

	assert.Equal(t, newVector.GetElementAt(idx), elt)
}

func TestEquals_ShouldReturnTrue(t *testing.T) {
	var newVector1 *Vector = New(3)
	var newVector2 *Vector = New(3)
	sameElt := 3.0

	for i := 0; i < newVector1.dim; i++ {
		newVector1.SetElementAt(i, sameElt)
		newVector2.SetElementAt(i, sameElt)
	}

	assert.Equal(t, newVector1.Equals(newVector2), true)
}

func TestEquals_ShouldReturnFalse_DifferentDimensions(t *testing.T) {
	var newVector1 *Vector = New(3)
	var newVector2 *Vector = New(2)

	assert.Equal(t, newVector1.Equals(newVector2), false)
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

	assert.Equal(t, newVector1.Equals(newVector2), false)
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
	assert.Equal(t, res.data, []float64{6.0, 6.0, 6.0})
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
	assert.Equal(t, res.data, []float64{0.0, 0.0, 0.0})
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
	assert.Equal(t, res.data, []float64{9.0, 9.0, 9.0})
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
	assert.Equal(t, res.data, []float64{1.0, 1.0, 1.0})
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

	assert.Equal(t, newVector.data, []float64{4.0, 4.0, 4.0})
}

func TestSubScalar(t *testing.T) {
	vector := New(3)
	newVector := vector.SubScalar(4.0)

	assert.Equal(t, newVector.data, []float64{-4.0, -4.0, -4.0})
}

func TestMulScalar(t *testing.T) {
	vector := New(3)
	for i := 0; i < vector.dim; i++ {
		vector.SetElementAt(i, 2.0)
	}
	newVector := vector.MulScalar(4.0)

	assert.Equal(t, newVector.data, []float64{8.0, 8.0, 8.0})
}

func TestDivScalar(t *testing.T) {
	vector := New(3)
	for i := 0; i < vector.dim; i++ {
		vector.SetElementAt(i, 2.0)
	}
	newVector := vector.DivScalar(2.0)

	assert.Equal(t, newVector.data, []float64{1.0, 1.0, 1.0})
}
