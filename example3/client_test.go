package example3_test

import (
	"bytes"
	"go-test-example/example3"
	"io"
	"net/http"
	"testing"
)

// Create mock Roundtrip function definition
type RoundTripFunc func(*http.Request) (*http.Response, error)

// Create mock Roundtrip function
func (fn RoundTripFunc) RoundTrip(r *http.Request) (*http.Response, error) {
	return fn(r)
}

// NewTestClient returns *http.Client with Transport replaced to avoid making real calls
func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: RoundTripFunc(fn),
	}
}

// func Test_Get(t *testing.T) {
// 	// Setup
// 	url := "https://devops.exzeo.io/bb/ping"
// 	expect := "Hello World!"

// 	c := test.NewTestClient(func(req *http.Request) (*http.Response, error) {

// 	})

// 	// Act
// 	result, err := example3.Get(url)

// 	// Assert
// 	if err != nil {
// 		t.Fatalf("Error is not nil: %v", err)
// 	}

// 	if *result != expect {
// 		t.Fatalf("%s does not equal %s", *result, expect)
// 	}
// }

func Test_Get(t *testing.T) {
	// Setup
	url := "https://devops.exzeo.io/bb/ping"
	expect := "Hello World!"

	c := NewTestClient(func(req *http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBuffer([]byte(expect))),
		}, nil
	})

	client := example3.Client{
		HttpClient: c,
	}

	// Act
	result, err := client.Get(url)

	// Assert
	if err != nil {
		t.Fatalf("Error is not nil: %v", err)
	}

	if *result != expect {
		t.Fatalf("%s does not equal %s", *result, expect)
	}
}
