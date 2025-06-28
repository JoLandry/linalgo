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

func TestDotProduct_ShouldSucceed(t *testing.T) {
	v1 := NewFromData([]float64{1.0, 2.0})
	v2 := NewFromData([]float64{1.0, 2.0})

	result, err := DotProduct(v1, v2)

	assert.NoError(t, err)
	assert.Equal(t, 5.0, result)
}

func TestDotProduct_ShouldFail_DifferentDimensions(t *testing.T) {
	v1 := NewFromData([]float64{1.0, 2.0})
	v2 := NewFromData([]float64{2.0})

	_, err := DotProduct(v1, v2)

	assert.Contains(t, err.Error(), "cannot compute dot products for vectors of different dimensions")
}

func TestLerp_ShouldFail_DifferentDimensions(t *testing.T) {
	v1 := NewFromData([]float64{1.0, 2.0})
	v2 := NewFromData([]float64{2.0})

	_, err := Lerp(v1, v2, 2.0)

	assert.Contains(t, err.Error(), "annot compute linear interpolation for vectors of different dimensions")
}

func TestLerp_ShouldSucceed(t *testing.T) {
	v1 := NewFromData([]float64{1.0, 2.0})
	v2 := NewFromData([]float64{2.0, 2.0})

	tValue := 0.0
	result, err := Lerp(v1, v2, tValue)
	expected := NewFromData([]float64{1.0, 2.0})

	assert.NoError(t, err)
	assert.True(t, result.Equals(expected), "expected %v, got %v", expected.data, result.data)
}

func TestCross2D_ShouldFail_DifferentDimensions(t *testing.T) {
	v1 := NewFromData([]float64{1.0, 2.0})
	v2 := NewFromData([]float64{2.0})

	_, err := Cross2D(v1, v2)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "cannot perform cross2D operation for vectors of different dimensions")
}

func TestCross2D_ShouldFail_Non2DVector(t *testing.T) {
	v1 := NewFromData([]float64{1.0, 2.0, 3.0})
	v2 := NewFromData([]float64{4.0, 5.0, 6.0})

	_, err := Cross2D(v1, v2)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "cross2D operation only supports 2D vectors.")
}

func TestCross2D_ShouldSucceed(t *testing.T) {
	v1 := NewFromData([]float64{1.0, 2.0})
	v2 := NewFromData([]float64{3.0, 4.0})

	result, err := Cross2D(v1, v2)
	expected := -2.0

	assert.NoError(t, err)
	assert.InEpsilon(t, expected, result, 1e-9, "expected %v, got %v", expected, result)
}

func TestAlmostEqual(t *testing.T) {
	epsilon := 1e-9

	// Exactly equal values
	assert.True(t, almostEqual(1.0, 1.0, epsilon))
	// Values within epsilon
	assert.True(t, almostEqual(1.0, 1.0+1e-10, epsilon))
	// Values outside epsilon
	assert.False(t, almostEqual(1.0, 1.0+1e-7, epsilon))
}

func TestVectorsAlmostEqual(t *testing.T) {
	epsilon := 1e-9

	v1 := NewFromData([]float64{1.0, 2.0, 3.0})
	v2 := NewFromData([]float64{1.0, 2.0, 3.0})
	v3 := NewFromData([]float64{1.0, 2.0, 3.0000000001})
	v4 := NewFromData([]float64{1.0, 2.0, 3.000001})
	v5 := NewFromData([]float64{1.0, 2.0}) // different size

	// Same vectors exactly
	assert.True(t, vectorsAlmostEqual(v1, v2, epsilon))

	// Slight difference within epsilon
	assert.True(t, vectorsAlmostEqual(v1, v3, epsilon))

	// Difference outside epsilon
	assert.False(t, vectorsAlmostEqual(v1, v4, epsilon))

	// Different dimensions
	assert.False(t, vectorsAlmostEqual(v1, v5, epsilon))
}

func TestAreOrthogonal_ShouldBeTrue(t *testing.T) {
	v1 := NewFromData([]float64{1.0, 0.0})
	v2 := NewFromData([]float64{0.0, 1.0})

	assert.Equal(t, true, AreOrthogonal(v1, v2))
}

func TestAreOrthogonal_ShouldBeFalse(t *testing.T) {
	v1 := NewFromData([]float64{1.0, 5.0})
	v2 := NewFromData([]float64{2.0, 1.0})

	assert.Equal(t, false, AreOrthogonal(v1, v2))
}

func TestAreOrthogonal_ShouldFail_DifferentDimensions(t *testing.T) {
	v1 := NewFromData([]float64{1.0, 0.0})
	v2 := NewFromData([]float64{5.0})

	assert.Equal(t, false, AreOrthogonal(v1, v2))
}
