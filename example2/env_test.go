package example2_test

import (
	"go-test-example/example2"
	"os"
	"testing"
)

// func Test_GetEnv(t *testing.T) {
// 	// Setup
// 	key := "MyKey"
// 	expected := "MyValue"

// 	// Act
// 	actual := example2.Get(key)

// 	// Assert
// 	if expected != actual {
// 		t.Fatalf("%s does not equal %s", expected, actual)
// 	}
// }

func Test_GetEnv(t *testing.T) {
	// Setup
	key := "MyKey"
	expected := "MyValue"

	// Reset Environment Variables
	os.Clearenv()
	os.Setenv(key, expected)

	// Act
	actual := example2.Get(key)

	// Assert
	if expected != actual {
		t.Fatalf("%s does not equal %s", expected, actual)
	}
}
