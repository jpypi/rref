package main

import (
	"fmt"
)

func tests() {
	m := newMatrix(3, 3)

	fmt.Println("Test matrix 1:")

	m.matrix[0][0] = Fraction{numerator: 2, denominator: 1}
	m.matrix[0][1] = Fraction{numerator: 3, denominator: 1}
	m.matrix[0][2] = Fraction{numerator: 3, denominator: 1}

	m.matrix[1][0] = Fraction{numerator: 3, denominator: 1}
	m.matrix[1][1] = Fraction{numerator: 0, denominator: 1}
	m.matrix[1][2] = Fraction{numerator: 5, denominator: 1}

	m.matrix[2][0] = Fraction{numerator: 3, denominator: 1}
	m.matrix[2][1] = Fraction{numerator: 5, denominator: 1}
	m.matrix[2][2] = Fraction{numerator: 5, denominator: 1}

	fmt.Println(m)
	fmt.Println("------------")
	fmt.Println(m.rref())

	// ----------------------------------------------------

	fmt.Println("\nTest matrix 2:")

	m.matrix[0][0] = Fraction{numerator: 1, denominator: 1}
	m.matrix[0][1] = Fraction{numerator: 2, denominator: 1}
	m.matrix[0][2] = Fraction{numerator: 3, denominator: 1}

	m.matrix[1][0] = Fraction{numerator: 4, denominator: 1}
	m.matrix[1][1] = Fraction{numerator: 5, denominator: 1}
	m.matrix[1][2] = Fraction{numerator: 6, denominator: 1}

	m.matrix[2][0] = Fraction{numerator: 7, denominator: 1}
	m.matrix[2][1] = Fraction{numerator: 8, denominator: 1}
	m.matrix[2][2] = Fraction{numerator: 9, denominator: 1}

	fmt.Println(m)
	fmt.Println("------------")
	fmt.Println(m.rref())

	fmt.Println("\nTest matrix 3:")

	m.matrix[0][0] = Fraction{numerator: 1, denominator: 1}
	m.matrix[0][1] = Fraction{numerator: 2, denominator: 1}
	m.matrix[0][2] = Fraction{numerator: 3, denominator: 1}

	m.matrix[1][0] = Fraction{numerator: 1, denominator: 1}
	m.matrix[1][1] = Fraction{numerator: 2, denominator: 1}
	m.matrix[1][2] = Fraction{numerator: 4, denominator: 1}

	m.matrix[2][0] = Fraction{numerator: 1, denominator: 1}
	m.matrix[2][1] = Fraction{numerator: 3, denominator: 1}
	m.matrix[2][2] = Fraction{numerator: 3, denominator: 1}

	fmt.Println(m)
	fmt.Println("------------")
	fmt.Println(m.rref())
}
