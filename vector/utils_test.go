package vector

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDistance_Success(t *testing.T) {
	v1 := NewFromData([]float64{3.0, 4.0})
	v2 := NewFromData([]float64{0.0, 0.0})

	dist, err := Distance(v1, v2)
	assert.NoError(t, err)
	assert.InDelta(t, 5.0, dist, 1e-9)
}

func TestDistance_DifferentDimensions(t *testing.T) {
	v1 := NewFromData([]float64{1.0, 2.0})
	v2 := NewFromData([]float64{1.0})

	_, err := Distance(v1, v2)
	assert.Error(t, err)
}
