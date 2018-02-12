/*
* @author Colton J. McCurdy
*	GitHub: mccurdyc
* Email:  mccurdyc22@gmail.com
* Date: 2018-02-15
*
*  usage:
*  go run main.go
*
*  OR
*
*  go build
*  ./example1
*
 */

// the first statement must be a package name.
// all files in a package must use the same package name
package main

import (
	"errors"
	"fmt"
)

func main() {
	// the := is not an operator at all; short variable declaration
	// https://golang.org/ref/spec#Short_variable_declarations
	sum := add(2, 3)
	fmt.Printf("sum: %d\n", sum)

	// _ just says "ignore" what is returned
	quotient1, _ := divide(4, 1)
	fmt.Printf("quotient1: %d\n", quotient1)

	quotient2, err := divide(2, 0)

	if err != nil {
		fmt.Printf("quotient2: %v\n", err)
	} else {
		fmt.Printf("quotient2: %d\n", quotient2)
	}
}

// add returns the integer sum of two integers
func add(x int, y int) int {
	return x + y
}

// divide returns the integer result of dividing x by y
// and returns an error if y is zero
func divide(x int, y int) (int, error) {
	if y == 0 {
		// we return 0 because this is the nil value of int
		return 0, errors.New("error dividing by zero")
	}

	return x / y, nil
}
