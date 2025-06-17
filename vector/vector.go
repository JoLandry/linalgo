package vector

// Vector type with dynamic size
type Vector struct {
	data []float64
	size int
}

// Get the data of the vector
func (v *Vector) GetData() []float64 {
	return v.data
}

// Get thesize of the vector
func (v *Vector) GetSize() int {
	return v.size
}

// Get the element at index idx
func (v *Vector) GetElementAt(idx int) float64 {
	return v.data[idx]
}

// Set the element called elt at index idx
func (v *Vector) SetElementAt(idx int, elt float64) {
	v.data[idx] = elt
}

// Create a new vector of a given size
func New(vSize int) *Vector {
	return &Vector{data: make([]float64, vSize), size: vSize}
}

// Create a new vector from a given slice (deep copy)
func NewFromData(vData []float64) *Vector {
	copiedVector := make([]float64, len(vData))
	copy(copiedVector, vData)

	return &Vector{
		data: copiedVector,
		size: len(copiedVector),
	}
}
