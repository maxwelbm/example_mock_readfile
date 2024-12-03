package examplemockfile

import (
	"errors"
	"testing"
)

func TestProcessFile(t *testing.T) {
	mockReader := MockFileReader{
		MockData: map[string][]byte{
			"test.txt": []byte("Hello, World!"),
		},
		MockError: nil,
	}

	err := ProcessFile(mockReader, "test.txt")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test error case
	mockReader.MockError = errors.New("file not found")
	err = ProcessFile(mockReader, "missing.txt")
	if err == nil {
		t.Fatal("expected an error but got nil")
	}
}
