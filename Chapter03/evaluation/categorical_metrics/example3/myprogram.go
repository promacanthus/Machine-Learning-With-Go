package main

import (
	"fmt"

	"gonum.org/v1/gonum/integrate"
	"gonum.org/v1/gonum/stat"
)

func main() {

	// Define our scores and classes.
	scores := []float64{0.1, 0.35, 0.4, 0.8}
	classes := []bool{true, false, true, false}

	// Calculate the true positive rates (recalls) and
	// false positive rates.
	tpr, fpr, _ := stat.ROC(nil, scores, classes, nil)

	// Compute the Area Under Curve.
	auc := integrate.Trapezoidal(tpr, fpr)

	// Output the results to standard out.
	fmt.Printf("true  positive rate: %v\n", tpr)
	fmt.Printf("false positive rate: %v\n", fpr)
	fmt.Printf("auc: %v\n", auc)
}
