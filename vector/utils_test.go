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

func TestAreColinear_ShouldReturnTrue_If_ZeroVector(t *testing.T) {
	v1 := NewFromData([]float64{0.0, 0.0})
	v2 := NewFromData([]float64{1.0, 1.0})

	assert.Equal(t, true, AreColinear(v1, v2))
}

func TestAreColinear_ShouldReturnFalse_DifferentDimensions(t *testing.T) {
	v1 := NewFromData([]float64{0.0, 0.0})
	v2 := NewFromData([]float64{1.0})

	assert.Equal(t, false, AreColinear(v1, v2))
}

func TestAreColinear_ShouldReturnTrue_If_ZeroSizeBoth(t *testing.T) {
	v1 := New(0)
	v2 := New(0)

	assert.Equal(t, true, AreColinear(v1, v2))
}

func TestAreColinear_ShouldReturnTrue(t *testing.T) {
	v1 := NewFromData([]float64{6.0, 8.0, 2.0})
	v2 := NewFromData([]float64{3.0, 4.0, 1.0})

	assert.Equal(t, true, AreColinear(v1, v2))
}

func TestAreColinear_ShouldReturnFalse_NotColinear(t *testing.T) {
	v1 := NewFromData([]float64{2.0, 3.0})
	v2 := NewFromData([]float64{1.0, 1.0})

	assert.False(t, AreColinear(v1, v2))
}

func TestAreColinear_ShouldReturnTrue_WithZeroCoordinates(t *testing.T) {
	v1 := NewFromData([]float64{0.0, 2.0, 4.0})
	v2 := NewFromData([]float64{0.0, 1.0, 2.0})

	assert.True(t, AreColinear(v1, v2))
}

func TestAreColinear_ShouldReturnFalse_When_B_IsZero_And_A_NotZero(t *testing.T) {
	v1 := NewFromData([]float64{1.0, 2.0})
	v2 := NewFromData([]float64{0.0, 2.0})

	assert.False(t, AreColinear(v1, v2))
}
