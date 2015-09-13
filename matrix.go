package main

import (
	"fmt"
)

type Matrix struct {
	matrix [][]Fraction
}

// Matrix constructor to initialize a matrix of slices (vectors)
func newMatrix(rows, cols int) Matrix {
	var m Matrix
	m.matrix = make([][]Fraction, rows)

	for i, _ := range m.matrix {
		m.matrix[i] = make([]Fraction, cols)
	}

	return m
}

// A nice print out method for matricies
func (m Matrix) String() string {
	// Initialize the string to build and return
	s := ""

	for _, row := range m.matrix {
		for _, value := range row {
			// Make sure the fractions are simplified
			value.simplify()
			if value.denominator == 1 {
				// This is the same as a plain ol' integer
				s += fmt.Sprintf("%3d ", value.numerator)
			} else if value.denominator == 0 {
				// Just output a 0
				s += fmt.Sprintf("%3d ", 0)
			} else {
				// Let the Fraction stringify itself
				s += value.String() + " "
			}
		}
		s += "\n"
	}

	// Return a slice of the string so as to chop off the last extra newline
	return s[:len(s)-1]
}
