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
		a := v1.GetElementAt(i)
		b := v2.GetElementAt(i)

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
