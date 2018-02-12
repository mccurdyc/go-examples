package main

/*
* @author Colton J. McCurdy
*	GitHub: mccurdyc
* Email:  mccurdyc22@gmail.com
* Date: 2018-02-15
*
*  usage:
*  go test -v
*  go test -v -run=XXX
 */

import (
	"errors"
	"testing"
)

func TestAdd(t *testing.T) {
	s := add(3, 6)
	if s != 9 {
		t.Error("Expected 9, got ", s)
	}
}

// this is called a struct
// it has four fields with their respective types
type testValues struct {
	x           int
	y           int
	expected    int
	expectedErr error
}

func TestDivide(t *testing.T) {
	// this is a slice of testValues (our struct)
	values := []testValues{
		{x: 6, y: 3, expected: 2, expectedErr: nil},
		{2, 0, 0, errors.New("error dividing by zero")},
		{4, 3, 1, nil},
		{4, 1, 4, nil},
	}

	// we are ranging (looping) over the values in our slice
	// ignoring the index
	for _, v := range values {
		// we are accessing fields x and y of our testValues struct
		q, err := divide(v.x, v.y)

		if q != v.expected {
			t.Errorf("Expected: %d : Received %d \n", v.expected, q)
		}

		if (err != nil && v.expectedErr == nil) || (err == nil && v.expectedErr != nil) {
			t.Errorf("Expected: %v : Received %v \n", v.expectedErr, err)
		}
	}
}
