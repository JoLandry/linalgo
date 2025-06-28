package main

import (
	"fmt"
	"log"

	"github.com/JoLandry/linalgo/matrix"
)

func main() {
	// Create a matrix from 2D slice
	m1, err := matrix.NewFromData([][]float64{
		{1, 2, 3},
		{4, 5, 6},
	})
	if err != nil {
		log.Fatal("Failed to create m1:", err)
	}

	// Create another matrix from flat data
	m2 := matrix.NewFromFlat(3, 2, []float64{
		7, 8,
		9, 10,
		11, 12,
	})

	fmt.Println("Matrix m1:")
	fmt.Println(m1.String())

	fmt.Println("Matrix m2:")
	fmt.Println(m2.String())

	// Multiply m1 and m2
	product, err := m1.Mul(m2)
	if err != nil {
		log.Fatal("Matrix multiplication failed:", err)
	}
	fmt.Println("Product m1 * m2:")
	fmt.Println(product.String())

	// Create and display identity matrix
	identity := matrix.NewIdentity(3)
	fmt.Println("Identity matrix (3x3):")
	fmt.Println(identity.String())
	fmt.Println("Is identity:", identity.IsIdentity())

	// Create and check zero matrix
	zero := matrix.New(2, 2)
	fmt.Println("Zero matrix (2x2):")
	fmt.Println(zero.String())
	fmt.Println("Is zero matrix:", zero.IsZero())

	// Matrix for determinant and rank
	square, _ := matrix.NewFromData([][]float64{
		{2, 3, 1},
		{4, 1, -3},
		{1, 2, 0},
	})
	fmt.Println("Square matrix for determinant and rank:")
	fmt.Println(square.String())

	det, err := square.Determinant()
	if err != nil {
		log.Fatal("Failed to compute determinant:", err)
	}
	fmt.Printf("Determinant: %.4f\n", det)

	rank := square.Rank()
	fmt.Printf("Rank: %d\n", rank)
}
