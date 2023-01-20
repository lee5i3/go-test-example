package example4_test

import (
	"bytes"
	"go-test-example/example4"
	"go-test-example/example4/mocks"
	"io"
	"net/http"
	"testing"
)

func Test_Get(t *testing.T) {
	// Setup
	url := "https://devops.exzeo.io/bb/ping"
	expect := "Hello World!"

	mockClient := new(mocks.IClientInterface)
	mockClient.On("Get", url).Return(&http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBuffer([]byte(expect))),
	}, nil)

	client := example4.Client{
		HttpClient: mockClient,
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
