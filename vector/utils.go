package vector

import (
	"fmt"
	"math"
)

// Epsilon for floating-point comparison
const epsilon = 1e-9

// Computes and returns the Euclidian distance between two vectors
func Distance(v1 *Vector, v2 *Vector) (float64, error) {
	if v1.dim != v2.dim {
		return 0.0, fmt.Errorf("cannot compute distance for vectors of different dimensions: %d vs %d", v1.dim, v2.dim)
	}

	var sumSquares float64
	for i := 0; i < v1.dim; i++ {
		diff := v1.data[i] - v2.data[i]
		sumSquares += diff * diff
	}

	return math.Sqrt(sumSquares), nil
}

// Tells if two vectors are colinear
//
// Returns true if so, else otherwise
func AreColinear(v1 *Vector, v2 *Vector) bool {
	if v1.dim == 0 && v2.dim == 0 {
		// Considered colinear if norm is zero
		return true
	} else if v1.dim != v2.dim {
		// Not colinear if not of same dimension
		return false
	} else if v1.IsZero() || v2.IsZero() {
		// Considered colinear if at least one of the vectors is the zero vector
		return true
	}

	// Check that it's the same factor (scalar) for every coordinate
	var referenceFactor *float64 = nil
	for i := 0; i < v1.dim; i++ {
		a := v1.data[i]
		b := v2.data[i]

		if math.Abs(b) < epsilon {
			if math.Abs(a) >= epsilon {
				return false
			}
			// both are zero, still valid
			continue
		}

		factor := a / b
		if referenceFactor == nil {
			referenceFactor = &factor
		} else if math.Abs(*referenceFactor-factor) > epsilon {
			return false
		}
	}

	return true
}

// Computes and returns the dot product between two vectors
func DotProduct(v1 *Vector, v2 *Vector) (float64, error) {
	if v1.dim != v2.dim {
		return 0.0, fmt.Errorf("cannot compute dot products for vectors of different dimensions: %d vs %d", v1.dim, v2.dim)
	}
	sum := 0.0
	for i := 0; i < v1.dim; i++ {
		sum += v1.data[i] * v2.data[i]
	}

	return sum, nil
}

// Computes and returns the linear interpolation for two vectors
// And a fraction, called t
// Returns an error if vectors of different dimensions
func Lerp(v1 *Vector, v2 *Vector, t float64) (*Vector, error) {
	if v1.dim != v2.dim {
		return nil, fmt.Errorf("cannot compute linear interpolation for vectors of different dimensions: %d vs %d", v1.dim, v2.dim)
	}

	result := make([]float64, v1.dim)
	for i := 0; i < v1.dim; i++ {
		result[i] = v1.data[i]*(1-t) + v2.data[i]*t
	}

	return NewFromData(result), nil
}

// Cross2D computes the 2D cross product (aka the perp dot product)
// Returns a scalar representing the norm (magnitude) of the vector perpendicular to the plane formed by v1 and v2
func Cross2D(v1 *Vector, v2 *Vector) (float64, error) {
	if v1.dim != v2.dim {
		return 0.0, fmt.Errorf("cannot perform cross2D operation for vectors of different dimensions: %d vs %d", v1.dim, v2.dim)
	}
	if v1.dim != 2 {
		return 0.0, fmt.Errorf("cross2D operation only supports 2D vectors. Got dimension: %d", v1.dim)
	}

	// 2D cross product (scalar): x1*y2 - y1*x2
	return v1.data[0]*v2.data[1] - v1.data[1]*v2.data[0], nil
}

// Utils function for floating verification
func almostEqual(a, b, epsilon float64) bool {
	return math.Abs(a-b) <= epsilon
}

// Utils function for floating verification
func vectorsAlmostEqual(v1, v2 *Vector, epsilon float64) bool {
	if v1.dim != v2.dim {
		return false
	}
	for i := 0; i < v1.dim; i++ {
		if !almostEqual(v1.data[i], v2.data[i], epsilon) {
			return false
		}
	}
	return true
}

// Tells if two vectors are orthogonal
// Returns true if so, false otherwise
func AreOrthogonal(v1 *Vector, v2 *Vector) bool {
	result, err := DotProduct(v1, v2)
	if err != nil {
		return false
	}
	return result == 0.0
}
