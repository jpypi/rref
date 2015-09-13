package main

import (
	"bufio"
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

func main() {
	var dimensions []string
	var rows, cols int

	for len(dimensions) != 2 {
		var err1, err2 error
		dim_string := (prompt("Matrix dimenstions: "))
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
	m.matrix[0][1] = Fraction{numerator: 1, denominator: 5}
	m.matrix[1][0] = Fraction{numerator: 5, denominator: 1}
	fmt.Println(m)
}
