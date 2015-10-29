package main

import (
	"fmt"
)

type Matrix struct {
	matrix [][]Fraction
	rows   int
	cols   int
}

// Matrix constructor to initialize a matrix of slices (vectors)
func newMatrix(rows, cols int) Matrix {
	var m Matrix
	m.matrix = make([][]Fraction, rows)
	m.rows = rows

	for i, _ := range m.matrix {
		m.matrix[i] = make([]Fraction, cols)
	}
	m.cols = cols

	return m
}

// Transform the matrix to rref form! :D
func (m Matrix) rref() Matrix {
	for r, row := range m.matrix {
		// Find the pivot point for the current row
		pivot_point := 0
		for pivot_point < m.cols && row[pivot_point].numerator == 0 {
			pivot_point += 1
		}

		// Make sure a pivot was actually found
		if pivot_point < m.cols {

			// Find the inverse (k*x = 1) [Basically, make first element 1.]
			inv := row[pivot_point].inv()

			// Multiply entire row times the inverse that was found
			for c := pivot_point; c < m.cols; c += 1 {
				m.matrix[r][c] = m.matrix[r][c].mul(inv)
			}

			// TODO: Work back upwards to clear numbers above the pivot
			/*
				for i := r - 1; i >= 0; i -= 1 {
					m.matrix[i][r] = NewIFrac(0)
				}
			*/

			for below_r := r + 1; below_r < m.rows; below_r += 1 {
				sub_val := m.matrix[below_r][pivot_point].mulInt(-1)
				for i := pivot_point; i < m.cols; i += 1 {
					m.matrix[below_r][i] = m.matrix[below_r][i].add(m.matrix[r][i].mul(sub_val))
				}
			}
		}
	}

	return m
}

func (m *Matrix) rowMul(row int, value Fraction) {
	for i, _ := range m.matrix[row] {
		m.matrix[row][i].imul(value)
	}
}

func (m Matrix) mulRow(row int, value Fraction) []Fraction {
	for i, _ := range m.matrix[row] {
		m.matrix[row][i].imul(value)
	}

	return m.matrix[row]
}

func (m *Matrix) rowAdd(src, dest int) {
	for i, value := range m.matrix[src] {
		m.matrix[dest][i].iadd(value)
	}
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
