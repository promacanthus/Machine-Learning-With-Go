package main

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

func main() {

	// Create a new vector value.
	myvector := mat.NewVecDense(2, []float64{11.0, 5.2})

	fmt.Println(myvector)
}
