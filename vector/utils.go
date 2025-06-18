package vector

import (
	"fmt"
	"math"
)

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
