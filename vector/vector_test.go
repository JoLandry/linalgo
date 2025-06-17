package vector

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	var newVector *Vector = New(2)

	assert.Equal(t, newVector.size, 2)
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
