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

// Adds the vector from which this function is called from
// To the vector given in paramater
//
// Returns an error if they are not of the same dimension
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

// Subs the vector from which this function is called from
// To the vector given in paramater
//
// Returns an error if they are not of the same dimension
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

// Multiplies the vector from which this function is called from
// With the vector given in paramater
//
// Returns an error if they are not of the same dimension
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

// Divides the coordinates of the vector this function
// Is called from by the coordinates of the vector
// Given in parameter
//
// Returns an error if they are not of the same dimension
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

// Adds the scalar value given in parameter to all of the
// Coordinates of the vector
func (v *Vector) AddScalar(scalar float64) *Vector {
	resultVector := NewFromData(v.data)
	for i := range v.data {
		resultVector.data[i] += scalar
	}

	return resultVector
}

// Subs the scalar value given in parameter from all of the
// Coordinates of the vector
func (v *Vector) SubScalar(scalar float64) *Vector {
	resultVector := NewFromData(v.data)
	for i := range v.data {
		resultVector.data[i] -= scalar
	}

	return resultVector
}

// Multiplies the scalar value given in parameter with all of the
// Coordinates of the vector
func (v *Vector) MulScalar(scalar float64) *Vector {
	resultVector := NewFromData(v.data)
	for i := range v.data {
		resultVector.data[i] *= scalar
	}

	return resultVector
}

// Divides all the coordinates of the vector by the scalar
// Value given in parameter
func (v *Vector) DivScalar(scalar float64) *Vector {
	resultVector := NewFromData(v.data)
	for i := range v.data {
		resultVector.data[i] /= scalar
	}

	return resultVector
}

// Retrieves the norm (magnitude) of a vector
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

// Normalizes a vector (making it a unit vector)
func (v *Vector) Normalize() *Vector {
	norm := v.Norm()
	return v.DivScalar(norm)
}

// Returns true if the vector calling this function is the zero vector
func (v *Vector) IsZero() bool {
	for i := 0; i < v.dim; i++ {
		if v.GetElementAt(i) != 0.0 {
			return false
		}
	}

	return true
}

// Performs the projection of a vector to
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
