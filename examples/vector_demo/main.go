package main

import (
	"fmt"
	"log"

	"github.com/JoLandry/linalgo/vector"
)

// Small demo of the vector package

func main() {
	// Create two vectors from slices
	v1 := vector.NewFromData([]float64{1, 2, 3})
	v2 := vector.NewFromData([]float64{4, 5, 6})

	// Create two new empty vectors, with dimension indications
	v3 := vector.New(2)
	v4 := vector.New(3)
	v5 := vector.New(16)

	v1String := v1.String()
	fmt.Println(v1String)

	v2String := v2.String()
	fmt.Println(v2String)

	v3String := v3.String()
	fmt.Println(v3String)

	v4String := v4.String()
	fmt.Println(v4String)

	v5String := v5.String()
	fmt.Println(v5String)

	// Vector addition
	sum, err := v1.Add(v2)
	if err != nil {
		log.Fatal("Add failed:", err)
	}
	fmt.Println("v1 + v2:", sum)

	// Scalar multiplication
	scaled := v1.MulScalar(2)
	fmt.Println("v1 * 2:", scaled)

	// Norm and normalization
	norm := v1.Norm()
	fmt.Printf("||v1|| = %.4f\n", norm)

	unit := v1.Normalize()
	fmt.Println("Normalized v1:", unit)
	fmt.Printf("||unit|| = %.4f\n", unit.Norm())

	// Projection of v1 onto v2
	proj, err := v1.ProjectOnto(v2)
	if err != nil {
		log.Fatal("Projection failed:", err)
	}
	fmt.Println("Projection of v1 onto v2:", proj)
}
