package examplemockfile

import (
	"fmt"
	"os"
)

type FileReader interface {
	ReadFile(filename string) ([]byte, error)
}

type OSFileReader struct{}

func (OSFileReader) ReadFile(filename string) ([]byte, error) {
	return os.ReadFile(filename)
}

func ProcessFile(reader FileReader, filename string) error {
	data, err := reader.ReadFile(filename)
	if err != nil {
		return err
	}
	fmt.Println("File content:", string(data))
	return nil
}
