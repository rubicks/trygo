// github.com/rubicks/trygo/dumbmath/sqrt.go

// Package dumbmath provides dumb implementations of common math functions
package dumbmath

import (
	"math"
)

func square(n float64) float64 {
	return n * n
}

func good_enough(guess, x float64) bool {
	return 0.001 > math.Abs(square(guess)-x)
}

func sum(arr []float64) (ret float64) {
	for _, e := range arr {
		ret += e
	}
	return
}

func average(arr []float64) float64 {
	return sum(arr) / float64(len(arr))
}

func improve(guess, x float64) float64 {
	return average([]float64{guess, x / guess})
}

// Sqrt returns the square root of its argument, calculated via Newton's method
// to an accuracy of 0.001
func Sqrt(x float64) float64 {
	guess := 1.0
	for ; !good_enough(guess, x); guess = improve(guess, x) {
	}
	return guess
}
