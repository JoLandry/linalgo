// Package vector provides a simple implementation of mathematical vectors
// with dynamic dimensions and common operations and features such as addition,
// subtraction, scalar operations, norm calculation, normalization,
// projection, and equality comparison.
//
// Vectors are represented as slices of float64 values, and all operations
// are implemented to work on vectors of the same dimension, returning
// errors when incompatible dimensions are encountered.
//
// The package supports:
//   - Creation of new vectors with a given size or from existing data
//   - Element-wise arithmetic operations (Add, Sub, Mul, Div)
//   - Scalar operations on each component (AddScalar, SubScalar, etc.)
//   - Computation of vector norm and normalization
//   - Dot product and projection onto another vector
//   - Basic utility checks like equality and zero-vector check
//
// This package can be very useful for numerical computations, simulations, physics,
// game engines and machine learning tasks that require vector math with customizable dimensions.
package vector

import (
	"fmt"
	"math"
)

// Epsilon for floating-point comparison.
const epsilon = 1e-9

// Computes and returns the Euclidian distance between two vectors.
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

// Determines whether two vectors are colinear.
//
// Two vectors are considered colinear if they "lie" on the same line,
// meaning one is a scalar multiple of the other.
//
// Special cases:
//   - If both vectors are zero vectors (dimension 0 or all elements are zero), they are considered colinear.
//   - If the vectors have different dimensions, they are not colinear.
//   - If either vector is a zero vector, the vectors are considered colinear.
//
// Returns true if the vectors are colinear, false otherwise.
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

// Computes and returns the dot product (scalar product) of two vectors.
//
// The dot product is calculated as the sum of the products of corresponding
// elements from the two vectors.
//
// Returns an error if the vectors have different dimensions.
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

// Performs and returns linear interpolation between two vectors using a given factor t.
//
// The interpolation is calculated as:
//
//	result = v1 * (1 - t) + v2 * t
//
// Where t is typically in the range [0,1], but values outside this range
// will extrapolate accordingly.
//
// Returns a new interpolated vector or an error if the input vectors have different dimensions.
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

// Computes and returns the 2D cross product (aka the perp dot product) of two 2-dimensional vectors.
//
// The result is a scalar equal to:
//
//	x1*y2 - y1*x2
//
// Returns an error if the vectors do not have exactly 2 dimensions or if their dimensions differ.
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

// Returns true if two float64 values are approximately equal,
// within a specified epsilon tolerance (constant variable at the start of this source file).
//
// This is used for comparing floating-point values while accounting for rounding errors.
func almostEqual(a, b, epsilon float64) bool {
	return math.Abs(a-b) <= epsilon
}

// Returns true if two vectors are approximately equal,
// comparing each corresponding element within a specified epsilon tolerance
// (constant variable at the start of this source file).
//
// Returns false if the vectors differ in dimension or any element differs
// beyond the allowed epsilon.
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

// Determines whether two vectors are orthogonal (perpendicular).
//
// Two vectors are orthogonal if their dot product is zero.
//
// Returns true if the vectors are orthogonal, false otherwise.
// Returns false if the dot product cannot be computed due to a dimension mismatch.
func AreOrthogonal(v1 *Vector, v2 *Vector) bool {
	result, err := DotProduct(v1, v2)
	if err != nil {
		return false
	}
	return result == 0.0
}
