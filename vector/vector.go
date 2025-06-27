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

// Vector type with dynamic size (dimension)
type Vector struct {
	data []float64
	dim  int
}

// Get the data of the vector
func (v *Vector) GetData() []float64 {
	return v.data
}

// Get the dimension of the vector
func (v *Vector) GetSize() int {
	return v.dim
}

// Get the element at index idx
func (v *Vector) GetElementAt(idx int) float64 {
	return v.data[idx]
}

// Set the element called elt at index idx
func (v *Vector) SetElementAt(idx int, elt float64) {
	v.data[idx] = elt
}

// Create a new vector of a given dimension
func New(vSize int) *Vector {
	return &Vector{data: make([]float64, vSize), dim: vSize}
}

// Create a new vector from a given slice (deep copy)
func NewFromData(vData []float64) *Vector {
	copiedVector := make([]float64, len(vData))
	copy(copiedVector, vData)

	return &Vector{
		data: copiedVector,
		dim:  len(copiedVector),
	}
}

// Tells if 2 vectors are "equal"
//
// Considered equal if they're of the same dimension
// And contain the same elements in the same order
func (v *Vector) Equals(v2 *Vector) bool {
	if v.dim != v2.dim {
		return false
	}
	for i := 0; i < v.dim; i++ {
		if v.data[i] != v2.data[i] {
			return false
		}
	}
	return true
}

// Returns a new vector resulting from the element-wise addition
// of the calling vector and the input vector v2.
//
// Both vectors must have the same dimension. If their dimensions differ,
// an error is returned.
func (v *Vector) Add(v2 *Vector) (*Vector, error) {
	if v.dim != v2.dim {
		return nil, fmt.Errorf("cannot add vectors of different dimensions: %d vs %d", v.dim, v2.dim)
	}

	result := make([]float64, v.dim)
	for i := range v.data {
		result[i] = v.data[i] + v2.data[i]
	}

	return NewFromData(result), nil
}

// Returns a new vector resulting from the element-wise subtraction
// of the input vector v2 from the calling vector.
//
// Both vectors must have the same dimension. If their dimensions differ,
// an error is returned.
func (v *Vector) Sub(v2 *Vector) (*Vector, error) {
	if v.dim != v2.dim {
		return nil, fmt.Errorf("cannot sub vectors of different dimensions: %d vs %d", v.dim, v2.dim)
	}

	result := make([]float64, v.dim)
	for i := range v.data {
		result[i] = v.data[i] - v2.data[i]
	}

	return NewFromData(result), nil
}

// Returns a new vector resulting from the element-wise
// multiplication of the calling vector and the input vector v2.
//
// Both vectors must have the same dimension. If their dimensions differ,
// an error is returned.
func (v *Vector) Mul(v2 *Vector) (*Vector, error) {
	if v.dim != v2.dim {
		return nil, fmt.Errorf("cannot multiply vectors of different dimensions: %d vs %d", v.dim, v2.dim)
	}

	result := make([]float64, v.dim)
	for i := range v.data {
		result[i] = v.data[i] * v2.data[i]
	}

	return NewFromData(result), nil
}

// Returns a new vector resulting from the element-wise division
// of the calling vector by the input vector v2.
//
// Both vectors must have the same dimension. If their dimensions differ,
// an error is returned.
func (v *Vector) Div(v2 *Vector) (*Vector, error) {
	if v.dim != v2.dim {
		return nil, fmt.Errorf("cannot divide vectors of different dimensions: %d vs %d", v.dim, v2.dim)
	}

	result := make([]float64, v.dim)
	for i := range v.data {
		result[i] = v.data[i] / v2.data[i]
	}

	return NewFromData(result), nil
}

// Returns a new vector resulting from the element-wise addition
// of the calling vector and the given scalar value.
func (v *Vector) AddScalar(scalar float64) *Vector {
	resultVector := NewFromData(v.data)
	for i := range v.data {
		resultVector.data[i] += scalar
	}

	return resultVector
}

// Returns a new vector resulting from the element-wise subtraction
// of the given scalar value from the calling vector .
func (v *Vector) SubScalar(scalar float64) *Vector {
	resultVector := NewFromData(v.data)
	for i := range v.data {
		resultVector.data[i] -= scalar
	}

	return resultVector
}

// Returns a new vector resulting from the element-wise
// multiplication of the calling vector and the given scalar value.
func (v *Vector) MulScalar(scalar float64) *Vector {
	resultVector := NewFromData(v.data)
	for i := range v.data {
		resultVector.data[i] *= scalar
	}

	return resultVector
}

// Returns a new vector resulting from the element-wise division
// of the calling vector by the given scalar value.
func (v *Vector) DivScalar(scalar float64) *Vector {
	resultVector := NewFromData(v.data)
	for i := range v.data {
		resultVector.data[i] /= scalar
	}

	return resultVector
}

// Computes and returns the norm (magnitude) of the calling vector
func (v *Vector) Norm() float64 {
	if v.dim == 0 {
		return 0.0
	}

	sumSquares := 0.0
	for i := 0; i < v.dim; i++ {
		sumSquares += v.data[i] * v.data[i]
	}

	return math.Sqrt(sumSquares)
}

// Returns a new vector resulting from normalization of the calling vector
// (making it a unit vector).
func (v *Vector) Normalize() *Vector {
	norm := v.Norm()
	return v.DivScalar(norm)
}

// Returns true if the calling vector is the zero vector
func (v *Vector) IsZero() bool {
	for i := 0; i < v.dim; i++ {
		if v.GetElementAt(i) != 0.0 {
			return false
		}
	}

	return true
}

// Returns the projection of the calling vector onto the target vector called `onto`.
//
// The projection is computed using the formula:
//
//	proj_v_onto_u = (v • u / ||u||²) * u
//
// This method requires that both vectors have the same dimension.
// If the dimensions do not match, or if the target vector is the zero vector
// (therefore has no defined direction), an error is returned.
func (v *Vector) ProjectOnto(onto *Vector) (*Vector, error) {
	if v.dim != onto.dim {
		return nil, fmt.Errorf("cannot project vectors of different dimensions: %d vs %d", v.dim, onto.dim)
	}

	normSquared := onto.Norm()
	if normSquared == 0.0 {
		return nil, fmt.Errorf("cannot project onto the zero vector")
	}
	normSquared *= normSquared

	dot, _ := DotProduct(v, onto)
	scalar := dot / normSquared

	return onto.MulScalar(scalar), nil
}
