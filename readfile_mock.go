package examplemockfile

type MockFileReader struct {
	MockData  map[string][]byte
	MockError error
}

func (m MockFileReader) ReadFile(filename string) ([]byte, error) {
	if m.MockError != nil {
		return nil, m.MockError
	}
	return m.MockData[filename], nil
}
