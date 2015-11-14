package main

import (
	"bufio"
	"fmt"
	"os"
	//	"strconv"
	"strings"
)

func prompt(prmpt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prmpt)
	text, _ := reader.ReadString('\n')

	return strings.Trim(text, "\n")
}

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

func main() {
	tests()
	/*
		var dimensions []string
		var rows, cols int

		for len(dimensions) != 2 {
			var err1, err2 error

			dim_string := prompt("Matrix dimenstions: ")

			// Replace "," with "x" so user can enter J,K or JxK
			dim_string = strings.Replace(dim_string, ",", "x", 1)
			dimensions = strings.Split(dim_string, "x")

			if len(dimensions) != 2 {
				fmt.Println("Please enter matrix dimensions as JxK.")
				continue
			}

			rows, err1 = strconv.Atoi(dimensions[0])
			cols, err2 = strconv.Atoi(dimensions[1])

			if err1 != nil || err2 != nil {
				fmt.Println("Please use integers for array dimensions.")
			}
		}

		for r := 0; r < rows; r+=1 {
			for c:=0; c< cols; c+=1 {


		m := newMatrix(rows, cols)
		fmt.Println(m)
	*/
}
