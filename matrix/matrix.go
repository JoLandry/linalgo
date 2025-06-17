package matrix

type Matrix struct {
	data   [][]float64
	nbRows int
	nbCols int
}

func (m *Matrix) GetData() [][]float64 {
	return m.data
}

func (m *Matrix) GetNbRows() int {
	return m.nbRows
}

func (m *Matrix) GetNbCols() int {
	return m.nbCols
}

func (m *Matrix) IsSquareMatrix() bool {
	return m.nbCols == m.nbRows
}

func (m *Matrix) GetElementAt(row int, col int) float64 {
	return m.data[row][col]
}

func (m *Matrix) SetElementAt(row int, col int, elt float64) {
	m.data[row][col] = elt
}

func New(nbRowsMat int, nbColsMat int) *Matrix {
	data := make([][]float64, nbRowsMat)
	for i := range data {
		data[i] = make([]float64, nbColsMat)
	}
	return &Matrix{
		data:   data,
		nbRows: nbRowsMat,
		nbCols: nbColsMat,
	}
}

func NewFromData(mData [][]float64) *Matrix {
	nbRows := len(mData)
	if nbRows == 0 {
		return &Matrix{
			data:   [][]float64{},
			nbRows: 0,
			nbCols: 0,
		}
	}

	nbCols := len(mData[0])

	// Make a deep copy
	copiedMatrix := make([][]float64, nbRows)
	for i := range mData {
		if len(mData[i]) != nbCols {
			panic("inconsistent number of columns in input data")
		}
		copiedMatrix[i] = make([]float64, nbCols)
		copy(copiedMatrix[i], mData[i])
	}

	return &Matrix{
		data:   copiedMatrix,
		nbRows: nbRows,
		nbCols: nbCols,
	}
}
