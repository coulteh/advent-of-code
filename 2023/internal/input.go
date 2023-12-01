package internal

import (
	"bufio"
	"io"
	"os"
	"path/filepath"
)

const inputPath = "2023/input/"

func GetInputPath(s string) (string, error) {
	base, err := filepath.Abs(inputPath)
	if err != nil {
		return "", err
	}
	return filepath.Join(base, s), nil
}

// ReadFile takes a filepath as an argument and returns a string
// slice of the contents of the file
func ReadFile(filepath string) (*os.File, error) {
	return os.Open(filepath)
}

func ReaderToStrings(reader io.Reader) []string {
	scanner := bufio.NewScanner(reader)
	output := make([]string, 0)
	for scanner.Scan() {
		output = append(output, scanner.Text())
	}
	return output
}
