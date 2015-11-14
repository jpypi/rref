package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func prompt(prmpt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prmpt)
	text, _ := reader.ReadString('\n')

	return strings.Trim(text, "\n")
}

func getInput() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	return strings.Trim(text, "\n")
}

func getFracInput() (Fraction, error) {
	fraction := NewFrac(0, 1)
	var err error

	input := getInput()
	values := strings.Split(input, "/")

	if len(values) > 0 {
		fraction.numerator, err = strconv.Atoi(values[0])
	}

	if err != nil {
		return fraction, err
	}

	if len(values) == 2 {
		fraction.denominator, err = strconv.Atoi(values[1])
	} else if len(values) > 2 {
		return fraction, errors.New("Value must be a scalar e.g. 2 or 1/2")
	}

	return fraction, err
}

func main() {
	var dimensions []string
	var rows, cols int

	// Get input from user for dimensions to use for matrix
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

	m := newMatrix(rows, cols)

	for r := 0; r < rows; r += 1 {
		for c := 0; c < cols; c += 1 {
			set := false
			for set != true {
				fmt.Printf("Fraction for %d, %d: ", r, c)
				value, err := getFracInput()
				if err == nil {
					m.matrix[r][c] = value
					set = true
				} else {
					fmt.Println(err)
				}
			}

		}
	}

	if rows == cols {
		m = m.rref()
	}
	fmt.Println(m)
}
