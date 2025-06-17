package vector

type Vector struct {
	data []float64
	size int
}

func (v *Vector) GetData() []float64 {
	return v.data
}

func (v *Vector) GetSize() int {
	return v.size
}

func (v *Vector) GetElementAt(idx int) float64 {
	return v.data[idx]
}

func (v *Vector) SetElementAt(idx int, elt float64) {
	v.data[idx] = elt
}

func New(vSize int) *Vector {
	return &Vector{data: make([]float64, vSize), size: vSize}
}

func NewFromData(vData []float64) *Vector {
	copiedVector := make([]float64, len(vData))
	copy(copiedVector, vData)

	return &Vector{
		data: copiedVector,
		size: len(copiedVector),
	}
}
