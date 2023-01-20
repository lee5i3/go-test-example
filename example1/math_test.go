package example1_test

import (
	"fmt"
	"go-test-example/example1"
	"testing"
)

func Test_Add(t *testing.T) {
	// Setup
	expected := 4

	// Act
	actual := example1.Add(2, 2)

	// Assert
	if expected != actual {
		t.Fatalf("%v does not equal %v", expected, actual)
	}
}

func Test_AddCase(t *testing.T) {
	type TestCase struct {
		A        int
		B        int
		Expected int
	}

	tests := []TestCase{
		{A: 1, B: 1, Expected: 2},
		{A: 2, B: 2, Expected: 4},
		{A: 20, B: 20, Expected: 40},
		{A: 36254, B: 11111, Expected: 47365},
	}

	// Setup
	for _, tc := range tests {
		t.Run(fmt.Sprintf("%v+%v=%v", tc.A, tc.B, tc.Expected), func(t *testing.T) {
			actual := example1.Add(tc.A, tc.B)

			// Assert
			if tc.Expected != actual {
				t.Fatalf("%v does not equal %v", tc.Expected, actual)
			}
		})
	}
}
